// Package fs implements operations with file system.
package fs

import (
	"io/ioutil"
	"os"
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

// MakeDirs creates necessary working directories (if directories exist they will overwritten).
func MakeDirs(basedir string) error {
	var err error

	if err = os.Mkdir(path.Join(basedir, workingdir), os.ModePerm); err != nil {
		return err
	}

	if err = os.Mkdir(path.Join(basedir, libdir), os.ModePerm); err != nil {
		return err
	}

	if err = os.Mkdir(path.Join(basedir, builddir), os.ModePerm); err != nil {
		return err
	}

	return nil
}

// RmDirs removes working directories.
func RmDirs(basedir string) error {
	return os.RemoveAll(path.Join(basedir, workingdir))
}
