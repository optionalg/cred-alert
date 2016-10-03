package revok

import (
	"cred-alert/db"
	"cred-alert/gitclient"
	"cred-alert/kolsch"
	"cred-alert/metrics"
	"cred-alert/scanners"
	"cred-alert/scanners/diffscanner"
	"cred-alert/sniff"
	"strings"

	"code.cloudfoundry.org/lager"
	git "github.com/libgit2/git2go"
)

//go:generate counterfeiter . Scanner

type Scanner interface {
	Scan(lager.Logger, string, string, string, string) error
}

type scanner struct {
	gitClient            gitclient.Client
	repositoryRepository db.RepositoryRepository
	scanRepository       db.ScanRepository
	sniffer              sniff.Sniffer
}

func NewScanner(
	gitClient gitclient.Client,
	repositoryRepository db.RepositoryRepository,
	scanRepository db.ScanRepository,
	sniffer sniff.Sniffer,
	emitter metrics.Emitter,
) Scanner {
	return &scanner{
		gitClient:            gitClient,
		repositoryRepository: repositoryRepository,
		scanRepository:       scanRepository,
		sniffer:              sniffer,
	}
}

func (s *scanner) Scan(
	logger lager.Logger,
	owner string,
	repository string,
	startSHA string,
	stopSHA string,
) error {
	dbRepository, err := s.repositoryRepository.Find(owner, repository)
	if err != nil {
		logger.Error("failed-to-find-db-repo", err)
		return err
	}

	repo, err := git.OpenRepository(dbRepository.Path)
	if err != nil {
		logger.Error("failed-to-open-repo", err)
		return err
	}

	startOid, err := git.NewOid(startSHA)
	if err != nil {
		logger.Error("failed-to-create-start-oid", err)
		return err
	}

	var stopOid *git.Oid
	if stopSHA != "" {
		var err error
		stopOid, err = git.NewOid(stopSHA)
		if err != nil {
			logger.Error("failed-to-create-stop-oid", err)
			return err
		}
	}

	quietLogger := kolsch.NewLogger()
	scan := s.scanRepository.Start(quietLogger, "repo-scan", &dbRepository, nil)
	scannedOids := map[git.Oid]struct{}{}
	scanFunc := func(child, parent *git.Oid) error {
		diff, err := s.gitClient.Diff(dbRepository.Path, parent, child)
		if err != nil {
			return err
		}

		s.sniffer.Sniff(
			quietLogger,
			diffscanner.NewDiffScanner(strings.NewReader(diff)),
			func(logger lager.Logger, violation scanners.Violation) error {
				scan.RecordCredential(db.NewCredential(
					dbRepository.Owner,
					dbRepository.Name,
					child.String(),
					violation.Line.Path,
					violation.Line.LineNumber,
					violation.Start,
					violation.End,
				))
				return nil
			},
		)

		scannedOids[*child] = struct{}{}

		return nil
	}

	err = s.scanAncestors(repo, scanFunc, scannedOids, startOid, stopOid)
	if err != nil {
		logger.Error("failed-to-scan", err)
	}

	err = scan.Finish()
	if err != nil {
		logger.Error("failed-to-finish-scan", err)
		return err
	}

	return nil
}

func (s *scanner) scanAncestors(
	repo *git.Repository,
	scanFunc func(*git.Oid, *git.Oid) error,
	scannedOids map[git.Oid]struct{},
	child *git.Oid,
	stopPoint *git.Oid,
) error {
	parents, err := s.gitClient.GetParents(repo, child)
	if err != nil {
		return err
	}

	if len(parents) == 0 {
		return scanFunc(child, nil)
	}

	for _, parent := range parents {
		if _, found := scannedOids[*parent]; found {
			continue
		}

		err = scanFunc(child, parent)
		if err != nil {
			return err
		}

		if stopPoint != nil && parent.Equal(stopPoint) {
			continue
		}

		return s.scanAncestors(repo, scanFunc, scannedOids, parent, stopPoint)
	}

	return nil
}