package npmsourcer_test

import (
	"fmt"
	"os"
	"testing"

	s "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/npm"
	tu "github.com/cloudrecipes/lambda-wrapper/internal/pkg/testutils"
)

var sourcer *s.NpmSourcer

func TestMain(m *testing.M) {
	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Expected to successfully clean up temporary directories before test")
		os.Exit(1)
	}

	if err := os.Mkdir(tu.Testdir, os.ModePerm); err != nil {
		fmt.Printf("\n>>> Expected err to be nil but got:\n%v", err)
		os.Exit(1)
	}

	sourcer = &s.NpmSourcer{}
	code := m.Run()

	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Temporary directories could not be cleaned")
	}

	os.Exit(code)
}

func TestLibGet(t *testing.T) {
	for _, test := range sourcerTestCases {
		_, err := sourcer.LibGet(test.libname, tu.Testdir)

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
	_, err := sourcer.LibTest(tu.Testdir)
	if err == nil {
		t.Fatalf("\n>>> Expected error not nil")
	}
}

func TestLibDeps(t *testing.T) {
	for _, test := range depsTestCases {
		_, err := sourcer.LibDeps(tu.Testdir, test.isprod)

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
