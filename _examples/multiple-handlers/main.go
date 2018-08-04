package main

import (
	"fmt"

	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path1 = "/webhooks1"
	path2 = "/webhooks2"
)

func main() {
	hook1, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecrect...?"))
	hook2, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecrect2...?"))

	http.HandleFunc(path1, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook1.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})

	http.HandleFunc(path2, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook2.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	http.ListenAndServe(":3000", nil)
}
