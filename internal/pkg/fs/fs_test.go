package fs_test

import (
	"os"
	"path"
	"testing"

	fs "github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
)

func TestReadFile(t *testing.T) {
	expected := "Hello Test!"
	filename := path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
		"lambda-wrapper", "test", "fixtures", "fs_readfile.txt")
	actual, err := fs.ReadFile(filename)

	if err != nil {
		t.Fatalf("Error %v", err)
	}

	if expected != string(actual) {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
