package fs_test

import (
	"bytes"
	"fmt"
	"os"
	"path"
	s "strings"
	"testing"

	f "github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
	tu "github.com/cloudrecipes/lambda-wrapper/internal/pkg/testutils"
)

var fs = &f.Fs{}

func createDummyFiles() error {
	for _, f := range filesToZip {
		if err := tu.CreateFile(f.filename, f.payload); err != nil {
			return err
		}
	}

	return nil
}

func TestMain(m *testing.M) {
	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Expected to successfully clean up temporary directories before test")
		os.Exit(1)
	}

	code := m.Run()

	if err := os.RemoveAll(tu.Testdir); err != nil {
		fmt.Println("\n>>> Temporary directories could not be cleaned")
	}

	os.Exit(code)
}

func TestReadFile(t *testing.T) {
	for _, test := range readFileTestCases {
		actual, err := fs.ReadFile(test.filename)

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
}

func TestReadFileAsBytes(t *testing.T) {
	for _, test := range readFileTestCases {
		actual, err := fs.ReadFileToBytes(test.filename)

		if test.err != nil {
			if err == nil || test.err.Error() != err.Error() {
				t.Fatalf("\n>>> Expected error:\n%v\n<<< but got:\n%v", test.err, err)
			}
			continue
		}

		if test.err == nil && err != nil {
			t.Fatalf("\n>>> Expected error:\nnil\n<<< but got:\n%v", err)
		}

		if !bytes.Equal([]byte(test.expected), actual) {
			t.Fatalf("\n>>> Bytes slices are not equal")
		}
	}
}

func TestDirGetters(t *testing.T) {
	var expected, actual string

	expected = ".lwtmp"
	actual = f.WorkingDir()
	if expected != actual {
		t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", expected, actual)
	}

	expected = ".lwtmp/lib"
	actual = f.LibDir()
	if expected != actual {
		t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", expected, actual)
	}

	expected = ".lwtmp/build"
	actual = f.BuildDir()
	if expected != actual {
		t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", expected, actual)
	}
}

func TestMakeDirs(t *testing.T) {
	if err := f.MakeDirs(tu.Basedir); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	if _, err := os.Stat(tu.Testdir); os.IsNotExist(err) {
		t.Fatal("\n>>> Expected working directory to be created")
	}

	if err := os.RemoveAll(tu.Testdir); err != nil {
		t.Fatal("\n>>> Expected to successfully clean up temporary directories")
	}
}

func TestMakeDirsErrorCase(t *testing.T) {
	if err := f.MakeDirs("blah"); err == nil {
		t.Fatal("\n>>> Expected err not to be nil")
	}
}

func TestRmDir(t *testing.T) {
	if err := tu.CreateTestDirStructure(); err != nil {
		t.Fatalf("\n>>> Expected err to be nil but got:\n%v", err)
	}

	if err := fs.RmDir(tu.Basedir); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	if _, err := os.Stat(tu.Testdir); os.IsExist(err) {
		t.Fatal("\n>>> Expected working directory to be deleted but it still exists")
	}
}

func TestZipDir(t *testing.T) {
	if err := tu.CreateTestDirStructure(); err != nil {
		t.Fatalf("\n>>> Expected err to be nil but got:\n%v", err)
	}

	if err := createDummyFiles(); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	source := tu.Testdir
	target := path.Join(tu.Basedir, "test.zip")
	if err := fs.ZipDir(source, target); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	// TODO: check archive

	if err := os.RemoveAll(source); err != nil {
		t.Fatal("\n>>> Expected to successfully clean up temporary directories")
	}

	if err := os.Remove(target); err != nil {
		t.Fatal("\n>>> Expected to successfully clean up archive file")
	}
}

func TestZipDirError(t *testing.T) {
	for _, test := range zipDirErrorTestCases {
		actual := fs.ZipDir(test.source, test.target)

		if s.Compare(test.expected.Error(), actual.Error()) != 0 {
			t.Fatalf("\n>>> Expected:\n%v\n<<< but got:\n%v", test.expected, actual)
		}

		if s.Compare("", test.target) != 0 {
			if err := os.Remove(test.target); err != nil {
				t.Fatal("\n>>> Expected to successfully clean up archive file")
			}
		}
	}
}
