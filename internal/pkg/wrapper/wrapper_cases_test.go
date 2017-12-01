package wrapper_test

import (
	"errors"
	"os"
	"path"

	o "github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
)

type testWrapper struct{}

func (w *testWrapper) Wrap(template string, opts *o.Options) (string, error) {
	return "template wrapped", nil
}

var wrapTestCases = []struct {
	options     *o.Options
	templatedir string
	err         error
	expected    string
}{
	{
		options:     &o.Options{},
		templatedir: "",
		err:         errors.New("open -: no such file or directory"),
		expected:    "",
	},
	{
		options: &o.Options{Cloud: "AWS", Engine: "Node"},
		templatedir: path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
			"lambda-wrapper", "test", "fixtures"),
		err:      nil,
		expected: "template wrapped",
	},
}
