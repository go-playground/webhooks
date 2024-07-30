package azuredevops

// this package receives Azure DevOps Server webhooks
// https://docs.microsoft.com/en-us/azure/devops/service-hooks/services/webhooks?view=azure-devops-2020

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// parse errors
var (
	ErrInvalidHTTPMethod           = errors.New("invalid HTTP Method")
	ErrParsingPayload              = errors.New("error parsing payload")
	ErrBasicAuthVerificationFailed = errors.New("basic auth verification failed")
)

// Event defines an Azure DevOps server hook event type
type Event string

// Azure DevOps Server hook types
const (
	BuildCompleteEventType         Event = "build.complete"
	GitPullRequestCreatedEventType Event = "git.pullrequest.created"
	GitPullRequestUpdatedEventType Event = "git.pullrequest.updated"
	GitPullRequestMergedEventType  Event = "git.pullrequest.merged"
	GitPushEventType               Event = "git.push"
)

// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// BasicAuth verifies payload using basic auth
func (WebhookOptions) BasicAuth(username, password string) Option {
	return func(hook *Webhook) error {
		hook.username = username
		hook.password = password
		return nil
	}
}

// Webhook instance contains all methods needed to process events
type Webhook struct {
	username string
	password string
}

// New creates and returns a WebHook instance
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
		_, _ = io.Copy(io.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if !hook.verifyBasicAuth(r) {
		return nil, ErrBasicAuthVerificationFailed
	}

	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	var pl BasicEvent
	err = json.Unmarshal([]byte(payload), &pl)
	if err != nil {
		return nil, ErrParsingPayload
	}

	switch pl.EventType {
	case GitPushEventType:
		var fpl GitPushEvent
		err = json.Unmarshal([]byte(payload), &fpl)
		return fpl, err
	case GitPullRequestCreatedEventType, GitPullRequestMergedEventType, GitPullRequestUpdatedEventType:
		var fpl GitPullRequestEvent
		err = json.Unmarshal([]byte(payload), &fpl)
		return fpl, err
	case BuildCompleteEventType:
		var fpl BuildCompleteEvent
		err = json.Unmarshal([]byte(payload), &fpl)
		return fpl, err
	default:
		return nil, fmt.Errorf("unknown event %s", pl.EventType)
	}
}

func (hook Webhook) verifyBasicAuth(r *http.Request) bool {
	// skip validation if username or password was not provided
	if hook.username == "" && hook.password == "" {
		return true
	}
	username, password, ok := r.BasicAuth()

	return ok && username == hook.username && password == hook.password
}
