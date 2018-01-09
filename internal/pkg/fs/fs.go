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

var libdir = path.Join(workingdir, "lib")
var builddir = path.Join(workingdir, "build")

// I interface provides methods to work with file system
type I interface {
	ReadFile(filename string) (string, error)
	ReadFileToBytes(filename string) ([]byte, error)
	WriteFile(filename string, payload []byte, perm os.FileMode) error
	RmDir(basedir string) error
	ZipDir(source, target string) error
}

// Fs is a generic file system operations structure
type Fs struct{}

type zipWriter interface {
	CreateHeader(*zip.FileHeader) (io.Writer, error)
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

// MakeDirs creates necessary working directories structure (if directories exist
// they will overwritten). This method is specific to the wrapper.
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

// ReadFile reterns file content or error.
func (fs *Fs) ReadFile(filename string) (string, error) {
	payload, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(payload), err
}

// ReadFileToBytes returns file content as bytes.
func (fs *Fs) ReadFileToBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// WriteFile writes file.
func (fs *Fs) WriteFile(filename string, payload []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, payload, perm)
}

// RmDir removes working directories.
func (fs *Fs) RmDir(basedir string) error {
	return os.RemoveAll(path.Join(basedir, workingdir))
}

// ZipDir archives directory.
func (fs *Fs) ZipDir(source, target string) error {
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
