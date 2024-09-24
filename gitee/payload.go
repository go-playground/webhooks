package gitee

import (
	"time"
)

type CommentEventPayload struct {
	Action        string          `json:"action"`
	Comment       NoteHook        `json:"comment"`
	Repository    ProjectHook     `json:"repository"`
	Project       ProjectHook     `json:"project"`
	Author        UserHook        `json:"author"`
	Sender        UserHook        `json:"sender"`
	URL           string          `json:"url"`
	Note          string          `json:"note"`
	NoteableType  string          `json:"noteable_type"`
	NoteableID    int64           `json:"noteable_id"`
	Title         string          `json:"title"`
	PerIID        string          `json:"per_iid"`
	ShortCommitID string          `json:"short_commit_id"`
	Enterprise    EnterpriseHook  `json:"enterprise"`
	PullRequest   PullRequestHook `json:"pull_request"`
	Issue         IssueHook       `json:"issue"`
	HookName      string          `json:"hook_name"`
	Password      string          `json:"password"`
}

type PushEventPayload struct {
	Ref                string         `json:"ref"`
	Before             string         `json:"before"`
	After              string         `json:"after"`
	TotalCommitsCount  int64          `json:"total_commits_count"`
	CommitsMoreThanTen bool           `json:"commits_more_than_ten"`
	Created            bool           `json:"created"`
	Deleted            bool           `json:"deleted"`
	Compare            string         `json:"compare"`
	Commits            []CommitHook   `json:"commits"`
	HeadCommit         CommitHook     `json:"head_commit"`
	Repository         ProjectHook    `json:"repository"`
	Project            ProjectHook    `json:"project"`
	UserID             int64          `json:"user_id"`
	UserName           string         `json:"user_name"`
	User               UserHook       `json:"user"`
	Pusher             UserHook       `json:"pusher"`
	Sender             UserHook       `json:"sender"`
	Enterprise         EnterpriseHook `json:"enterprise"`
	HookName           string         `json:"hook_name"`
	Password           string         `json:"password"`
}

type IssueEventPayload struct {
	Action      string         `json:"action"`
	Issue       IssueHook      `json:"issue"`
	Repository  ProjectHook    `json:"repository"`
	Project     ProjectHook    `json:"project"`
	Sender      UserHook       `json:"sender"`
	TargetUser  UserHook       `json:"target_user"`
	User        UserHook       `json:"user"`
	Assignee    UserHook       `json:"assignee"`
	UpdatedBy   UserHook       `json:"updated_by"`
	IID         string         `json:"iid"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	State       string         `json:"state"`
	Milestone   string         `json:"milestone"`
	URL         string         `json:"url"`
	Enterprise  EnterpriseHook `json:"enterprise"`
	HookName    string         `json:"hook_name"`
	Password    string         `json:"password"`
}

type MergeRequestEventPayload struct {
	Action         string          `json:"action"`
	ActionDesc     string          `json:"action_desc"`
	PullRequest    PullRequestHook `json:"pull_request"`
	Number         int64           `json:"number"`
	IID            int64           `json:"iid"`
	Title          string          `json:"title"`
	Body           string          `json:"body"`
	State          string          `json:"state"`
	MergeStatus    string          `json:"merge_status"`
	MergeCommitSha string          `json:"merge_commit_sha"`
	URL            string          `json:"url"`
	SourceBranch   string          `json:"source_branch"`
	SourceRepo     RepoInfo        `json:"source_repo"`
	TargetBranch   string          `json:"target_branch"`
	TargetRepo     RepoInfo        `json:"target_repo"`
	Project        ProjectHook     `json:"project"`
	Repository     ProjectHook     `json:"repository"`
	Author         UserHook        `json:"author"`
	UpdatedBy      UserHook        `json:"updated_by"`
	Sender         UserHook        `json:"sender"`
	TargetUser     UserHook        `json:"target_user"`
	Enterprise     EnterpriseHook  `json:"enterprise"`
	HookName       string          `json:"hook_name"`
	Password       string          `json:"password"`
}

type TagEventPayload struct {
	Action string `json:"action"`
}

// RepoInfo : Repository information
type RepoInfo struct {
	Project    ProjectHook `json:"project"`
	Repository ProjectHook `json:"repository"`
}

// LabelHook : Label, issue and pull request labels
type LabelHook struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// EnterpriseHook : Enterprise information
type EnterpriseHook struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// NoteHook : comment information
type NoteHook struct {
	Id        int64    `json:"id"`
	Body      string   `json:"body"`
	User      UserHook `json:"user"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	HtmlUrl   string   `json:"html_url"`
	Position  string   `json:"position"`
	CommitId  string   `json:"commit_id"`
}

// UserHook : user information
type UserHook struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	UserName  string    `json:"user_name"`
	Url       string    `json:"url"`
	Login     string    `json:"login"`
	AvatarUrl string    `json:"avatar_url"`
	HtmlUrl   string    `json:"html_url"`
	Type_     string    `json:"type"`
	SiteAdmin bool      `json:"site_admin"`
	Time      time.Time `json:"time"`
	Remark    string    `json:"remark"`
}

// CommitHook : git commit information
type CommitHook struct {
	Id        string    `json:"id"`
	TreeId    string    `json:"tree_id"`
	ParentIds []string  `json:"parent_ids"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Url       string    `json:"url"`
	Author    UserHook  `json:"author"`
	Committer UserHook  `json:"committer"`
	Distinct  bool      `json:"distinct"`
	Added     []string  `json:"added"`
	Removed   []string  `json:"removed"`
	Modified  []string  `json:"modified"`
}

// MilestoneHook : milestone information
type MilestoneHook struct {
	Id           int64     `json:"id"`
	HtmlUrl      string    `json:"html_url"`
	Number       int64     `json:"number"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	OpenIssues   int64     `json:"open_issues"`
	ClosedIssues int64     `json:"closed_issues"`
	State        string    `json:"state"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DueOn        string    `json:"due_on"`
}

// IssueHook : issue information
type IssueHook struct {
	Id            int64         `json:"id"`
	HtmlUrl       string        `json:"html_url"`
	Number        string        `json:"number"`
	Title         string        `json:"title"`
	User          UserHook      `json:"user"`
	Labels        []LabelHook   `json:"labels"`
	State         string        `json:"state"`
	StateName     string        `json:"state_name"`
	TypeName      string        `json:"type_name"`
	Assignee      UserHook      `json:"assignee"`
	Collaborators []UserHook    `json:"collaborators"`
	Milestone     MilestoneHook `json:"milestone"`
	Comments      int64         `json:"comments"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Body          string        `json:"body"`
}

// ProjectHook : project information
type ProjectHook struct {
	Id              int64    `json:"id"`
	Name            string   `json:"name"`
	Path            string   `json:"path"`
	FullName        string   `json:"full_name"`
	Owner           UserHook `json:"owner"`
	Private         bool     `json:"private"`
	HtmlUrl         string   `json:"html_url"`
	Url             string   `json:"url"`
	Description     string   `json:"description"`
	Fork            bool     `json:"fork"`
	PushedAt        string   `json:"pushed_at"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	SshUrl          string   `json:"ssh_url"`
	GitUrl          string   `json:"git_url"`
	CloneUrl        string   `json:"clone_url"`
	SvnUrl          string   `json:"svn_url"`
	GitHttpUrl      string   `json:"git_http_url"`
	GitSshUrl       string   `json:"git_ssh_url"`
	GitSvnUrl       string   `json:"git_svn_url"`
	Homepage        string   `json:"homepage"`
	StargazersCount int64    `json:"stargazers_count"`
	WatchersCount   int64    `json:"watchers_count"`
	ForksCount      int64    `json:"forks_count"`
	Language        string   `json:"language"`

	HasIssues bool   `json:"has_issues"`
	HasWiki   bool   `json:"has_wiki"`
	HasPage   bool   `json:"has_pages"`
	License   string `json:"license"`

	OpenIssuesCount int64  `json:"open_issues_count"`
	DefaultBranch   string `json:"default_branch"`
	Namespace       string `json:"namespace"`

	NameWithNamespace string `json:"name_with_namespace"`
	PathWithNamespace string `json:"path_with_namespace"`
}

// BranchHook : branch information
type BranchHook struct {
	Label string       `json:"label"`
	Ref   string       `json:"ref"`
	Sha   string       `json:"sha"`
	User  *UserHook    `json:"user"`
	Repo  *ProjectHook `json:"repo"`
}

// PullRequestHook : PR information
type PullRequestHook struct {
	Id                 int64         `json:"id"`
	Number             int64         `json:"number"`
	State              string        `json:"state"`
	HtmlUrl            string        `json:"html_url"`
	DiffUrl            string        `json:"diff_url"`
	PatchUrl           string        `json:"patch_url"`
	Title              string        `json:"title"`
	Body               string        `json:"body"`
	StaleLabels        []LabelHook   `json:"stale_labels"`
	Labels             []LabelHook   `json:"labels"`
	CreatedAt          string        `json:"created_at"`
	UpdatedAt          string        `json:"updated_at"`
	ClosedAt           string        `json:"closed_at"`
	MergedAt           string        `json:"merged_at"`
	MergeCommitSha     string        `json:"merge_commit_sha"`
	MergeReferenceName string        `json:"merge_reference_name"`
	User               UserHook      `json:"user"`
	Assignee           UserHook      `json:"assignee"`
	Assignees          []UserHook    `json:"assignees"`
	Tester             []UserHook    `json:"tester"`
	Testers            []UserHook    `json:"testers"`
	NeedTest           bool          `json:"need_test"`
	NeedReview         bool          `json:"need_review"`
	Milestone          MilestoneHook `json:"milestone"`
	Head               BranchHook    `json:"head"`
	Base               BranchHook    `json:"base"`
	Merged             bool          `json:"merged"`
	Mergeable          bool          `json:"mergeable"`
	MergeStatus        string        `json:"merge_status"`
	UpdatedBy          UserHook      `json:"updated_by"`
	Comments           int64         `json:"comments"`
	Commits            int64         `json:"commits"`
	Additions          int64         `json:"additions"`
	Deletions          int64         `json:"deletions"`
	ChangedFiles       int64         `json:"changed_files"`
}
