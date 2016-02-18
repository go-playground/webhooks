Library webhooks
================
<img align="right" src="https://raw.githubusercontent.com/go-playground/webhooks/v1/logo.png">
![Project status](https://img.shields.io/badge/version-1.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/projects/5b9e2eda-8f8d-40aa-8cb4-e3f6120171fe/587820/badge.svg)](https://semaphoreci.com/joeybloggs/webhooks)
[![Coverage Status](https://coveralls.io/repos/go-playground/webhooks/badge.svg?branch=v1&service=github)](https://coveralls.io/github/go-playground/webhooks?branch=v1)
[![Go Report Card](https://goreportcard.com/badge/go-playground/webhooks)](https://goreportcard.com/report/go-playground/webhooks)
[![GoDoc](https://godoc.org/gopkg.in/go-playground/webhooks.v1?status.svg)](https://godoc.org/gopkg.in/go-playground/webhooks.v1)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Library webhooks allows for easy recieving and parsing of GitHub & Bitbucket Webhook Events

Features:

* Parses the entire payload, not just a few fields.
* Fields + Schema directly lines up with webhook posted json

Notes:

* Github - Currently only accepting json payloads.

Installation
------------

Use go get.

	go get gopkg.in/go-playground/webhooks.v1

or to update

	go get -u gopkg.in/go-playground/webhooks.v1

Then import the validator package into your own code.

	import "gopkg.in/go-playground/webhooks.v1"

Usage and documentation
------

Please see http://godoc.org/gopkg.in/go-playground/webhooks.v1 for detailed usage docs.

##### Examples:

Multiple Handlers for each event you subscribe to
```go
package main

import (
	"fmt"
	"strconv"

	"gopkg.in/go-playground/webhooks.v1"
	"gopkg.in/go-playground/webhooks.v1/github"
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
func HandleRelease(payload interface{}) {

	fmt.Println("Handling Release")

	pl := payload.(github.ReleasePayload)

	// only want to compile on full releases
	if pl.Release.Draft || pl.Release.Prelelease || pl.Release.TargetCommitish != "master" {
		return
	}

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}

// HandlePullRequest handles GitHub pull_request events
func HandlePullRequest(payload interface{}) {

	fmt.Println("Handling Pull Request")

	pl := payload.(github.PullRequestPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}
```

Single receiver for events you subscribe to
```go
package main

import (
	"fmt"
	"strconv"

	"gopkg.in/go-playground/webhooks.v1"
	"gopkg.in/go-playground/webhooks.v1/github"
)

const (
	path = "/webhooks"
	port = 3016
)

func main() {
	hook := github.New(&github.Config{Secret: "MyGitHubSuperSecretSecrect...?"})
	hook.RegisterEvents(HandleMultiple, github.ReleaseEvent, github.PullRequestEvent) // Add as many as you want

	err := webhooks.Run(hook, ":"+strconv.Itoa(port), path)
	if err != nil {
		fmt.Println(err)
	}
}

// HandleMultiple handles multiple GitHub events
func HandleMultiple(payload interface{}) {

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
```

Contributing
------

Pull requests for other service like BitBucket are welcome!

There will always be a development branch for each version i.e. `v1-development`. In order to contribute, 
please make your pull requests against those branches.

If the changes being proposed or requested are breaking changes, please create an issue, for discussion
or create a pull request against the highest development branch for example this package has a
v1 and v1-development branch however, there will also be a v2-development branch even though v2 doesn't exist yet.

License
------
Distributed under MIT License, please see license file in code for more details.
