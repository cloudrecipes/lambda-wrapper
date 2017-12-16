package testutils

import (
	"os"
	"path"
)

// starting directory of dummy file structure
const Headdir string = ".lwtmp"

var Basedir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
	"lambda-wrapper", "test", "tmp")

var Testdir = path.Join(Basedir, Headdir)
