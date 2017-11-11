package options_test

import (
	"reflect"
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
)

func TestValidate(t *testing.T) {
	for _, test := range validateTestCases {
		testOptions := &options.Options{
			Cloud:     test.cloud,
			Engine:    test.engine,
			LibSource: test.libsource,
			LibName:   test.libname,
			Output:    test.output,
		}

		actual := testOptions.Validate()

		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("\n>>> Expected:\n%v\n<<< but got:\n%v", test.expected, actual)
		}
	}
}
