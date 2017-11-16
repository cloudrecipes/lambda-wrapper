package wrapper_test

import (
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper"
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
