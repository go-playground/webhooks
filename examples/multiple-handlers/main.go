package main

import (
	"fmt"
	"strconv"

	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

const (
	path = "/webhooks"
	port = 3016
)

func main() {
	hook := github.New(&github.Config{Secret: "MyGitHubSuperSecretSecrect...?"})
	hook.RegisterEvents(HandleRelease, github.ReleaseEvent)
	hook.RegisterEvents(HandlePullRequest, github.PullRequestEvent)

	err := webhooks.Run(hook, ":"+strconv.Itoa(port), path)
	if err != nil {
		fmt.Println(err)
	}
}

// HandleRelease handles GitHub release events
func HandleRelease(payload interface{}, header webhooks.Header) {
	fmt.Println("Handling Release")

	pl := payload.(github.ReleasePayload)

	// only want to compile on full releases
	if pl.Release.Draft || pl.Release.Prerelease || pl.Release.TargetCommitish != "master" {
		return
	}

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}

// HandlePullRequest handles GitHub pull_request events
func HandlePullRequest(payload interface{}, header webhooks.Header) {

	fmt.Println("Handling Pull Request")

	pl := payload.(github.PullRequestPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}
