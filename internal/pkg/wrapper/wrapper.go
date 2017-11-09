// Package wrapper implements operations with lambda wrapper.
package wrapper

import s "strings"
import "fmt"

// TODO: implement ReadTemplateFile method
// TODO: implement BuildWrapper method
// TODO: implement InjectServicesIntoTemplate method

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
func BuildWrapper(template, engine, libraryName string, services []string) string {
	return ""
}

// InjectLibraryIntoTemplate injects libraryName into template.
func injectLibraryIntoTemplate(template, libraryName string) string {
	return s.Replace(template, "{{lib}}", libraryName, -1)
}

// InjectServicesIntoTemplate injects services into template.
func InjectServicesIntoTemplate(services []string) string {
	return ""
}
