package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
	c "github.com/cloudrecipes/lambda-wrapper/internal/pkg/commander"
	fs "github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
	s "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer"
	git "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/git"
	npm "github.com/cloudrecipes/lambda-wrapper/internal/pkg/sourcer/npm"
	w "github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper"
	aws "github.com/cloudrecipes/lambda-wrapper/internal/pkg/wrapper/aws"
)

func getSourcer(source string) (s.Sourcer, error) {
	switch source {
	case "npm":
		return &npm.NpmSourcer{}, nil
	case "git":
		return &git.GitSourcer{}, nil
	default:
		return nil, fmt.Errorf("Unsupported libsource: %s\nCurrently supported: [npm, git]", source)
	}
}

func getWrapper(cloud string) (w.Wrapper, error) {
	switch cloud {
	case "AWS":
		return &aws.AwsWrapper{}, nil
	default:
		return nil, fmt.Errorf("Unsupported cloud provider: %s\nCurrently supported: [AWS]", cloud)
	}
}

func main() {
	opts, _ := options.FromYamlFile(options.DefaultOptionsFileName)

	action := func(opts *options.Options) error {
		var err error
		var lambda string
		var sourcer s.Sourcer
		var wrapper w.Wrapper
		var workingdir string
		commander := &c.RealCommander{}

		if os.Getenv("DEBUG") != "" {
			fmt.Printf("%v\n", opts)
		}

		// Prepate library sourcer
		if sourcer, err = getSourcer(opts.LibSource); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// Create `.lwtmp`, `.lwtmp/lib`, `.lwtmp/build` directories in Output directory
		if err = fs.MakeDirs(opts.Output); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// Install library into `.lwtmp/lib`
		workingdir = path.Join(opts.Output, fs.LibDir())
		if _, err = sourcer.LibGet(commander, opts.LibName, workingdir); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// Install dependencies to `.lwtmp/lib`
		if _, err = sourcer.LibDeps(commander, workingdir, false); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// Run unit tests at `.lwtmp/lib`
		if opts.TestRequired {
			if _, err = sourcer.LibTest(commander, workingdir); err != nil {
				fmt.Printf("%v\n", err)
				return err
			}
		}

		// Install library into `.lwtmp/build`
		workingdir = path.Join(opts.Output, fs.BuildDir())
		if _, err = sourcer.LibGet(commander, opts.LibName, workingdir); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// Install production dependencies to `.lwtmp/build`
		if _, err = sourcer.LibDeps(commander, workingdir, true); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// Prepare lambda wrapper
		if wrapper, err = getWrapper(opts.Cloud); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// Wrap code into lambda
		if lambda, err = w.Wrap(wrapper, opts, w.DefaultTemplateDir()); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// TODO: this code does not belong to main.go.
		// Move it to either wrapper or fs package.
		filename := path.Join(workingdir, "index.js")
		if err = ioutil.WriteFile(filename, []byte(lambda), 0644); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		// Zip package for deploy
		if err := fs.ZipDir(workingdir, "tmp.zip"); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		return nil
	}

	cli.NewCliApp(opts, action).Run(os.Args)
}
