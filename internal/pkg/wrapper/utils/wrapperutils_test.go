package wrapperutils_test

import "testing"
import "github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper/utils"
import f "github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"

func TestBuildTemplateFileName(t *testing.T) {
	for _, test := range templateFileNameTestCases {
		actual := wrapperutils.TemplateFileName(test.cloud, test.engine)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestReadTemplateFile(t *testing.T) {
	fs := &f.Fs{}
	for _, test := range readTemplateFileTestCases {
		actual, err := wrapperutils.ReadTemplateFile(test.templatedir, test.filename, fs)

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
