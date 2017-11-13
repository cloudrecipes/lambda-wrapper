package cli_test

import (
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
)

func TestNewCliApp(t *testing.T) {
	action := func(opts *options.Options) error {
		return nil
	}

	testApp := cli.NewCliApp(action)
	if testApp == nil {
		t.Fatal("Expected Application to be not nil")
	}

	if testApp.App == nil {
		t.Fatal("Expected Cli Application handler to be not nil")
	}

	if testApp.Options == nil {
		t.Fatal("Expected Application Options to be not nil")
	}
}

func TestRun(t *testing.T) {
	action := func(opts *options.Options) error {
		return nil
	}
	testApp := cli.NewCliApp(action)
	args := []string{"lambda-wrapper", "--help"}
	err := testApp.Run(args)

	// TODO: Add check of app options based on arguments
	// TODO: separately check help message
	// TODO: separately check version message

	if err != nil {
		t.Fatal("Expected Application to be successfully run")
	}
}
