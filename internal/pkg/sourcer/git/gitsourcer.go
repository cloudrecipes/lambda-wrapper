// Package gitsourcer implements library download operations from git based sources.
package gitsourcer

import cmd "github.com/cloudrecipes/lambda-wrapper/internal/pkg/commander"

// GitSourcer structure.
type GitSourcer struct{}

// LibGet gets library source from Git.
func (s *GitSourcer) LibGet(c cmd.Commander, libname, workingdir string) ([]byte, error) {
	var command = "git"
	var args = []string{workingdir, "clone", libname, "."}
	if _, err := c.CombinedOutput(command, args...); err != nil {
		return nil, err
	}

	// Remove .git directory
	command = "rm"
	args = []string{workingdir, "-rf", ".git"}
	return c.CombinedOutput(command, args...)
}

// LibTest runs tests defined in library.
func (s *GitSourcer) LibTest(c cmd.Commander, workingdir string) ([]byte, error) {
	// TODO: update command in relation to the language/engine
	// npm is applicable to NodeJS only
	command := "npm"
	args := []string{workingdir, "test"}
	return c.CombinedOutput(command, args...)
}

// LibDeps installs library's dependencies.
func (s *GitSourcer) LibDeps(c cmd.Commander, workingdir string, isprod bool) ([]byte, error) {
	// TODO: update command in relation to the language/engine
	// npm is applicable to NodeJS only
	command := "npm"
	args := []string{workingdir, "install"}
	if isprod {
		args = append(args, "--prod")
	}

	return c.CombinedOutput(command, args...)
}

// INFO: Only for NodeJS based lambdas
// TODO: when cloning code it can contain 'index.js' on top level. This is an
// entry point to library, not lambda. And MUST not be overwritten by lambda wrapper.
// TODO: wrapper should parse package.json to find an entry point and test
