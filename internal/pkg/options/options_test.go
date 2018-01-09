package options_test

import (
	"path"
	"reflect"
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
	tu "github.com/cloudrecipes/lambda-wrapper/internal/pkg/testutils"
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

func TestFromYamlFile(t *testing.T) {
	filename := path.Join(tu.Fixturesdir, ".lwrc.yaml")
	opts, err := options.FromYamlFile(filename)

	if err != nil {
		t.Fatal("Expected no errors")
	}

	if opts == nil {
		t.Fatal("Expected options not to be null")
	}

	if "AWS" != opts.Cloud {
		t.Fatalf("\n>>> Expected Cloud:\n%s\n<<< but got:\n%s", "AWS", opts.Cloud)
	}

	if "node" != opts.Engine {
		t.Fatalf("\n>>> Expected Engine:\n%s\n<<< but got:\n%s", "node", opts.Engine)
	}

	services := []string{"s3", "sqs"}
	if !reflect.DeepEqual(services, opts.Services) {
		t.Fatalf("\n>>> Expected Service:\n%s\n<<< but got:\n%s", services, opts.Services)
	}

	if "npm" != opts.LibSource {
		t.Fatalf("\n>>> Expected LibSource:\n%s\n<<< but got:\n%s", "npm", opts.LibSource)
	}

	if "@foo/bar" != opts.LibName {
		t.Fatalf("\n>>> Expected LibName:\n%s\n<<< but got:\n%s", "@foo/bar", opts.LibName)
	}

	if "lambda.zip" != opts.Output {
		t.Fatalf("\n>>> Expected Output:\n%s\n<<< but got:\n%s", "lambda.zip", opts.Output)
	}

	if true != opts.TestRequired {
		t.Fatal("\n>>> Expected TestRequired to be true")
	}
}
