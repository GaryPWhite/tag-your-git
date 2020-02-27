package tagger

import (
	context "context"
	"fmt"

	github "github.com/google/go-github/github"
)

// PostTagsToPullRequest will post the tags found in t tagger to pull request pr
func (t Tagger) PostTagsToPullRequest(pr *github.PullRequest) error {
	prRepo := pr.Base.GetRepo()
	files, _, err := t.Client.PullRequests.ListFiles(
		context.Background(),
		*prRepo.Owner.Name,
		*prRepo.Name,
		*pr.Number,
		&github.ListOptions{PerPage: 3000},
	)
	if err != nil {
		fmt.Println("ERROR: Failed to fetch Pull Request files")
		return err
	}
	tagList := []string{}
	for _, tagMatch := range t.TagMapping {
		for _, commitFile := range files {
			if tagMatch.Match(*commitFile.Filename) {
				// if file matches, we post back to PR with the appropriate label(s)
				for _, tag := range tagMatch.GetTags() {
					tagList = append(tagList, tag)
				}
			}
		}
	}
	// after finding all appropriate tags, post them to PR
	t.Client.Issues.AddLabelsToIssue(
		context.Background(),
		*prRepo.Owner.Name,
		*prRepo.Name,
		*pr.Number,
		tagList,
	) // here you might notice that t-y-g doesn't care about duplicates. Waiting to see if that's an issue or not with the API.
	return nil
}
