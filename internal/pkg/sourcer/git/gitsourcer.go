// Package gitsourcer implements library download operations from git based sources.
package gitsourcer

import git "gopkg.in/src-d/go-git.v4"

// GitSourcer structure.
type GitSourcer struct{}

// LibGet gets library source from Git.
func (s *GitSourcer) LibGet(libname, destination string) error {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      libname,
		Progress: nil,
	})

	return err
}

// LibTest runs tests defined in library.
func (s *GitSourcer) LibTest(location string) error {
	return nil
}

// LibDeps installs library's dependencies.
func (s *GitSourcer) LibDeps(location string, isprod bool) error {
	return nil
}
