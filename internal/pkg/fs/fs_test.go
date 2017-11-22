package fs_test

import (
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
