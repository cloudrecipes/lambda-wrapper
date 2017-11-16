package wrapper_test

import (
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper"
)

func TestBuildWrapper(t *testing.T) {
	for _, test := range buildWrapperTestCases {
		actual := wrapper.BuildWrapper(test.template, test.libraryname, test.services)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}
