package main

import (
	"os"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
)

func main() {
	cli.NewCliApp().Run(os.Args)
}
