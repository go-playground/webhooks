package gitea

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"code.gitea.io/gitea/modules/structs"
)

// parse errors
var (
	ErrEventNotSpecifiedToParse = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod        = errors.New("invalid HTTP Method")
	ErrMissingGiteaEventHeader  = errors.New("missing X-Gitea-Event Header")
	ErrEventNotFound            = errors.New("event not defined to be parsed")
	ErrParsingPayload           = errors.New("error parsing payload")
	ErrSecretNotMatch           = errors.New("secret not match")
)

// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// Secret registers the GitLab secret
func (WebhookOptions) Secret(secret string) Option {
	return func(hook *Webhook) error {
		hook.secret = secret
		return nil
	}
}

// Webhook instance contains all methods needed to process events
type Webhook struct {
	secret string
}

// Event defines a Gitea hook event type
type Event string

// Gitea hook types
const (
	CreateEvent       Event = "create"
	DeleteEvent       Event = "delete"
	ForkEvent         Event = "fork"
	PushEvent         Event = "push"
	IssuesEvent       Event = "issues"
	IssueCommentEvent Event = "issue_comment"
	PullRequestEvent  Event = "pull_request"
	ReleaseEvent      Event = "release"
)

// New creates and returns a WebHook instance denoted by the Provider type
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("Error applying Option")
		}
	}
	return hook, nil
}

func (hook Webhook) verifySecret(secret string) error {
	if len(hook.secret) > 0 && hook.secret != secret {
		return ErrSecretNotMatch
	}

	return nil
}

// Parse verifies and parses the events specified and returns the payload object or an error
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}
	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	event := r.Header.Get("X-Gitea-Event")
	if len(event) == 0 {
		return nil, ErrMissingGiteaEventHeader
	}

	giteaEvent := Event(event)

	var found bool
	for _, evt := range events {
		if evt == giteaEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, ErrEventNotFound
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	switch giteaEvent {
	case CreateEvent:
		var pl structs.CreatePayload
		err = json.Unmarshal([]byte(payload), &pl)

		if err == nil {
			err = hook.verifySecret(pl.Secret)
		}

		return pl, err

	case ReleaseEvent:
		var pl structs.ReleasePayload
		err = json.Unmarshal([]byte(payload), &pl)

		if err == nil {
			err = hook.verifySecret(pl.Secret)
		}

		return pl, err

	case PushEvent:
		var pl structs.PushPayload
		err = json.Unmarshal([]byte(payload), &pl)

		if err == nil {
			err = hook.verifySecret(pl.Secret)
		}

		return pl, err

	case DeleteEvent:
		var pl structs.DeletePayload
		err = json.Unmarshal([]byte(payload), &pl)

		if err == nil {
			err = hook.verifySecret(pl.Secret)
		}

		return pl, err

	case ForkEvent:
		var pl structs.ForkPayload
		err = json.Unmarshal([]byte(payload), &pl)

		if err == nil {
			err = hook.verifySecret(pl.Secret)
		}

		return pl, err

	case IssuesEvent:
		var pl structs.IssuePayload
		err = json.Unmarshal([]byte(payload), &pl)

		if err == nil {
			err = hook.verifySecret(pl.Secret)
		}

		return pl, err

	case IssueCommentEvent:
		var pl structs.IssueCommentPayload
		err = json.Unmarshal([]byte(payload), &pl)

		if err == nil {
			err = hook.verifySecret(pl.Secret)
		}

		return pl, err

	case PullRequestEvent:
		var pl structs.PullRequestPayload
		err = json.Unmarshal([]byte(payload), &pl)

		if err == nil {
			err = hook.verifySecret(pl.Secret)
		}

		return pl, err

	default:
		return nil, fmt.Errorf("unknown event %s", giteaEvent)
	}
}
