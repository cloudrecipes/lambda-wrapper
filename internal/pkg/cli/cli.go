// Package cli contains methods to work with command line arguments.
package cli

import (
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
	clihandler "gopkg.in/urfave/cli.v1"
)

// Name is the applicaion name.
const Name string = "lambda-wrapper"

// Version is the current applicaion version.
const Version string = "0.1.0"

// Description is the applicaion description.
const Description = "Wrap the library into lambda function archive."

// ShortDescription is the applicaion description which comes in the NAME section.
const ShortDescription string = ""

// Usage is the applicaion usage text.
const Usage string = "lambda-wrapper [options]"

// App is a structure for the CLI applicaion.
type App struct {
	App     *clihandler.App
	Options *options.Options
}

// NewCliApp creates an instance of CLI applicaion.
func NewCliApp() *App {
	opts := &options.Options{}

	app := clihandler.NewApp()
	app.Name = Name
	app.Version = Version
	app.Description = Description
	app.Usage = ShortDescription
	app.UsageText = Usage

	app.Flags = []clihandler.Flag{
		clihandler.StringFlag{
			Name:        "cloud",
			Value:       "AWS",
			Usage:       "cloud provider name",
			Destination: &opts.Cloud,
		},
		clihandler.StringFlag{
			Name:        "engine",
			Value:       "node",
			Usage:       "lambda function engine",
			Destination: &opts.Engine,
		},
		clihandler.StringSliceFlag{
			Name:  "services",
			Usage: "a list of cloud services to initiate in the wrapper",
		},
		clihandler.StringFlag{
			Name:        "libsource",
			Value:       "npm",
			Usage:       "the source where to find library's code",
			Destination: &opts.LibSource,
		},
		clihandler.StringFlag{
			Name:        "libname",
			Usage:       "the name of the library in the source",
			Destination: &opts.LibName,
		},
		clihandler.StringFlag{
			Name:        "output",
			Usage:       "path to save deployable lambda archive",
			Destination: &opts.Output,
		},
		clihandler.BoolFlag{
			Name:        "test",
			Usage:       "flag to run library's unit tests before wrapping into lambda package",
			Destination: &opts.TestRequired,
		},
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
