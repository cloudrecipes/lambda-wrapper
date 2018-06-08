package wrapper_test

import (
	"errors"

	o "github.com/cloudrecipes/lambda-wrapper/pkg/options"
	tu "github.com/cloudrecipes/lambda-wrapper/pkg/testutils"
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
		options:     &o.Options{Cloud: "AWS", Engine: "Node"},
		templatedir: tu.Fixturesdir,
		err:         nil,
		expected:    "template wrapped",
	},
}
