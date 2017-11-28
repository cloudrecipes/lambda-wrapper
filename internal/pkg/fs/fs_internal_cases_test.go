package fs

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

type BrokenZipWriter struct{}

func (w *BrokenZipWriter) CreateHeader(*zip.FileHeader) (io.Writer, error) {
	return nil, errors.New("Create Header")
}

// make sure it satisfies the interface
var _ zipWriter = (*BrokenZipWriter)(nil)

var brokenWriter = &BrokenZipWriter{}

var filepathWalkTestCases = []struct {
	basedir  string
	source   string
	archive  zipWriter
	path     string
	info     os.FileInfo
	err      error
	expected error
}{
	{basedir: "", source: "", archive: nil, path: "", info: nil, err: errors.New("Test Error"), expected: errors.New("Test Error")},
	// {basedir: "", source: "", archive: &BrokenWriter{}, path: "blah", info: nil, err: nil, expected: errors.New("Test Error")},
}
