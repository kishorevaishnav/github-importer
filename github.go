package main

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	ghAccessToken, ghOwner, ghRepo string
)

func ghInit(accessToken, userName, repoName string) {
	// Replace these values with your GitHub credentials and repository information.
	// ghAccessToken = os.Getenv("GITHUB_ACCESSTOKEN")
	// ghOwner = os.Getenv("GITHUB_USERNAME")
	// ghRepo = os.Getenv("GITHUB_REPONAME")

	ghAccessToken, ghOwner, ghRepo = accessToken, userName, repoName

	// TODO#: Need to handle error.
}

func ghSubmitIssue(title, description, labels, assignee string) (string, error) {
	// Create a GitHub client with authentication.
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Create an issue request.
	issueRequest := &github.IssueRequest{
		Title:    github.String(title),
		Body:     github.String(description),
		Labels:   &[]string{labels},
		Assignee: github.String(assignee),
	}

	// Create the issue in the repository.
	issue, _, err := client.Issues.Create(ctx, ghOwner, ghRepo, issueRequest)
	if err != nil {
		return "", err
	}

	return *issue.HTMLURL, nil
}
