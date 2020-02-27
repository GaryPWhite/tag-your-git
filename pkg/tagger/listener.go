package tagger

import (
	"fmt"
	"net/http"

	tagMatcher "github.com/garypwhite/tag-your-git/pkg/tagmatcher"
	github "github.com/google/go-github/github"
)

// Tagger is the struct to include
type Tagger struct {
	Client     *github.Client
	TagMapping []tagMatcher.TagMatch
}

// Listen will begin listening for incoming webhooks
func (l Tagger) Listen(GithubSecretKey []byte) error {
	// config
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Here is my current configuration:")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
		payload, err := github.ValidatePayload(r, GithubSecretKey)
		if err != nil {
			fmt.Printf("ERROR: Failed to parse payload: %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Could not parse payload")
		}
		event, err := github.ParseWebHook(github.WebHookType(r), payload)
		if err != nil {
			fmt.Printf("ERROR: Failed to parse payload: %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Could not parse payload")
		}
		switch event := event.(type) {
		case *github.PullRequestEvent:
			err = l.processPullRequestEvent(event)
			break
		default:
			fmt.Fprintf(w, "I only care about pull request events!")
		}
	})

	return http.ListenAndServe(":8000", nil)
}

func (l Tagger) processPullRequestEvent(event *github.PullRequestEvent) error {
	switch event.GetAction() {
	case "opened":
	case "edited": // add tags as needed to the PR
		// get files from pull request
		return l.PostTagsToPullRequest(event.GetPullRequest())
	default: // do nothing
		return nil
	}
	return nil
}
