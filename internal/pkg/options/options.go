// Package options contains structures and methods to work with application options.
package options

import (
	"errors"
	"fmt"
	s "strings"
)

// Options is a structure to store application options
type Options struct {
	Cloud        string
	Engine       string
	Services     []string
	LibSource    string
	LibName      string
	Output       string
	TestRequired bool
}

// Validate method checks if all the required options are set.
func (o *Options) Validate() error {
	errs := []string{}

	if o.Cloud == "" {
		errs = append(errs, "cloud")
	}

	if o.Engine == "" {
		errs = append(errs, "engine")
	}

	if o.LibSource == "" {
		errs = append(errs, "libsource")
	}

	if o.LibName == "" {
		errs = append(errs, "libname")
	}

	if o.Output == "" {
		errs = append(errs, "output")
	}

	if len(errs) > 0 {
		errmessage := fmt.Sprintf("Missing some of the required options: %s", s.Join(errs, ", "))
		return errors.New(errmessage)
	}

	return nil
}
