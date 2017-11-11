package options_test

import "errors"

var validateTestCases = []struct {
	cloud     string // cloud provider
	engine    string // lambda function engine
	libsource string // library source
	libname   string // library name
	output    string // output file name
	expected  error  // expected validation error
}{
	{
		"", "", "", "", "",
		errors.New(`Cloud provider required.
Engine required.
Library source required.
Library name required.
Output file name required.`),
	},
	{
		"AWS", "", "", "", "",
		errors.New(`Engine required.
Library source required.
Library name required.
Output file name required.`),
	},
	{
		"AWS", "node", "", "", "",
		errors.New(`Library source required.
Library name required.
Output file name required.`),
	},
	{
		"AWS", "node", "npm", "", "",
		errors.New(`Library name required.
Output file name required.`),
	},
	{
		"AWS", "node", "npm", "foo/bar", "",
		errors.New("Output file name required."),
	},
	{
		"AWS", "node", "npm", "foo/bar", "lambda.zip", nil,
	},
}
