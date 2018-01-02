package npmsourcer_test

import "errors"

var sourcerTestCases = []struct {
	libname  string
	expected string
	err      error
}{
	{libname: "package-not-found", expected: "package not found", err: errors.New("exit status 1")},
	{libname: "@antklim/api-to-cloud", expected: "", err: nil},
}

var depsTestCases = []struct {
	isprod   bool
	expected string
	err      error
}{
	{isprod: false, expected: "", err: nil},
	{isprod: true, expected: "", err: nil},
}
