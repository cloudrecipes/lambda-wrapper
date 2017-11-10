// Package wrapper implements operations with lambda wrapper.
package wrapper

import s "strings"
import "fmt"

// TODO: implement ReadTemplateFile method
// TODO: implement InjectServicesIntoTemplate method
// TODO: add check on supported services in initialeServiceHandlers
// TODO: currently this code explicitly works with AWS/Node lambdas
//       only. Restructure package in a way, where every package
//       works with it's own cloud and engine.
//       Refactoring: define common interface, and all other modules
//       should implement an interface

// BuildTemplateFileName by cloud provider name and engine.
func BuildTemplateFileName(cloud, engine string) string {
	return fmt.Sprintf("%s-%s", cloud, engine)
}

// ReadTemplateFile reads teamplate file and returns it's content or error.
func ReadTemplateFile(templateHome, fileName string) (string, error) {
	return "", nil
}

// BuildWrapper takes teamplate payload and injects necessary dependencies into it
// to build wrapper code.
func BuildWrapper(template, libraryName string, services []string) string {
	resultStr := template
	resultStr = injectLibraryIntoTemplate(resultStr, libraryName)
	resultStr = injectServicesIntoTemplate(resultStr, services)
	return resultStr
}

// injectLibraryIntoTemplate injects libraryName into template.
func injectLibraryIntoTemplate(template, libraryName string) string {
	return s.Replace(template, "{{lib}}", libraryName, -1)
}

// injectServicesIntoTemplate injects services into template.
func injectServicesIntoTemplate(template string, services []string) string {
	resultStr := template
	resultStr = s.Replace(resultStr, "{{aws}}", initiateAwsHandler(services), -1)
	resultStr = s.Replace(resultStr, "{{services}}", initialeServiceHandlers(services), -1)
	return resultStr
}

// initiateAwsHandler adds
func initiateAwsHandler(services []string) string {
	if len(services) == 0 {
		return ""
	}

	return "const aws = require('aws-sdk')"
}

func initialeServiceHandlers(services []string) string {
	if len(services) == 0 {
		return ""
	}

	return "services.s3 = new aws.S3({apiVersion: 'latest'})"
}
