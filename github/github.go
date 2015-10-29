package github

import "github.com/joeybloggs/webhooks"

// Webhook instance contains all methods needed to process events
type Webhook struct {
	provider webhooks.Provider
}

// Config defines the configuration to create a new GitHubWebhook instance
type Config struct {
	Provider webhooks.Provider
}

// Event defines a GitHub hook event type
type Event string

// GitHub hook types
const (
	// AnyEvent                      Event = "*"
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

// Provider returns the Webhook's provider
func (w Webhook) Provider() webhooks.Provider {
	return w.provider
}

// UnderlyingProvider returns the Config's Provider
func (c Config) UnderlyingProvider() webhooks.Provider {
	return c.Provider
}

// New creates and returns a WebHook instance denoted by the Provider type
func New(config *Config) *Webhook {
	return &Webhook{
		provider: config.Provider,
	}
}
