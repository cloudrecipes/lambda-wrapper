package wrapper_test

import (
	"os"
	"path"
	"testing"

	tu "github.com/cloudrecipes/lambda-wrapper/pkg/testutils"
	"github.com/cloudrecipes/lambda-wrapper/pkg/wrapper"
)

func TestWrap(t *testing.T) {
	w := &testWrapper{}
	for _, test := range wrapTestCases {
		actual, err := wrapper.Wrap(w, test.options, test.templatedir)

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

func TestDefaultTemplateDir(t *testing.T) {
	expected := path.Join(os.Getenv("GOPATH"), "src", "github.com",
		"cloudrecipes", "lambda-wrapper", "assets", "templates")
	actual := wrapper.DefaultTemplateDir()

	if expected != actual {
		t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", expected, actual)
	}
}

func TestSave(t *testing.T) {
	expected := "test error"
	os.Setenv(tu.GoTestFsWriteFileError, expected)
	err := wrapper.Save("file.txt", "payload", &tu.TestFs{})

	if err == nil || err.Error() != expected {
		t.Fatalf("\n>>> Expected error:\n%s\n<<< but got:\n%v", expected, err)
	}
}
