// Package awswrapper implements wrapper operations specific to AWS lambda
// infrastructure.
package awswrapper

import (
	"fmt"
	"path"
	s "strings"

	f "github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
	gs "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/git"
	p "github.com/cloudrecipes/packagejson/pkg/packagejson"
)

// awsservices is the map of supported services and handler initiators.
var awsservices = map[string]string{
	"s3":  "new aws.S3({apiVersion: 'latest'})",
	"sns": "new aws.SNS()",
}

// AwsWrapper structure.
type AwsWrapper struct{}

// Wrap methods creates AWS lambda wrapper.
func (w *AwsWrapper) Wrap(template string, opts *options.Options) (string, error) {
	var err error
	resultstr := template

	if resultstr, err = injectLibraryIntoTemplate(resultstr, opts); err != nil {
		return "", err
	}

	resultstr = injectServicesIntoTemplate(resultstr, opts.Services)
	return resultstr, nil
}

// injectLibraryIntoTemplate injects library into template.
func injectLibraryIntoTemplate(template string, opts *options.Options) (string, error) {
	switch opts.LibSource {
	case "git":
		return injectGitLibraryIntoTemplate(template, opts, &f.Fs{})
	default:
		return s.Replace(template, "{{lib}}", opts.LibName, -1), nil
	}
}

// injectGitLibraryIntoTemplate injects git based library into template.
func injectGitLibraryIntoTemplate(template string, opts *options.Options, fs f.I) (string, error) {
	filepath := path.Join(opts.Output, f.LibDir(), gs.GitSourceDir, "package.json")
	payload, err := fs.ReadFileToBytes(filepath)
	if err != nil {
		return "", err
	}

	packagejson, err := p.Parse(payload)
	if err != nil {
		return "", err
	}

	if err := packagejson.Validate(); err != nil {
		return "", err
	}

	entrypoint := packagejson.Main

	return s.Replace(template, "{{lib}}", fmt.Sprintf("./_git/%s", entrypoint), -1), nil
}

// injectServicesIntoTemplate injects services into template.
func injectServicesIntoTemplate(template string, services []string) string {
	resultstr := template
	resultstr = s.Replace(resultstr, "{{aws}}", initiateAwsHandler(services), -1)
	resultstr = s.Replace(resultstr, "{{services}}", initiateServiceHandlers(services), -1)
	return resultstr
}

// initiateAwsHandler adds aws 'require' if necessary.
func initiateAwsHandler(services []string) string {
	if len(services) == 0 {
		return ""
	}

	return "const aws = require('aws-sdk')"
}

// initiateServiceHandlers creates required service handlers which will be passed
// to the library.
func initiateServiceHandlers(services []string) string {
	if len(services) == 0 {
		return ""
	}

	handlers := make([]string, len(services))
	for i, v := range services {
		handler, exists := awsservices[v]
		if exists == true {
			handlers[i] = fmt.Sprintf("services.%s = %s", v, handler)
		} else {
			handlers[i] = ""
		}
	}

	return s.Join(handlers, "\n")
}
