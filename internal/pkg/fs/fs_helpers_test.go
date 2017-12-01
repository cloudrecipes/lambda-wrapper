package fs_test

import (
	"fmt"
	"os"
	"path"
)

func createDummyDirStructure(basedir, headdir string) error {
	if err := os.Mkdir(path.Join(basedir, headdir), os.ModePerm); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(basedir, headdir, "blah"), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func createDummyFiles() error {
	for _, f := range filesToZip {
		if err := createFile(f.filename, f.payload); err != nil {
			return err
		}
	}

	return nil
}

func createFile(filename, payload string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprint(f, payload)

	return nil
}
