// Package gitsourcer implements library download operations from git based sources.
package gitsourcer

import (
	"path"

	cmd "github.com/cloudrecipes/lambda-wrapper/pkg/commander"
)

// GitSourcer structure.
type GitSourcer struct{}

// GitSourceDir is a subdirectory in working directory where to clone library.
// It's requered to avoid overwriting of the librarie's entry point 'index.js'
// with lambda handler 'index.js'.
const GitSourceDir = "_git"

// LibGet gets library source from Git.
func (s *GitSourcer) LibGet(c cmd.Commander, libname, workingdir string) ([]byte, error) {
	var command = "git"
	var args = []string{workingdir, "clone", libname, GitSourceDir}
	if _, err := c.CombinedOutput(command, args...); err != nil {
		return nil, err
	}

	// Remove .git directory
	command = "rm"
	args = []string{path.Join(workingdir, GitSourceDir), "-rf", ".git"}
	return c.CombinedOutput(command, args...)
}

// LibTest runs tests defined in library.
func (s *GitSourcer) LibTest(c cmd.Commander, workingdir string) ([]byte, error) {
	// TODO: update command in relation to the language/engine
	// npm is applicable to NodeJS only
	// TODO: add check if 'test' is available in package.json
	command := "npm"
	args := []string{path.Join(workingdir, GitSourceDir), "test"}
	return c.CombinedOutput(command, args...)
}

// LibDeps installs library's dependencies.
func (s *GitSourcer) LibDeps(c cmd.Commander, workingdir string, isprod bool) ([]byte, error) {
	// TODO: update command in relation to the language/engine
	// npm is applicable to NodeJS only
	command := "npm"
	args := []string{path.Join(workingdir, GitSourceDir), "install"}
	if isprod {
		args = append(args, "--prod")
	}

	return c.CombinedOutput(command, args...)
}

// VerifySourcerCommands checks if git, npm commands are available on the host OS.
func (s *GitSourcer) VerifySourcerCommands(c cmd.Commander) error {
	args := []string{"", "--version"}
	commands := []string{"git", "npm"}

	for _, command := range commands {
		if _, err := c.CombinedOutput(command, args...); err != nil {
			return err
		}
	}

	return nil
}

// INFO: Only for NodeJS based lambdas
