package revok

import (
	"cred-alert/db"
	"cred-alert/gitclient"
	"cred-alert/kolsch"
	"cred-alert/metrics"
	"cred-alert/notifications"
	"cred-alert/scanners"
	"cred-alert/scanners/diffscanner"
	"cred-alert/sniff"
	"strings"

	"code.cloudfoundry.org/lager"
	git "github.com/libgit2/git2go"
)

//go:generate counterfeiter . Scanner

type Scanner interface {
	Scan(lager.Logger, string, string, map[git.Oid]struct{}, string, string) error
	ScanNoNotify(lager.Logger, string, string, string, string) ([]db.Credential, error)
}

type scanner struct {
	gitClient            gitclient.Client
	repositoryRepository db.RepositoryRepository
	scanRepository       db.ScanRepository
	credentialRepository db.CredentialRepository
	sniffer              sniff.Sniffer
	notifier             notifications.Notifier
}

func NewScanner(
	gitClient gitclient.Client,
	repositoryRepository db.RepositoryRepository,
	scanRepository db.ScanRepository,
	credentialRepository db.CredentialRepository,
	sniffer sniff.Sniffer,
	notifier notifications.Notifier,
	emitter metrics.Emitter,
) Scanner {
	return &scanner{
		gitClient:            gitClient,
		repositoryRepository: repositoryRepository,
		scanRepository:       scanRepository,
		credentialRepository: credentialRepository,
		sniffer:              sniffer,
		notifier:             notifier,
	}
}

func (s *scanner) Scan(
	logger lager.Logger,
	owner string,
	repository string,
	scannedOids map[git.Oid]struct{},
	startSHA string,
	stopSHA string,
) error {
	dbRepository, err := s.repositoryRepository.Find(owner, repository)
	if err != nil {
		logger.Error("failed-to-find-db-repo", err)
		return err
	}

	credentials, err := s.scan(logger, dbRepository, scannedOids, startSHA, stopSHA)
	if err != nil {
		return err
	}

	var batch []notifications.Notification

	for _, credential := range credentials {
		batch = append(batch, notifications.Notification{
			Owner:      credential.Owner,
			Repository: credential.Repository,
			SHA:        credential.SHA,
			Path:       credential.Path,
			LineNumber: credential.LineNumber,
			Private:    dbRepository.Private,
		})
	}

	if batch != nil {
		err = s.notifier.SendBatchNotification(logger, batch)
		if err != nil {
			logger.Error("failed", err)
			return err
		}
	}

	return nil
}

func (s *scanner) ScanNoNotify(
	logger lager.Logger,
	owner string,
	repository string,
	startSHA string,
	stopSHA string,
) ([]db.Credential, error) {
	dbRepository, err := s.repositoryRepository.Find(owner, repository)
	if err != nil {
		logger.Error("failed-to-find-db-repo", err)
		return nil, err
	}

	scannedOids := map[git.Oid]struct{}{}
	credentials, err := s.scan(logger, dbRepository, scannedOids, startSHA, stopSHA)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func (s *scanner) scan(
	logger lager.Logger,
	dbRepository db.Repository,
	scannedOids map[git.Oid]struct{},
	startSHA string,
	stopSHA string,
) ([]db.Credential, error) {
	repo, err := git.OpenRepository(dbRepository.Path)
	if err != nil {
		logger.Error("failed-to-open-repo", err)
		return nil, err
	}

	startOid, err := git.NewOid(startSHA)
	if err != nil {
		logger.Error("failed-to-create-start-oid", err)
		return nil, err
	}

	var stopOid *git.Oid
	if stopSHA != "" {
		var err error
		stopOid, err = git.NewOid(stopSHA)
		if err != nil {
			logger.Error("failed-to-create-stop-oid", err)
			return nil, err
		}
	}

	quietLogger := kolsch.NewLogger()
	scan := s.scanRepository.Start(quietLogger, "repo-scan", startSHA, stopSHA, &dbRepository, nil)

	var credentials []db.Credential

	scanFunc := func(child, parent *git.Oid) error {
		diff, err := s.gitClient.Diff(dbRepository.Path, parent, child)
		if err != nil {
			return err
		}

		s.sniffer.Sniff(
			quietLogger,
			diffscanner.NewDiffScanner(strings.NewReader(diff)),
			func(logger lager.Logger, violation scanners.Violation) error {
				credential := db.NewCredential(
					dbRepository.Owner,
					dbRepository.Name,
					child.String(),
					violation.Line.Path,
					violation.Line.LineNumber,
					violation.Start,
					violation.End,
					dbRepository.Private,
				)

				scan.RecordCredential(credential)
				credentials = append(credentials, credential)

				return nil
			},
		)

		scannedOids[*child] = struct{}{}

		return nil
	}

	knownSHAs := map[string]struct{}{}
	shas, err := s.credentialRepository.UniqueSHAsForRepoAndRulesVersion(dbRepository, sniff.RulesVersion)
	for i := range shas {
		knownSHAs[shas[i]] = struct{}{}
	}

	err = s.scanAncestors(repo, scanFunc, scannedOids, knownSHAs, startOid, stopOid)
	if err != nil {
		logger.Error("failed-to-scan-ancestors", err, lager.Data{
			"start":      startOid.String(),
			"stop":       stopOid.String(),
			"repository": dbRepository.Name,
		})
	}

	err = scan.Finish()
	if err != nil {
		logger.Error("failed-to-finish-scan", err)
		return nil, err
	}

	return credentials, nil
}

func (s *scanner) scanAncestors(
	repo *git.Repository,
	scanFunc func(*git.Oid, *git.Oid) error,
	scannedOids map[git.Oid]struct{},
	knownSHAs map[string]struct{},
	child *git.Oid,
	stopPoint *git.Oid,
) error {
	if _, found := scannedOids[*child]; found {
		return nil
	}

	if _, found := knownSHAs[child.String()]; found {
		return nil
	}

	parents, err := s.gitClient.GetParents(repo, child)
	if err != nil {
		return err
	}

	if len(parents) == 0 {
		return scanFunc(child, nil)
	}

	if len(parents) == 1 {
		err = scanFunc(child, parents[0])
		if err != nil {
			return err
		}
	}

	for _, parent := range parents {
		if stopPoint != nil && parent.Equal(stopPoint) {
			continue
		}

		err = s.scanAncestors(repo, scanFunc, scannedOids, knownSHAs, parent, stopPoint)
		if err != nil {
			return err
		}
	}

	return nil
}
