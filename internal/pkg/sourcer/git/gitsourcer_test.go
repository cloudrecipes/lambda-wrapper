package gitsourcer_test

import (
	"os"
	"path"
	"testing"

	s "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/git"
)

// starting directory of dummy file structure
const headdir string = ".lwtmp"

var basedir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
	"lambda-wrapper", "test", "tmp")

func TestLibGet(t *testing.T) {
	sourcer := &s.GitSourcer{}

	destination := path.Join(basedir, headdir)

	// create temporary directory to store library
	if err := os.Mkdir(destination, os.ModePerm); err != nil {
		t.Fatalf("\n>>> Expected err to be nil but got:\n%v", err)
	}

	for _, test := range sourcerTestCases {
		err := sourcer.LibGet(test.libname, destination)

		// remove .git directory
		if err := os.RemoveAll(path.Join(destination, ".git")); err != nil {
			t.Fatal("\n>>> Expected to successfully clean up temporary directories")
		}

		if test.err != nil {
			if err == nil || test.err.Error() != err.Error() {
				t.Fatalf("\n>>> Expected error:\n%v\n<<< but got:\n%v", test.err, err)
			}
			continue
		}

		if test.err == nil && err != nil {
			t.Fatalf("\n>>> Expected error:\nnil\n<<< but got:\n%v", err)
		}
	}

	// remove temporary directory
	if err := os.RemoveAll(destination); err != nil {
		t.Fatal("\n>>> Expected to successfully clean up temporary directories")
	}
}
