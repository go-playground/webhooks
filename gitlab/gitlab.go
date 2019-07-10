package gitlab

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
	ErrEventNotSpecifiedToParse      = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod             = errors.New("invalid HTTP Method")
	ErrMissingGitLabEventHeader      = errors.New("missing X-Gitlab-Event Header")
	ErrGitLabTokenVerificationFailed = errors.New("X-Gitlab-Token validation failed")
	ErrEventNotFound                 = errors.New("event not defined to be parsed")
	ErrParsingPayload                = errors.New("error parsing payload")
	ErrParsingSystemPayload          = errors.New("error parsing system payload")
	// ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

// GitLab hook types
const (
	PushEvents               Event = "Push Hook"
	TagEvents                Event = "Tag Push Hook"
	IssuesEvents             Event = "Issue Hook"
	ConfidentialIssuesEvents Event = "Confidential Issue Hook"
	CommentEvents            Event = "Note Hook"
	MergeRequestEvents       Event = "Merge Request Hook"
	WikiPageEvents           Event = "Wiki Page Hook"
	PipelineEvents           Event = "Pipeline Hook"
	BuildEvents              Event = "Build Hook"
	JobEvents                Event = "Job Hook"
	SystemHookEvents         Event = "System Hook"

	objectPush         string = "push"
	objectTag          string = "tag_push"
	objectMergeRequest string = "merge_request"
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

// Event defines a GitLab hook event type by the X-Gitlab-Event Header
type Event string

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

	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	return hook.ParsePayload(
		payload,
		r.Header.Get("X-Gitlab-Event"),
		r.Header.Get("X-Gitlab-Token"),
		events...,
	)
}

func eventParsing(gitLabEvent Event, events []Event, payload []byte) (interface{}, error) {
	var found bool
	for _, evt := range events {
		if evt == gitLabEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, ErrEventNotFound
	}

	switch gitLabEvent {
	case PushEvents:
		var pl PushEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case TagEvents:
		var pl TagEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case ConfidentialIssuesEvents:
		var pl ConfidentialIssueEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case IssuesEvents:
		var pl IssueEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case CommentEvents:
		var pl CommentEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case MergeRequestEvents:
		var pl MergeRequestEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case WikiPageEvents:
		var pl WikiPageEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case PipelineEvents:
		var pl PipelineEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case BuildEvents:
		var pl BuildEventPayload
		return pl, json.Unmarshal(payload, &pl)
	case JobEvents:
		var pl JobEventPayload
		return pl, json.Unmarshal(payload, &pl)

	case SystemHookEvents:
		var pl SystemHookPayload
		if err := json.Unmarshal(payload, &pl); err != nil {
			return nil, err
		}
		switch pl.ObjectKind {
		case objectPush:
			return eventParsing(PushEvents, events, payload)
		case objectTag:
			return eventParsing(TagEvents, events, payload)
		case objectMergeRequest:
			return eventParsing(MergeRequestEvents, events, payload)
		default:
			return nil, fmt.Errorf("unknown system hook event %s", gitLabEvent)
		}
	default:
		return nil, fmt.Errorf("unknown event %s", gitLabEvent)
	}
}

// ParsePayload verifies and parses the events from a payload and string
// metadata (event type and token), and returns the payload object or an error.
//
// Similar to Parse (which uses this method under the hood), this is useful in
// cases where payloads are not represented as HTTP requests - for example are
// put on a queue for pull processing.
func (hook Webhook) ParsePayload(payload []byte, eventType, token string, events ...Event) (interface{}, error) {
	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		if token != hook.secret {
			return nil, ErrGitLabTokenVerificationFailed
		}
	}

	if len(eventType) == 0 {
		return nil, ErrMissingGitLabEventHeader
	}

	gitLabEvent := Event(eventType)

	return eventParsing(gitLabEvent, events, payload)
}
