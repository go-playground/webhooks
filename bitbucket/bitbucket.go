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

	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}
	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	uuid := r.Header.Get("X-Hook-UUID")
	if hook.uuid != "" && uuid == "" {
		return nil, ErrMissingHookUUIDHeader
	}

	event := r.Header.Get("X-Event-Key")
	if event == "" {
		return nil, ErrMissingEventKeyHeader
	}

	if len(hook.uuid) > 0 && uuid != hook.uuid {
		return nil, ErrUUIDVerificationFailed
	}

	bitbucketEvent := Event(event)

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

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	switch bitbucketEvent {
	case RepoPushEvent:
		var pl RepoPushPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoForkEvent:
		var pl RepoForkPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoUpdatedEvent:
		var pl RepoUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoCommitCommentCreatedEvent:
		var pl RepoCommitCommentCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoCommitStatusCreatedEvent:
		var pl RepoCommitStatusCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoCommitStatusUpdatedEvent:
		var pl RepoCommitStatusUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueCreatedEvent:
		var pl IssueCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueUpdatedEvent:
		var pl IssueUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueCommentCreatedEvent:
		var pl IssueCommentCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCreatedEvent:
		var pl PullRequestCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestUpdatedEvent:
		var pl PullRequestUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestApprovedEvent:
		var pl PullRequestApprovedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestUnapprovedEvent:
		var pl PullRequestUnapprovedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestMergedEvent:
		var pl PullRequestMergedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestDeclinedEvent:
		var pl PullRequestDeclinedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentCreatedEvent:
		var pl PullRequestCommentCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentUpdatedEvent:
		var pl PullRequestCommentUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentDeletedEvent:
		var pl PullRequestCommentDeletedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event %s", bitbucketEvent)
	}
}
