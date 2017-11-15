package cli_test

import (
	"bytes"
	"io"
	"os"
	"reflect"
	s "strings"
	"testing"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
)

func TestNewCliApp(t *testing.T) {
	action := func(opts *options.Options) error {
		return nil
	}

	testApp := cli.NewCliApp(&options.Options{}, action)
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

func TestRunHelp(t *testing.T) {
	action := func(opts *options.Options) error {
		t.Fail()
		return nil
	}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	testApp := cli.NewCliApp(&options.Options{}, action)
	args := []string{"lambda-wrapper", "--help"}
	err := testApp.Run(args)

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if err != nil {
		t.Fatal("Expected Application to be successfully run")
	}

	prefix := "\nUsage: lambda-wrapper [options]"

	if !s.HasPrefix(out, prefix) {
		t.Fatalf("\n>>> Expected:\n%s\n<<< to have prefix:\n%s", out, prefix)
	}
}

func TestRunVersion(t *testing.T) {
	action := func(opts *options.Options) error {
		t.Fail()
		return nil
	}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	testApp := cli.NewCliApp(&options.Options{}, action)
	args := []string{"lambda-wrapper", "--version"}
	err := testApp.Run(args)

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if err != nil {
		t.Fatal("Expected Application to be successfully run")
	}

	prefix := "lambda-wrapper version"

	if !s.HasPrefix(out, prefix) {
		t.Fatalf("\n>>> Expected:\n%s\n<<< to have prefix:\n%s", out, prefix)
	}
}

func TestRunDefault(t *testing.T) {
	action := func(opts *options.Options) error {
		return nil
	}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	testApp := cli.NewCliApp(&options.Options{}, action)
	args := []string{"lambda-wrapper"}
	err := testApp.Run(args)

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if err != nil {
		t.Fatal("Expected Application to be successfully run")
	}

	prefix := "Missing some of the required options:"

	if !s.HasPrefix(out, prefix) {
		t.Fatalf("\n>>> Expected:\n%s\n<<< to have prefix:\n%s", out, prefix)
	}
}

func TestRun(t *testing.T) {
	action := func(opts *options.Options) error {
		if "AWS" != opts.Cloud {
			t.Fatalf("\n>>> Expected Cloud:\n%s\n<<< but got:\n%s", "AWS", opts.Cloud)
		}

		if "node" != opts.Engine {
			t.Fatalf("\n>>> Expected Engine:\n%s\n<<< but got:\n%s", "node", opts.Engine)
		}

		expectedServices := []string{"s3"}
		if !reflect.DeepEqual(expectedServices, opts.Services) {
			t.Fatalf("\n>>> Expected Services:\n%s\n<<< but got:\n%s", expectedServices, opts.Services)
		}

		if "npm" != opts.LibSource {
			t.Fatalf("\n>>> Expected LibSource:\n%s\n<<< but got:\n%s", "npm", opts.LibSource)
		}

		if "@foo/bar" != opts.LibName {
			t.Fatalf("\n>>> Expected LibName:\n%s\n<<< but got:\n%s", "@foo/bar", opts.LibName)
		}

		if "lambda.zip" != opts.Output {
			t.Fatalf("\n>>> Expected Output:\n%s\n<<< but got:\n%s", "lambda.zip", opts.Output)
		}

		if true != opts.TestRequired {
			t.Fatal("\n>>> Expected TestRequired to be true")
		}

		return nil
	}

	opts := &options.Options{
		Cloud:     "AWS",
		Engine:    "node",
		Services:  []string{"s3"},
		LibSource: "npm",
	}

	testApp := cli.NewCliApp(opts, action)
	args := []string{"lambda-wrapper", "-N", "@foo/bar", "--output", "lambda.zip", "-t"}
	err := testApp.Run(args)

	if err != nil {
		t.Fatal("Expected Application to be successfully run")
	}
}
