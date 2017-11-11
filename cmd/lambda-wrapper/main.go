package main

import (
	"fmt"
	"os"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/cli"
)

func main() {
	fmt.Println("Hello, Lambda Wrapper")
	cli.NewCliApp().Run(os.Args)
}
