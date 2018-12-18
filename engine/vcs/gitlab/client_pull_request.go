package gitlab

import (
	"context"
	"fmt"

	"github.com/ovh/cds/sdk"
)

// PullRequests fetch all the pull request for a repository
func (c *gitlabClient) PullRequests(context.Context, string) ([]sdk.VCSPullRequest, error) {
	return []sdk.VCSPullRequest{}, nil
}

// PullRequestComment push a new comment on a pull request
func (c *gitlabClient) PullRequestComment(context.Context, string, int, string) error {
	return nil
}

// PullRequestCreate create a new pullrequest
func (c *gitlabClient) PullRequestCreate(repo string, branchName string, msg string) (sdk.VCSPullRequest, error) {
	return sdk.VCSPullRequest{}, fmt.Errorf("not yet implemented")
}
