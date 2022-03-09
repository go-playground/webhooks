package gitea

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

// parse errors
var (
	ErrEventNotSpecifiedToParse    = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod           = errors.New("invalid HTTP Method")
	ErrMissingGiteaEventHeader     = errors.New("missing X-Gitea-Event Header")
	ErrMissingGiteaSignatureHeader = errors.New("missing X-Gitea-Signature Header")
	ErrEventNotFound               = errors.New("event not defined to be parsed")
	ErrParsingPayload              = errors.New("error parsing payload")
	ErrHMACVerificationFailed      = errors.New("HMAC verification failed")
)

// Gitea hook types
// https://github.com/go-gitea/gitea/blob/bf7b083cfe47cc922090ce7922b89f7a5030a44d/models/webhook/hooktask.go#L31
const (
	CreateEvent               Event = "create"
	DeleteEvent               Event = "delete"
	ForkEvent                 Event = "fork"
	IssuesEvent               Event = "issues"
	IssueAssignEvent          Event = "issue_assign"
	IssueLabelEvent           Event = "issue_label"
	IssueMilestoneEvent       Event = "issue_milestone"
	IssueCommentEvent         Event = "issue_comment"
	PushEvent                 Event = "push"
	PullRequestEvent          Event = "pull_request"
	PullRequestAssignEvent    Event = "pull_request_assign"
	PullRequestLabelEvent     Event = "pull_request_label"
	PullRequestMilestoneEvent Event = "pull_request_milestone"
	PullRequestCommentEvent   Event = "pull_request_comment"
	PullRequestReviewEvent    Event = "pull_request_review"
	PullRequestSyncEvent      Event = "pull_request_sync"
	RepositoryEvent           Event = "repository"
	ReleaseEvent              Event = "release"
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

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Gitea-Signature")
		if len(signature) == 0 {
			return nil, ErrMissingGiteaSignatureHeader
		}
		sig256 := hmac.New(sha256.New, []byte(hook.secret))
		_, _ = io.Writer(sig256).Write([]byte(payload))
		expectedMAC := hex.EncodeToString(sig256.Sum(nil))

		if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
			return nil, ErrHMACVerificationFailed
		}
	}

	// https://github.com/go-gitea/gitea/blob/33fca2b537d36cf998dd27425b2bb8ed5b0965f3/services/webhook/payloader.go#L27
	switch giteaEvent {
	case CreateEvent:
		var pl CreatePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case DeleteEvent:
		var pl DeletePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ForkEvent:
		var pl ForkPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PushEvent:
		var pl PushPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssuesEvent, IssueAssignEvent, IssueLabelEvent, IssueMilestoneEvent:
		var pl IssuePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueCommentEvent, PullRequestCommentEvent:
		var pl IssueCommentPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestEvent, PullRequestAssignEvent, PullRequestLabelEvent, PullRequestMilestoneEvent, PullRequestReviewEvent, PullRequestSyncEvent:
		var pl PullRequestPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryEvent:
		var pl RepositoryPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ReleaseEvent:
		var pl ReleasePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event %s", giteaEvent)
	}
}
