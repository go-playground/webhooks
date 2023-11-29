package gitee

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// parse errors
var (
	ErrMethodNotAllowed             = errors.New("method not allowed")
	ErrMissingEvents                = errors.New("missing X-Gitee-Events")
	ErrMissingEventHeader           = errors.New("missing X-Gitee-Event Header")
	ErrMissingTimestampHeader       = errors.New("missing X-Gitee-Timestamp Header")
	ErrMissingToken                 = errors.New("missing X-Gitee-Token")
	ErrContentType                  = errors.New("hook only accepts content-type: application/json")
	ErrRequestBody                  = errors.New("failed to read request body")
	ErrGiteeTokenVerificationFailed = errors.New("failed to verify token")
	ErrParsingPayload               = errors.New("failed to parsing payload")
	ErrEventNotFound                = errors.New("failed to find event")
	// ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

// Gitee hook types
const (
	PushEvents         Event = "Push Hook"
	TagEvents          Event = "Tag Push Hook"
	IssuesEvents       Event = "Issue Hook"
	CommentEvents      Event = "Note Hook"
	MergeRequestEvents Event = "Merge Request Hook"
)

// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// Secret registers the Gitee secret
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

// Event defines a Gitee hook event type by the X-Gitee-Event Header
type Event string

// New creates and returns a WebHook instance denoted by the Provider type
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("error applying Option")
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
		return nil, ErrMissingEvents
	}
	if r.Method != http.MethodPost {
		return nil, ErrMethodNotAllowed
	}

	timeStamp := r.Header.Get("X-Gitee-Timestamp")
	if len(timeStamp) == 0 {
		return nil, ErrMissingTimestampHeader
	}

	contentType := r.Header.Get("content-type")
	if contentType != "application/json" {
		return nil, ErrContentType
	}

	event := r.Header.Get("X-Gitee-Event")
	if len(event) == 0 {
		return nil, ErrMissingEventHeader
	}

	giteeEvent := Event(event)

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Gitee-Token")
		if signature != hook.secret {
			return nil, ErrGiteeTokenVerificationFailed
		}
	}

	return eventParsing(giteeEvent, events, payload)
}

func eventParsing(giteeEvent Event, events []Event, payload []byte) (interface{}, error) {

	var found bool
	for _, evt := range events {
		if evt == giteeEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, ErrEventNotFound
	}

	switch giteeEvent {
	case PushEvents:
		var pl PushEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TagEvents:
		var pl TagEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case IssuesEvents:
		var pl IssueEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case CommentEvents:
		var pl CommentEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case MergeRequestEvents:
		var pl MergeRequestEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	default:
		return nil, fmt.Errorf("unknown event %s", giteeEvent)
	}
}
