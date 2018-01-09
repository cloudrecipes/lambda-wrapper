package awswrapper_test

import "testing"
import w "github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper/aws"

func TestWrapper(t *testing.T) {
	wrapper := &w.AwsWrapper{}

	for _, test := range wrapperTestCases {
		actual, err := wrapper.Wrap(test.template, test.opts)

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
