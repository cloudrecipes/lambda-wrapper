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
		cloud: "", engine: "", libsource: "", libname: "", output: "",
		expected: errors.New("Missing some of the required options: cloud, engine, libsource, libname, output"),
	},
	{
		cloud: "AWS", engine: "", libsource: "", libname: "", output: "",
		expected: errors.New("Missing some of the required options: engine, libsource, libname, output"),
	},
	{
		cloud: "AWS", engine: "node", libsource: "", libname: "", output: "",
		expected: errors.New("Missing some of the required options: libsource, libname, output"),
	},
	{
		cloud: "AWS", engine: "node", libsource: "npm", libname: "", output: "",
		expected: errors.New("Missing some of the required options: libname, output"),
	},
	{
		cloud: "AWS", engine: "node", libsource: "npm", libname: "foo/bar", output: "",
		expected: errors.New("Missing some of the required options: output"),
	},
	{
		cloud: "AWS", engine: "node", libsource: "npm", libname: "foo/bar", output: "lambda.zip", expected: nil,
	},
}
