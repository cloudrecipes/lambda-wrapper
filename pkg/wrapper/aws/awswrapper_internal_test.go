package awswrapper

import (
	"os"
	"testing"

	tu "github.com/cloudrecipes/lambda-wrapper/pkg/testutils"
)

func TestInjectLibraryIntoTemplate(t *testing.T) {
	for _, test := range injectLibraryIntoTemplateTestCases {
		actual, err := injectLibraryIntoTemplate(test.template, test.opts)

		if err != nil {
			t.Fatalf("\n>>> Expected: err to be nil\n<<< but got:\n%v", err)
		}

		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestInjectGitLibraryIntoTemplate(t *testing.T) {
	var fs = &tu.TestFs{}
	for _, test := range injectGitLibraryIntoTemplateTestCases {
		os.Setenv(tu.GoTestFsReadFileToBytesExpected, test.payload)
		os.Setenv(tu.GoTestFsReadFileToBytesError, test.readFileErr)

		actual, err := injectGitLibraryIntoTemplate(test.template, test.opts, fs)

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

func TestInjectServicesIntoTemplate(t *testing.T) {
	for _, test := range injectServicesIntoTemplateTestCases {
		actual := injectServicesIntoTemplate(test.template, test.services)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestInitiateAwsHandler(t *testing.T) {
	for _, test := range initiateAwsHandlerTestCases {
		actual := initiateAwsHandler(test.services)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestInitiateServiceHandlers(t *testing.T) {
	for _, test := range initiateServiceHandlersTestCases {
		actual := initiateServiceHandlers(test.services)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}
