Library webhooks
================
<img align="right" src="https://raw.githubusercontent.com/go-playground/webhooks/v6/logo.png">![Project status](https://img.shields.io/badge/version-6.2.0-green.svg)
[![Test](https://github.com/go-playground/webhooks/workflows/Test/badge.svg?branch=master)](https://github.com/go-playground/webhooks/actions)
[![Coverage Status](https://coveralls.io/repos/go-playground/webhooks/badge.svg?branch=master&service=github)](https://coveralls.io/github/go-playground/webhooks?branch=master)
[![Go Report Card](https://goreportcard.com/badge/go-playground/webhooks)](https://goreportcard.com/report/go-playground/webhooks)
[![GoDoc](https://godoc.org/github.com/go-playground/webhooks/v6?status.svg)](https://godoc.org/github.com/go-playground/webhooks/v6)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Library webhooks allows for easy receiving and parsing of GitHub, Bitbucket, GitLab, Docker Hub and Gogs Webhook Events

Features:

* Parses the entire payload, not just a few fields.
* Fields + Schema directly lines up with webhook posted json

Notes:

* Currently only accepting json payloads.

Installation
------------

Use go get.

```shell
go get -u github.com/go-playground/webhooks/v6
```

Then import the package into your own code.

	import "github.com/go-playground/webhooks/v6"

Usage and Documentation
------

Please see http://godoc.org/github.com/go-playground/webhooks/v6 for detailed usage docs.

##### Examples:
```go
package main

import (
	"fmt"

	"net/http"

	"github.com/go-playground/webhooks/v6/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecret...?"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn't one of the ones asked to be parsed
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
