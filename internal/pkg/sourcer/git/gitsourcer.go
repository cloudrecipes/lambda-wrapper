// Package gitsourcer implements library download operations from git based sources.
package gitsourcer

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

// GitSourcer structure.
type GitSourcer struct{}

// LibGet gets library source from Git.
func (s *GitSourcer) LibGet(libname, destination string) error {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      libname,
		Progress: os.Stdout,
	})

	return err
}
