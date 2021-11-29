package cmd

import (
	"fmt"

	"github.com/gabehoban/go-cli-template/core"
	"github.com/urfave/cli/v2"
)

// Hello command
var Hello = cli.Command{
	Name:        "hello",
	Usage:       "Hello prints a message to the console",
	Description: "Hello prints a message to the console",
	Action:      hello,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "name, n",
			Usage:       "Name to print to the console",
			Value:       "world",
			Destination: &name,
		},
	},
}

func hello(c *cli.Context) error {
	author := core.AppConfig.GetString("app", "author")
	fmt.Printf("%s says: Hello, %s", author, name)

	return nil
}
