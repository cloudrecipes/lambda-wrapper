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
