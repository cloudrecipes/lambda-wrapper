// Package awswrapper implements wrapper operations specific to AWS lambda
// infrastructure.
package awswrapper

import (
	"fmt"
	s "strings"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
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
	resultstr := template
	resultstr = injectLibraryIntoTemplate(resultstr, opts.LibName, opts.LibSource)
	resultstr = injectServicesIntoTemplate(resultstr, opts.Services)
	return resultstr, nil
}

// injectLibraryIntoTemplate injects library into template.
func injectLibraryIntoTemplate(template, libname, libsource string) string {
	switch libsource {
	case "git":
		return s.Replace(template, "{{lib}}", "./_git", -1)
	default:
		return s.Replace(template, "{{lib}}", libname, -1)
	}
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
