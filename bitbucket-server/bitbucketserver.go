package bitbucketserver

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	ErrEventNotSpecifiedToParse  = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod         = errors.New("invalid HTTP Method")
	ErrMissingEventKeyHeader     = errors.New("missing X-Event-Key Header")
	ErrMissingHubSignatureHeader = errors.New("missing X-Hub-Signature Header")
	ErrEventNotFound             = errors.New("event not defined to be parsed")
	ErrParsingPayload            = errors.New("error parsing payload")
	ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

type Event string

const (
	RepositoryReferenceChangedEvent Event = "repo:refs_changed"
	RepositoryModifiedEvent         Event = "repo:modified"
	RepositoryForkedEvent           Event = "repo:forked"
	RepositoryCommentAddedEvent     Event = "repo:comment:added"
	RepositoryCommentEditedEvent    Event = "repo:comment:edited"
	RepositoryCommentDeletedEvent   Event = "repo:comment:deleted"

	PullRequestOpenedEvent   Event = "pr:opened"
	PullRequestModifiedEvent Event = "pr:modified"
	PullRequestMergedEvent   Event = "pr:merged"
	PullRequestDeclinedEvent Event = "pr:declined"
	PullRequestDeletedEvent  Event = "pr:deleted"

	PullRequestReviewerUpdatedEvent    Event = "pr:reviewer:updated"
	PullRequestReviewerApprovedEvent   Event = "pr:reviewer:approved"
	PullRequestReviewerUnapprovedEvent Event = "pr:reviewer:unapproved"
	PullRequestReviewerNeedsWorkEvent  Event = "pr:reviewer:needs_work"

	PullRequestCommentAddedEvent   Event = "pr:comment:added"
	PullRequestCommentEditedEvent  Event = "pr:comment:edited"
	PullRequestCommentDeletedEvent Event = "pr:comment:deleted"

	DiagnosticsPingEvent Event = "diagnostics:ping"
)

// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// Secret registers the GitHub secret
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
func (hook *Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
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
		r.Header.Get("X-Hub-Signature"),
		events...,
	)
}

// ParsePayload verifies and parses the events from a payload and string
// metadata (event type and signature), and returns the payload object or an
// error.
//
// Similar to Parse (which uses this method under the hood), this is useful in
// cases where payloads are not represented as HTTP requests - for example are
// put on a queue for pull processing.
func (hook Webhook) ParsePayload(payload []byte, eventType, signature string, events ...Event) (interface{}, error) {
	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}

	event := eventType
	if event == "" {
		return nil, ErrMissingEventKeyHeader
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

	if bitbucketEvent == DiagnosticsPingEvent {
		return DiagnosticsPingPayload{}, nil
	}

	if len(hook.secret) > 0 {
		if len(signature) == 0 {
			return nil, ErrMissingHubSignatureHeader
		}
		mac := hmac.New(sha256.New, []byte(hook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[7:]), []byte(expectedMAC)) {
			return nil, ErrHMACVerificationFailed
		}
	}

	switch bitbucketEvent {
	case RepositoryReferenceChangedEvent:
		var pl RepositoryReferenceChangedPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepositoryModifiedEvent:
		var pl RepositoryModifiedPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepositoryForkedEvent:
		var pl RepositoryForkedPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepositoryCommentAddedEvent:
		var pl RepositoryCommentAddedPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepositoryCommentEditedEvent:
		var pl RepositoryCommentEditedPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepositoryCommentDeletedEvent:
		var pl RepositoryCommentDeletedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestOpenedEvent:
		var pl PullRequestOpenedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestModifiedEvent:
		var pl PullRequestModifiedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestMergedEvent:
		var pl PullRequestMergedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestDeclinedEvent:
		var pl PullRequestDeclinedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestDeletedEvent:
		var pl PullRequestDeletedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestReviewerUpdatedEvent:
		var pl PullRequestReviewerUpdatedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestReviewerApprovedEvent:
		var pl PullRequestReviewerApprovedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestReviewerUnapprovedEvent:
		var pl PullRequestReviewerUnapprovedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestReviewerNeedsWorkEvent:
		var pl PullRequestReviewerNeedsWorkPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestCommentAddedEvent:
		var pl PullRequestCommentAddedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestCommentEditedEvent:
		var pl PullRequestCommentEditedPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestCommentDeletedEvent:
		var pl PullRequestCommentDeletedPayload
		return pl, json.Unmarshal(payload, &pl)
	default:
		return nil, fmt.Errorf("unknown event %s", bitbucketEvent)
	}
}
