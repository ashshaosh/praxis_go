package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Just print HW"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name, n",
			Usage:       "Who are you?",
			EnvVar:      "",
			Hidden:      false,
			Value:       "World",
			Destination: new(string),
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s\n", name)
		return nil
	}
	app.Run(os.Args)
}
