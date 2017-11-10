package fs_test

import (
	"testing"

	fs "github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
)

func TestReadFile(t *testing.T) {
	for _, test := range readFileTestCases {
		actual, err := fs.ReadFile(test.filename)

		if test.err != nil {
			if err == nil || test.err.Error() != err.Error() {
				t.Fatalf("Expected error to be %v but got %v", test.err, err)
			}
			continue
		}

		if test.expected != actual {
			t.Fatalf("Expected %s but got %s", test.expected, actual)
		}
	}
}
