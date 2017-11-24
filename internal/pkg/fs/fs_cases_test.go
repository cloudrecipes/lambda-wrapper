package fs_test

import (
	"errors"
	"os"
	"path"
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
		filename: path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
			"lambda-wrapper", "test", "fixtures", "fs_readfile.txt"),
		err:      nil,
		expected: "Hello Test!",
	},
}

var filesToZip = []struct {
	filename string
	payload  string
}{
	{filename: path.Join(basedir, headdir, "file1.txt"), payload: "test file1"},
	{filename: path.Join(basedir, headdir, "blah", "file2.txt"), payload: "test file2"},
}
