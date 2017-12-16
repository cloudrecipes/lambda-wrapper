package testutils

import (
	"fmt"
	"os"
	"path"
)

// Basedir is a base temporary directory.
var Basedir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
	"lambda-wrapper", "test", "tmp")

// Headdir is a starting directory of dummy file structure.
const Headdir string = ".lwtmp"

// Testdir is a testing directory in Basedir.
var Testdir = path.Join(Basedir, Headdir)

// CreateDummyDirStructure creates directory structure in Testdir.
func CreateDummyDirStructure() error {
	if err := os.Mkdir(Testdir, os.ModePerm); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(Testdir, "blah"), os.ModePerm); err != nil {
		return err
	}

	return nil
}

// CreateFile creates dummy files for tests.
func CreateFile(filename, payload string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprint(f, payload)

	return nil
}
