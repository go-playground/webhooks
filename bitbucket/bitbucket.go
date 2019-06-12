package bitbucket

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
	ErrEventNotSpecifiedToParse = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod        = errors.New("invalid HTTP Method")
	ErrMissingHookUUIDHeader    = errors.New("missing X-Hook-UUID Header")
	ErrMissingEventKeyHeader    = errors.New("missing X-Event-Key Header")
	ErrEventNotFound            = errors.New("event not defined to be parsed")
	ErrParsingPayload           = errors.New("error parsing payload")
	ErrUUIDVerificationFailed   = errors.New("UUID verification failed")
)

// Webhook instance contains all methods needed to process events
type Webhook struct {
	uuid string
}

// Event defines a Bitbucket hook event type
type Event string

// Bitbucket hook types
const (
	RepoPushEvent                  Event = "repo:push"
	RepoForkEvent                  Event = "repo:fork"
	RepoUpdatedEvent               Event = "repo:updated"
	RepoCommitCommentCreatedEvent  Event = "repo:commit_comment_created"
	RepoCommitStatusCreatedEvent   Event = "repo:commit_status_created"
	RepoCommitStatusUpdatedEvent   Event = "repo:commit_status_updated"
	IssueCreatedEvent              Event = "issue:created"
	IssueUpdatedEvent              Event = "issue:updated"
	IssueCommentCreatedEvent       Event = "issue:comment_created"
	PullRequestCreatedEvent        Event = "pullrequest:created"
	PullRequestUpdatedEvent        Event = "pullrequest:updated"
	PullRequestApprovedEvent       Event = "pullrequest:approved"
	PullRequestUnapprovedEvent     Event = "pullrequest:unapproved"
	PullRequestMergedEvent         Event = "pullrequest:fulfilled"
	PullRequestDeclinedEvent       Event = "pullrequest:rejected"
	PullRequestCommentCreatedEvent Event = "pullrequest:comment_created"
	PullRequestCommentUpdatedEvent Event = "pullrequest:comment_updated"
	PullRequestCommentDeletedEvent Event = "pullrequest:comment_deleted"
)

// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// UUID registers the BitBucket secret
func (WebhookOptions) UUID(uuid string) Option {
	return func(hook *Webhook) error {
		hook.uuid = uuid
		return nil
	}
}

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
		r.Header.Get("X-Event-Key"),
		r.Header.Get("X-Hook-UUID"),
		events...,
	)
}

// ParsePayload verifies and parses the events from a payload and string
// metadata (event type and UUID), and returns the payload object or an
// error.
//
// Similar to Parse (which uses this method under the hood), this is useful in
// cases where payloads are not represented as HTTP requests - for example are
// put on a queue for pull processing.
func (hook Webhook) ParsePayload(payload []byte, eventType, uuid string, events ...Event) (interface{}, error) {
	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}

	if hook.uuid != "" && uuid == "" {
		return nil, ErrMissingHookUUIDHeader
	}

	if eventType == "" {
		return nil, ErrMissingEventKeyHeader
	}

	if len(hook.uuid) > 0 && uuid != hook.uuid {
		return nil, ErrUUIDVerificationFailed
	}

	bitbucketEvent := Event(eventType)

	var found bool
	for _, evt := range events {
		if evt == bitbucketEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, ErrEventNotFound
	}

	switch bitbucketEvent {
	case RepoPushEvent:
		var pl RepoPushPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepoForkEvent:
		var pl RepoForkPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepoUpdatedEvent:
		var pl RepoUpdatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepoCommitCommentCreatedEvent:
		var pl RepoCommitCommentCreatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepoCommitStatusCreatedEvent:
		var pl RepoCommitStatusCreatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepoCommitStatusUpdatedEvent:
		var pl RepoCommitStatusUpdatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case IssueCreatedEvent:
		var pl IssueCreatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case IssueUpdatedEvent:
		var pl IssueUpdatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case IssueCommentCreatedEvent:
		var pl IssueCommentCreatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestCreatedEvent:
		var pl PullRequestCreatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestUpdatedEvent:
		var pl PullRequestUpdatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestApprovedEvent:
		var pl PullRequestApprovedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestUnapprovedEvent:
		var pl PullRequestUnapprovedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestMergedEvent:
		var pl PullRequestMergedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestDeclinedEvent:
		var pl PullRequestDeclinedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestCommentCreatedEvent:
		var pl PullRequestCommentCreatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestCommentUpdatedEvent:
		var pl PullRequestCommentUpdatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestCommentDeletedEvent:
		var pl PullRequestCommentDeletedPayload
		return pl, json.Unmarshal(payload, &pl)
	default:
		return nil, fmt.Errorf("unknown event %s", bitbucketEvent)
	}
}
