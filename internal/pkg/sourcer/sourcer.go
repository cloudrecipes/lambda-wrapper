// Package sourcer implements operations (download, test, etc.) with different library sources.
package sourcer

import cmd "github.com/cloudrecipes/lambda-wrapper/internal/pkg/commander"

// Sourcer is a generic interface for all types of sources.
type Sourcer interface {
	LibGet(c cmd.Commander, libname, workingdir string) ([]byte, error)
	LibTest(c cmd.Commander, workingdir string) ([]byte, error)
	LibDeps(c cmd.Commander, workingdir string, isprod bool) ([]byte, error)
	VerifySourcerCommands(c cmd.Commander) error
}

// LibGet downloads/gets library using different source types.
func LibGet(s Sourcer, libname, workingdir string) ([]byte, error) {
	return s.LibGet(&cmd.RealCommander{}, libname, workingdir)
}

// LibTest runs tests defined at library.
func LibTest(s Sourcer, workingdir string) ([]byte, error) {
	return s.LibTest(&cmd.RealCommander{}, workingdir)
}

// LibDeps downloads librarie's dependencies.
func LibDeps(s Sourcer, workingdir string, isprod bool) ([]byte, error) {
	return s.LibDeps(&cmd.RealCommander{}, workingdir, isprod)
}

// VerifySourcerCommands checks if all the necessary tools are available on the
// host OS.
func VerifySourcerCommands(s Sourcer) error {
	return s.VerifySourcerCommands(&cmd.RealCommander{})
}

// TODO: Check if host OS has installed git and npm
