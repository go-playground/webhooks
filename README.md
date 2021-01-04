Library webhooks
================
<img align="right" src="https://raw.githubusercontent.com/go-playground/webhooks/v5/logo.png">![Project status](https://img.shields.io/badge/version-5.17.0-green.svg)
[![Build Status](https://travis-ci.org/go-playground/webhooks.svg?branch=v5)](https://travis-ci.org/go-playground/webhooks)
[![Coverage Status](https://coveralls.io/repos/go-playground/webhooks/badge.svg?branch=v5&service=github)](https://coveralls.io/github/go-playground/webhooks?branch=v5)
[![Go Report Card](https://goreportcard.com/badge/go-playground/webhooks)](https://goreportcard.com/report/go-playground/webhooks)
[![GoDoc](https://godoc.org/gopkg.in/go-playground/webhooks.v5?status.svg)](https://godoc.org/gopkg.in/go-playground/webhooks.v5)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Library webhooks allows for easy receiving and parsing of GitHub, Bitbucket and GitLab Webhook Events

Features:

* Parses the entire payload, not just a few fields.
* Fields + Schema directly lines up with webhook posted json

Notes:

* Currently only accepting json payloads.

Installation
------------

Use go get.

```shell
go get -u gopkg.in/go-playground/webhooks.v5
```

Then import the package into your own code.

	import "gopkg.in/go-playground/webhooks.v5"

Usage and Documentation
------

Please see http://godoc.org/gopkg.in/go-playground/webhooks.v5 for detailed usage docs.

##### Examples:
```go
package main

import (
	"fmt"

	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecrect...?"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
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

```

Contributing
------

Pull requests for other services are welcome!

If the changes being proposed or requested are breaking changes, please create an issue for discussion.

License
------
Distributed under MIT License, please see license file in code for more details.
