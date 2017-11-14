// Package cli contains methods to work with command line arguments.
package cli

import (
	"fmt"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
	clihandler "gopkg.in/urfave/cli.v1"
)

// Name is the applicaion name.
const Name string = "lambda-wrapper"

// Version is the current applicaion version.
const Version string = "0.1.0"

// Usage is the application usage text template.
const Usage string = `
Usage: lambda-wrapper [options]


  Options:

		{{range .VisibleFlags}}{{.}}
		{{end}}
`

// App is a structure for the CLI applicaion.
type App struct {
	App     *clihandler.App
	Options *options.Options
}

// NewCliApp creates an instance of CLI applicaion.
func NewCliApp(action func(o *options.Options) error) *App {
	opts := &options.Options{}

	app := clihandler.NewApp()
	app.Name = Name
	app.Version = Version
	app.CustomAppHelpTemplate = Usage
	app.Flags = flags(opts)

	app.Action = func(c *clihandler.Context) error {
		if err := opts.Validate(); err != nil {
			fmt.Println(err)
			clihandler.ShowAppHelp(c)
			return nil
		}
		return action(opts)
	}

	return &App{
		App:     app,
		Options: opts,
	}
}

// Run CLI applicaion.
func (cli *App) Run(args []string) error {
	return cli.App.Run(args)
}

// flags configures list of applicaion flags.
func flags(opts *options.Options) []clihandler.Flag {
	return []clihandler.Flag{
		clihandler.StringFlag{
			Name:        "cloud, c",
			Value:       "AWS",
			Usage:       "cloud provider name",
			Destination: &opts.Cloud,
		},
		clihandler.StringFlag{
			Name:        "engine, e",
			Value:       "node",
			Usage:       "lambda function engine",
			Destination: &opts.Engine,
		},
		clihandler.StringSliceFlag{
			Name: "service, s",
			Usage: `a list of cloud services, the wrapper will automatically
			initiate handlers to these services and pass then to
			the library`,
		},
		clihandler.StringFlag{
			Name:        "libsource, S",
			Value:       "npm",
			Usage:       "the source where to find library's code",
			Destination: &opts.LibSource,
		},
		clihandler.StringFlag{
			Name:        "libname, N",
			Usage:       "the name of the library in the source",
			Destination: &opts.LibName,
		},
		clihandler.StringFlag{
			Name:        "output, o",
			Usage:       "path to save deployable lambda archive",
			Destination: &opts.Output,
		},
		clihandler.BoolFlag{
			Name: "test, t",
			Usage: `flag to run library's unit tests before wrapping
			into lambda package`,
			Destination: &opts.TestRequired,
		},
	}
}
