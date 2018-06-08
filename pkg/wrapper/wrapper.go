// Package wrapper implements operations with lambda wrapper.
package wrapper

import (
	"os"
	"path"

	f "github.com/cloudrecipes/lambda-wrapper/pkg/fs"
	"github.com/cloudrecipes/lambda-wrapper/pkg/options"
	utils "github.com/cloudrecipes/lambda-wrapper/pkg/wrapper/utils"
)

// templatedir is the default template directory
var templatedir = path.Join(os.Getenv("GOPATH"), "src", "github.com",
	"cloudrecipes", "lambda-wrapper", "assets", "templates")

// Wrapper generic interface for all types of wrappers.
type Wrapper interface {
	Wrap(template string, opts *options.Options) (string, error)
}

// Wrap reads template and wraps library into it.
func Wrap(w Wrapper, opts *options.Options, templatedir string) (string, error) {
	templatefile := utils.TemplateFileName(opts.Cloud, opts.Engine)
	template, err := utils.ReadTemplateFile(templatedir, templatefile, &f.Fs{})

	if err != nil {
		return "", err
	}

	return w.Wrap(template, opts)
}

// Save writes lambda handler in a file 'filename'.
func Save(filename, payload string, fs f.I) error {
	return fs.WriteFile(filename, []byte(payload), 0644)
}

// DefaultTemplateDir returns default template directory.
func DefaultTemplateDir() string {
	return templatedir
}
