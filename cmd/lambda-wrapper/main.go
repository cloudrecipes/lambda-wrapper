package main

import (
	"fmt"
	"os"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
	c "github.com/cloudrecipes/lambda-wrapper/internal/pkg/commander"
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
	s "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer"
	git "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/git"
	npm "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/npm"
)

func main() {
	opts, _ := options.FromYamlFile(options.DefaultOptionsFileName)

	action := func(opts *options.Options) error {
		fmt.Printf("%v\n", opts)

		var err error
		var sourcer s.Sourcer
		commander := &c.RealCommander{}

		switch opts.LibSource {
		case "npm":
			sourcer = &npm.NpmSourcer{}
		case "git":
			sourcer = &git.GitSourcer{}
		default:
			err = fmt.Errorf("Unsupported libsource: %s\nCurrently supported: [npm, git]", opts.LibSource)
			fmt.Printf("%v\n", err)
			return err
		}

		if _, err = sourcer.LibGet(commander, opts.LibName, opts.Output); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		if _, err = sourcer.LibDeps(commander, opts.Output, true); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		return nil
	}

	cli.NewCliApp(opts, action).Run(os.Args)
}
