// Package wrapperutils provides utilities shared across wrappers.
package wrapperutils

import (
	"fmt"
	"path"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
)

// TemplateFileName returns template file name in regards to cloud provider name
// and a lambda engine.
func TemplateFileName(cloud, engine string) string {
	return fmt.Sprintf("%s-%s", cloud, engine)
}

// ReadTemplateFile reads teamplate file and returns it's content or error.
func ReadTemplateFile(templatedir, filename string) (string, error) {
	templatefile := path.Join(templatedir, filename)
	return fs.ReadFile(templatefile)
}
