// Package npmsourcer implements library download/test operations from npm based sources.
package npmsourcer

import (
	cmd "github.com/cloudrecipes/lambda-wrapper/internal/pkg/commander"
)

// NpmSourcer structure.
type NpmSourcer struct{}

var command = "npm"

// LibGet gets library source from Npm.
func (s *NpmSourcer) LibGet(c cmd.Commander, libname, workingdir string) ([]byte, error) {
	args := []string{workingdir, "install", libname}
	return c.CombinedOutput(command, args...)
}

// LibTest runs tests defined in library.
func (s *NpmSourcer) LibTest(c cmd.Commander, workingdir string) ([]byte, error) {
	args := []string{workingdir, "test"}
	return c.CombinedOutput(command, args...)
}

// LibDeps installs library's dependencies via npm.
func (s *NpmSourcer) LibDeps(c cmd.Commander, workingdir string, isprod bool) ([]byte, error) {
	args := []string{workingdir, "install"}
	if isprod {
		args = append(args, "--prod")
	}

	return c.CombinedOutput(command, args...)
}

// VerifySourcerCommands checks if npm command is available on the host OS.
func (s *NpmSourcer) VerifySourcerCommands(c cmd.Commander) error {
	args := []string{"", "--version"}
	_, err := c.CombinedOutput(command, args...)
	return err
}
