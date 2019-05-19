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

	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}
	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Gitlab-Token")
		if signature != hook.secret {
			return nil, ErrGitLabTokenVerificationFailed
		}
	}

	event := r.Header.Get("X-Gitlab-Event")
	if len(event) == 0 {
		return nil, ErrMissingGitLabEventHeader
	}

	gitLabEvent := Event(event)

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	return eventParsing(gitLabEvent, events, payload)
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
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TagEvents:
		var pl TagEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case ConfidentialIssuesEvents:
		var pl ConfidentialIssueEventPayload
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

	case WikiPageEvents:
		var pl WikiPageEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case PipelineEvents:
		var pl PipelineEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case BuildEvents:
		var pl BuildEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case JobEvents:
		var p1 JobEventPayload
		err := json.Unmarshal([]byte(payload), &p1)
		return p1, err

	case SystemHookEvents:
		var pl SystemHookPayload
		err := json.Unmarshal([]byte(payload), &pl)
		if err != nil {
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
