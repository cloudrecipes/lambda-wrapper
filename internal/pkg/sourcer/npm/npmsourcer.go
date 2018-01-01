// Package npmsourcer implements library download/test operations from npm based sources.
package npmsourcer

import "os/exec"

// NpmSourcer structure.
type NpmSourcer struct{}

// LibGet gets library source from Npm.
func (s *NpmSourcer) LibGet(libname, destination string) ([]byte, error) {
	cmd := exec.Command("npm", "install", libname)
	cmd.Dir = destination
	return cmd.CombinedOutput()
}

// LibTest runs tests defined in library.
func (s *NpmSourcer) LibTest(location string) ([]byte, error) {
	cmd := exec.Command("npm", "test")
	cmd.Dir = location
	return cmd.CombinedOutput()
}

// LibDeps installs library's dependencies via npm.
func (s *NpmSourcer) LibDeps(location string, isprod bool) ([]byte, error) {
	if isprod {
		return libDepsProd(location)
	}

	return libDepsDefault(location)
}

func libDepsDefault(location string) ([]byte, error) {
	cmd := exec.Command("npm", "install")
	cmd.Dir = location
	return cmd.CombinedOutput()
}

func libDepsProd(location string) ([]byte, error) {
	cmd := exec.Command("npm", "install", "--prod")
	cmd.Dir = location
	return cmd.CombinedOutput()
}
