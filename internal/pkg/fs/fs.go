// Package fs implements operations with file system.
package fs

import (
	"io/ioutil"
	"path"
)

const workingdir string = ".lwtmp"

var libdir string = path.Join(workingdir, "lib")
var builddir string = path.Join(workingdir, "build")

// ReadFile reterns fila content or error.
func ReadFile(filename string) (string, error) {
	payload, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(payload), err
}

// WorkingDir returns working directory name.
func WorkingDir() string {
	return workingdir
}

// LibDir returns library directory name.
func LibDir() string {
	return libdir
}

// BuildDir returns build directory name.
func BuildDir() string {
	return builddir
}
