package github

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/go-playground/webhooks.v3"
)

// Webhook instance contains all methods needed to process events
type Webhook struct {
	provider   webhooks.Provider
	secret     string
	eventFuncs map[Event]webhooks.ProcessPayloadFunc
}

// Config defines the configuration to create a new GitHub Webhook instance
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
	LabelEvent                    Event = "label"
	MemberEvent                   Event = "member"
	MembershipEvent               Event = "membership"
	MilestoneEvent                Event = "milestone"
	OrganizationEvent             Event = "organization"
	OrgBlockEvent                 Event = "org_block"
	PageBuildEvent                Event = "page_build"
	ProjectCardEvent              Event = "project_card"
	ProjectColumnEvent            Event = "project_column"
	ProjectEvent                  Event = "project"
	PublicEvent                   Event = "public"
	PullRequestEvent              Event = "pull_request"
	PullRequestReviewEvent        Event = "pull_request_review"
	PullRequestReviewCommentEvent Event = "pull_request_review_comment"
	PushEvent                     Event = "push"
	ReleaseEvent                  Event = "release"
	RepositoryEvent               Event = "repository"
	StatusEvent                   Event = "status"
	TeamEvent                     Event = "team"
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

	log.Println("Gettting X-GitHub-Event")
	event := r.Header.Get("X-GitHub-Event")
	if len(event) == 0 {
		http.Error(w, "400 Bad Request - Missing X-GitHub-Event Header", http.StatusBadRequest)
		return
	}

	gitHubEvent := Event(event)

	log.Println("Looking for Hook:", gitHubEvent)
	fn, ok := hook.eventFuncs[gitHubEvent]
	// if no event registered
	if !ok {
		return
	}

	log.Println("READING PAYLOAD FROM BODY")

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		http.Error(w, "Error reading Body", http.StatusInternalServerError)
		return
	}

	log.Println("Checking GitHub secret")

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {

		log.Println("Get GitHub signature")
		signature := r.Header.Get("X-Hub-Signature")

		if len(signature) == 0 {
			http.Error(w, "403 Forbidden - Missing X-Hub-Signature required for HMAC verification", http.StatusForbidden)
			return
		}

		mac := hmac.New(sha1.New, []byte(hook.secret))
		mac.Write(payload)

		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		log.Println("Checking HMAC Equality")

		if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
			http.Error(w, "403 Forbidden - HMAC verification failed", http.StatusForbidden)
			return
		}

		log.Println("HMAC Equal")
	}

	// Make headers available to ProcessPayloadFunc as a webhooks type
	hd := webhooks.Header(r.Header)
	var pl interface{}

	log.Println("Unmarshal based on GitHub event:", gitHubEvent)

	switch gitHubEvent {
	case CommitCommentEvent:

		var cc CommitCommentPayload

		err = json.Unmarshal([]byte(payload), &cc)
		pl = cc

	case CreateEvent:

		var c CreatePayload

		err = json.Unmarshal([]byte(payload), &c)
		pl = c

	case DeleteEvent:
		var d DeletePayload
		json.Unmarshal([]byte(payload), &d)
		hook.runProcessPayloadFunc(fn, d, hd)

	case DeploymentEvent:

		var d DeploymentPayload

		err = json.Unmarshal([]byte(payload), &d)
		pl = d

	case DeploymentStatusEvent:

		var d DeploymentStatusPayload

		err = json.Unmarshal([]byte(payload), &d)
		pl = d

	case ForkEvent:

		var f ForkPayload

		err = json.Unmarshal([]byte(payload), &f)
		pl = f

	case GollumEvent:

		var g GollumPayload

		err = json.Unmarshal([]byte(payload), &g)
		pl = g

	case IssueCommentEvent:

		var i IssueCommentPayload

		err = json.Unmarshal([]byte(payload), &i)
		pl = i

	case IssuesEvent:

		var i IssuesPayload

		err = json.Unmarshal([]byte(payload), &i)
		pl = i

	case LabelEvent:

		var l LabelPayload

		err = json.Unmarshal([]byte(payload), &l)
		pl = l

	case MemberEvent:

		var m MemberPayload

		err = json.Unmarshal([]byte(payload), &m)
		pl = m

	case MembershipEvent:

		var m MembershipPayload

		err = json.Unmarshal([]byte(payload), &m)
		pl = m

	case MilestoneEvent:

		var m MilestonePayload

		err = json.Unmarshal([]byte(payload), &m)
		pl = m

	case OrganizationEvent:

		var o OrganizationPayload

		err = json.Unmarshal([]byte(payload), &o)
		pl = o

	case OrgBlockEvent:

		var o OrgBlockPayload

		err = json.Unmarshal([]byte(payload), &o)
		pl = o

	case PageBuildEvent:

		var p PageBuildPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case ProjectCardEvent:

		var p ProjectCardPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case ProjectColumnEvent:

		var p ProjectColumnPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case ProjectEvent:

		var p ProjectPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case PublicEvent:

		var p PublicPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case PullRequestEvent:

		var p PullRequestPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case PullRequestReviewEvent:

		var p PullRequestReviewPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case PullRequestReviewCommentEvent:

		var p PullRequestReviewCommentPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case PushEvent:

		var p PushPayload

		err = json.Unmarshal([]byte(payload), &p)
		pl = p

	case ReleaseEvent:

		var r ReleasePayload

		err = json.Unmarshal([]byte(payload), &r)
		pl = r

	case RepositoryEvent:

		var r RepositoryPayload

		err = json.Unmarshal([]byte(payload), &r)
		pl = r

	case StatusEvent:

		var s StatusPayload

		err = json.Unmarshal([]byte(payload), &s)
		pl = s

	case TeamEvent:

		var t TeamPayload

		err = json.Unmarshal([]byte(payload), &t)
		pl = t

	case TeamAddEvent:

		var t TeamAddPayload

		err = json.Unmarshal([]byte(payload), &t)
		pl = t

	case WatchEvent:

		var w WatchPayload

		err = json.Unmarshal([]byte(payload), &w)
		pl = w
	}

	if err != nil {
		log.Println("There was an erro parsing JSON:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Println("Running runProcessPayloadFunc")
	hook.runProcessPayloadFunc(fn, pl, hd)
}

func (hook Webhook) runProcessPayloadFunc(fn webhooks.ProcessPayloadFunc, results interface{}, header webhooks.Header) {
	go func(fn webhooks.ProcessPayloadFunc, results interface{}, header webhooks.Header) {
		log.Println("Calling hook function")
		fn(results, header)
	}(fn, results, header)
}
