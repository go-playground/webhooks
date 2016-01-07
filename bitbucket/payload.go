package bitbucket

import "time"

// PullRequestCommentDeletedPayload is the BitBucket pull_request:comment_deleted payload
type PullRequestCommentDeletedPayload struct {
	Actor       User        `json:"actor"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pullrequest"`
	Comment     Comment     `json:"comment"`
}

// PullRequestCommentUpdatedPayload is the BitBucket pullrequest:comment_updated payload
type PullRequestCommentUpdatedPayload struct {
	Actor       User        `json:"actor"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pullrequest"`
	Comment     Comment     `json:"comment"`
}

// PullRequestCommentCreatedPayload is the BitBucket pullrequest:comment_created payload
type PullRequestCommentCreatedPayload struct {
	Actor       User        `json:"actor"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pullrequest"`
	Comment     Comment     `json:"comment"`
}

// PullRequestDeclinedPayload is the BitBucket pullrequest:rejected payload
type PullRequestDeclinedPayload struct {
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
}

// PullRequestMergedPayload is the BitBucket pullrequest:fulfilled payload
type PullRequestMergedPayload struct {
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
}

// PullRequestUnapprovedPayload is the BitBucket pullrequest:unapproved payload
type PullRequestUnapprovedPayload struct {
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
	Approval    Approval    `json:"approval"`
}

// PullRequestApprovedPayload is the BitBucket pullrequest:approved payload
type PullRequestApprovedPayload struct {
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
	Approval    Approval    `json:"approval"`
}

// PullRequestUpdatedPayload is the BitBucket pullrequest:updated payload
type PullRequestUpdatedPayload struct {
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
}

// PullRequestCreatedPayload is the BitBucket pullrequest:created payload
type PullRequestCreatedPayload struct {
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
}

// IssueCommentCreatedPayload is the BitBucket issue:comment_created payload
type IssueCommentCreatedPayload struct {
	Actor      User       `json:"actor"`
	Repository Repository `json:"repository"`
	Issue      Issue      `json:"issue"`
	Comment    Comment    `json:"comment"`
}

// IssueUpdatedPayload is the BitBucket issue:updated payload
type IssueUpdatedPayload struct {
	Actor      User         `json:"actor"`
	Issue      Issue        `json:"issue"`
	Repository Repository   `json:"repository"`
	Comment    Comment      `json:"comment"`
	Changes    IssueChanges `json:"changes"`
}

// IssueCreatedPayload is the BitBucket issue:created payload
type IssueCreatedPayload struct {
	Actor      User       `json:"actor"`
	Issue      Issue      `json:"issue"`
	Repository Repository `json:"repository"`
}

// RepoCommitStatusUpdatedPayload is the BitBucket repo:commit_status_updated payload
type RepoCommitStatusUpdatedPayload struct {
	Actor        User         `json:"actor"`
	Repository   Repository   `json:"repository"`
	CommitStatus CommitStatus `json:"commit_status"`
}

// RepoCommitStatusCreatedPayload is the BitBucket repo:commit_status_created payload
type RepoCommitStatusCreatedPayload struct {
	Actor        User         `json:"actor"`
	Repository   Repository   `json:"repository"`
	CommitStatus CommitStatus `json:"commit_status"`
}

// RepoCommitCommentedPayload is the BitBucket repo:commit_comment_created payload
type RepoCommitCommentedPayload struct {
	Actor      User       `json:"actor"`
	Comment    Comment    `json:"comment"`
	Repository Repository `json:"repository"`
	Commit     CommitHash `json:"commit"`
}

// RepoForkPayload is the BitBucket repo:fork payload
type RepoForkPayload struct {
	Actor      User       `json:"actor"`
	Repository Repository `json:"repository"`
	Fork       Repository `json:"fork"`
}

// RepoPushPayload is the BitBucket repo:push payload
type RepoPushPayload struct {
	Actor      User       `json:"actor"`
	Repository Repository `json:"repository"`
	Push       Push       `json:"push"`
}

// Approval is the common BitBucket Issue Approval Sub Entity
type Approval struct {
	Date time.Time `json:"date"`
	User User      `json:"user"`
}

// IssueChanges is the common BitBucket Issue Changes Sub Entity
type IssueChanges struct {
	Status IssueChangeStatus `json:"status"`
}

// IssueChangeStatus is the common BitBucket Issue Change Status Sub Entity
type IssueChangeStatus struct {
	Old string `json:"old"`
	New string `json:"new"`
}

// CommitStatus is the common BitBucket CommitStatus Sub Entity
type CommitStatus struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	State       string          `json:"state"`
	Key         string          `json:"key"`
	URL         string          `json:"url"`
	Type        string          `json:"type"`
	CreatedOn   time.Time       `json:"created_on"`
	UpdatedOn   time.Time       `json:"updated_on"`
	Links       LinksSelfCommit `json:"links"`
}

// Push is the common BitBucket Push Sub Entity
type Push struct {
	Changes []Change `json:"changes"`
}

// Change is the common BitBucket Change Sub Entity
type Change struct {
	New       ChangeData           `json:"new"`
	Old       ChangeData           `json:"old"`
	Links     LinksHTMLDiffCommits `json:"links"`
	Created   bool                 `json:"created"`
	Forced    bool                 `json:"forced"`
	Closed    bool                 `json:"closed"`
	Commits   []Commit             `json:"commits"`
	Truncated bool                 `json:"truncated"`
}

// ChangeData is the common BitBucket ChangeData Sub Entity
type ChangeData struct {
	Type   string               `json:"type"`
	Name   string               `json:"name"`
	Target Target               `json:"target"`
	Links  LinksHTMLSelfCommits `json:"links"`
}

// Target is the common BitBucket Target Sub Entity
type Target struct {
	Type    string        `json:"type"`
	Hash    string        `json:"hash"`
	Author  User          `json:"author"`
	Message string        `json:"message"`
	Date    time.Time     `json:"date"`
	Parents []Parent      `json:"parents"`
	Links   LinksHTMLSelf `json:"links"`
}

// Parent is the common BitBucket Parent Sub Entity
type Parent struct {
	Type  string        `json:"type"`
	Hash  string        `json:"hash"`
	Links LinksHTMLSelf `json:"links"`
}

// Commit is the common BitBucket Commit Sub Entity
type Commit struct {
	Hash    string        `json:"hash"`
	Type    string        `json:"type"`
	Message string        `json:"message"`
	Author  User          `json:"author"`
	Links   LinksHTMLSelf `json:"links"`
}

// User is the common BitBucket User Entity
type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	UUID        string `json:"uuid"`
	Links       Links  `json:"links"`
}

// Repository is the common BitBucket Repository Entity
type Repository struct {
	Links     Links  `json:"links"`
	UUID      string `json:"uuid"`
	FullName  string `json:"full_name"`
	Name      string `json:"name"`
	Scm       string `json:"scm"`
	IsPrivate bool   `json:"is_private"`
}

// Issue is the common BitBucket Issue Entity
type Issue struct {
	ID        int64         `json:"id"`
	Component string        `json:"component"`
	Title     string        `json:"title"`
	Content   Content       `json:"content"`
	Priority  string        `json:"priority"`
	State     string        `json:"state"`
	Type      string        `json:"type"`
	Milestone Milestone     `json:"milestone"`
	Version   Version       `json:"version"`
	CreatedOn time.Time     `json:"created_on"`
	UpdatedOn time.Time     `json:"updated_on"`
	Links     LinksHTMLSelf `json:"links"`
}

// Comment is the common BitBucket Comment Entity
type Comment struct {
	ID        int64         `json:"id"`
	Parent    ParentID      `json:"parent"`
	Content   Content       `json:"content"`
	Inline    Inline        `json:"inline"`
	CreatedOn time.Time     `json:"created_on"`
	UpdatedOn time.Time     `json:"updated_on"`
	Links     LinksHTMLSelf `json:"links"`
}

// PullRequest is the common BitBucket PullRequest Entity
type PullRequest struct {
	ID                int64         `json:"id"`
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	State             string        `json:"state"`
	Author            User          `json:"author"`
	Source            Source        `json:"source"`
	Destination       Destination   `json:"destination"`
	MergeCommit       CommitHash    `json:"merge_commit"`
	Participants      []User        `json:"participants"`
	Reviewers         []User        `json:"reviewers"`
	CloseSourceBranch bool          `json:"close_source_branch"`
	ClosedBy          User          `json:"closed_by"`
	Reason            string        `json:"reason"`
	CreatedOn         time.Time     `json:"created_on"`
	UpdatedOn         time.Time     `json:"updated_on"`
	Links             LinksHTMLSelf `json:"links"`
}

// Destination is the common BitBucket Destination Sub Entity
type Destination struct {
	Branch     Branch     `json:"branch"`
	Commit     CommitHash `json:"commit"`
	Repository Repository `json:"repository"`
}

// Source is the common BitBucket Source Sub Entity
type Source struct {
	Branch     Branch     `json:"branch"`
	Commit     CommitHash `json:"commit"`
	Repository Repository `json:"repository"`
}

// Branch is the common BitBucket Branch Sub Entity
type Branch struct {
	Name string `json:"name"`
}

// CommitHash is the common BitBucket CommitHash Sub Entity
type CommitHash struct {
	Hash string `json:"hash"`
}

// Inline is the common BitBucket Inline Sub Entity
type Inline struct {
	Path string `json:"path"`
	From *int64 `json:"from"`
	To   int64  `json:"to"`
}

// ParentID is the common BitBucket ParentID Sub Entity
type ParentID struct {
	ID int64 `json:"id"`
}

// Avatar is the common BitBucket Avatar Sub Entity
type Avatar struct {
	HREF string `json:"href"`
}

// HTML is the common BitBucket HTML Sub Entity
type HTML struct {
	HREF string `json:"href"`
}

// Self is the common BitBucket Self Sub Entity
type Self struct {
	HREF string `json:"href"`
}

// Diff is the common BitBucket Diff Sub Entity
type Diff struct {
	HREF string `json:"href"`
}

// Commits is the common BitBucket Commits Sub Entity
type Commits struct {
	HREF string `json:"href"`
}

// LinksSelfCommit is the common BitBucket LinksSelfCommit Sub Entity
type LinksSelfCommit struct {
	Self   Self    `json:"self"`
	Commit Commits `json:"commit"`
}

// LinksHTMLSelfCommits is the common BitBucket LinksHTMLSelfCommits Sub Entity
type LinksHTMLSelfCommits struct {
	Self    Self    `json:"self"`
	Commits Commits `json:"commits"`
	HTML    HTML    `json:"html"`
}

// LinksHTMLDiffCommits is the common BitBucket LinksHTMLDiffCommits Sub Entity
type LinksHTMLDiffCommits struct {
	HTML    HTML    `json:"html"`
	Diff    Diff    `json:"diff"`
	Commits Commits `json:"commits"`
}

// Links is the common BitBucket Links Sub Entity
type Links struct {
	Avatar Avatar `json:"avatar"`
	HTML   HTML   `json:"html"`
	Self   Self   `json:"self"`
}

// LinksHTMLSelf is the common BitBucket LinksHTMLSelf Sub Entity
type LinksHTMLSelf struct {
	HTML HTML `json:"html"`
	Self Self `json:"self"`
}

// Content is the common BitBucket Content Sub Entity
type Content struct {
	HTML   string `json:"html"`
	Markup string `json:"markup"`
	Raw    string `json:"raw"`
}

// Milestone is the common BitBucket Milestone Sub Entity
type Milestone struct {
	Name string `json:"name"`
}

// Version is the common BitBucket Version Sub Entity
type Version struct {
	Name string `json:"name"`
}
