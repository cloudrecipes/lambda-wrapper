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
