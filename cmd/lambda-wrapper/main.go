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
	"github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
)

// AppName is an application name to print to CLI
var AppName = "labda-wrapper"

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

func initBar(steps []string, app string) *uiprogress.Bar {
	bar := uiprogress.AddBar(len(steps)).AppendCompleted().PrependElapsed()
	bar.Width = 50
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return strutil.Resize(app+": "+steps[b.Current()-1], 40)
	})
	return bar
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

		// TODO: refactor steps
		// IDEA: move steps to a separate package
		var steps = []string{
			"get sourcer",
			"verify sourcer",
			"make dirs",
			"get library",
			"get library deps",
			"test library",
			"get library for build",
			"get library deps for build",
			"get wrapper",
			"wrap library",
			"save wrapper handler",
			"create deployables",
		}

		bar := initBar(steps, AppName)
		uiprogress.Start()

		// Prepate library sourcer
		if sourcer, err = getSourcer(opts.LibSource); err != nil {
			fmt.Printf("[1] get sourcer: %v\n", err)
			return err
		}
		bar.Incr()

		if err = sourcer.VerifySourcerCommands(commander); err != nil {
			fmt.Printf("[2] sourcer commands verification failed: %v\n", err)
			return err
		}
		bar.Incr()

		// Create `.lwtmp`, `.lwtmp/lib`, `.lwtmp/build` directories in Output directory
		if err = f.MakeDirs(opts.Output); err != nil {
			fmt.Printf("[3] make dirs: %v\n", err)
			return err
		}
		bar.Incr()

		// Install library into `.lwtmp/lib`
		workingdir = path.Join(opts.Output, f.LibDir())
		if _, err = sourcer.LibGet(commander, opts.LibName, workingdir); err != nil {
			fmt.Printf("[4] library get: %v\n", err)
			return err
		}
		bar.Incr()

		// Install dependencies to `.lwtmp/lib`
		if _, err = sourcer.LibDeps(commander, workingdir, false); err != nil {
			fmt.Printf("[5] library dependencies: %v\n", err)
			return err
		}
		bar.Incr()

		// Run unit tests at `.lwtmp/lib`
		if opts.TestRequired {
			if _, err = sourcer.LibTest(commander, workingdir); err != nil {
				fmt.Printf("[6] library test: %v\n", err)
				return err
			}
		}
		bar.Incr()

		// Install library into `.lwtmp/build`
		workingdir = path.Join(opts.Output, f.BuildDir())
		if _, err = sourcer.LibGet(commander, opts.LibName, workingdir); err != nil {
			fmt.Printf("[7] build library get: %v\n", err)
			return err
		}
		bar.Incr()

		// Install production dependencies to `.lwtmp/build`
		if _, err = sourcer.LibDeps(commander, workingdir, true); err != nil {
			fmt.Printf("[8] build library dependencies: %v\n", err)
			return err
		}
		bar.Incr()

		// Prepare lambda wrapper
		if wrapper, err = getWrapper(opts.Cloud); err != nil {
			fmt.Printf("[9] get wrapper: %v\n", err)
			return err
		}
		bar.Incr()

		// Wrap code into lambda
		if lambda, err = w.Wrap(wrapper, opts, w.DefaultTemplateDir()); err != nil {
			fmt.Printf("[10] wrap: %v\n", err)
			return err
		}
		bar.Incr()

		filename := path.Join(workingdir, "index.js")
		if err = w.Save(filename, lambda, fs); err != nil {
			fmt.Printf("[11] write wrapper: %v\n", err)
			return err
		}
		bar.Incr()

		// Zip package for deploy
		if err := fs.ZipDir(workingdir, "tmp.zip"); err != nil {
			fmt.Printf("[12] save deployables: %v\n", err)
			return err
		}
		bar.Incr()
		uiprogress.Stop()

		fmt.Printf("%s: successfully finished\n", AppName)

		return nil
	}

	cli.NewCliApp(opts, action).Run(os.Args)
}
