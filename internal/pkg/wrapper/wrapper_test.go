package wrapper_test

import (
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper"
)

func TestBuildTemplateFileName(t *testing.T) {
	for _, test := range buildTemplateFileNameTestCases {
		actual := wrapper.BuildTemplateFileName(test.cloud, test.engine)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestReadTemplateFile(t *testing.T) {
	for _, test := range readTemplateFileTestCases {
		actual, err := wrapper.ReadTemplateFile(test.templatedir, test.filename)

		if test.err != nil {
			if err == nil || test.err.Error() != err.Error() {
				t.Fatalf("\n>>> Expected error to be:\n%v\n<<< but got:\n%v", test.err, err)
			}
			continue
		}

		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestBuildWrapper(t *testing.T) {
	for _, test := range buildWrapperTestCases {
		actual := wrapper.BuildWrapper(test.template, test.libraryname, test.services)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}
