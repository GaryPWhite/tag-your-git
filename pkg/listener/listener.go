package pkg

import (
	context "context"
	"fmt"
	"net/http"

	tagMatcher "github.com/garypwhite/tag-your-git/pkg/tagMatcher"
	github "github.com/google/go-github/github"
)

// Listener is the struct to include
type Listener struct {
	Client     *github.Client
	SecretKey  []byte
	TagMapping []tagMatcher.TagMatch
}

// Start will begin listening for incoming webhooks
func (l Listener) Start() error {
	// config
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Here is my current configuration:")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
		payload, err := github.ValidatePayload(r, l.SecretKey)
		if err != nil {
			fmt.Errorf("Failed to parse payload: %+v", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Could not parse payload")
		}
		event, err := github.ParseWebHook(github.WebHookType(r), payload)
		if err != nil {
			fmt.Errorf("Failed to parse payload: %+v", err)
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

func (l Listener) processPullRequestEvent(event *github.PullRequestEvent) error {
	switch event.GetAction() {
	case "opened":
	case "edited":
		// get files from pull request
		pr := event.GetPullRequest()
		prRepo := event.GetRepo()
		files, _, err := l.Client.PullRequests.ListFiles(
			context.Background(),
			*prRepo.Owner.Name,
			*prRepo.Name,
			*pr.Number,
			&github.ListOptions{PerPage: 3000},
		)
		if err != nil {
			fmt.Errorf("failed to fetch Pull Request")
			return err
		}
		for _, tagMatch := range l.TagMapping {
			for _, commitFile := range files {
				if tagMatch.Match(*commitFile.Filename) {
					// if file matches, we post back to PR with the appropriate label(s)
					l.Client.Issues.AddLabelsToIssue(
						context.Background(),
						*prRepo.Owner.Name,
						*prRepo.Name,
						*pr.Number,
						tagMatch.GetTags(),
					)
				}
			}
		}
		break
	default: // do nothing
		return nil
	}
	return nil
}
