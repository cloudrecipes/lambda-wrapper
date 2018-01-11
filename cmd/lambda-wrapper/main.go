package main

import (
	"fmt"
	"os"
	"path"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
	c "github.com/cloudrecipes/lambda-wrapper/internal/pkg/commander"
	f "github.com/cloudrecipes/lambda-wrapper/internal/pkg/fs"
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
		var fs = &f.Fs{}
		var lambda string
		var sourcer s.Sourcer
		var wrapper w.Wrapper
		var workingdir string
		commander := &c.RealCommander{}

		if os.Getenv("DEBUG") != "" {
			fmt.Printf("%v\n", opts)
		}

		// Prepate library sourcer
		fmt.Println("[1] getting sourcer...")
		if sourcer, err = getSourcer(opts.LibSource); err != nil {
			fmt.Printf("[1] get sourcer: %v\n", err)
			return err
		}

		fmt.Println("[2] verifying all required tools are installed on a host OS...")
		if err = sourcer.VerifySourcerCommands(commander); err != nil {
			fmt.Printf("[2] sourcer commands verification failed: %v\n", err)
			return err
		}

		// Create `.lwtmp`, `.lwtmp/lib`, `.lwtmp/build` directories in Output directory
		fmt.Println("[3] making dirs...")
		if err = f.MakeDirs(opts.Output); err != nil {
			fmt.Printf("[3] make dirs: %v\n", err)
			return err
		}

		// Install library into `.lwtmp/lib`
		fmt.Println("[4] getting library...")
		workingdir = path.Join(opts.Output, f.LibDir())
		if _, err = sourcer.LibGet(commander, opts.LibName, workingdir); err != nil {
			fmt.Printf("[4] library get: %v\n", err)
			return err
		}

		// Install dependencies to `.lwtmp/lib`
		fmt.Println("[5] installing library dependencies...")
		if _, err = sourcer.LibDeps(commander, workingdir, false); err != nil {
			fmt.Printf("[5] library dependencies: %v\n", err)
			return err
		}

		// Run unit tests at `.lwtmp/lib`
		if opts.TestRequired {
			fmt.Println("[6] running library tests...")
			if _, err = sourcer.LibTest(commander, workingdir); err != nil {
				fmt.Printf("[6] library test: %v\n", err)
				return err
			}
		}

		// Install library into `.lwtmp/build`
		fmt.Println("[7] getting library for build...")
		workingdir = path.Join(opts.Output, f.BuildDir())
		if _, err = sourcer.LibGet(commander, opts.LibName, workingdir); err != nil {
			fmt.Printf("[7] build library get: %v\n", err)
			return err
		}

		// Install production dependencies to `.lwtmp/build`
		fmt.Println("[8] installing library dependencies...")
		if _, err = sourcer.LibDeps(commander, workingdir, true); err != nil {
			fmt.Printf("[8] build library dependencies: %v\n", err)
			return err
		}

		// Prepare lambda wrapper
		fmt.Println("[9] geting wrapper...")
		if wrapper, err = getWrapper(opts.Cloud); err != nil {
			fmt.Printf("[9] get wrapper: %v\n", err)
			return err
		}

		// Wrap code into lambda
		fmt.Println("[10] wrapping...")
		if lambda, err = w.Wrap(wrapper, opts, w.DefaultTemplateDir()); err != nil {
			fmt.Printf("[10] wrap: %v\n", err)
			return err
		}

		fmt.Println("[11] writing lambda handler...")
		filename := path.Join(workingdir, "index.js")
		if err = w.Save(filename, lambda, fs); err != nil {
			fmt.Printf("[11] write wrapper: %v\n", err)
			return err
		}

		// Zip package for deploy
		fmt.Println("[12] save deployables...")
		if err := fs.ZipDir(workingdir, "tmp.zip"); err != nil {
			fmt.Printf("[12] save deployables: %v\n", err)
			return err
		}

		fmt.Println("[*] done!")

		return nil
	}

	cli.NewCliApp(opts, action).Run(os.Args)
}
