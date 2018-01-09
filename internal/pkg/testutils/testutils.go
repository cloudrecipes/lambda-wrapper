package testutils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
)

// Fixturesdir is a test fixtures dir.
var Fixturesdir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
	"lambda-wrapper", "test", "fixtures")

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

// EnvVarsForTestMocks returns an array of additional environment variables
// for use in TestCommander.
func EnvVarsForTestMocks(namespace, expected string, err error) []string {
	code := "0"
	if err != nil {
		code = "1"
	}

	return []string{
		fmt.Sprintf("GO_TEST_%s_EXPECTED=%s", namespace, expected),
		fmt.Sprintf("GO_TEST_%s_EXIT_CODE=%s", namespace, code),
	}
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

// TestFs is a test implementation of fs.I interface.
type TestFs struct{}

// GoTestFsReadFileExpected env variable name.
const GoTestFsReadFileExpected = "GO_TEST_FS_READ_FILE_EXPECTED"

// GoTestFsReadFileError env variable name.
const GoTestFsReadFileError = "GO_TEST_FS_READ_FILE_ERROR"

// ReadFile is test implementation of read file to string.
func (fs *TestFs) ReadFile(filename string) (string, error) {
	var err error
	payload := os.Getenv(GoTestFsReadFileExpected)
	errmessage := os.Getenv(GoTestFsReadFileError)

	if len(errmessage) > 0 {
		err = errors.New(errmessage)
	}

	return payload, err
}

// GoTestFsReadFileToBytesExpected env variable name.
const GoTestFsReadFileToBytesExpected = "GO_TEST_FS_READ_FILE_TO_BYTES_EXPECTED"

// GoTestFsReadFileToBytesError env variable name.
const GoTestFsReadFileToBytesError = "GO_TEST_FS_READ_FILE_TO_BYTES_ERROR"

// ReadFileToBytes is a test implementation of read file to bytes.
func (fs *TestFs) ReadFileToBytes(filename string) ([]byte, error) {
	var payload []byte
	var err error

	payloadstr := os.Getenv(GoTestFsReadFileToBytesExpected)
	errmessage := os.Getenv(GoTestFsReadFileToBytesError)

	if len(payloadstr) > 0 {
		payload = []byte(payloadstr)
	}

	if len(errmessage) > 0 {
		err = errors.New(errmessage)
	}

	return payload, err
}

// GoTestFsRmDirError env variable name.
const GoTestFsRmDirError = "GO_TEST_FS_RM_DIR_ERROR"

// RmDir is a test implementation of remove directory.
func (fs *TestFs) RmDir(basedir string) error {
	var err error

	errmessage := os.Getenv(GoTestFsRmDirError)

	if len(errmessage) > 0 {
		err = errors.New(errmessage)
	}

	return err
}

// GoTestFsZipDirError env variable name.
const GoTestFsZipDirError = "GO_TEST_FS_ZIP_DIR_ERROR"

// ZipDir is a test implementaiton of zip directory.
func (fs *TestFs) ZipDir(source, target string) error {
	var err error

	errmessage := os.Getenv(GoTestFsZipDirError)

	if len(errmessage) > 0 {
		err = errors.New(errmessage)
	}

	return err
}
