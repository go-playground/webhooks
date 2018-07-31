package gogs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	client "github.com/gogits/go-gogs-client"
)

// parse errors
var (
	ErrEventNotSpecifiedToParse   = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod          = errors.New("invalid HTTP Method")
	ErrMissingGogsEventHeader     = errors.New("missing X-Gogs-Event Header")
	ErrMissingGogsSignatureHeader = errors.New("missing X-Gogs-Signature Header")
	ErrEventNotFound              = errors.New("event not defined to be parsed")
	ErrParsingPayload             = errors.New("error parsing payload")
	ErrHMACVerificationFailed     = errors.New("HMAC verification failed")
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

// Event defines a Gogs hook event type
type Event string

// Gogs hook types
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

	event := r.Header.Get("X-Gogs-Event")
	if len(event) == 0 {
		return nil, ErrMissingGogsEventHeader
	}

	gogsEvent := Event(event)

	var found bool
	for _, evt := range events {
		if evt == gogsEvent {
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

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Gogs-Signature")
		if len(signature) == 0 {
			return nil, ErrMissingGogsSignatureHeader
		}

		mac := hmac.New(sha256.New, []byte(hook.secret))
		_, _ = mac.Write(payload)

		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
			return nil, ErrHMACVerificationFailed
		}
	}

	switch gogsEvent {
	case CreateEvent:
		var pl client.CreatePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case ReleaseEvent:
		var pl client.ReleasePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case PushEvent:
		var pl client.PushPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case DeleteEvent:
		var pl client.DeletePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case ForkEvent:
		var pl client.ForkPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case IssuesEvent:
		var pl client.IssuesPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case IssueCommentEvent:
		var pl client.IssueCommentPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case PullRequestEvent:
		var pl client.PullRequestPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err

	default:
		return nil, fmt.Errorf("unknown event %s", gogsEvent)
	}
}
