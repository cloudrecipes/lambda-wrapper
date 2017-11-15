package main

import (
	"fmt"
	"os"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
)

func main() {
	opts, _ := options.FromYamlFile(options.DefaultOptionsFileName)

	action := func(opts *options.Options) error {
		fmt.Printf("%v\n", opts)
		return nil
	}

	cli.NewCliApp(opts, action).Run(os.Args)
}
