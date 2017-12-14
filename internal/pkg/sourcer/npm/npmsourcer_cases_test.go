package npmsourcer_test

import "errors"

var sourcerTestCases = []struct {
	libname string
	err     error
}{
	{libname: "package-not-found", err: errors.New("exit status 1")},
	// Please, do not commit it, use this test case only locally
	// {libname: "@antklim/api-to-cloud", err: nil},
}
