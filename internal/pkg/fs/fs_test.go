package fs_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
)

// starting directory of dummy file structure
const headdir string = ".lwtmp"

var basedir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
	"lambda-wrapper", "test", "tmp")

func createDummyDirStructure(basedir, headdir string) error {
	if err := os.Mkdir(path.Join(basedir, headdir), os.ModePerm); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(basedir, headdir, "blah"), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func createFile(filename, payload string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprint(f, payload)

	return nil
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

func TestDirGetters(t *testing.T) {
	var expected, actual string

	expected = ".lwtmp"
	actual = fs.WorkingDir()
	if expected != actual {
		t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", expected, actual)
	}

	expected = ".lwtmp/lib"
	actual = fs.LibDir()
	if expected != actual {
		t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", expected, actual)
	}

	expected = ".lwtmp/build"
	actual = fs.BuildDir()
	if expected != actual {
		t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", expected, actual)
	}
}

func TestMakeDirs(t *testing.T) {
	if err := fs.MakeDirs(basedir); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	if _, err := os.Stat(path.Join(basedir, headdir)); os.IsNotExist(err) {
		t.Fatal("\n>>> Expected working directory to be created")
	}

	if err := os.RemoveAll(path.Join(basedir, headdir)); err != nil {
		t.Fatal("\n>>> Expected to successfully clean up temporary directories")
	}
}

func TestRmDirs(t *testing.T) {
	if err := createDummyDirStructure(basedir, headdir); err != nil {
		t.Fatalf("\n>>> Expected err to be nil but got:\n%v", err)
	}

	if err := fs.RmDirs(basedir); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	if _, err := os.Stat(path.Join(basedir, headdir)); os.IsExist(err) {
		t.Fatal("\n>>> Expected working directory to be deleted but it still exists")
	}
}

func TestZipDir(t *testing.T) {
	if err := createDummyDirStructure(basedir, headdir); err != nil {
		t.Fatalf("\n>>> Expected err to be nil but got:\n%v", err)
	}

	for _, f := range filesToZip {
		if err := createFile(f.filename, f.payload); err != nil {
			t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
		}
	}

	if err := fs.ZipDir(basedir, "test.zip"); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	// TODO: check archive

	if err := os.RemoveAll(path.Join(basedir, headdir)); err != nil {
		t.Fatal("\n>>> Expected to successfully clean up temporary directories")
	}
}
