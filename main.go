package main

import (
	"context"
	"log"
	"os"

	tagger "github.com/garypwhite/tag-your-git/pkg/tagger"
	tagmatcher "github.com/garypwhite/tag-your-git/pkg/tagmatcher"
	github "github.com/google/go-github/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

func makeClient(url, apikey string) (*github.Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apikey},
	)
	tc := oauth2.NewClient(ctx, ts)
	if len(url) != 0 {
		return github.NewEnterpriseClient(url, url, tc) // use default client (nil)
	}
	return github.NewClient(tc), nil
}

func makeTagger(enterpriseURL, apikey, tagsJSON string) (*tagger.Tagger, error) {
	githubClient, err := makeClient(enterpriseURL, apikey)
	if err != nil {
		return nil, err
	}
	mapping, err := tagmatcher.Unmarshal([]byte(tagsJSON))
	if err != nil {
		return nil, err
	}
	tagger := tagger.Tagger{
		Client:     githubClient,
		TagMapping: mapping,
	}
	return &tagger, nil
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "enterprise-URL",
				Aliases: []string{"enterprise-url"},
				EnvVars: []string{"TYG_ENTERPRISE_URL"},
			},
			&cli.StringFlag{
				Name:     "tags",
				Aliases:  []string{"t"},
				Usage:    "JSON struct of tags to post. Uses the format [ { [patterns], [tags] } ]",
				EnvVars:  []string{"TYG_TAGS"},
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
					tagger, err := makeTagger(c.String("enterprise-URL"), c.String("git-api-key"), c.String("tags"))
					if err != nil {
						return err
					}

					return tagger.PostTagsToPullRequest(nil)
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
				Usage: "Listen for incoming webhooks to process + post tags according to (--tags, /etc/tag-your-git/tags.json, TYG_TAGS) specified",
				Action: func(c *cli.Context) error {
					// use CLI context to make client
					tagger, err := makeTagger(c.String("enterprise-URL"), c.String("git-api-key"), c.String("tags"))
					if err != nil {
						return err
					}
					secretKey := []byte(c.String("webhook-secret-key"))
					return tagger.Listen(secretKey)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "webhook-secret-key",
						Usage:   "secret provided from GitHub webhook",
						EnvVars: []string{"TYG_WEBHOOK_SECRET_KEY"},
					},
				},
			},
		},
		Name:    "tag-your-git (tyg)",
		Version: "v0.1.0",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Gary White Jr.",
				Email: "garypwhite@github",
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
