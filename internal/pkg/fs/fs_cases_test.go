package fs_test

import (
	"errors"
	"path"

	tu "github.com/cloudrecipes/lambda-wrapper/internal/pkg/testutils"
)

var readFileTestCases = []struct {
	filename string // filename to read payload
	err      error  // an error flag
	expected string // expected file payload
}{
	{
		filename: "no_such_file.txt",
		err:      errors.New("open no_such_file.txt: no such file or directory"),
		expected: "",
	},
	{
		filename: path.Join(tu.Fixturesdir, "fs_readfile.txt"),
		err:      nil,
		expected: "Hello Test!",
	},
}

var filesToZip = []struct {
	filename string
	payload  string
}{
	{filename: path.Join(tu.Testdir, "file1.txt"), payload: "test file1"},
	{filename: path.Join(tu.Testdir, "blah", "file2.txt"), payload: "test file2"},
}

var zipDirErrorTestCases = []struct {
	source   string
	target   string
	expected error
}{
	{source: "", target: "", expected: errors.New("open : no such file or directory")},
	{source: "", target: path.Join(tu.Basedir, "test.zip"), expected: errors.New("stat : no such file or directory")},
}
