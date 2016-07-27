package gitscanner

import (
	"cred-alert/scanners"
	"strings"

	"github.com/pivotal-golang/lager"
)

type DiffScanner struct {
	diff              []string
	cursor            int
	currentHunk       *Hunk
	currentPath       string
	currentLineNumber int
}

func NewDiffScanner(diff string) *DiffScanner {
	return &DiffScanner{
		cursor: -1,
		diff:   strings.Split(diff, "\n"),
	}
}

func (d *DiffScanner) Scan(logger lager.Logger) bool {
	logger = logger.Session("diff-scanner").Session("scan")
	logger.Info("starting")

	// read information about hunk
	var isContentLine bool
	for !isContentLine {
		logger = logger.WithData(lager.Data{
			"line-number": d.cursor,
		})

		d.cursor++

		if d.cursor >= len(d.diff) {
			logger.Debug("passed-last-line")
			logger.Info("done")
			return false
		}

		logger.Debug("considering-line")
		rawLine := d.diff[d.cursor]

		d.scanHeader(logger, rawLine)
		isContentLine = d.scanHunk(logger, rawLine)
	}

	logger.Info("done")
	return true
}

func (d *DiffScanner) Line(logger lager.Logger) *scanners.Line {
	lineNumber := d.currentLineNumber
	path := d.currentHunk.path

	logger = logger.Session("line", lager.Data{
		"liner-number": lineNumber,
		"path":         path,
	})
	logger.Info("starting")

	content, err := content(d.diff[d.cursor])
	if err != nil {
		logger.Error("setting content to ''", err)
	}

	logger.Info("done")
	return &scanners.Line{
		Content:    content,
		LineNumber: lineNumber,
		Path:       path,
	}
}

func (d *DiffScanner) scanHeader(logger lager.Logger, rawLine string) {
	logger = logger.Session("scan-header", lager.Data{
		"current-line-number": d.currentLineNumber,
	})
	logger.Info("starting")

	nextLineNumber := d.currentLineNumber + 1

	if !isInHeader(nextLineNumber, d.currentHunk) {
		logger.Info("done")
		return
	}

	path, err := fileHeader(rawLine, nextLineNumber, d.currentHunk)
	if err == nil {
		logger.Debug("detected-file-header")
		d.currentPath = path
		d.currentHunk = nil
	}

	startLine, length, err := hunkHeader(logger, rawLine)
	if err == nil {
		logger.Debug("detected-hunk-header")
		d.currentHunk = newHunk(d.currentPath, startLine, length)

		// the hunk header exists immeidately before the first line
		d.currentLineNumber = startLine - 1
	}
}

func (d *DiffScanner) scanHunk(logger lager.Logger, rawLine string) bool {
	logger = logger.Session("scan-hunk", lager.Data{
		"current-line-number": d.currentLineNumber,
	})
	logger.Info("starting")
	nextLineNumber := d.currentLineNumber + 1

	if isInHeader(nextLineNumber, d.currentHunk) {
		logger.Info("done")
		return false
	}

	if contextOrAddedLine(rawLine) {
		logger.Debug("detected-content-line")
		d.currentLineNumber = nextLineNumber
		logger.Info("done")
		return true
	}

	logger.Info("done")
	return false
}