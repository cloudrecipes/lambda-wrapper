package options_test

import "errors"

var validateTestCases = []struct {
	cloud     string // cloud provider
	engine    string // lambda function engine
	libsource string // library source
	libname   string // library name
	expected  error  // expected validation error
}{
	{
		"", "", "", "",
		errors.New(`Cloud provider required.
Engine required.
Library source required.
Library name required.`),
	},
	{
		"AWS", "", "", "",
		errors.New(`Engine required.
Library source required.
Library name required.`),
	},
	{
		"AWS", "node", "", "",
		errors.New(`Library source required.
Library name required.`),
	},
	{
		"AWS", "node", "npm", "",
		errors.New("Library name required."),
	},
	{
		"AWS", "node", "npm", "foo/bar", nil,
	},
}
