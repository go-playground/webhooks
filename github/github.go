package github

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/joeybloggs/webhooks"
)

// Webhook instance contains all methods needed to process events
type Webhook struct {
	provider   webhooks.Provider
	secret     string
	eventFuncs map[Event]webhooks.ProcessPayloadFunc
}

// Config defines the configuration to create a new GitHubWebhook instance
type Config struct {
	Secret string
}

// Event defines a GitHub hook event type
type Event string

// GitHub hook types
const (
	CommitCommentEvent            Event = "commit_comment"
	CreateEvent                   Event = "create"
	DeleteEvent                   Event = "delete"
	DeploymentEvent               Event = "deployment"
	DeploymentStatusEvent         Event = "deployment_status"
	ForkEvent                     Event = "fork"
	GollumEvent                   Event = "gollum"
	IssueCommentEvent             Event = "issue_comment"
	IssuesEvent                   Event = "issues"
	MemberEvent                   Event = "member"
	MembershipEvent               Event = "membership"
	PageBuildEvent                Event = "page_build"
	PublicEvent                   Event = "public"
	PullRequestReviewCommentEvent Event = "pull_request_review_comment"
	PullRequestEvent              Event = "pull_request"
	PushEvent                     Event = "push"
	RepositoryEvent               Event = "repository"
	ReleaseEvent                  Event = "release"
	StatusEvent                   Event = "status"
	TeamAddEvent                  Event = "team_add"
	WatchEvent                    Event = "watch"
)

// EventSubtype defines a GitHub Hook Event subtype
type EventSubtype string

// GitHub hook event subtypes
const (
	NoSubtype     EventSubtype = ""
	BranchSubtype EventSubtype = "branch"
	TagSubtype    EventSubtype = "tag"
	PullSubtype   EventSubtype = "pull"
	IssueSubtype  EventSubtype = "issues"
)

// New creates and returns a WebHook instance denoted by the Provider type
func New(config *Config) *Webhook {
	return &Webhook{
		provider:   webhooks.GitHub,
		secret:     config.Secret,
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

	event := r.Header.Get("X-GitHub-Event")
	if len(event) == 0 {
		http.Error(w, "400 Bad Request - Missing X-GitHub-Event Header", http.StatusBadRequest)
		return
	}

	gitHubEvent := Event(event)

	fn, ok := hook.eventFuncs[gitHubEvent]
	// if no event registered
	if !ok {
		return
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {

		signature := r.Header.Get("X-Hub-Signature")

		if len(signature) == 0 {
			http.Error(w, "403 Forbidden - Missing X-Hub-Signature required for HMAC verification", http.StatusForbidden)
			return
		}

		mac := hmac.New(sha1.New, []byte(hook.secret))
		_, err := mac.Write(payload)
		if err != nil {
			http.Error(w, "400 Bad Request - HMAC verification failed with body parsing", http.StatusBadRequest)
			return
		}

		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
			http.Error(w, "403 Forbidden - HMAC verification failed", http.StatusForbidden)
			return
		}
	}

	var results interface{}

	switch gitHubEvent {
	case ReleaseEvent:
		var release ReleasePayload
		json.Unmarshal([]byte(payload), &release)
		results = release
	}

	go func(fn webhooks.ProcessPayloadFunc, results interface{}) {

		// put in recovery here!

		fn(results)
	}(fn, results)
}
