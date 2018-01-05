// Package gitsourcer implements library download operations from git based sources.
package gitsourcer

import cmd "github.com/cloudrecipes/lambda-wrapper/internal/pkg/commander"

// GitSourcer structure.
type GitSourcer struct{}

var command = "git"

// LibGet gets library source from Git.
func (s *GitSourcer) LibGet(c cmd.Commander, libname, workingdir string) ([]byte, error) {
	args := []string{workingdir, "clone", libname, "."}
	return c.CombinedOutput(command, args...)
}

// LibTest runs tests defined in library.
func (s *GitSourcer) LibTest(c cmd.Commander, workingdir string) ([]byte, error) {
	return nil, nil
}

// LibDeps installs library's dependencies.
func (s *GitSourcer) LibDeps(c cmd.Commander, workingdir string, isprod bool) ([]byte, error) {
	return nil, nil
}

// INFO: Only for NodeJS based lambdas
// TODO: Run LibDeps via npm
// TODO: Run LibTest via npm
// TODO: when cloning code it can contain 'index.js' on top level. This is an
// entry point to library, not lambda. And MUST not be overwritten by lambda wrapper.
// TODO: wrapper should parse package.json to find an entry point and test
