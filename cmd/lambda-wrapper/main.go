package main

import (
	"os"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
)

func main() {
	action := func(opts *options.Options) error {
		return nil
	}

	cli.NewCliApp(action).Run(os.Args)
}
