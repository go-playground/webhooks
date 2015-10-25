package webhooks

// GitHubWebhook instance contains all methods needed to process events
type GitHubWebhook struct {
	provider Provider
}

// GitHubConfig defines the configuration to create a new GitHubWebhook instance
type GitHubConfig struct {
	Provider Provider
}

// GitHubHook defines a GitHub hook type
type GitHubHook string

// GitHub hook types
const (
	Any                      GitHubHook = "*"
	CommitComment            GitHubHook = "commit_comment"
	Create                   GitHubHook = "create"
	Delete                   GitHubHook = "delete"
	Deployment               GitHubHook = "deployment"
	DeploymentStatus         GitHubHook = "deployment_status"
	Fork                     GitHubHook = "fork"
	Gollum                   GitHubHook = "gollum"
	IssueComment             GitHubHook = "issue_comment"
	Issues                   GitHubHook = "issues"
	Member                   GitHubHook = "member"
	Membership               GitHubHook = "membership"
	PageBuild                GitHubHook = "page_build"
	Public                   GitHubHook = "public"
	PullRequestReviewComment GitHubHook = "pull_request_review_comment"
	PullRequest              GitHubHook = "pull_request"
	Push                     GitHubHook = "push"
	Repository               GitHubHook = "repository"
	Release                  GitHubHook = "release"
	Status                   GitHubHook = "status"
	TeamAdd                  GitHubHook = "team_add"
	Watch                    GitHubHook = "watch"
)

// GitHubHookSubtype defines a GitHub Hook subtype
type GitHubHookSubtype string

// GitHub hook subtypes
const (
	Branch GitHubHookSubtype = "branch"
	Tag    GitHubHookSubtype = "tag"
	Pull   GitHubHookSubtype = "pull"
	Issue  GitHubHookSubtype = "issues"
)

// Provider returns the GitHubWebhook's provider
func (w GitHubWebhook) Provider() Provider {
	return w.provider
}

// UnderlyingProvider returns the GitHubConfig's Provider
func (c GitHubConfig) UnderlyingProvider() Provider {
	return c.Provider
}
