// Package options contains structures and methods to work with application options.
package options

import (
	"errors"
	"fmt"
	"io/ioutil"
	s "strings"

	yaml "gopkg.in/yaml.v2"
)

// DefaultOptionsFileName default file name with options to read.
const DefaultOptionsFileName string = "lwrc.yaml"

// Options is a structure to store application options
type Options struct {
	Cloud        string   `yaml:"cloud"`
	Engine       string   `yaml:"engine"`
	Services     []string `yaml:"service"`
	LibSource    string   `yaml:"libsource"`
	LibName      string   `yaml:"libname"`
	Output       string   `yaml:"output"`
	TestRequired bool     `yaml:"test"`
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

// FromYamlFile reads options from the YAML file
func FromYamlFile(filename string) (*Options, error) {
	opts := &Options{}
	file, err := ioutil.ReadFile(filename)

	if err == nil {
		err = yaml.Unmarshal(file, opts)
	}

	return opts, err
}
