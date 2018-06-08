package gitsourcer_test

import "errors"

var sourcerTestCases = []struct {
	libname  string
	expected string
	err      error
}{
	{libname: "test", expected: "repository not found", err: errors.New("exit status 1")},
	{libname: "https://github.com/cloudrecipes/lambda-wrapper", expected: "", err: nil},
}

var depsTestCases = []struct {
	isprod   bool
	expected string
	err      error
}{
	{isprod: false, expected: "", err: nil},
	{isprod: true, expected: "", err: nil},
}

var verifyCommandsTestCases = []struct {
	isprod   bool
	expected string
	err      error
}{
	{expected: "1.0.0", err: nil},
	{expected: "command not found", err: errors.New("exit status 1")},
}
