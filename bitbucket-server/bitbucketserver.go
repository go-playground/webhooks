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

	PullRequestOpenedEvent               Event = "pr:opened"
	PullRequestFromReferenceUpdatedEvent Event = "pr:from_ref_updated"
	PullRequestModifiedEvent             Event = "pr:modified"
	PullRequestMergedEvent               Event = "pr:merged"
	PullRequestDeclinedEvent             Event = "pr:declined"
	PullRequestDeletedEvent              Event = "pr:deleted"

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

func (hook *Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
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

	event := r.Header.Get("X-Event-Key")
	if event == "" {
		return nil, ErrMissingEventKeyHeader
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

	if bitbucketEvent == DiagnosticsPingEvent {
		return DiagnosticsPingPayload{}, nil
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Hub-Signature")
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
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryModifiedEvent:
		var pl RepositoryModifiedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryForkedEvent:
		var pl RepositoryForkedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryCommentAddedEvent:
		var pl RepositoryCommentAddedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryCommentEditedEvent:
		var pl RepositoryCommentEditedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryCommentDeletedEvent:
		var pl RepositoryCommentDeletedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestOpenedEvent:
		var pl PullRequestOpenedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestFromReferenceUpdatedEvent:
		var pl PullRequestFromReferenceUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestModifiedEvent:
		var pl PullRequestModifiedPayload
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
	case PullRequestDeletedEvent:
		var pl PullRequestDeletedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewerUpdatedEvent:
		var pl PullRequestReviewerUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewerApprovedEvent:
		var pl PullRequestReviewerApprovedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewerUnapprovedEvent:
		var pl PullRequestReviewerUnapprovedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewerNeedsWorkEvent:
		var pl PullRequestReviewerNeedsWorkPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentAddedEvent:
		var pl PullRequestCommentAddedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentEditedEvent:
		var pl PullRequestCommentEditedPayload
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
