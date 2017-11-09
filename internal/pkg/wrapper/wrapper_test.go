package wrapper_test

import (
	"testing"

	wrapper "github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper"
)

func TestBuildTemplateFileName(t *testing.T) {
	for _, test := range buildTemplateFileNameTestCases {
		actual := wrapper.BuildTemplateFileName(test.cloud, test.engine)
		if test.expected != actual {
			t.Fatalf("Expected %s but got %s", test.expected, actual)
		}
	}
}
