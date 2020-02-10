package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "pull-request",
				Aliases: []string{"pr"},
				Usage:   "Pull Request to tag",
				EnvVars: []string{"TYG_PULL_REQUEST"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "tag",
				Usage: "Tag one Pull Request (--pr, --pull-request) with a specified tag (-t, --tag)",
				Action: func(c *cli.Context) error {
					//TODO: Call Git API, after figuring out arguments
					fmt.Printf("nothing yet")
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
