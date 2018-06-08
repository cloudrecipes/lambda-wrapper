// Package wrapperutils provides utilities shared across wrappers.
package wrapperutils

import (
	"fmt"
	"path"
	s "strings"

	f "github.com/cloudrecipes/lambda-wrapper/pkg/fs"
)

// TemplateFileName returns template file name in regards to cloud provider name
// and a lambda engine.
func TemplateFileName(cloud, engine string) string {
	return fmt.Sprintf("%s-%s", s.ToLower(cloud), s.ToLower(engine))
}

// ReadTemplateFile reads teamplate file and returns it's content or error.
func ReadTemplateFile(templatedir, filename string, fs *f.Fs) (string, error) {
	templatefile := path.Join(templatedir, filename)
	return fs.ReadFile(templatefile)
}
