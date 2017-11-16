package awswrapper_test

import "testing"
import w "github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper/aws"

func TestWrapper(t *testing.T) {
	wrapper := &w.AwsWrapper{}

	for _, test := range wrapperTestCases {
		actual, err := wrapper.Wrap(test.template, test.options)

		if err != nil {
			t.Fatalf("\n>>> Expected: err to be nil\n<<< but got:\n%v", err)
		}

		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}
