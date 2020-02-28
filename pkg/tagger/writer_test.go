package tagger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fixutures() (owner string, repo string, pullRequestID int) {
	owner = "garypwhite"
	repo = "tag-your-git"
	pullRequestID = 10
	return
}

func TestGetPullRequestDetailsNoTrailing(t *testing.T) {
	owner, repo, pullRequestID := fixutures()
	fullURL := fmt.Sprintf("https://github.com/%s/%s/pull/%d", owner, repo, pullRequestID)
	cOwner, cRepo, cPullRequestID, err := GetPullRequestDetails(fullURL)
	assert.Equal(t, owner, cOwner)
	assert.Equal(t, repo, cRepo)
	assert.Equal(t, pullRequestID, cPullRequestID)
	assert.Empty(t, err)
}

func TestGetPullRequestDetailsWithTrailing(t *testing.T) {
	owner, repo, pullRequestID := fixutures()
	fullURL := fmt.Sprintf("https://github.com/%s/%s/pull/%d/", owner, repo, pullRequestID)
	cOwner, cRepo, cPullRequestID, err := GetPullRequestDetails(fullURL)
	assert.Equal(t, owner, cOwner)
	assert.Equal(t, repo, cRepo)
	assert.Equal(t, pullRequestID, cPullRequestID)
	assert.Empty(t, err)
}

func TestGetPullRequestDetailsWithNoHTTPS(t *testing.T) {
	owner, repo, pullRequestID := fixutures()
	fullURL := fmt.Sprintf("github.com/%s/%s/pull/%d/", owner, repo, pullRequestID)
	cOwner, cRepo, cPullRequestID, err := GetPullRequestDetails(fullURL)
	assert.Equal(t, owner, cOwner)
	assert.Equal(t, repo, cRepo)
	assert.Equal(t, pullRequestID, cPullRequestID)
	assert.Empty(t, err)
}

func TestGetPullRequestDetailsWithSSHPrefix(t *testing.T) {
	owner, repo, pullRequestID := fixutures()
	fullURL := fmt.Sprintf("ssh://github.com/%s/%s/pull/%d/", owner, repo, pullRequestID)
	cOwner, cRepo, cPullRequestID, err := GetPullRequestDetails(fullURL)
	assert.Equal(t, owner, cOwner)
	assert.Equal(t, repo, cRepo)
	assert.Equal(t, pullRequestID, cPullRequestID)
	assert.Empty(t, err)
}

func TestGetPullRequestDetailsWithGitPrefix(t *testing.T) {
	owner, repo, pullRequestID := fixutures()
	fullURL := fmt.Sprintf("git://github.com/%s/%s/pull/%d/", owner, repo, pullRequestID)
	cOwner, cRepo, cPullRequestID, err := GetPullRequestDetails(fullURL)
	assert.Equal(t, owner, cOwner)
	assert.Equal(t, repo, cRepo)
	assert.Equal(t, pullRequestID, cPullRequestID)
	assert.Empty(t, err)
}

func TestErrorOnGarbageInput(t *testing.T) {
	_, _, _, err := GetPullRequestDetails("heyo this is a bunch of CRUSJKDHJSAIUHDAISJDKLFJLDKJLJF no sasdflklsjah!JKLSD/.12,./,/3.4,/23,./4,23er7d867c78scy bjkn434n 9y89y3r ")
	assert.NotEmpty(t, err)
}

func TestErrorOnEmptyInput(t *testing.T) {
	_, _, _, err := GetPullRequestDetails("")
	assert.NotEmpty(t, err)
}
