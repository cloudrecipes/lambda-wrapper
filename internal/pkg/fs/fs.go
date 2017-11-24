// Package fs implements operations with file system.
package fs

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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
	dirs := []string{
		path.Join(basedir, workingdir),
		path.Join(basedir, libdir),
		path.Join(basedir, builddir),
	}

	for _, dir := range dirs {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

// RmDirs removes working directories.
func RmDirs(basedir string) error {
	return os.RemoveAll(path.Join(basedir, workingdir))
}

// ZipDir archives directory.
func ZipDir(source, target string) error {
	var err error
	// zipfile, err := os.Create(target)
	// if err != nil {
	// 	return err
	// }
	// defer zipfile.Close()
	//
	// archive := zip.NewWriter(zipfile)
	// defer archive.Close()
	//
	if _, err = os.Stat(source); err != nil {
		return err
	}

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
