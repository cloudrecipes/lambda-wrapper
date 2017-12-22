// Package sourcer implements operations (download, test, etc.) with different library sources.
package sourcer

// Sourcer is a generic interface for all types of sources.
type Sourcer interface {
	LibGet(libname string) error
	LibTest(location string) error
}

// LibGet downloads/gets library using different source types.
func LibGet(s Sourcer, libname, destination string) error {
	return s.LibGet(libname)
}

// LibTest runs tests defined at library.
func LibTest(s Sourcer, location string) error {
	return s.LibTest(location)
}
