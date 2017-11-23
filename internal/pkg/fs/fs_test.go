package fs_test

import (
	"os"
	"path"
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
)

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
	var basedir string
	var err error

	if basedir, err = os.Getwd(); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	if err = fs.MakeDirs(basedir); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	if _, err = os.Stat(path.Join(basedir, ".lwtmp")); os.IsNotExist(err) {
		t.Fatal("\n>>> Expected working directory to be created")
	}

	if err = os.RemoveAll(path.Join(basedir, ".lwtmp")); err != nil {
		t.Fatal("\n>>> Expected to successfully clean up temporary directories")
	}
}

func TestRmDirs(t *testing.T) {
	var basedir string
	var err error

	if basedir, err = os.Getwd(); err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	if err = os.Mkdir(path.Join(basedir, ".lwtmp"), os.ModePerm); err != nil {
		t.Fatalf("\n>>> Expected err to be nil (.lwtmp) but got:\n%v", err)
	}

	if err = os.Mkdir(path.Join(basedir, ".lwtmp", "blah"), os.ModePerm); err != nil {
		t.Fatalf("\n>>> Expected err to be nil (.lwtmp/blah) but got:\n%v", err)
	}

	err = fs.RmDirs(basedir)
	if err != nil {
		t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
	}

	if _, err = os.Stat(path.Join(basedir, ".lwtmp")); os.IsExist(err) {
		t.Fatal("\n>>> Expected working directory to be deleted but it still exists")
	}
}
