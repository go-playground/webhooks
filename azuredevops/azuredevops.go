package azuredevops

// this package receives Azure DevOps Server webhooks
// https://docs.microsoft.com/en-us/azure/devops/service-hooks/services/webhooks?view=azure-devops-2020

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
	ErrInvalidHTTPMethod = errors.New("invalid HTTP Method")
	ErrParsingPayload    = errors.New("error parsing payload")
)

// Event defines an Azure DevOps server hook event type
type Event string

// Azure DevOps Server hook types
const (
	BuildCompleteEventType         Event = "build.complete"
	GitPullRequestCreatedEventType Event = "git.pullrequest.created"
	GitPullRequestUpdatedEventType Event = "git.pullrequest.updated"
	GitPullRequestMergedEventType  Event = "git.pullrequest.merged"
)

// Webhook instance contains all methods needed to process events
type Webhook struct {
}

// New creates and returns a WebHook instance
func New() (*Webhook, error) {
	hook := new(Webhook)
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

	var pl BasicEvent
	err = json.Unmarshal([]byte(payload), &pl)
	if err != nil {
		return nil, ErrParsingPayload
	}

	switch pl.EventType {
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
