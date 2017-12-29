// Package npmsourcer implements library download/test operations from npm based sources.
package npmsourcer

import "os/exec"

// NpmSourcer structure.
type NpmSourcer struct{}

// LibGet gets library source from Npm.
func (s *NpmSourcer) LibGet(libname, destination string) error {
	cmd := exec.Command("npm", "install", libname)
	cmd.Dir = destination
	err := cmd.Run()
	return err
}

// LibTest runs tests defined in library.
func (s *NpmSourcer) LibTest(location string) error {
	cmd := exec.Command("npm", "test")
	cmd.Dir = location
	err := cmd.Run()
	return err
}

// LibDeps installs library's dependencies via npm.
func (s *NpmSourcer) LibDeps(location string, isprod bool) error {
	var cmd *exec.Cmd

	if isprod {
		cmd = exec.Command("npm", "install", "--prod")
	} else {
		cmd = exec.Command("npm", "install")
	}

	cmd.Dir = location
	err := cmd.Run()
	return err
}