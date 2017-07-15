package main

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

const (
	path = "/webhooks"
	port = 3016
)

type myLogger struct {
	PrintDebugs bool
}

func (l *myLogger) Info(msg string) {
	log.Println(msg)
}

func (l *myLogger) Error(msg string) {
	log.Println(msg)
}

func (l *myLogger) Debug(msg string) {
	if !l.PrintDebugs {
		return
	}
	log.Println(msg)
}

func main() {
	// webhooks.DefaultLog=webhooks.NewLogger(true)
	//
	// or override with your own
	webhooks.DefaultLog = &myLogger{PrintDebugs: true}

	hook := github.New(&github.Config{Secret: "MyGitHubSuperSecretSecrect...?"})
	hook.RegisterEvents(HandleMultiple, github.ReleaseEvent, github.PullRequestEvent) // Add as many as you want

	err := webhooks.Run(hook, ":"+strconv.Itoa(port), path)
	if err != nil {
		fmt.Println(err)
	}
}

// HandleMultiple handles multiple GitHub events
func HandleMultiple(payload interface{}, header webhooks.Header) {
	fmt.Println("Handling Payload..")

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
}
