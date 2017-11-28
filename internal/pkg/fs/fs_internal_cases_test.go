package fs

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"time"
)

type BrokenZipWriter struct{}

func (w *BrokenZipWriter) CreateHeader(*zip.FileHeader) (io.Writer, error) {
	return nil, errors.New("Create Header Error")
}

// make sure it satisfies the interface
var _ zipWriter = (*BrokenZipWriter)(nil)

var brokenZipWriter = &BrokenZipWriter{}

type TestFileInfo struct{}

func (i *TestFileInfo) Name() string {
	return ""
}

func (i *TestFileInfo) IsDir() bool {
	return false
}

func (i *TestFileInfo) Mode() os.FileMode {
	return os.ModePerm
}

func (i *TestFileInfo) ModTime() time.Time {
	return time.Now()
}

func (i *TestFileInfo) Size() int64 {
	return 0
}

func (i *TestFileInfo) Sys() interface{} {
	return nil
}

var _ os.FileInfo = (*TestFileInfo)(nil)

var testFileInfo = &TestFileInfo{}

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
	{basedir: "", source: "", archive: brokenZipWriter, path: "blah", info: testFileInfo, err: nil, expected: errors.New("Create Header Error")},
}
