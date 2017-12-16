package gitsourcer_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	s "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/git"
	tu "github.com/cloudrecipes/lambda-wrapper/internal/pkg/testutils"
)

func TestMain(m *testing.M) {
	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Expected to successfully clean up temporary directories before test")
		os.Exit(1)
	}

	if err := os.Mkdir(tu.Testdir, os.ModePerm); err != nil {
		fmt.Printf("\n>>> Expected err to be nil but got:\n%v", err)
		os.Exit(1)
	}

	code := m.Run()

	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Temporary directories could not be cleaned")
	}

	os.Exit(code)
}

func TestLibGet(t *testing.T) {
	sourcer := &s.GitSourcer{}

	for _, test := range sourcerTestCases {
		err := sourcer.LibGet(test.libname, tu.Testdir)

		// remove .git directory to avoid git clone errors
		if err := os.RemoveAll(path.Join(tu.Testdir, ".git")); err != nil {
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
}
