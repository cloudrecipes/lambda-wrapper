package gitsourcer_test

import (
	"fmt"
	"os"
	"testing"

	s "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/git"
	tu "github.com/cloudrecipes/lambda-wrapper/internal/pkg/testutils"
)

var sourcer *s.GitSourcer

func TestMain(m *testing.M) {
	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Expected to successfully clean up temporary directories before test")
		os.Exit(1)
	}

	if err := os.Mkdir(tu.Testdir, os.ModePerm); err != nil {
		fmt.Printf("\n>>> Expected err to be nil but got:\n%v", err)
		os.Exit(1)
	}

	sourcer = &s.GitSourcer{}
	code := m.Run()

	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Temporary directories could not be cleaned")
	}

	os.Exit(code)
}

func TestLibGet(t *testing.T) {
	for _, test := range sourcerTestCases {
		// // remove .git directory to avoid git clone errors
		// if err := os.RemoveAll(path.Join(tu.Testdir, ".git")); err != nil {
		// 	t.Fatal("\n>>> Expected to successfully clean up temporary directories")
		// }

		err := sourcer.LibGet(test.libname, tu.Testdir)

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

func TestLibTest(t *testing.T) {
	err := sourcer.LibTest(tu.Testdir)
	if err != nil {
		t.Fatalf("\n>>> Expected error:\nnil\n<<< but got:\n%v", err)
	}
}

func TestLibDeps(t *testing.T) {
	err := sourcer.LibDeps(tu.Testdir, false)
	if err != nil {
		t.Fatalf("\n>>> Expected error:\nnil\n<<< but got:\n%v", err)
	}
}
