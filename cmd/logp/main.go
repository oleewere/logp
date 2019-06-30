package main

import (
	"fmt"
	"os"

	logpcli "github.com/oleewere/logp/cli"
	"github.com/urfave/cli"
)

// Version that will be generated during the build as a constant
var Version string

// GitRevString that will be generated during the build as a constant - represents git revision value
var GitRevString string

func main() {
	app := cli.NewApp()
	app.Name = "logp"
	app.Usage = "CLI tool for processing local or remote logs"
	app.EnableBashCompletion = true
	app.UsageText = "logp command [command options] [arguments...]"
	if len(Version) > 0 {
		app.Version = Version
	} else {
		app.Version = "0.1.0"
	}
	if len(GitRevString) > 0 {
		app.Version = app.Version + fmt.Sprintf(" (git short hash: %v)", GitRevString)
	}

	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, logpcli.ProcessCommand())
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
