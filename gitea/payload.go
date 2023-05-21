package gitea

import "time"

// PayloadUser represents the author or committer of a commit
type PayloadUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

// PayloadCommitVerification represents the GPG verification of a commit
type PayloadCommitVerification struct {
	Verified  bool         `json:"verified"`
	Reason    string       `json:"reason"`
	Signature string       `json:"signature"`
	Signer    *PayloadUser `json:"signer"`
	Payload   string       `json:"payload"`
}

// PayloadCommit represents a commit
type PayloadCommit struct {
	// sha1 hash of the commit
	ID           string                     `json:"id"`
	Message      string                     `json:"message"`
	URL          string                     `json:"url"`
	Author       *PayloadUser               `json:"author"`
	Committer    *PayloadUser               `json:"committer"`
	Verification *PayloadCommitVerification `json:"verification"`
	// swagger:strfmt date-time
	Timestamp time.Time `json:"timestamp"`
	Added     []string  `json:"added"`
	Removed   []string  `json:"removed"`
	Modified  []string  `json:"modified"`
}

// Repository represents a repository
type Repository struct {
	ID                        int64            `json:"id"`
	Owner                     *User            `json:"owner"`
	Name                      string           `json:"name"`
	FullName                  string           `json:"full_name"`
	Description               string           `json:"description"`
	Empty                     bool             `json:"empty"`
	Private                   bool             `json:"private"`
	Fork                      bool             `json:"fork"`
	Template                  bool             `json:"template"`
	Parent                    *Repository      `json:"parent"`
	Mirror                    bool             `json:"mirror"`
	Size                      int              `json:"size"`
	HTMLURL                   string           `json:"html_url"`
	SSHURL                    string           `json:"ssh_url"`
	CloneURL                  string           `json:"clone_url"`
	OriginalURL               string           `json:"original_url"`
	Website                   string           `json:"website"`
	Stars                     int              `json:"stars_count"`
	Forks                     int              `json:"forks_count"`
	Watchers                  int              `json:"watchers_count"`
	OpenIssues                int              `json:"open_issues_count"`
	OpenPulls                 int              `json:"open_pr_counter"`
	Releases                  int              `json:"release_counter"`
	DefaultBranch             string           `json:"default_branch"`
	Archived                  bool             `json:"archived"`
	Created                   time.Time        `json:"created_at"`
	Updated                   time.Time        `json:"updated_at"`
	Permissions               *Permission      `json:"permissions,omitempty"`
	HasIssues                 bool             `json:"has_issues"`
	InternalTracker           *InternalTracker `json:"internal_tracker,omitempty"`
	ExternalTracker           *ExternalTracker `json:"external_tracker,omitempty"`
	HasWiki                   bool             `json:"has_wiki"`
	ExternalWiki              *ExternalWiki    `json:"external_wiki,omitempty"`
	HasPullRequests           bool             `json:"has_pull_requests"`
	HasProjects               bool             `json:"has_projects"`
	IgnoreWhitespaceConflicts bool             `json:"ignore_whitespace_conflicts"`
	AllowMerge                bool             `json:"allow_merge_commits"`
	AllowRebase               bool             `json:"allow_rebase"`
	AllowRebaseMerge          bool             `json:"allow_rebase_explicit"`
	AllowSquash               bool             `json:"allow_squash_merge"`
	DefaultMergeStyle         string           `json:"default_merge_style"`
	AvatarURL                 string           `json:"avatar_url"`
	Internal                  bool             `json:"internal"`
	MirrorInterval            string           `json:"mirror_interval"`
}

// User represents a user
type User struct {
	ID            int64     `json:"id"`
	UserName      string    `json:"login"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	AvatarURL     string    `json:"avatar_url"`
	Language      string    `json:"language"`
	IsAdmin       bool      `json:"is_admin"`
	LastLogin     time.Time `json:"last_login,omitempty"`
	Created       time.Time `json:"created,omitempty"`
	Restricted    bool      `json:"restricted"`
	IsActive      bool      `json:"active"`
	ProhibitLogin bool      `json:"prohibit_login"`
	Location      string    `json:"location"`
	Website       string    `json:"website"`
	Description   string    `json:"description"`
	Visibility    string    `json:"visibility"`
	Followers     int       `json:"followers_count"`
	Following     int       `json:"following_count"`
	StarredRepos  int       `json:"starred_repos_count"`
}

// Permission represents a set of permissions
type Permission struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

// InternalTracker represents settings for internal tracker
type InternalTracker struct {
	EnableTimeTracker                bool `json:"enable_time_tracker"`
	AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time"`
	EnableIssueDependencies          bool `json:"enable_issue_dependencies"`
}

// ExternalTracker represents settings for external tracker
type ExternalTracker struct {
	ExternalTrackerURL    string `json:"external_tracker_url"`
	ExternalTrackerFormat string `json:"external_tracker_format"`
	ExternalTrackerStyle  string `json:"external_tracker_style"`
}

// ExternalWiki represents setting for external wiki
type ExternalWiki struct {
	ExternalWikiURL string `json:"external_wiki_url"`
}

// PusherType define the type to push
type PusherType string

// HookIssueCommentAction defines hook issue comment action
type HookIssueCommentAction string

// Issue represents an issue in a repository
type Issue struct {
	ID               int64            `json:"id"`
	URL              string           `json:"url"`
	HTMLURL          string           `json:"html_url"`
	Index            int64            `json:"number"`
	Poster           *User            `json:"user"`
	OriginalAuthor   string           `json:"original_author"`
	OriginalAuthorID int64            `json:"original_author_id"`
	Title            string           `json:"title"`
	Body             string           `json:"body"`
	Ref              string           `json:"ref"`
	Labels           []*Label         `json:"labels"`
	Milestone        *Milestone       `json:"milestone"`
	Assignee         *User            `json:"assignee"`
	Assignees        []*User          `json:"assignees"`
	State            StateType        `json:"state"`
	IsLocked         bool             `json:"is_locked"`
	Comments         int              `json:"comments"`
	Created          time.Time        `json:"created_at"`
	Updated          time.Time        `json:"updated_at"`
	Closed           *time.Time       `json:"closed_at"`
	Deadline         *time.Time       `json:"due_date"`
	PullRequest      *PullRequestMeta `json:"pull_request"`
	Repo             *RepositoryMeta  `json:"repository"`
}

// Label a label to an issue or a pr
type Label struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Milestone milestone is a collection of issues on one repository
type Milestone struct {
	ID           int64      `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	State        StateType  `json:"state"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	Created      time.Time  `json:"created_at"`
	Updated      *time.Time `json:"updated_at"`
	Closed       *time.Time `json:"closed_at"`
	Deadline     *time.Time `json:"due_on"`
}

// StateType issue state type
type StateType string

// PullRequestMeta PR info if an issue is a PR
type PullRequestMeta struct {
	HasMerged bool       `json:"merged"`
	Merged    *time.Time `json:"merged_at"`
}

// RepositoryMeta basic repository information
type RepositoryMeta struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	FullName string `json:"full_name"`
}

// CreatePayload represents a payload create reposoitory
type CreatePayload struct {
	Sha     string      `json:"sha"`
	Ref     string      `json:"ref"`
	RefType string      `json:"ref_type"`
	Repo    *Repository `json:"repository"`
	Sender  *User       `json:"sender"`
}

// DeletePayload represents delete payload
type DeletePayload struct {
	Ref        string      `json:"ref"`
	RefType    string      `json:"ref_type"`
	PusherType PusherType  `json:"pusher_type"`
	Repo       *Repository `json:"repository"`
	Sender     *User       `json:"sender"`
}

// ForkPayload represents fork payload
type ForkPayload struct {
	Forkee *Repository `json:"forkee"`
	Repo   *Repository `json:"repository"`
	Sender *User       `json:"sender"`
}

// IssueCommentPayload represents a payload information of issue comment event.
type IssueCommentPayload struct {
	Action     HookIssueCommentAction `json:"action"`
	Issue      *Issue                 `json:"issue"`
	Comment    *Comment               `json:"comment"`
	Changes    *ChangesPayload        `json:"changes,omitempty"`
	Repository *Repository            `json:"repository"`
	Sender     *User                  `json:"sender"`
	IsPull     bool                   `json:"is_pull"`
}

// ChangesPayload represents the payload information of issue change
type ChangesPayload struct {
	Title *ChangesFromPayload `json:"title,omitempty"`
	Body  *ChangesFromPayload `json:"body,omitempty"`
	Ref   *ChangesFromPayload `json:"ref,omitempty"`
}

// ChangesFromPayload FIXME
type ChangesFromPayload struct {
	From string `json:"from"`
}

// Comment represents a comment on a commit or issue
type Comment struct {
	ID               int64     `json:"id"`
	HTMLURL          string    `json:"html_url"`
	PRURL            string    `json:"pull_request_url"`
	IssueURL         string    `json:"issue_url"`
	Poster           *User     `json:"user"`
	OriginalAuthor   string    `json:"original_author"`
	OriginalAuthorID int64     `json:"original_author_id"`
	Body             string    `json:"body"`
	Created          time.Time `json:"created_at"`
	Updated          time.Time `json:"updated_at"`
}

// PushPayload represents a payload information of push event.
type PushPayload struct {
	Ref        string           `json:"ref"`
	Before     string           `json:"before"`
	After      string           `json:"after"`
	CompareURL string           `json:"compare_url"`
	Commits    []*PayloadCommit `json:"commits"`
	HeadCommit *PayloadCommit   `json:"head_commit"`
	Repo       *Repository      `json:"repository"`
	Pusher     *User            `json:"pusher"`
	Sender     *User            `json:"sender"`
}

// PullRequestPayload represents a payload information of pull request event.
type PullRequestPayload struct {
	Action      HookIssueAction `json:"action"`
	Index       int64           `json:"number"`
	Changes     *ChangesPayload `json:"changes,omitempty"`
	PullRequest *PullRequest    `json:"pull_request"`
	Repository  *Repository     `json:"repository"`
	Sender      *User           `json:"sender"`
	Review      *ReviewPayload  `json:"review"`
}

// HookIssueAction FIXME
type HookIssueAction string

// PullRequest represents a pull request
type PullRequest struct {
	ID             int64         `json:"id"`
	URL            string        `json:"url"`
	Index          int64         `json:"number"`
	Poster         *User         `json:"user"`
	Title          string        `json:"title"`
	Body           string        `json:"body"`
	Labels         []*Label      `json:"labels"`
	Milestone      *Milestone    `json:"milestone"`
	Assignee       *User         `json:"assignee"`
	Assignees      []*User       `json:"assignees"`
	State          StateType     `json:"state"`
	IsLocked       bool          `json:"is_locked"`
	Comments       int           `json:"comments"`
	HTMLURL        string        `json:"html_url"`
	DiffURL        string        `json:"diff_url"`
	PatchURL       string        `json:"patch_url"`
	Mergeable      bool          `json:"mergeable"`
	HasMerged      bool          `json:"merged"`
	Merged         *time.Time    `json:"merged_at"`
	MergedCommitID *string       `json:"merge_commit_sha"`
	MergedBy       *User         `json:"merged_by"`
	Base           *PRBranchInfo `json:"base"`
	Head           *PRBranchInfo `json:"head"`
	MergeBase      string        `json:"merge_base"`
	Deadline       *time.Time    `json:"due_date"`
	Created        *time.Time    `json:"created_at"`
	Updated        *time.Time    `json:"updated_at"`
	Closed         *time.Time    `json:"closed_at"`
}

// PRBranchInfo information about a branch
type PRBranchInfo struct {
	Name       string      `json:"label"`
	Ref        string      `json:"ref"`
	Sha        string      `json:"sha"`
	RepoID     int64       `json:"repo_id"`
	Repository *Repository `json:"repo"`
}

// ReviewPayload FIXME
type ReviewPayload struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// RepositoryPayload payload for repository webhooks
type RepositoryPayload struct {
	Action       HookRepoAction `json:"action"`
	Repository   *Repository    `json:"repository"`
	Organization *User          `json:"organization"`
	Sender       *User          `json:"sender"`
}

// HookRepoAction an action that happens to a repo
type HookRepoAction string

// ReleasePayload represents a payload information of release event.
type ReleasePayload struct {
	Action     HookReleaseAction `json:"action"`
	Release    *Release          `json:"release"`
	Repository *Repository       `json:"repository"`
	Sender     *User             `json:"sender"`
}

// Release represents a repository release
type Release struct {
	ID           int64         `json:"id"`
	TagName      string        `json:"tag_name"`
	Target       string        `json:"target_commitish"`
	Title        string        `json:"name"`
	Note         string        `json:"body"`
	URL          string        `json:"url"`
	HTMLURL      string        `json:"html_url"`
	TarURL       string        `json:"tarball_url"`
	ZipURL       string        `json:"zipball_url"`
	IsDraft      bool          `json:"draft"`
	IsPrerelease bool          `json:"prerelease"`
	CreatedAt    time.Time     `json:"created_at"`
	PublishedAt  time.Time     `json:"published_at"`
	Publisher    *User         `json:"author"`
	Attachments  []*Attachment `json:"assets"`
}

// HookReleaseAction defines hook release action type
type HookReleaseAction string

// Attachment a generic attachment
type Attachment struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Size          int64     `json:"size"`
	DownloadCount int64     `json:"download_count"`
	Created       time.Time `json:"created_at"`
	UUID          string    `json:"uuid"`
	DownloadURL   string    `json:"browser_download_url"`
}

// IssuePayload represents the payload information that is sent along with an issue event.
type IssuePayload struct {
	Action     HookIssueAction `json:"action"`
	Index      int64           `json:"number"`
	Changes    *ChangesPayload `json:"changes,omitempty"`
	Issue      *Issue          `json:"issue"`
	Repository *Repository     `json:"repository"`
	Sender     *User           `json:"sender"`
}
