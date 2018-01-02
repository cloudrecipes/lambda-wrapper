package testutils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

// Basedir is a base temporary directory.
var Basedir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
	"lambda-wrapper", "test", "tmp")

// Headdir is a starting directory of dummy file structure.
const Headdir string = ".lwtmp"

// Testdir is a testing directory in Basedir.
var Testdir = path.Join(Basedir, Headdir)

// CreateTestDirStructure creates directory structure in Testdir.
func CreateTestDirStructure() error {
	if err := os.RemoveAll(Testdir); err != nil {
		return err
	}

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

// TestCommander is a test implementation of the Commander interface.
type TestCommander struct {
	EnvVars []string
}

// CombinedOutput creates mock of Commander.
func (c TestCommander) CombinedOutput(command string, args ...string) ([]byte, error) {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	env := []string{"GO_TEST_HELPER_PROCESS=1"}

	for _, v := range c.EnvVars {
		env = append(env, v)
	}

	cmd.Env = env
	out, err := cmd.CombinedOutput()
	return out, err
}

// EnvVarsForCommander returns an array of additional environment variables
// for use in TestCommander.
func EnvVarsForCommander(namespace, expected string, err error) []string {
	code := "0"
	if err != nil {
		code = "1"
	}

	return []string{
		fmt.Sprintf("GO_TEST_%s_EXPECTED=%s", namespace, expected),
		fmt.Sprintf("GO_TEST_%s_EXIT_CODE=%s", namespace, code),
	}
}
