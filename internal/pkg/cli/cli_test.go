package cli_test

import (
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
)

func TestNewCliApp(t *testing.T) {
	testApp := cli.NewCliApp()
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
	testApp := cli.NewCliApp()
	args := []string{"lambda-wrapper", "--help"}
	err := testApp.Run(args)

	// TODO: Add check of app options based on arguments
	// TODO: separately check help message
	// TODO: separately check version message

	if err != nil {
		t.Fatal("Expected Application to be successfully run")
	}
}
