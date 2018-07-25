package bitbucket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	webhooks "gopkg.in/go-playground/webhooks.v4"
)

// Webhook instance contains all methods needed to process events
type Webhook struct {
	provider   webhooks.Provider
	uuid       string
	eventFuncs map[Event]webhooks.ProcessPayloadFunc
}

// Config defines the configuration to create a new Bitbucket Webhook instance
type Config struct {
	UUID string
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

// New creates and returns a WebHook instance denoted by the Provider type
func New(config *Config) *Webhook {
	return &Webhook{
		provider:   webhooks.Bitbucket,
		uuid:       config.UUID,
		eventFuncs: map[Event]webhooks.ProcessPayloadFunc{},
	}
}

// Provider returns the current hooks provider ID
func (hook Webhook) Provider() webhooks.Provider {
	return hook.provider
}

// RegisterEvents registers the function to call when the specified event(s) are encountered
func (hook Webhook) RegisterEvents(fn webhooks.ProcessPayloadFunc, events ...Event) {

	for _, event := range events {
		hook.eventFuncs[event] = fn
	}
}

// ParsePayload parses and verifies the payload and fires off the mapped function, if it exists.
func (hook Webhook) ParsePayload(w http.ResponseWriter, r *http.Request) {
	webhooks.DefaultLog.Info("Parsing Payload...")

	uuid := r.Header.Get("X-Hook-UUID")
	if uuid == "" {
		webhooks.DefaultLog.Error("Missing X-Hook-UUID Header")
		http.Error(w, "400 Bad Request - Missing X-Hook-UUID Header", http.StatusBadRequest)
		return
	}
	webhooks.DefaultLog.Debug(fmt.Sprintf("X-Hook-UUID:%s", uuid))

	if len(hook.uuid) > 0 {
		if uuid != hook.uuid {
			// means this is a different hook - ignore
			return
			// webhooks.DefaultLog.Error(fmt.Sprintf("X-Hook-UUID %s does not match configured uuid of %s", uuid, hook.uuid))
			// http.Error(w, "403 Forbidden - X-Hook-UUID does not match", http.StatusForbidden)
			// return
		}
	} else {
		webhooks.DefaultLog.Debug("hook uuid not defined - recommend setting for improved security")
	}

	event := r.Header.Get("X-Event-Key")
	if event == "" {
		webhooks.DefaultLog.Error("Missing X-Event-Key Header")
		http.Error(w, "400 Bad Request - Missing X-Event-Key Header", http.StatusBadRequest)
		return
	}
	webhooks.DefaultLog.Debug(fmt.Sprintf("X-Event-Key:%s", event))

	bitbucketEvent := Event(event)

	fn, ok := hook.eventFuncs[bitbucketEvent]
	// if no event registered
	if !ok {
		webhooks.DefaultLog.Info(fmt.Sprintf("Webhook Event %s not registered, it is recommended to setup only events in bitbucket that will be registered in the webhook to avoid unnecessary traffic and reduce potential attack vectors.", event))
		return
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		webhooks.DefaultLog.Error("Issue reading Payload")
		http.Error(w, "Issue reading Payload", http.StatusInternalServerError)
		return
	}
	webhooks.DefaultLog.Debug(fmt.Sprintf("Payload:%s", string(payload)))
	hd := webhooks.Header(r.Header)

	switch bitbucketEvent {
	case RepoPushEvent:
		var pl RepoPushPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case RepoForkEvent:
		var pl RepoForkPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case RepoUpdatedEvent:
		var pl RepoUpdatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case RepoCommitCommentCreatedEvent:
		var pl RepoCommitCommentCreatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case RepoCommitStatusCreatedEvent:
		var pl RepoCommitStatusCreatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case RepoCommitStatusUpdatedEvent:
		var pl RepoCommitStatusUpdatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case IssueCreatedEvent:
		var pl IssueCreatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case IssueUpdatedEvent:
		var pl IssueUpdatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case IssueCommentCreatedEvent:
		var pl IssueCommentCreatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestCreatedEvent:
		var pl PullRequestCreatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestUpdatedEvent:
		var pl PullRequestUpdatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestApprovedEvent:
		var pl PullRequestApprovedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestUnapprovedEvent:
		var pl PullRequestUnapprovedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestMergedEvent:
		var pl PullRequestMergedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestDeclinedEvent:
		var pl PullRequestDeclinedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestCommentCreatedEvent:
		var pl PullRequestCommentCreatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestCommentUpdatedEvent:
		var pl PullRequestCommentUpdatedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	case PullRequestCommentDeletedEvent:
		var pl PullRequestCommentDeletedPayload
		json.Unmarshal([]byte(payload), &pl)
		hook.runProcessPayloadFunc(fn, pl, hd)
	}
}

func (hook Webhook) runProcessPayloadFunc(fn webhooks.ProcessPayloadFunc, results interface{}, header webhooks.Header) {
	go func(fn webhooks.ProcessPayloadFunc, results interface{}, header webhooks.Header) {
		fn(results, header)
	}(fn, results, header)
}
