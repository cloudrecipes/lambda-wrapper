package npmsourcer_test

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	s "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/npm"
	tu "github.com/cloudrecipes/lambda-wrapper/internal/pkg/testutils"
)

var sourcer *s.NpmSourcer

// func TestMain(m *testing.M) {
// 	if err := os.RemoveAll(tu.Testdir); err != nil {
// 		fmt.Println("\n>>> Expected to successfully clean up temporary directories before test")
// 		os.Exit(1)
// 	}
//
// 	if err := os.Mkdir(tu.Testdir, os.ModePerm); err != nil {
// 		fmt.Printf("\n>>> Expected err to be nil but got:\n%v", err)
// 		os.Exit(1)
// 	}
//
// 	sourcer = &s.NpmSourcer{}
// 	code := m.Run()
//
// 	if err := os.RemoveAll(tu.Testdir); err != nil {
// 		fmt.Println("\n>>> Temporary directories could not be cleaned")
// 	}
//
// 	os.Exit(code)
// }

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_TEST_HELPER_PROCESS") != "1" {
		return
	}

	expected := os.Getenv("GO_TEST_NPMSOURCER_EXPECTED")
	code := 0
	i, err := strconv.Atoi(os.Getenv("GO_TEST_NPMSOURCER_EXIT_CODE"))

	if err == nil {
		code = i
	}

	if code == 0 {
		fmt.Print(expected)
	} else {
		fmt.Fprint(os.Stderr, expected)
	}

	defer os.Exit(code)
}

func TestLibGet(t *testing.T) {
	if err := tu.CreateTestDirStructure(); err != nil {
		t.Fatalf("\n>>> Expected err to be nil but got:\n%v", err)
	}

	for _, test := range sourcerTestCases {
		envvars := tu.EnvVarsForCommander("NPMSOURCER", test.expected, test.err)
		commander := &tu.TestCommander{EnvVars: envvars}
		out, err := sourcer.LibGet(commander, test.libname, tu.Testdir)
		actual := string(out[:])

		if test.err != nil {
			if err == nil || test.err.Error() != err.Error() {
				t.Fatalf("\n>>> Expected error:\n%v\n<<< but got:\n%v", test.err, err)
			}
			continue
		}

		if test.err == nil && err != nil {
			t.Fatalf("\n>>> Expected error:\nnil\n<<< but got:\n%v", err)
		}

		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}

	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Temporary directories could not be cleaned")
	}
}

func TestLibTest(t *testing.T) {
	envvars := tu.EnvVarsForCommander("NPMSOURCER", "", errors.New("LibTest error"))
	commander := &tu.TestCommander{EnvVars: envvars}
	_, err := sourcer.LibTest(commander, tu.Testdir)

	if err == nil {
		t.Fatalf("\n>>> Expected error not nil")
	}
}

func TestLibDeps(t *testing.T) {
	for _, test := range depsTestCases {
		envvars := tu.EnvVarsForCommander("NPMSOURCER", test.expected, test.err)
		commander := &tu.TestCommander{EnvVars: envvars}
		_, err := sourcer.LibDeps(commander, tu.Testdir, test.isprod)

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
