package cli

import (
	"github.com/urfave/cli"
)

func ProcessCommand() cli.Command {
	return cli.Command{
		Name:  "process",
		Usage: "Execute a logp plan.",
		Action: func(c *cli.Context) error {
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "plan, p", Usage: "Plan to be executed."},
		},
	}
}
