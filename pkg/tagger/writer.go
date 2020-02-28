package tagger

import (
	context "context"
	"fmt"
	"strconv"
	"strings"

	github "github.com/google/go-github/github"
)

// GetPullRequestDetails will provide the necessary fields for fetching from the GH client through parsing a URL
func GetPullRequestDetails(prURL string) (owner, repo string, pullRequestNumber int, err error) {
	// assuming a url with trailing ..../owner/repo/pull/number....
	splitStrings := strings.Split(strings.Trim(prURL, "/"), "/")
	lenStrings := len(splitStrings)
	if lenStrings < 4 { // sanity check -- even partial URL will have at least 4
		err = fmt.Errorf("Invalid PR input, should be like https://github.com/garypwhite/tag-your-git/pull/10, got %s", prURL)
		return
	}
	owner = splitStrings[lenStrings-4]
	repo = splitStrings[lenStrings-3]
	pullRequestNumber, err = strconv.Atoi(splitStrings[lenStrings-1])
	return
}

// PostTagsToPullRequest will post the tags found in t tagger to pull request pr
func (t Tagger) PostTagsToPullRequest(pr *github.PullRequest) error {
	owner, repo, number, err := GetPullRequestDetails(pr.GetIssueURL())
	if err != nil {
		return err
	}
	files, _, err := t.Client.PullRequests.ListFiles(
		context.Background(),
		owner,
		repo,
		number,
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
		owner,
		repo,
		number,
		tagList,
	) // here you might notice that t-y-g doesn't care about duplicates. Waiting to see if that's an issue or not with the API.
	return nil
}
