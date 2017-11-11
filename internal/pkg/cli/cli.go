// Package cli contains methods to work with command line arguments.
package cli

import clihandler "gopkg.in/urfave/cli.v1"

// Name is the applicaion name.
const Name string = "lambda-wrapper"

// Version is the current applicaion version.
const Version string = "0.1.0"

// Description is the applicaion description.
const Description string = "Lambda wrapper application. Used to wrap libraries into lambda functions to be run on cloud."

// ShortDescription is the applicaion description which comes in the NAME section.
const ShortDescription string = ""

// Usage is the applicaion usage text.
const Usage string = "[global options] command [command options]"

type CliApp struct {
	App *clihandler.App
}

func NewCliApp() *CliApp {
	app := clihandler.NewApp()
	app.Name = Name
	app.Version = Version
	app.Description = Description
	app.Usage = ShortDescription
	app.UsageText = Usage
	app.Flags = []clihandler.Flag{
		clihandler.StringFlag{Name: "cloud", Value: "AWS", Usage: "cloud provider name"},
		clihandler.StringFlag{Name: "engine", Value: "node", Usage: "lambda function engine"},
		clihandler.StringSliceFlag{Name: "services", Usage: "a list of cloud services to initiate in the wrapper"},
		clihandler.StringFlag{Name: "libsource", Value: "npm", Usage: "the source where to find library's code"},
		clihandler.StringFlag{Name: "libname", Usage: "the name of the library in the source"},
		clihandler.StringFlag{Name: "output", Usage: "path to save deployable lambda archive"},
	}

	app.Commands = []clihandler.Command{
		{
			Name:        "testlib",
			Aliases:     []string{"t"},
			Usage:       "Runs library unit tests",
			Description: "This command runs unit tests of the library.",
		},
		{
			Name:        "wraplib",
			Aliases:     []string{"w"},
			Usage:       "Wraps library into labmda function",
			Description: "This command wraps library into the archive to deploy to cloud as a lambda function.",
		},
	}

	return &CliApp{App: app}
}

func (cli *CliApp) Run(args []string) error {
	return cli.App.Run(args)
}

// GetRuntimeFlags returns the list of flags and it's values with which applicaion
// has been started.
func (cli *CliApp) GetRuntimeFlags() {

}
