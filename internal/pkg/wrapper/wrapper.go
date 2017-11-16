// Package wrapper implements operations with lambda wrapper.
package wrapper

import (
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
	utils "github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper/utils"
)

// Wrapper generic interface for all types of wrappers.
type Wrapper interface {
	Wrap(template string, opts *options.Options) (string, error)
}

// Wrap reads template and wraps library into it.
func Wrap(w Wrapper, opts *options.Options, templatedir string) (string, error) {
	templatefile := utils.TemplateFileName(opts.Cloud, opts.Engine)
	template, err := utils.ReadTemplateFile(templatedir, templatefile)

	if err != nil {
		return "", err
	}

	return w.Wrap(template, opts)
}
