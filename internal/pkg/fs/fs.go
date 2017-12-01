// Package fs implements operations with file system.
package fs

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const workingdir string = ".lwtmp"

var libdir string = path.Join(workingdir, "lib")
var builddir string = path.Join(workingdir, "build")

type zipWriter interface {
	CreateHeader(*zip.FileHeader) (io.Writer, error)
}

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
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	var basedir string
	if info.IsDir() {
		basedir = filepath.Base(source)
	}

	return filepath.Walk(source, filepathWalk(basedir, source, archive))
}

func filepathWalk(basedir, source string, archive zipWriter) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, _ := zip.FileInfoHeader(info)

		if basedir != "" {
			header.Name = filepath.Join(basedir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)

		return err
	}
}
