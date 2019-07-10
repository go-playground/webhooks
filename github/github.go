package github

import (
	"crypto/hmac"
	"crypto/sha1"
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
	ErrEventNotSpecifiedToParse  = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod         = errors.New("invalid HTTP Method")
	ErrMissingGithubEventHeader  = errors.New("missing X-GitHub-Event Header")
	ErrMissingHubSignatureHeader = errors.New("missing X-Hub-Signature Header")
	ErrEventNotFound             = errors.New("event not defined to be parsed")
	ErrParsingPayload            = errors.New("error parsing payload")
	ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

// Event defines a GitHub hook event type
type Event string

// GitHub hook types
const (
	CheckRunEvent                     Event = "check_run"
	CheckSuiteEvent                   Event = "check_suite"
	CommitCommentEvent                Event = "commit_comment"
	CreateEvent                       Event = "create"
	DeleteEvent                       Event = "delete"
	DeploymentEvent                   Event = "deployment"
	DeploymentStatusEvent             Event = "deployment_status"
	ForkEvent                         Event = "fork"
	GollumEvent                       Event = "gollum"
	InstallationEvent                 Event = "installation"
	InstallationRepositoriesEvent     Event = "installation_repositories"
	IntegrationInstallationEvent      Event = "integration_installation"
	IssueCommentEvent                 Event = "issue_comment"
	IssuesEvent                       Event = "issues"
	LabelEvent                        Event = "label"
	MemberEvent                       Event = "member"
	MembershipEvent                   Event = "membership"
	MilestoneEvent                    Event = "milestone"
	OrganizationEvent                 Event = "organization"
	OrgBlockEvent                     Event = "org_block"
	PageBuildEvent                    Event = "page_build"
	PingEvent                         Event = "ping"
	ProjectCardEvent                  Event = "project_card"
	ProjectColumnEvent                Event = "project_column"
	ProjectEvent                      Event = "project"
	PublicEvent                       Event = "public"
	PullRequestEvent                  Event = "pull_request"
	PullRequestReviewEvent            Event = "pull_request_review"
	PullRequestReviewCommentEvent     Event = "pull_request_review_comment"
	PushEvent                         Event = "push"
	ReleaseEvent                      Event = "release"
	RepositoryEvent                   Event = "repository"
	RepositoryVulnerabilityAlertEvent Event = "repository_vulnerability_alert"
	SecurityAdvisoryEvent             Event = "security_advisory"
	StatusEvent                       Event = "status"
	TeamEvent                         Event = "team"
	TeamAddEvent                      Event = "team_add"
	WatchEvent                        Event = "watch"
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
		r.Header.Get("X-GitHub-Event"),
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

	if eventType == "" {
		return nil, ErrMissingGithubEventHeader
	}
	gitHubEvent := Event(eventType)

	var found bool
	for _, evt := range events {
		if evt == gitHubEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, ErrEventNotFound
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		if len(signature) == 0 {
			return nil, ErrMissingHubSignatureHeader
		}
		mac := hmac.New(sha1.New, []byte(hook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
			return nil, ErrHMACVerificationFailed
		}
	}

	switch gitHubEvent {
	case CheckRunEvent:
		var pl CheckRunPayload
		return pl, json.Unmarshal(payload, &pl)
	case CheckSuiteEvent:
		var pl CheckSuitePayload
		return pl, json.Unmarshal(payload, &pl)
	case CommitCommentEvent:
		var pl CommitCommentPayload
		return pl, json.Unmarshal(payload, &pl)
	case CreateEvent:
		var pl CreatePayload
		return pl, json.Unmarshal(payload, &pl)
	case DeleteEvent:
		var pl DeletePayload
		return pl, json.Unmarshal(payload, &pl)
	case DeploymentEvent:
		var pl DeploymentPayload
		return pl, json.Unmarshal(payload, &pl)
	case DeploymentStatusEvent:
		var pl DeploymentStatusPayload
		return pl, json.Unmarshal(payload, &pl)
	case ForkEvent:
		var pl ForkPayload
		return pl, json.Unmarshal(payload, &pl)
	case GollumEvent:
		var pl GollumPayload
		return pl, json.Unmarshal(payload, &pl)
	case InstallationEvent, IntegrationInstallationEvent:
		var pl InstallationPayload
		return pl, json.Unmarshal(payload, &pl)
	case InstallationRepositoriesEvent:
		var pl InstallationRepositoriesPayload
		return pl, json.Unmarshal(payload, &pl)
	case IssueCommentEvent:
		var pl IssueCommentPayload
		return pl, json.Unmarshal(payload, &pl)
	case IssuesEvent:
		var pl IssuesPayload
		return pl, json.Unmarshal(payload, &pl)
	case LabelEvent:
		var pl LabelPayload
		return pl, json.Unmarshal(payload, &pl)
	case MemberEvent:
		var pl MemberPayload
		return pl, json.Unmarshal(payload, &pl)
	case MembershipEvent:
		var pl MembershipPayload
		return pl, json.Unmarshal(payload, &pl)
	case MilestoneEvent:
		var pl MilestonePayload
		return pl, json.Unmarshal(payload, &pl)
	case OrganizationEvent:
		var pl OrganizationPayload
		return pl, json.Unmarshal(payload, &pl)
	case OrgBlockEvent:
		var pl OrgBlockPayload
		return pl, json.Unmarshal(payload, &pl)
	case PageBuildEvent:
		var pl PageBuildPayload
		return pl, json.Unmarshal(payload, &pl)
	case PingEvent:
		var pl PingPayload
		return pl, json.Unmarshal(payload, &pl)
	case ProjectCardEvent:
		var pl ProjectCardPayload
		return pl, json.Unmarshal(payload, &pl)
	case ProjectColumnEvent:
		var pl ProjectColumnPayload
		return pl, json.Unmarshal(payload, &pl)
	case ProjectEvent:
		var pl ProjectPayload
		return pl, json.Unmarshal(payload, &pl)
	case PublicEvent:
		var pl PublicPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestEvent:
		var pl PullRequestPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestReviewEvent:
		var pl PullRequestReviewPayload
		return pl, json.Unmarshal(payload, &pl)
	case PullRequestReviewCommentEvent:
		var pl PullRequestReviewCommentPayload
		return pl, json.Unmarshal(payload, &pl)
	case PushEvent:
		var pl PushPayload
		return pl, json.Unmarshal(payload, &pl)
	case ReleaseEvent:
		var pl ReleasePayload
		return pl, json.Unmarshal(payload, &pl)
	case RepositoryEvent:
		var pl RepositoryPayload
		return pl, json.Unmarshal(payload, &pl)
	case RepositoryVulnerabilityAlertEvent:
		var pl RepositoryVulnerabilityAlertPayload
		return pl, json.Unmarshal(payload, &pl)
	case SecurityAdvisoryEvent:
		var pl SecurityAdvisoryPayload
		return pl, json.Unmarshal(payload, &pl)
	case StatusEvent:
		var pl StatusPayload
		return pl, json.Unmarshal(payload, &pl)
	case TeamEvent:
		var pl TeamPayload
		return pl, json.Unmarshal(payload, &pl)
	case TeamAddEvent:
		var pl TeamAddPayload
		return pl, json.Unmarshal(payload, &pl)
	case WatchEvent:
		var pl WatchPayload
		return pl, json.Unmarshal(payload, &pl)
	default:
		return nil, fmt.Errorf("unknown event %s", gitHubEvent)
	}
}
