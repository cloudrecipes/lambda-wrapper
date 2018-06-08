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

type ValidZipWriter struct{}

func (w *ValidZipWriter) CreateHeader(*zip.FileHeader) (io.Writer, error) {
	return nil, nil
}

// make sure it satisfies the interface
var _ zipWriter = (*ValidZipWriter)(nil)

var validZipWriter = &ValidZipWriter{}

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
	source   string
	archive  zipWriter
	path     string
	info     os.FileInfo
	err      error
	expected error
}{
	{source: "", archive: nil, path: "", info: nil, err: errors.New("Test Error"), expected: errors.New("Test Error")},
	{source: "", archive: brokenZipWriter, path: "blah", info: testFileInfo, err: nil, expected: errors.New("Create Header Error")},
	{source: "", archive: validZipWriter, path: "blah", info: testFileInfo, err: nil, expected: errors.New("open blah: no such file or directory")},
}
