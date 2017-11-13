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
		errors.New("Missing some of the required options: cloud, engine, libsource, libname, output"),
	},
	{
		"AWS", "", "", "", "",
		errors.New("Missing some of the required options: engine, libsource, libname, output"),
	},
	{
		"AWS", "node", "", "", "",
		errors.New("Missing some of the required options: libsource, libname, output"),
	},
	{
		"AWS", "node", "npm", "", "",
		errors.New("Missing some of the required options: libname, output"),
	},
	{
		"AWS", "node", "npm", "foo/bar", "",
		errors.New("Missing some of the required options: output"),
	},
	{
		"AWS", "node", "npm", "foo/bar", "lambda.zip", nil,
	},
}
