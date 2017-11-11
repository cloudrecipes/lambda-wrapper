// Package options contains structures and methods to work with application options.
package options

import "errors"
import s "strings"

// Options is a structure to store application options
type Options struct {
	Cloud        string
	Engine       string
	Services     []string
	LibSource    string
	LibName      string
	TestRequired bool
}

// Validate method checks if all the required options are set.
func (o *Options) Validate() error {
	errs := []string{}

	if o.Cloud == "" {
		errs = append(errs, "Cloud provider required.")
	}

	if o.Engine == "" {
		errs = append(errs, "Engine required.")
	}

	if o.LibSource == "" {
		errs = append(errs, "Library source required.")
	}

	if o.LibName == "" {
		errs = append(errs, "Library name required.")
	}

	if len(errs) > 0 {
		return errors.New(s.Join(errs, "\n"))
	}

	return nil
}
