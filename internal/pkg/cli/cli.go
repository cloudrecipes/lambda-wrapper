// Package cli contains methods to work with command line arguments.
package cli

import clihandler "gopkg.in/urfave/cli.v1"

// Name is the applicaion name.
const Name string = "lambda-wrapper"

// Version is the current applicaion version.
const Version string = "0.1.0"

// Description is the applicaion description.
const Description string = "Lambda wrapper application. Used to wrap libraries into lambda functions to be run on cloud."

type CliApp struct {
	App *clihandler.App
}

func NewCliApp() *CliApp {
	app := clihandler.NewApp()
	app.Name = Name
	app.Version = Version
	app.Description = Description
	app.Flags = []clihandler.Flag{
		clihandler.StringFlag{Name: "cloud", Value: "AWS", Usage: "cloud provider name"},
		clihandler.StringFlag{Name: "engine", Value: "node", Usage: "lambda function engine"},
		clihandler.StringSliceFlag{Name: "services", Usage: "a list of cloud services to initiate in the wrapper"},
		clihandler.StringFlag{Name: "source", Value: "npm", Usage: "the source where to find library's code"},
		clihandler.StringFlag{Name: "name", Usage: "the name of the library in the source"},
		clihandler.StringFlag{Name: "output", Usage: "path to save deployable lambda archive"},
		clihandler.BoolFlag{Name: "test", Usage: "lag to run library's unit tests before wrapping into lambda package"},
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
