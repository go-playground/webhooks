package github

import "time"

// WatchPayload contains the information for GitHub's watch hook event
type WatchPayload struct {
	Action     string     `json:"action"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// TeamAddPayload contains the information for GitHub's team_add hook event
type TeamAddPayload struct {
	Team         Team         `json:"team"`
	Repository   Repository   `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       Sender       `json:"sender"`
}

// StatusPayload contains the information for GitHub's status hook event
type StatusPayload struct {
	ID          int          `json:"id"`
	SHA         string       `json:"sha"`
	Name        string       `json:"name"`
	TragetURL   string       `json:"target_url"`
	Context     string       `json:"context"`
	Desctiption string       `json:"description"`
	State       string       `json:"state"`
	Commit      StatusCommit `json:"commit"`
	Branches    []Branch     `json:"branches"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Repository  Repository   `json:"repository"`
	Sender      Sender       `json:"sender"`
}

// ReleasePayload contains the information for GitHub's release hook event
type ReleasePayload struct {
	Action     string     `json:"action"`
	Release    Release    `json:"release"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// RepositoryPayload contains the information for GitHub's repository hook event
type RepositoryPayload struct {
	Action       string       `json:"action"`
	Repository   Repository   `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       Sender       `json:"sender"`
}

// PushPayload contains the information for GitHub's push hook event
type PushPayload struct {
	Ref        string     `json:"ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Created    bool       `json:"created"`
	Deleted    bool       `json:"deleted"`
	Forced     bool       `json:"forced"`
	BaseRef    string     `json:"base_ref"`
	Compare    string     `json:"compare"`
	Commits    []Commit   `json:"commits"`
	HeadCommit HeadCommit `json:"head_commit"`
	Repository Repository `json:"repository"`
	Pusher     PusherPush `json:"pusher"`
	Sender     Sender     `json:"sender"`
}

// PullRequestPayload contains the information for GitHub's pull_request hook event
type PullRequestPayload struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
	Sender      Sender      `json:"sender"`
}

// PullRequestReviewCommentPayload contains the information for GitHub's pull_request_review_comment hook event
type PullRequestReviewCommentPayload struct {
	Action     string             `json:"action"`
	Comment    PullRequestComment `json:"comment"`
	Repository Repository         `json:"repository"`
	Sender     Sender             `json:"sender"`
}

// PublicPayload contains the information for GitHub's public hook event
type PublicPayload struct {
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// PageBuildPayload contains the information for GitHub's page_build hook event
type PageBuildPayload struct {
	ID         int        `json:"id"`
	Build      Build      `json:"build"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// MembershipPayload contains the information for GitHub's membership hook event
type MembershipPayload struct {
	Action       string       `json:"action"`
	Scope        string       `json:"scope"`
	Member       Member       `json:"member"`
	Sender       Sender       `json:"sender"`
	Team         Team         `json:"team"`
	Organization Organization `json:"organization"`
}

// MemberPayload contains the information for GitHub's member hook event
type MemberPayload struct {
	Action     string     `json:"action"`
	Member     Member     `json:"member"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// IssuesPayload contains the information for GitHub's issues hook event
type IssuesPayload struct {
	Action     string     `json:"action"`
	Issue      Issue      `json:"issue"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// IssueCommentPayload contains the information for GitHub's issue_comment hook event
type IssueCommentPayload struct {
	IssuesPayload
	Comment Comment `json:"comment"`
}

// GollumPayload contains the information for GitHub's gollum hook event
type GollumPayload struct {
	Pages      []Page     `json:"pages"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// ForkPayload contains the information for GitHub's fork hook event
type ForkPayload struct {
	Forkee     Forkee     `json:"forkee"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// DeploymentStatusPayload contains the information for GitHub's deployment_status hook event
type DeploymentStatusPayload struct {
	Deployment       Deployment       `json:"deployment"`
	DeploymentStatus DeploymentStatus `json:"deployment_status"`
	Repository       Repository       `json:"repository"`
	Sender           Sender           `json:"sender"`
}

// DeploymentPayload contains the information for GitHub's deployment hook
type DeploymentPayload struct {
	Deployment Deployment `json:"deployment"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// CommitCommentPayload contains the information for GitHub's commit_comment hook event
type CommitCommentPayload struct {
	Action     string     `json:"action"`
	RefType    string     `json:"ref_type"`
	Comment    Comment    `json:"comment"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// DeletePayload contains the information for GitHub's delete hook event
type DeletePayload struct {
	Ref        string     `json:"ref"`
	RefType    string     `json:"ref_type"`
	PusherType string     `json:"pusher_type"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// CreatePayload contains the information for GitHub's create hook event
type CreatePayload struct {
	Ref          string     `json:"ref"`
	RefType      string     `json:"ref_type"`
	MasterBranch string     `json:"master_branch"`
	Description  string     `json:"description"`
	PusherType   string     `json:"pusher_type"`
	Repository   Repository `json:"repository"`
	Sender       Sender     `json:"sender"`
}

// Repository contais all of the GitHub repository information
type Repository struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	FullName         string    `json:"full_name"`
	Owner            Owner     `json:"owner"`
	Private          bool      `json:"private"`
	HTMLURL          string    `json:"html_url"`
	Description      string    `json:"description"`
	Fork             bool      `json:"fork"`
	URL              string    `json:"url"`
	ForksURL         string    `json:"forks_url"`
	KeysURL          string    `json:"keys_url"`
	CollaboratorsURL string    `json:"collaborators_url"`
	TeamsURL         string    `json:"teams_url"`
	HooksURL         string    `json:"hooks_url"`
	IssueEventsURL   string    `json:"issue_events_url"`
	EventsURL        string    `json:"events_url"`
	AssigneesURL     string    `json:"assignees_url"`
	BranchesURL      string    `json:"branches_url"`
	TagsURL          string    `json:"tags_url"`
	BlobsURL         string    `json:"blobs_url"`
	GitTagsURL       string    `json:"git_tags_url"`
	GitRefsURL       string    `json:"git_refs_url"`
	TreesURL         string    `json:"trees_url"`
	StatusesURL      string    `json:"statuses_url"`
	LanguagesURL     string    `json:"languages_url"`
	StargazersURL    string    `json:"stargazers_url"`
	ContributorsURL  string    `json:"contributors_url"`
	SubscribersURL   string    `json:"subscribers_url"`
	SubscriptionURL  string    `json:"subscription_url"`
	CommitsURL       string    `json:"commits_url"`
	GitCommitsURL    string    `json:"git_commits_url"`
	CommentsURL      string    `json:"comments_url"`
	IssueCommentURL  string    `json:"issue_comment_url"`
	ContentsURL      string    `json:"contents_url"`
	CompareURL       string    `json:"compare_url"`
	MergesURL        string    `json:"merges_url"`
	ArchiveURL       string    `json:"archive_url"`
	DownloadsURL     string    `json:"downloads_url"`
	IssuesURL        string    `json:"issues_url"`
	PullsURL         string    `json:"pulls_url"`
	MilestonesURL    string    `json:"milestones_url"`
	NotificationsURL string    `json:"notifications_url"`
	LabelsURL        string    `json:"labels_url"`
	ReleasesURL      string    `json:"releases_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	PushedAt         time.Time `json:"pushed_at"`
	GitURL           string    `json:"git_url"`
	SSHURL           string    `json:"ssh_url"`
	CloneURL         string    `json:"clone_url"`
	SVNURL           string    `json:"svn_url"`
	Homepage         string    `json:"homepage"`
	Size             int       `json:"size"`
	StargazersCount  int       `json:"stargazers_count"`
	WatchersCount    int       `json:"watchers_count"`
	Language         string    `json:"language"`
	HasIssues        bool      `json:"has_issues"`
	HasDownloads     bool      `json:"has_downloads"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	ForksCount       int       `json:"forks_count"`
	MirrorURL        string    `json:"mirror_url"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	Forks            int       `json:"forks"`
	OpenIssues       int       `json:"open_issues"`
	Watchers         int       `json:"watchers"`
	DefaultBranch    string    `json:"default_branch"`
}

// // Repo contains GitHub's repo information
// type Repo struct {
// 	Repository
// }

// User contains GitHub's user information
type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Member contains GitHub's member information
type Member struct {
	User
}

// Owner contains GitHub's owner information
type Owner struct {
	User
}

// Sender contains GitHub's sender information
type Sender struct {
	User
}

// Creator contains GitHub's creator information
type Creator struct {
	User
}

// Pusher contains GitHub's pusher information
type Pusher struct {
	User
}

// Author contains GitHub's author information
type Author struct {
	User
}

// Commiter contains GitHub's commiter information
type Commiter struct {
	User
}

// Comment contains GitHub's comment information
type Comment struct {
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	ID        int       `json:"id"`
	User      User      `json:"user"`
	Position  int       `json:"position"`
	Line      int       `json:"line"`
	Path      string    `json:"path"`
	CommitID  string    `json:"commit_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
}

// Deployment contains GitHub's deployment information
type Deployment struct {
	URL  string `json:"url"`
	ID   int    `json:"id"`
	SHA  string `json:"sha"`
	Ref  string `json:"ref"`
	Task string `json:"task"`
	//paylod
	Environment   string    `json:"environment"`
	Description   string    `json:"description"`
	Creator       Creator   `json:"creator"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	StatusesURL   string    `json:"statuses_url"`
	RepositoryURL string    `json:"repository_url"`
}

// DeploymentStatus contains GitHub's deployment_status information
type DeploymentStatus struct {
	URL           string    `json:"url"`
	ID            int       `json:"id"`
	State         string    `json:"state"`
	Creator       Creator   `json:"creator"`
	Description   string    `json:"description"`
	TargetURL     string    `json:"target_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeploymentURL string    `json:"deployment_url"`
	RepositoryURL string    `json:"repository_url"`
}

// Forkee contains GitHub's forkee information
type Forkee struct {
	Repository
	Public bool `json:"public"`
}

// Page contains GitHub's page information
type Page struct {
	PageName string `json:"page_name"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Action   string `json:"action"`
	SHA      string `json:"sha"`
	HTMLURL  string `json:"html_url"`
}

// Label contains GitHub's label information
type Label struct {
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// Issue contains GitHub's issue information
type Issue struct {
	URL         string    `json:"url"`
	LabelsURL   string    `json:"labels_url"`
	CommentsURL string    `json:"comments_url"`
	EventsURL   string    `json:"events_url"`
	HTMLURL     string    `json:"html_url"`
	ID          int       `json:"id"`
	Number      int       `json:"number"`
	Title       string    `json:"title"`
	User        User      `json:"user"`
	Labels      []Label   `json:"labels"`
	State       string    `json:"state"`
	Locked      bool      `json:"locked"`
	Assignee    string    `json:"assignee"`
	Milestone   string    `json:"milestone"`
	Comments    int       `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ClosedAt    time.Time `json:"closed_at"`
	Body        string    `json:"body"`
}

// Team contains GitHub's team information
type Team struct {
	Name            string `json:"name"`
	ID              int    `json:"id"`
	Slug            string `json:"slug"`
	Permission      string `json:"permission"`
	URL             string `json:"url"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
}

// Organization contains GitHub's organization information
type Organization struct {
	Login            string `json:"login"`
	ID               int    `json:"id"`
	URL              string `json:"url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	MembersURL       string `json:"members_url"`
	PublicMembersURL string `json:"public_members_url"`
	AvatarURL        string `json:"avatar_url"`
}

// Error contains GitHub's error information
type Error struct {
	Message string `json:"message"`
}

// Build contains GitHub's build information
type Build struct {
	URL       string    `json:"url"`
	Status    string    `json:"status"`
	Error     Error     `json:"error"`
	Pusher    Pusher    `json:"pusher"`
	Commit    string    `json:"commit"`
	Duration  int       `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PullRequestHREF contains GitHub's pull_request href information
type PullRequestHREF struct {
	HREF string `json:"href"`
}

// HTML contains GitHub's html information
type HTML struct {
	HREF string `json:"href"`
}

// Self contains GitHub's self information
type Self struct {
	HREF string `json:"href"`
}

// Links contains GitHub's link information
type Links struct {
	Self        Self            `json:"self"`
	HTML        HTML            `json:"html"`
	PullRequest PullRequestHREF `json:"pull_request"`
}

// PullRequestComment contains GitHub's pull request comment information
type PullRequestComment struct {
	URL              string    `json:"url"`
	ID               int       `json:"id"`
	DiffHunk         string    `json:"diff_hunk"`
	Path             string    `json:"path"`
	Position         int       `json:"position"`
	OriginalPosition int       `json:"original_position"`
	CommitID         string    `json:"commit_id"`
	OriginalCommitID string    `json:"original_commit_id"`
	User             User      `json:"user"`
	Body             string    `json:"body"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	HTMLURL          string    `json:"html_url"`
	PullRequestURL   string    `json:"pull_request_url"`
	Links            Links     `json:"links"`
}

// Head contains GitHub's head information
type Head struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	SHA   string     `json:"sha"`
	User  User       `json:"user"`
	Repo  Repository `json:"repo"`
}

// Base contains GitHub's base information
type Base struct {
	Head
}

// IssueHREF contains GitHub's issue href information
type IssueHREF struct {
	HREF string `json:"href"`
}

// CommentsHREF contains GitHub's comments href information
type CommentsHREF struct {
	HREF string `json:"href"`
}

// ReviewCommentsHREF contains GitHub's review comments href information
type ReviewCommentsHREF struct {
	HREF string `json:"href"`
}

// ReviewCommentHREF contains GitHub's review comment href information
type ReviewCommentHREF struct {
	HREF string `json:"href"`
}

// CommitsHREF contains GitHub's commits href information
type CommitsHREF struct {
	HREF string `json:"href"`
}

// StatusesHREF contains GitHub's statuses href information
type StatusesHREF struct {
	HREF string `json:"href"`
}

// LinksPullRequest contains GitHub's pull request link information
type LinksPullRequest struct {
	Self           Self               `json:"self"`
	HTML           HTML               `json:"html"`
	Issue          IssueHREF          `json:"issue"`
	Comments       CommentsHREF       `json:"comments"`
	ReviewComments ReviewCommentsHREF `json:"review_comments"`
	ReviewComment  ReviewCommentHREF  `json:"review_comment"`
	Commits        CommitsHREF        `json:"commits"`
	Statuses       StatusesHREF       `json:"statuses"`
}

// PullRequest contains GitHub's pull_request information
type PullRequest struct {
	URL               string           `json:"url"`
	ID                int              `json:"id"`
	HTMLURL           string           `json:"html_url"`
	DiffURL           string           `json:"diff_url"`
	PatchURL          string           `json:"patch_url"`
	IssueURL          string           `json:"issue_url"`
	Number            int              `json:"number"`
	State             string           `json:"state"`
	Locked            bool             `json:"locked"`
	Title             string           `json:"title"`
	User              User             `json:"user"`
	Body              string           `json:"body"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
	ClosedAt          time.Time        `json:"closed_at"`
	MergedAt          time.Time        `json:"merged_at"`
	MergeCommitSHA    string           `json:"merge_commit_sha"`
	Assignee          string           `json:"assignee"`
	Milestone         string           `json:"milestone"`
	CommitsURL        string           `json:"commits_url"`
	ReviewCommentsURL string           `json:"review_comments_url"`
	ReviewCommentURL  string           `json:"review_comment_url"`
	CommentsURL       string           `json:"comments_url"`
	StatusesURL       string           `json:"statuses_url"`
	Head              Head             `json:"head"`
	Base              Base             `json:"base"`
	Links             LinksPullRequest `json:"_links"`
	Merged            bool             `json:"merged"`
	Mergable          bool             `json:"mergeable"`
	MergableState     string           `json:"mergeable_state"`
	MergedBy          string           `json:"merged_by"`
	Comments          int              `json:"comments"`
	ReviewComments    int              `json:"review_comments"`
	Commits           int              `json:"commits"`
	Additions         int              `json:"additions"`
	Deletions         int              `json:"deletions"`
	ChangedFiles      int              `json:"changed_files"`
}

// PusherPush contains GitHub's push pusher information
type PusherPush struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CommitAuthor contains GitHub's commit author information
type CommitAuthor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// CommitCommitter contains GitHub's commit commiter information
type CommitCommitter struct {
	CommitAuthor
}

// Commit contains GitHub's commit information
type Commit struct {
	ID        string          `json:"id"`
	Distinct  bool            `json:"distinct"`
	Message   string          `json:"message"`
	Timestamp time.Time       `json:"timestamp"`
	URL       string          `json:"url"`
	Author    CommitAuthor    `json:"author"`
	Committer CommitCommitter `json:"committer"`
	Added     []string        `json:"added"`
	Removed   []string        `json:"removed"`
	Modified  []string        `json:"modified"`
}

// HeadCommit contains GitHub's head_commit information
type HeadCommit struct {
	Commit
}

// Release contains GitHub's release information
type Release struct {
	URL             string    `json:"url"`
	AssetsURL       string    `json:"assets_url"`
	UploadURL       string    `json:"upload_url"`
	HTMLURL         string    `json:"html_url"`
	ID              int       `json:"id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Author          Author    `json:"author"`
	Prelelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []string  `json:"assets"`
	TarballURL      string    `json:"tarball_url"`
	ZipballURL      string    `json:"zipball_url"`
	Body            string    `json:"body"`
}

// BranchCommit contains GitHub's branch commit information
type BranchCommit struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

// Branch contains GitHub's branch information
type Branch struct {
	Name   string       `json:"name"`
	Commit BranchCommit `json:"commit"`
}

// StatusCommitAuthor contains GitHub's status commit author information
type StatusCommitAuthor struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

// StatusCommitCommiter contains GitHub's status commit committer information
type StatusCommitCommiter struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

// Tree contains GitHub's tree information
type Tree struct {
	BranchCommit
}

// StatusCommitInner contains GitHub's inner status commit information
type StatusCommitInner struct {
	Author       StatusCommitAuthor   `json:"author"`
	Commiter     StatusCommitCommiter `jsons:"committer"`
	Message      string               `json:"message"`
	Tree         Tree                 `json:"tree"`
	URL          string               `json:"url"`
	CommentCount int                  `json:"comment_count"`
}

// StatusCommit contains GitHub's status commit information
type StatusCommit struct {
	SHA         string            `json:"sha"`
	Commit      StatusCommitInner `json:"commit"`
	URL         string            `json:"url"`
	HTMLURL     string            `json:"html_url"`
	CommentsURL string            `json:"comments_url"`
	Author      Author            `json:"author"`
	Committer   Commiter          `json:"committer"`
	Parents     []string          `json:"parents"`
}
