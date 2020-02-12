# Tag your git!

This program is intended to help you apply labels (tags) to a github repo's pull requests.

Under development now, this appliance would require:

```zsh
export TYG_TAGS = '{ { paths: ["some-directory/*","maybe-more-but-one-works-too"], tags: ["tag-as-some-directory", "anothertag"] },...}'
# ^ TYG_TAGS could (should?) be a /etc/tag-your-git/tags.json file.
export TYG_GIT_API_KEY =<A git API Key>
# /etc/tag-your-git/git-api-key"
```

you can override files with environment variables, and environment variables with command line arguments.

then you may use the command line to either `tag` a particular PR, or `listen` for incoming webhooks. `listen` as a command will be something to deploy into a docker image / using a hosting service. Stay tuned for a docker image on docker hub and/or a preferred deployment mechanism.

After deployment, we'd set an incoming webhook from git to hit the `listen`ing service, and expect that tag-your-git would read from the provided tags, and post tags appropriately.