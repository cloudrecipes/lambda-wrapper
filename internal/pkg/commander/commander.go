package commander

import "os/exec"

// Commander is an interface to an OS level command executor.
type Commander interface {
	CombinedOutput(string, ...string) ([]byte, error)
}

// RealCommander is a concrete implementation of the Commander interface.
type RealCommander struct{}

// CombinedOutput returns a result of the execution of a command on OS.
func (c RealCommander) CombinedOutput(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args[1:]...)
	workingdir := args[0]
	if len(workingdir) > 0 {
		cmd.Dir = workingdir
	}
	return cmd.CombinedOutput()
}
