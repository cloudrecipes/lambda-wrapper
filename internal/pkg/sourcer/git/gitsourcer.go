// Package gitsourcer implements library download operations from git based sources.
package gitsourcer

import "os/exec"

// GitSourcer structure.
type GitSourcer struct{}

// LibGet gets library source from Git.
func (s *GitSourcer) LibGet(libname, destination string) ([]byte, error) {
	cmd := exec.Command("git", "clone", libname, ".")
	cmd.Dir = destination
	return cmd.CombinedOutput()
}

// LibTest runs tests defined in library.
func (s *GitSourcer) LibTest(location string) ([]byte, error) {
	return nil, nil
}

// LibDeps installs library's dependencies.
func (s *GitSourcer) LibDeps(location string, isprod bool) ([]byte, error) {
	return nil, nil
}
