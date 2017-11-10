// Package wrapper implements operations with lambda wrapper.
package wrapper

import (
	"fmt"
	"path"
	s "strings"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
)

// TODO: currently this code explicitly works with AWS/Node lambdas
//       only. Restructure package in a way, where every package
//       works with it's own cloud and engine.
//       Refactoring: define common interface, and all other modules
//       should implement an interface

var AWS_SERVICES = map[string]string{
	"s3":  "new aws.S3({apiVersion: 'latest'})",
	"sns": "new aws.SNS()",
}

// BuildTemplateFileName by cloud provider name and engine.
func BuildTemplateFileName(cloud, engine string) string {
	return fmt.Sprintf("%s-%s", cloud, engine)
}

// ReadTemplateFile reads teamplate file and returns it's content or error.
func ReadTemplateFile(templatedir, filename string) (string, error) {
	templatefile := path.Join(templatedir, filename)
	return fs.ReadFile(templatefile)
}

// BuildWrapper takes teamplate payload and injects necessary dependencies into it
// to build wrapper code.
func BuildWrapper(template, libraryname string, services []string) string {
	resultstr := template
	resultstr = injectLibraryIntoTemplate(resultstr, libraryname)
	resultstr = injectServicesIntoTemplate(resultstr, services)
	return resultstr
}

// injectLibraryIntoTemplate injects libraryname into template.
func injectLibraryIntoTemplate(template, libraryname string) string {
	return s.Replace(template, "{{lib}}", libraryname, -1)
}

// injectServicesIntoTemplate injects services into template.
func injectServicesIntoTemplate(template string, services []string) string {
	resultstr := template
	resultstr = s.Replace(resultstr, "{{aws}}", initiateAwsHandler(services), -1)
	resultstr = s.Replace(resultstr, "{{services}}", initiateServiceHandlers(services), -1)
	return resultstr
}

// initiateAwsHandler adds
func initiateAwsHandler(services []string) string {
	if len(services) == 0 {
		return ""
	}

	return "const aws = require('aws-sdk')"
}

func initiateServiceHandlers(services []string) string {
	if len(services) == 0 {
		return ""
	}

	handlers := make([]string, len(services))
	for i, v := range services {
		handler, exists := AWS_SERVICES[v]
		if exists == true {
			handlers[i] = fmt.Sprintf("services.%s = %s", v, handler)
		} else {
			handlers[i] = ""
		}
	}

	return s.Join(handlers, "\n")
}
