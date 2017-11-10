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
		"no_such_file.txt",
		errors.New("open no_such_file.txt: no such file or directory"),
		"",
	},
	{
		path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
			"lambda-wrapper", "test", "fixtures", "fs_readfile.txt"),
		nil,
		"Hello Test!",
	},
}
