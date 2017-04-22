package gitlab

import (
	"strings"
	"time"
)

type customTime struct {
	time.Time
}

func (t *customTime) UnmarshalJSON(b []byte) (err error) {
	layout := []string{
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05 Z07:00",
		"2006-01-02 15:04:05 Z0700",
		time.RFC3339,
	}
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	for _, l := range layout {
		t.Time, err = time.Parse(l, s)
		if err == nil {
			break
		}
	}
	return
}

// IssueEventPayload contains the information for GitLab's issue event
type IssueEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	Repository       Repository       `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Assignee         Assignee         `json:"assignee"`
}

// MergeRequestEventPayload contains the information for GitLab's merge request event
type MergeRequestEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

// PushEventPayload contains the information for GitLab's push event
type PushEventPayload struct {
	ObjectKind        string     `json:"object_kind"`
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	CheckoutSHA       string     `json:"checkout_sha"`
	UserID            int64      `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserEmail         string     `json:"user_email"`
	UserAvatar        string     `json:"user_avatar"`
	ProjectID         int64      `json:"project_id"`
	Project           Project    `json:"Project"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int64      `json:"total_commits_count"`
}

// TagEventPayload contains the information for GitLab's tag push event
type TagEventPayload struct {
	ObjectKind        string     `json:"object_kind"`
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	CheckoutSHA       string     `json:"checkout_sha"`
	UserID            int64      `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserAvatar        string     `json:"user_avatar"`
	ProjectID         int64      `json:"project_id"`
	Project           Project    `json:"Project"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int64      `json:"total_commits_count"`
}

// WikiPageEventPayload contains the information for GitLab's wiki created/updated event
type WikiPageEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	Wiki             Wiki             `json:"wiki"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

// PipelineEventPayload contains the information for GitLab's pipeline status change event
type PipelineEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	Commit           Commit           `json:"commit"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Builds           []Build          `json:"builds"`
}

// CommentEventPayload contains the information for GitLab's comment event
type CommentEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	ProjectID        int64            `json:"project_id"`
	Project          Project          `json:"project"`
	Repository       Repository       `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	MergeRequest     MergeRequest     `json:"merge_request"`
	Commit           Commit           `json:"commit"`
	Issue            Issue            `json:"issue"`
	Snippet          Snippet          `json:"snippet"`
}

// BuildEventPayload contains the information for GitLab's build status change event
type BuildEventPayload struct {
	ObjectKind        string      `json:"object_kind"`
	Ref               string      `json:"ref"`
	Tag               bool        `json:"tag"`
	BeforeSHA         string      `json:"before_sha"`
	SHA               string      `json:"sha"`
	BuildID           int64       `json:"build_id"`
	BuildName         string      `json:"build_name"`
	BuildStage        string      `json:"build_stage"`
	BuildStatus       string      `json:"build_status"`
	BuildStartedAt    customTime  `json:"build_started_at"`
	BuildFinishedAt   customTime  `json:"build_finished_at"`
	BuildDuration     int64       `json:"build_duration"`
	BuildAllowFailure bool        `json:"build_allow_failure"`
	ProjectID         int64       `json:"project_id"`
	ProjectName       string      `json:"project_name"`
	User              User        `json:"user"`
	Commit            BuildCommit `json:"commit"`
	Repository        Repository  `json:"repository"`
}

// Issue contains all of the GitLab issue information
type Issue struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	AssigneeID  int64      `json:"assignee_id"`
	AuthorID    int64      `json:"author_id"`
	ProjectID   int64      `json:"project_id"`
	CreatedAt   customTime `json:"created_at"`
	UpdatedAt   customTime `json:"updated_at"`
	Position    int64      `json:"position"`
	BranchName  string     `json:"branch_name"`
	Description string     `json:"description"`
	MilestoneID int64      `json:"milestone_id"`
	State       string     `json:"state"`
	IID         int64      `json:"iid"`
}

// Build contains all of the GitLab build information
type Build struct {
	ID            int64         `json:"id"`
	Stage         string        `json:"stage"`
	Name          string        `json:"name"`
	Status        string        `json:"status"`
	CreatedAt     customTime    `json:"created_at"`
	StartedAt     customTime    `json:"started_at"`
	FinishedAt    customTime    `json:"finished_at"`
	When          string        `json:"when"`
	Manual        bool          `json:"manual"`
	User          User          `json:"user"`
	Runner        string        `json:"runner"`
	ArtifactsFile ArtifactsFile `json:"artifactsfile"`
}

// ArtifactsFile contains all of the GitLab artifact information
type ArtifactsFile struct {
	Filename string `json:"filename"`
	Size     string `json:"size"`
}

// Wiki contains all of the GitLab wiki information
type Wiki struct {
	WebURL            string `json:"web_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
}

// Commit contains all of the GitLab commit information
type Commit struct {
	ID        string     `json:"id"`
	Message   string     `json:"message"`
	Timestamp customTime `json:"timestamp"`
	URL       string     `json:"url"`
	Author    Author     `json:"author"`
	Added     []string   `json:"added"`
	Modified  []string   `json:"modified"`
	Removed   []string   `json:"removed"`
}

// BuildCommit contains all of the GitLab build commit information
type BuildCommit struct {
	ID          int64      `json:"id"`
	SHA         string     `json:"sha"`
	Message     string     `json:"message"`
	AuthorName  string     `json:"auuthor_name"`
	AuthorEmail string     `json:"author_email"`
	Status      string     `json:"status"`
	Duration    int64      `json:"duration"`
	StartedAt   customTime `json:"started_at"`
	FinishedAt  customTime `json:"finished_at"`
}

// Snippet contains all of the GitLab snippet information
type Snippet struct {
	ID              int64      `json:"id"`
	Title           string     `json:"title"`
	Content         string     `json:"content"`
	AuthorID        int64      `json:"author_id"`
	ProjectID       int64      `json:"project_id"`
	CreatedAt       customTime `json:"created_at"`
	UpdatedAt       customTime `json:"updated_at"`
	FileName        string     `json:"file_name"`
	ExpiresAt       customTime `json:"expires_at"`
	Type            string     `json:"type"`
	VisibilityLevel int64      `json:"visibility_level"`
}

// User contains all of the GitLab user information
type User struct {
	Name      string `json:"name"`
	UserName  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// Project contains all of the GitLab project information
type Project struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSSHURL        string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int64  `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

// Repository contains all of the GitLab repository information
type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

// ObjectAttributes contains all of the GitLab object attributes information
type ObjectAttributes struct {
	ID              int64      `json:"id"`
	Title           string     `json:"title"`
	AssigneeID      int64      `json:"assignee_id"`
	AuthorID        int64      `json:"author_id"`
	ProjectID       int64      `json:"project_id"`
	CreatedAt       customTime `json:"created_at"`
	UpdatedAt       customTime `json:"updated_at"`
	Position        int64      `json:"position"`
	BranchName      string     `json:"branch_name"`
	Description     string     `json:"description"`
	MilestoneID     int64      `json:"milestone_id"`
	State           string     `json:"state"`
	IID             int64      `json:"iid"`
	URL             string     `json:"url"`
	Action          string     `json:"action"`
	TargetBranch    string     `json:"target_branch"`
	SourceBranch    string     `json:"source_branch"`
	SourceProjectID int64      `json:"source_project_id"`
	TargetProjectID int64      `json:"target_project_id"`
	StCommits       string     `json:"st_commits"`
	MergeStatus     string     `json:"merge_status"`
	Content         string     `json:"content"`
	Format          string     `json:"format"`
	Message         string     `json:"message"`
	Slug            string     `json:"slug"`
	Ref             string     `json:"ref"`
	Tag             bool       `json:"tag"`
	SHA             string     `json:"sha"`
	BeforeSHA       string     `json:"before_sha"`
	Status          string     `json:"status"`
	Stages          []string   `json:"stages"`
	Duration        int64      `json:"duration"`
	Note            string     `json:"note"`
	NotebookType    string     `json:"noteable_type"`
	At              customTime `json:"attachment"`
	LineCode        string     `json:"line_code"`
	CommitID        string     `json:"commit_id"`
	NoteableID      int64      `json:"noteable_id"`
	System          bool       `json:"system"`
	WorkInProgress  bool       `json:"work_in_progress"`
	StDiffs         []StDiff   `json:"st_diffs"`
	Source          Source     `json:"source"`
	Target          Target     `json:"target"`
	LastCommit      LastCommit `json:"last_commit"`
	Assignee        Assignee   `json:"assignee"`
}

// MergeRequest contains all of the GitLab merge request information
type MergeRequest struct {
	ID              int64      `json:"id"`
	TargetBranch    string     `json:"target_branch"`
	SourceBranch    string     `json:"source_branch"`
	SourceProjectID string     `json:"source_project_id"`
	AssigneeID      int64      `json:"assignee_id"`
	AuthorID        int64      `json:"author_id"`
	Title           string     `json:"title"`
	CreatedAt       customTime `json:"created_at"`
	UpdatedAt       customTime `json:"updated_at"`
	MilestoneID     int64      `json:"milestone_id"`
	State           string     `json:"state"`
	MergeStatus     string     `json:"merge_status"`
	TargetProjectID int64      `json:"target_project_id"`
	IID             int64      `json:"iid"`
	Description     string     `json:"description"`
	Position        int64      `json:"position"`
	LockedAt        customTime `json:"locked_at"`
	Source          Source     `json:"source"`
	Target          Target     `json:"target"`
	LastCommit      LastCommit `json:"last_commit"`
	WorkInProgress  bool       `json:"work_in_progress"`
	Assignee        Assignee   `json:"assignee"`
}

// Assignee contains all of the GitLab assignee information
type Assignee struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// StDiff contains all of the GitLab diff information
type StDiff struct {
	Diff        string `json:"diff"`
	NewPath     string `json:"new_path"`
	OldPath     string `json:"old_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
}

// Source contains all of the GitLab source information
type Source struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int64  `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

// Target contains all of the GitLab target information
type Target struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int64  `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

// LastCommit contains all of the GitLab last commit information
type LastCommit struct {
	ID        string     `json:"id"`
	Message   string     `json:"message"`
	Timestamp customTime `json:"timestamp"`
	URL       string     `json:"url"`
	Author    Author     `json:"author"`
}

// Author contains all of the GitLab author information
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
