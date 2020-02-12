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
				Name:     "tags-to-post",
				Aliases:  []string{"tags-list"},
				Usage:    "JSON struct of tags to post. Uses the format { { [directories], [tags] },  }",
				EnvVars:  []string{"TYG_TAG_LIST"},
				FilePath: "/etc/tag-your-git/tags.json", // Recommended to use a docker image
			},
			&cli.StringFlag{
				Name:     "git-api-key",
				Aliases:  []string{"gapi-key"},
				Usage:    "API key to post to specified Git API",
				EnvVars:  []string{"TYG_GIT_API_KEY"},
				FilePath: "/etc/tag-your-git/git-api-key", // Recommended to use a docker image
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "tag",
				Usage: "Tag one Pull Request (--pr, --pull-request) with a specified tags JSON array (--tags, --t)",
				Action: func(c *cli.Context) error {
					//TODO: Call Git API, after figuring out arguments
					fmt.Printf("nothing yet")
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "pull-request",
						Aliases:  []string{"pr"},
						Usage:    "Pull Request URL to tag",
						EnvVars:  []string{"TYG_PULL_REQUEST"},
						Required: true,
					},
				},
			},
			{
				Name:  "listen",
				Usage: "Listen for incoming webhooks to process + post tags according to (--tags, ) specified",
				Action: func(c *cli.Context) error {
					//TODO:

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
