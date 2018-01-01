package gitsourcer_test

import "errors"

var sourcerTestCases = []struct {
	libname string
	err     error
}{
	{libname: "test", err: errors.New("exit status 128")},
	// Please, do not commit it, use this test case only locally
	// {libname: "https://github.com/cloudrecipes/lambda-wrapper", err: nil},
}
