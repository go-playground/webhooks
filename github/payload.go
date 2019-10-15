package github

import "time"

// CheckPullRequest contains the information about a pull request in the context of a check-suite or check-run payload
type CheckPullRequest struct {
	URL    string `json:"url"`
	ID     int64  `json:"id"`
	Number int    `json:"number"`
	Head   struct {
		Ref  string `json:"ref"`
		SHA  string `json:"sha"`
		Repo struct {
			ID   int64  `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"repo"`
	} `json:"head"`
	Base struct {
		Ref  string `json:"ref"`
		SHA  string `json:"sha"`
		Repo struct {
			ID   int64  `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"repo"`
	} `json:"base"`
}

// CheckRunPayload contains the information for GitHub's check_run hook event
type CheckRunPayload struct {
	Action   string `json:"action"`
	CheckRun struct {
		ID          int64      `json:"id"`
		NodeID      string     `json:"node_id"`
		Name        string     `json:"name"`
		HeadSHA     string     `json:"head_sha"`
		Status      string     `json:"status"`
		Conclusion  *string    `json:"conclusion"`
		URL         string     `json:"url"`
		HtmlURL     string     `json:"html_url"`
		StarterAt   time.Time  `json:"started_at"`
		CompletedAt *time.Time `json:"completed_at"`
		DetailsURL  string     `json:"details_url"`
		ExternalID  string     `json:"external_id"`
		Output      struct {
			Title            *string `json:"title"`
			Summary          *string `json:"summary"`
			Text             *string `json:"text"`
			AnnotationsCount int64   `json:"annotations_count"`
			AnnotationsURL   string  `json:"annotations_url"`
		} `json:"output"`
		CheckSuite struct {
			ID           int64              `json:"id"`
			NodeID       string             `json:"node_id"`
			HeadBranch   string             `json:"head_branch"`
			HeadSHA      string             `json:"head_sha"`
			Status       string             `json:"status"`
			Conclusion   *string            `json:"conclusion"`
			URL          string             `json:"url"`
			Before       string             `json:"before"`
			After        string             `json:"after"`
			PullRequests []CheckPullRequest `json:"pull_requests"`
			App          App                `json:"app"`
			CreatedAt    time.Time          `json:"created_at"`
			UpdatedAt    time.Time          `json:"updated_at"`
		} `json:"check_suite"`
		App          App                `json:"app"`
		PullRequests []CheckPullRequest `json:"pull_requests"`
	} `json:"check_run"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// CheckSuitePayload contains the information for GitHub's check_suite hook event
type CheckSuitePayload struct {
	Action     string `json:"action"`
	CheckSuite struct {
		ID                   int64              `json:"id"`
		NodeID               string             `json:"node_id"`
		HeadBranch           string             `json:"head_branch"`
		HeadSHA              string             `json:"head_sha"`
		Status               string             `json:"status"`
		Conclusion           *string            `json:"conclusion"`
		URL                  string             `json:"url"`
		Before               string             `json:"before"`
		After                string             `json:"after"`
		PullRequests         []CheckPullRequest `json:"pull_requests"`
		App                  App                `json:"app"`
		CreatedAt            time.Time          `json:"created_at"`
		UpdatedAt            time.Time          `json:"updated_at"`
		LatestCheckRunsCount int64              `json:"latest_check_runs_count"`
		CheckRunsURL         string             `json:"check_runs_url"`
		HeadCommit           struct {
			ID        string    `json:"id"`
			TreeID    string    `json:"tree_id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
			Committer struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"committer"`
		} `json:"head_commit"`
	} `json:"check_suite"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// CommitCommentPayload contains the information for GitHub's commit_comment hook event
type CommitCommentPayload struct {
	Action  string `json:"action"`
	Comment struct {
		URL               string    `json:"url"`
		HTMLURL           string    `json:"html_url"`
		ID                int64     `json:"id"`
		NodeID            string    `json:"node_id"`
		User              User      `json:"user"`
		Position          *int64    `json:"position"`
		Line              *int64    `json:"line"`
		Path              *string   `json:"path"`
		CommitID          string    `json:"commit_id"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Body              string    `json:"body"`
		AuthorAssociation string    `json:"author_association"`
	} `json:"comment"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// CreatePayload contains the information for GitHub's create hook event
type CreatePayload struct {
	Ref          string       `json:"ref"`
	RefType      string       `json:"ref_type"`
	MasterBranch string       `json:"master_branch"`
	Description  *string      `json:"description"` //TODO: check can it be ref?
	PusherType   string       `json:"pusher_type"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// DeletePayload contains the information for GitHub's delete hook event
type DeletePayload struct {
	Ref          string       `json:"ref"`
	RefType      string       `json:"ref_type"`
	PusherType   string       `json:"pusher_type"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// DeploymentPayload contains the information for GitHub's deployment hook
type DeploymentPayload struct {
	Deployment   Deployment   `json:"deployment"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// DeploymentStatusPayload contains the information for GitHub's deployment_status hook event
type DeploymentStatusPayload struct {
	DeploymentStatus struct {
		URL           string    `json:"url"`
		ID            int64     `json:"id"`
		NodeID        string    `json:"node_id"`
		State         string    `json:"state"`
		Creator       User      `json:"creator"`
		Description   *string   `json:"description"`
		Environment   string    `json:"environment"`
		TargetURL     *string   `json:"target_url"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		DeploymentURL string    `json:"deployment_url"`
		RepositoryURL string    `json:"repository_url"`
	} `json:"deployment_status"`
	Deployment   Deployment   `json:"deployment"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// ForkPayload contains the information for GitHub's fork hook event
type ForkPayload struct {
	// TODO: Forkee has a `Public bool` that is not in the Repository object, but
	// it is redundant with the `Private bool` property, is it worth it to duplicate
	// thes struct to add this one property?
	Forkee       Repository   `json:"forkee"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// GollumPayload contains the information for GitHub's gollum hook event
type GollumPayload struct {
	Pages []struct {
		PageName string  `json:"page_name"`
		Title    string  `json:"title"`
		Summary  *string `json:"summary"`
		Action   string  `json:"action"`
		Sha      string  `json:"sha"`
		HTMLURL  string  `json:"html_url"`
	} `json:"pages"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// InstallationPayload contains the information for GitHub's installation and integration_installation hook events
type InstallationPayload struct {
	Action       string          `json:"action"`
	Installation InstallationObj `json:"installation"`
	Repositories []struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	} `json:"repositories"`
	Sender User `json:"sender"`
}

// InstallationRepositoriesPayload contains the information for GitHub's installation_repositories hook events
type InstallationRepositoriesPayload struct {
	Action              string          `json:"action"`
	RepositorySelection string          `json:"repository_selection"`
	Installation        InstallationObj `json:"installation"`
	RepositoriesAdded   []struct {
		ID       int64  `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	} `json:"repositories_added"`
	RepositoriesRemoved []struct {
		ID       int64  `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	} `json:"repositories_removed"`
	Sender User `json:"sender"`
}

// IssueCommentPayload contains the information for GitHub's issue_comment hook event
type IssueCommentPayload struct {
	Action  string `json:"action"`
	Issue   Issue  `json:"issue"`
	Comment struct {
		URL               string    `json:"url"`
		HTMLURL           string    `json:"html_url"`
		IssueURL          string    `json:"issue_url"`
		ID                int64     `json:"id"`
		NodeID            string    `json:"node_id"`
		User              User      `json:"user"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Body              string    `json:"body"`
		AuthorAssociation string    `json:"author_association"`
	} `json:"comment"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// IssuesPayload contains the information for GitHub's issues hook event
type IssuesPayload struct {
	Action       string       `json:"action"`
	Issue        Issue        `json:"issue"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Changes      interface{}  `json:"changes"` //TODO not sure what goes here.
	Installation Installation `json:"installation"`
}

// LabelPayload contains the information for GitHub's label hook event
type LabelPayload struct {
	Action       string       `json:"action"`
	Label        Label        `json:"label"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// MemberPayload contains the information for GitHub's member hook event
type MemberPayload struct {
	Action       string       `json:"action"`
	Member       User         `json:"member"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// MembershipPayload contains the information for GitHub's membership hook event
type MembershipPayload struct {
	Action       string       `json:"action"`
	Scope        string       `json:"scope"`
	Member       User         `json:"member"`
	Sender       User         `json:"sender"`
	Team         Team         `json:"team"`
	Organization Organization `json:"organization"`
	Installation Installation `json:"installation"`
}

// MetaPayload contains the information for GitHub's meta hook event
type MetaPayload struct {
	HookID int    `json:"hook_id"`
	Action string `json:"action"`
	Hook   struct {
		Type   string   `json:"type"`
		ID     int64    `json:"id"`
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
			URL         string `json:"url"`
		} `json:"config"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"hook"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// MilestonePayload contains the information for GitHub's milestone hook event
type MilestonePayload struct {
	Action       string       `json:"action"`
	Milestone    Milestone    `json:"milestone"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// OrganizationPayload contains the information for GitHub's organization hook event
type OrganizationPayload struct {
	Action     string `json:"action"`
	Membership struct {
		URL             string `json:"url"`
		State           string `json:"state"`
		Role            string `json:"role"`
		OrganizationURL string `json:"organization_url"`
		User            User   `json:"user"`
	} `json:"membership"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// OrgBlockPayload contains the information for GitHub's org_block hook event
type OrgBlockPayload struct {
	Action       string       `json:"action"`
	BlockedUser  User         `json:"blocked_user"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// PageBuildPayload contains the information for GitHub's page_build hook event
type PageBuildPayload struct {
	ID    int64 `json:"id"`
	Build struct {
		URL    string `json:"url"`
		Status string `json:"status"`
		Error  struct {
			Message *string `json:"message"`
		} `json:"error"`
		Pusher    User      `json:"pusher"`
		Commit    string    `json:"commit"`
		Duration  int64     `json:"duration"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"build"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// PingPayload contains the information for GitHub's ping hook event
type PingPayload struct {
	HookID int `json:"hook_id"`
	Hook   struct {
		Type   string   `json:"type"`
		ID     int64    `json:"id"`
		NodeID string   `json:"node_id"`
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		AppID  int      `json:"app_id"`
		Config struct {
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
			Secret      string `json:"secret"`
			URL         string `json:"url"`
		} `json:"config"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"hook"`
	Repository Repository `json:"repository"`
	Sender     User       `json:"sender"`
}

// ProjectCardPayload contains the information for GitHub's project_payload hook event
type ProjectCardPayload struct {
	Action      string `json:"action"`
	ProjectCard struct {
		URL        string    `json:"url"`
		ColumnURL  string    `json:"column_url"`
		ColumnID   int64     `json:"column_id"`
		ID         int64     `json:"id"`
		NodeID     string    `json:"node_id"`
		Note       *string   `json:"note"`
		Archived   bool      `json:"archived"`
		ProjectURL string    `json:"project_url"`
		Creator    User      `json:"creator"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"project_card"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// ProjectColumnPayload contains the information for GitHub's project_column hook event
type ProjectColumnPayload struct {
	Action        string `json:"action"`
	ProjectColumn struct {
		URL        string    `json:"url"`
		ProjectURL string    `json:"project_url"`
		CardsURL   string    `json:"cards_url"`
		ID         int64     `json:"id"`
		NodeID     string    `json:"node_id"`
		Name       string    `json:"name"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"project_column"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// ProjectPayload contains the information for GitHub's project hook event
type ProjectPayload struct {
	Action  string `json:"action"`
	Project struct {
		OwnerURL   string    `json:"owner_url"`
		URL        string    `json:"url"`
		ColumnsURL string    `json:"columns_url"`
		ID         int64     `json:"id"`
		NodeID     string    `json:"node_id"`
		Name       string    `json:"name"`
		Body       string    `json:"body"`
		Number     int64     `json:"number"`
		State      string    `json:"state"`
		HtmlURL    string    `json:"html_url"`
		Creator    User      `json:"creator"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"project"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// PublicPayload contains the information for GitHub's public hook event
type PublicPayload struct {
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// PullRequestPayload contains the information for GitHub's pull_request hook event
type PullRequestPayload struct {
	Action      string `json:"action"`
	Number      int64  `json:"number"`
	PullRequest struct {
		URL                 string     `json:"url"`
		ID                  int64      `json:"id"`
		NodeID              string     `json:"node_id"`
		HTMLURL             string     `json:"html_url"`
		DiffURL             string     `json:"diff_url"`
		PatchURL            string     `json:"patch_url"`
		IssueURL            string     `json:"issue_url"`
		Number              int64      `json:"number"`
		State               string     `json:"state"`
		Locked              bool       `json:"locked"`
		Title               string     `json:"title"`
		AuthorAssociation   string     `json:"author_association"`
		Draft               bool       `json:"draft"`
		MaintainerCanModify bool       `json:"maintainer_can_modify"`
		User                User       `json:"user"`
		Body                string     `json:"body"`
		CreatedAt           time.Time  `json:"created_at"`
		UpdatedAt           time.Time  `json:"updated_at"`
		ClosedAt            *time.Time `json:"closed_at"`
		MergedAt            *time.Time `json:"merged_at"`
		MergeCommitSha      *string    `json:"merge_commit_sha"`
		Assignee            *User      `json:"assignee"`
		Assignees           []*User    `json:"assignees"`
		Milestone           *Milestone `json:"milestone"`
		CommitsURL          string     `json:"commits_url"`
		ReviewCommentsURL   string     `json:"review_comments_url"`
		ReviewCommentURL    string     `json:"review_comment_url"`
		CommentsURL         string     `json:"comments_url"`
		StatusesURL         string     `json:"statuses_url"`
		RequestedReviewers  []*User    `json:"requested_reviewers"`
		RequestedTeams      []*Team    `json:"requested_teams"`
		Labels              []*Label   `json:"labels"`
		Head                struct {
			Label string     `json:"label"`
			Ref   string     `json:"ref"`
			Sha   string     `json:"sha"`
			User  User       `json:"user"`
			Repo  Repository `json:"repo"`
		} `json:"head"`
		Base struct {
			Label string     `json:"label"`
			Ref   string     `json:"ref"`
			Sha   string     `json:"sha"`
			User  User       `json:"user"`
			Repo  Repository `json:"repo"`
		} `json:"base"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Issue struct {
				Href string `json:"href"`
			} `json:"issue"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			ReviewComments struct {
				Href string `json:"href"`
			} `json:"review_comments"`
			ReviewComment struct {
				Href string `json:"href"`
			} `json:"review_comment"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"_links"`
		Merged         bool   `json:"merged"`
		Mergeable      *bool  `json:"mergeable"`
		Rebaseable     *bool  `json:"rebaseable"`
		MergeableState string `json:"mergeable_state"`
		MergedBy       *User  `json:"merged_by"`
		Comments       int64  `json:"comments"`
		ReviewComments int64  `json:"review_comments"`
		Commits        int64  `json:"commits"`
		Additions      int64  `json:"additions"`
		Deletions      int64  `json:"deletions"`
		ChangedFiles   int64  `json:"changed_files"`
	} `json:"pull_request"`
	Repository Repository `json:"repository"`
	Sender     User       `json:"sender"`

	Installation Installation `json:"installation"`
}

// ReviewPullRequest is the PullRequest object contained in a pullrequest review payload
type ReviewPullRequest struct {
	URL                string     `json:"url"`
	ID                 int64      `json:"id"`
	NodeID             string     `json:"node_id"`
	HTMLURL            string     `json:"html_url"`
	DiffURL            string     `json:"diff_url"`
	PatchURL           string     `json:"patch_url"`
	IssueURL           string     `json:"issue_url"`
	Number             int64      `json:"number"`
	State              string     `json:"state"`
	Locked             bool       `json:"locked"`
	Title              string     `json:"title"`
	AuthorAssociation  string     `json:"author_association"`
	User               User       `json:"user"`
	Body               string     `json:"body"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	ClosedAt           *time.Time `json:"closed_at"`
	MergedAt           *time.Time `json:"merged_at"`
	MergeCommitSha     string     `json:"merge_commit_sha"`
	Assignee           *User      `json:"assignee"`
	Assignees          []User     `json:"assignees"`
	RequestedReviewers []*User    `json:"requested_reviewers"`
	RequestedTeams     []Team     `json:"requested_teams"`
	Milestone          *Milestone `json:"milestone"`
	Labels             []*Label   `json:"labels"`
	CommitsURL         string     `json:"commits_url"`
	ReviewCommentsURL  string     `json:"review_comments_url"`
	ReviewCommentURL   string     `json:"review_comment_url"`
	CommentsURL        string     `json:"comments_url"`
	StatusesURL        string     `json:"statuses_url"`
	Head               struct {
		Label string     `json:"label"`
		Ref   string     `json:"ref"`
		Sha   string     `json:"sha"`
		User  User       `json:"user"`
		Repo  Repository `json:"repo"`
	} `json:"head"`
	Base struct {
		Label string     `json:"label"`
		Ref   string     `json:"ref"`
		Sha   string     `json:"sha"`
		User  User       `json:"user"`
		Repo  Repository `json:"repo"`
	} `json:"base"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Issue struct {
			Href string `json:"href"`
		} `json:"issue"`
		Comments struct {
			Href string `json:"href"`
		} `json:"comments"`
		ReviewComments struct {
			Href string `json:"href"`
		} `json:"review_comments"`
		ReviewComment struct {
			Href string `json:"href"`
		} `json:"review_comment"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Statuses struct {
			Href string `json:"href"`
		} `json:"statuses"`
	} `json:"_links"`
}

// PullRequestReviewPayload contains the information for GitHub's pull_request_review hook event
type PullRequestReviewPayload struct {
	Action string `json:"action"`
	Review struct {
		ID                int64     `json:"id"`
		NodeID            string    `json:"node_id"`
		CommitID          string    `json:"commit_id"`
		AuthorAssociation string    `json:"author_association"`
		User              User      `json:"user"`
		Body              *string   `json:"body"`
		SubmittedAt       time.Time `json:"submitted_at"`
		State             string    `json:"state"`
		HTMLURL           string    `json:"html_url"`
		PullRequestURL    string    `json:"pull_request_url"`
		Links             struct {
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			PullRequest struct {
				Href string `json:"href"`
			} `json:"pull_request"`
		} `json:"_links"`
	} `json:"review"`
	PullRequest  ReviewPullRequest `json:"pull_request"`
	Repository   Repository        `json:"repository"`
	Sender       User              `json:"sender"`
	Installation Installation      `json:"installation"`
}

// PullRequestReviewCommentPayload contains the information for GitHub's pull_request_review_comments hook event
type PullRequestReviewCommentPayload struct {
	Action  string `json:"action"`
	Comment struct {
		URL               string    `json:"url"`
		ID                int64     `json:"id"`
		NodeID            string    `json:"node_id"`
		DiffHunk          string    `json:"diff_hunk"`
		Path              string    `json:"path"`
		Position          int64     `json:"position"`
		OriginalPosition  int64     `json:"original_position"`
		CommitID          string    `json:"commit_id"`
		OriginalCommitID  string    `json:"original_commit_id"`
		User              User      `json:"user"`
		Body              string    `json:"body"`
		AuthorAssociation string    `json:"author_association"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		HTMLURL           string    `json:"html_url"`
		PullRequestURL    string    `json:"pull_request_url"`
		Links             struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			PullRequest struct {
				Href string `json:"href"`
			} `json:"pull_request"`
		} `json:"_links"`
		//TODO this is not in the payload example on github docs
		// InReplyToID         int64 `json:"in_reply_to_id"`
		PullRequestReviewID int64 `json:"pull_request_review_id"`
	} `json:"comment"`
	PullRequest  ReviewPullRequest `json:"pull_request"`
	Repository   Repository        `json:"repository"`
	Sender       User              `json:"sender"`
	Installation Installation      `json:"installation"`
}

// PushPayload contains the information for GitHub's push hook event
type PushPayload struct {
	Ref        string         `json:"ref"`
	Before     string         `json:"before"`
	After      string         `json:"after"`
	Created    bool           `json:"created"`
	Deleted    bool           `json:"deleted"`
	Forced     bool           `json:"forced"`
	BaseRef    *string        `json:"base_ref"`
	Compare    string         `json:"compare"`
	Commits    []*Commit      `json:"commits"`
	HeadCommit *Commit        `json:"head_commit"`
	Repository PushRepository `json:"repository"`
	Pusher     struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// ReleasePayload contains the information for GitHub's release hook event
type ReleasePayload struct {
	Action  string `json:"action"`
	Release struct {
		URL             string    `json:"url"`
		AssetsURL       string    `json:"assets_url"`
		UploadURL       string    `json:"upload_url"`
		HTMLURL         string    `json:"html_url"`
		ID              int64     `json:"id"`
		NodeID          string    `json:"node_id"`
		TagName         string    `json:"tag_name"`
		TargetCommitish string    `json:"target_commitish"`
		Name            *string   `json:"name"`
		Draft           bool      `json:"draft"`
		Author          User      `json:"author"`
		Prerelease      bool      `json:"prerelease"`
		CreatedAt       time.Time `json:"created_at"`
		PublishedAt     time.Time `json:"published_at"`
		Assets          []Asset   `json:"assets"`
		TarballURL      string    `json:"tarball_url"`
		ZipballURL      string    `json:"zipball_url"`
		Body            *string   `json:"body"`
	} `json:"release"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// RepositoryPayload contains the information for GitHub's repository hook event
type RepositoryPayload struct {
	Action       string       `json:"action"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// RepositoryVulnerabilityAlertPayload contains the information for GitHub's repository_vulnerability_alert hook event.
type RepositoryVulnerabilityAlertPayload struct {
	Action string `json:"action"`
	Alert  struct {
		ID int64 `json:"id"`
		// TODO: is it really gone?
		// Summary             string `json:"summary"`
		AffectedRange       string `json:"affected_range"`
		AffectedPackageName string `json:"affected_package_name"`
		ExternalReference   string `json:"external_reference"`
		ExternalIdentifier  string `json:"external_identifier"`
		FixedIn             string `json:"fixed_in"`
	} `json:"alert"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// SecurityAdvisoryPayload contains the information for GitHub's security_advisory hook event.
type SecurityAdvisoryPayload struct {
	Action           string `json:"action"`
	SecurityAdvisory struct {
		GHSAID      string `json:"ghsa_id"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Severity    string `json:"severity"`
		Identifiers []struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"identifiers"`
		References []struct {
			URL string `json:"url"`
		} `json:"references"`
		PublishedAt     time.Time  `json:"published_at"`
		UpdatedAt       time.Time  `json:"updated_at"`
		WithdrawnAt     *time.Time `json:"withdrawn_at"`
		Vulnerabilities []struct {
			Package struct {
				Ecosystem string `json:"ecosystem"`
				Name      string `json:"name"`
			} `json:"package"`
			Severity               string `json:"severity"`
			VulnerableVersionRange string `json:"vulnerable_version_range"`
			FirstPatchedVersion    *struct {
				Identifier string `json:"identifier"`
			} `json:"first_patched_version"`
		} `json:"vulnerabilities"`
	} `json:"security_advisory"`
	Installation Installation `json:"installation"`
}

// StatusPayload contains the information for GitHub's status hook event
type StatusPayload struct {
	ID          int64   `json:"id"`
	Sha         string  `json:"sha"`
	Name        string  `json:"name"`
	TargetURL   *string `json:"target_url"`
	Context     string  `json:"context"`
	Description *string `json:"description"`
	State       string  `json:"state"`
	Commit      struct {
		Sha    string `json:"sha"`
		NodeID string `json:"node_id"`
		Commit struct {
			Author struct {
				Name  string    `json:"name"`
				Email string    `json:"email"`
				Date  time.Time `json:"date"`
			} `json:"author"`
			Committer struct {
				Name  string    `json:"name"`
				Email string    `json:"email"`
				Date  time.Time `json:"date"`
			} `json:"committer"`
			Message string `json:"message"`
			Tree    struct {
				Sha string `json:"sha"`
				URL string `json:"url"`
			} `json:"tree"`
			URL          string `json:"url"`
			CommentCount int64  `json:"comment_count"`
			Verification struct {
				Payload   string `json:"payload"`
				Signature string `json:"signature"`
				Reason    string `json:"reason"`
				Verified  bool   `json:"verified"`
			} `json:"verification"`
		} `json:"commit"`
		URL         string   `json:"url"`
		HTMLURL     string   `json:"html_url"`
		CommentsURL string   `json:"comments_url"`
		Author      User     `json:"author"`
		Committer   User     `json:"committer"`
		Parents     []Parent `json:"parents"`
	} `json:"commit"`
	Branches []struct {
		Name      string `json:"name"`
		Protected bool   `json:"protected"`
		Commit    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"commit"`
	} `json:"branches"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// TeamPayload contains the information for GitHub's team hook event
type TeamPayload struct {
	Action       string       `json:"action"`
	Team         Team         `json:"team"`
	Organization Organization `json:"organization"`
	Repository   struct {
		ID               int64     `json:"id"`
		NodeID           string    `json:"node_id"`
		Name             string    `json:"name"`
		FullName         string    `json:"full_name"`
		Owner            User      `json:"owner"`
		Private          bool      `json:"private"`
		Archived         bool      `json:"archived"`
		Disabled         bool      `json:"disabled"`
		HTMLURL          string    `json:"html_url"`
		Description      *string   `json:"description"`
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
		SvnURL           string    `json:"svn_url"`
		DeploymentsURL   string    `json:"deployments_url"`
		License          *string   `json:"license"`
		Homepage         *string   `json:"homepage"`
		Size             int64     `json:"size"`
		StargazersCount  int64     `json:"stargazers_count"`
		WatchersCount    int64     `json:"watchers_count"`
		Language         *string   `json:"language"`
		HasIssues        bool      `json:"has_issues"`
		HasDownloads     bool      `json:"has_downloads"`
		HasWiki          bool      `json:"has_wiki"`
		HasPages         bool      `json:"has_pages"`
		HasProjects      bool      `json:"has_projects"`
		ForksCount       int64     `json:"forks_count"`
		MirrorURL        *string   `json:"mirror_url"`
		OpenIssuesCount  int64     `json:"open_issues_count"`
		Forks            int64     `json:"forks"`
		OpenIssues       int64     `json:"open_issues"`
		Watchers         int64     `json:"watchers"`
		DefaultBranch    string    `json:"default_branch"`
		Permissions      struct {  //TODO: what other permissions exist?
			Admin bool `json:"admin"`
			Pull  bool `json:"pull"`
			Push  bool `json:"push"`
		} `json:"permissions"`
	} `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// TeamAddPayload contains the information for GitHub's team_add hook event
type TeamAddPayload struct {
	Team         Team         `json:"team"`
	Repository   Repository   `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// WatchPayload contains the information for GitHub's watch hook event
type WatchPayload struct {
	Action       string       `json:"action"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Installation Installation `json:"installation"`
}

// Organization contains Github's organization information
type Organization struct {
	Login            string `json:"login"`
	ID               int64  `json:"id"`
	NodeID           string `json:"node_id"`
	URL              string `json:"url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	HooksURL         string `json:"hooks_url"`
	IssuesURL        string `json:"issues_url"`
	MembersURL       string `json:"members_url"`
	PublicMembersURL string `json:"public_members_url"`
	AvatarURL        string `json:"avatar_url"`
	Description      string `json:"description"`
}

// Team containst Github's team information
type Team struct {
	Name            string `json:"name"`
	ID              int64  `json:"id"`
	NodeID          string `json:"node_id"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	Privacy         string `json:"privacy"`
	URL             string `json:"url"`
	HTMLURL         string `json:"html_url"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
	Permission      string `json:"permission"`
}

// User contains GitHub's user information
type User struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
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

// PushRepositoryOwner represents a user in a PushPayload which differs slightly from other user objects.
type PushRepositoryOwner struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
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
	Email             string `json:"email"`
	Name              string `json:"name"`
}

// Milestone contains GitHub's milestone information
type Milestone struct {
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	ID           int64      `json:"id"`
	NodeID       string     `json:"node_id"`
	Number       int64      `json:"number"`
	State        string     `json:"state"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Creator      User       `json:"creator"`
	OpenIssues   int64      `json:"open_issues"`
	ClosedIssues int64      `json:"closed_issues"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	ClosedAt     *time.Time `json:"closed_at"`
	DueOn        time.Time  `json:"due_on"`
}

// Asset contains GitHub's asset information
type Asset struct {
	URL                string    `json:"url"`
	BrowserDownloadURL string    `json:"browser_download_url"`
	ID                 int64     `json:"id"`
	NodeID             string    `json:"node_id"`
	Name               string    `json:"name"`
	Label              string    `json:"label"`
	State              string    `json:"state"`
	ContentType        string    `json:"content_type"`
	Size               int64     `json:"size"`
	DownloadCount      int64     `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Uploader           User      `json:"uploader"`
}

// Parent contains GitHub's parent information
type Parent struct {
	URL string `json:"url"`
	Sha string `json:"sha"`
}

// Label contains Issue's Label information
type Label struct {
	ID      int64  `json:"id"`
	NodeID  string `json:"node_id"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Default bool   `json:"default"`
}

// Repository contains Github's repository information
type Repository struct {
	ID               int64     `json:"id"`
	NodeID           string    `json:"node_id"`
	Name             string    `json:"name"`
	FullName         string    `json:"full_name"`
	Owner            User      `json:"owner"`
	Private          bool      `json:"private"`
	Archived         bool      `json:"archived"`
	Disabled         bool      `json:"disabled"`
	HTMLURL          string    `json:"html_url"`
	Description      *string   `json:"description"`
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
	SvnURL           string    `json:"svn_url"`
	DeploymentsURL   string    `json:"deployments_url"`
	License          *string   `json:"license"`
	Homepage         *string   `json:"homepage"`
	Size             int64     `json:"size"`
	StargazersCount  int64     `json:"stargazers_count"`
	WatchersCount    int64     `json:"watchers_count"`
	Language         *string   `json:"language"`
	HasIssues        bool      `json:"has_issues"`
	HasDownloads     bool      `json:"has_downloads"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	HasProjects      bool      `json:"has_projects"`
	ForksCount       int64     `json:"forks_count"`
	MirrorURL        *string   `json:"mirror_url"`
	OpenIssuesCount  int64     `json:"open_issues_count"`
	Forks            int64     `json:"forks"`
	OpenIssues       int64     `json:"open_issues"`
	Watchers         int64     `json:"watchers"`
	DefaultBranch    string    `json:"default_branch"`
}

// PushRepository represents the repository object of a PushPayload,
// CreatedAt and PushedAt are integer timestamps instead of formatted dates.
type PushRepository struct {
	ID               int64               `json:"id"`
	NodeID           string              `json:"node_id"`
	Name             string              `json:"name"`
	FullName         string              `json:"full_name"`
	Owner            PushRepositoryOwner `json:"owner"`
	Private          bool                `json:"private"`
	Archived         bool                `json:"archived"`
	Disabled         bool                `json:"disabled"`
	HTMLURL          string              `json:"html_url"`
	Description      *string             `json:"description"`
	Fork             bool                `json:"fork"`
	URL              string              `json:"url"`
	ForksURL         string              `json:"forks_url"`
	KeysURL          string              `json:"keys_url"`
	CollaboratorsURL string              `json:"collaborators_url"`
	TeamsURL         string              `json:"teams_url"`
	HooksURL         string              `json:"hooks_url"`
	IssueEventsURL   string              `json:"issue_events_url"`
	EventsURL        string              `json:"events_url"`
	AssigneesURL     string              `json:"assignees_url"`
	BranchesURL      string              `json:"branches_url"`
	TagsURL          string              `json:"tags_url"`
	BlobsURL         string              `json:"blobs_url"`
	GitTagsURL       string              `json:"git_tags_url"`
	GitRefsURL       string              `json:"git_refs_url"`
	TreesURL         string              `json:"trees_url"`
	StatusesURL      string              `json:"statuses_url"`
	LanguagesURL     string              `json:"languages_url"`
	StargazersURL    string              `json:"stargazers_url"`
	ContributorsURL  string              `json:"contributors_url"`
	SubscribersURL   string              `json:"subscribers_url"`
	SubscriptionURL  string              `json:"subscription_url"`
	CommitsURL       string              `json:"commits_url"`
	GitCommitsURL    string              `json:"git_commits_url"`
	CommentsURL      string              `json:"comments_url"`
	IssueCommentURL  string              `json:"issue_comment_url"`
	ContentsURL      string              `json:"contents_url"`
	CompareURL       string              `json:"compare_url"`
	MergesURL        string              `json:"merges_url"`
	ArchiveURL       string              `json:"archive_url"`
	DownloadsURL     string              `json:"downloads_url"`
	IssuesURL        string              `json:"issues_url"`
	PullsURL         string              `json:"pulls_url"`
	MilestonesURL    string              `json:"milestones_url"`
	NotificationsURL string              `json:"notifications_url"`
	LabelsURL        string              `json:"labels_url"`
	ReleasesURL      string              `json:"releases_url"`
	CreatedAt        int64               `json:"created_at"`
	UpdatedAt        time.Time           `json:"updated_at"`
	PushedAt         int64               `json:"pushed_at"`
	GitURL           string              `json:"git_url"`
	SSHURL           string              `json:"ssh_url"`
	CloneURL         string              `json:"clone_url"`
	SvnURL           string              `json:"svn_url"`
	DeploymentsURL   string              `json:"deployments_url"`
	License          *string             `json:"license"`
	Homepage         *string             `json:"homepage"`
	Size             int64               `json:"size"`
	StargazersCount  int64               `json:"stargazers_count"`
	WatchersCount    int64               `json:"watchers_count"`
	Language         *string             `json:"language"`
	HasIssues        bool                `json:"has_issues"`
	HasDownloads     bool                `json:"has_downloads"`
	HasWiki          bool                `json:"has_wiki"`
	HasPages         bool                `json:"has_pages"`
	HasProjects      bool                `json:"has_projects"`
	ForksCount       int64               `json:"forks_count"`
	MirrorURL        *string             `json:"mirror_url"`
	OpenIssuesCount  int64               `json:"open_issues_count"`
	Forks            int64               `json:"forks"`
	OpenIssues       int64               `json:"open_issues"`
	Watchers         int64               `json:"watchers"`
	DefaultBranch    string              `json:"default_branch"`
	MasterBranch     string              `json:"master_branch"`
	Stargazers       int                 `json:"stargazers"`
}

// Installation contains the ID of the installation of the app that recieved the payload
type Installation struct {
	ID     int64  `json:"id"`
	NodeID string `json:"node_id"`
}

// AppPermissions contains information about a Github app's permissions
type AppPermissions struct {
	Administration            string `json:"administration"`
	Checks                    string `json:"checks"`
	Contents                  string `json:"contents"`
	Deployments               string `json:"deployments"`
	Issues                    string `json:"issues"`
	Members                   string `json:"members"`
	Metadata                  string `json:"metadata"`
	OrganizatonAdministration string `json:"organization_administration"`
	OrganizationHooks         string `json:"organization_hooks"`
	OrganizationPlan          string `json:"organization_plan"`
	OrganizationProjects      string `json:"organization_projects"`
	OrganizationUserBlocking  string `json:"organization_user_blocking"`
	Pages                     string `json:"pages"`
	PullRequests              string `json:"pull_requests"`
	RepositoryHooks           string `json:"repository_hooks"`
	RepositoryProjects        string `json:"repository_projects"`
	Statuses                  string `json:"statuses"`
	TeamDiscussion            string `json:"team_discussions"`
	VulnerabilityAlerts       string `json:"vulnerability_alerts"`
}

// App containst the information about a Github app used in check payloads
type App struct {
	ID     int64  `json:"id"`
	NodeID string `json:"node_id"`
	Owner  struct {
		Login             string `json:"login"`
		ID                int64  `json:"id"`
		NodeID            string `json:"node_id"`
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
	} `json:"owner"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ExternalURL string         `json:"external_url"`
	HtmlURL     string         `json:"html_url"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	Permissions AppPermissions `json:"permissions"`
	Events      []string       `json:"events"` //TODO: check
}

// Issue contains the information about a Github issue
type Issue struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	HTMLURL       string `json:"html_url"`
	ID            int64  `json:"id"`
	NodeID        string `json:"node_id"`
	Number        int64  `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login             string `json:"login"`
		ID                int64  `json:"id"`
		NodeID            string `json:"node_id"`
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
	} `json:"user"`
	Labels            []Label    `json:"labels"`
	AuthorAssociation string     `json:"author_association"`
	State             string     `json:"state"`
	Locked            bool       `json:"locked"`
	Assignee          *User      `json:"assignee"`
	Assignees         []*User    `json:"assignees"`
	Milestone         *Milestone `json:"milestone"`
	Comments          int64      `json:"comments"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	ClosedAt          *time.Time `json:"closed_at"`
	Body              string     `json:"body"`
}

// Commit contains information about a commit in a push payload
type Commit struct {
	ID        string `json:"id"`
	TreeID    string `json:"tree_id"`
	Distinct  bool   `json:"distinct"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"author"`
	Committer struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"committer"`
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
	Modified []string `json:"modified"`
}

// Deployment contains Github's deployment information
type Deployment struct {
	URL                 string    `json:"url"`
	ID                  int64     `json:"id"`
	NodeID              string    `json:"node_id"`
	Sha                 string    `json:"sha"`
	Ref                 string    `json:"ref"`
	Task                string    `json:"task"`
	Payload             struct{}  `json:"payload"`
	Environment         string    `json:"environment"`
	OriginalEnvironment string    `json:"original_environment"`
	Description         *string   `json:"description"`
	Creator             User      `json:"creator"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	StatusesURL         string    `json:"statuses_url"`
	RepositoryURL       string    `json:"repository_url"`
}

// InstallationObj contains Github's installation information
type InstallationObj struct {
	ID                  int64          `json:"id"`
	Account             User           `json:"account"`
	RepositorySelection string         `json:"repository_selection"`
	AccessTokensURL     string         `json:"access_tokens_url"`
	RepositoriesURL     string         `json:"repositories_url"`
	HTMLURL             string         `json:"html_url"`
	AppID               int            `json:"app_id"`
	TargetID            int            `json:"target_id"`
	TargetType          string         `json:"target_type"`
	Permissions         AppPermissions `json:"permissions"`
	Events              []string       `json:"events"`
	CreatedAt           int64          `json:"created_at"`
	UpdatedAt           int64          `json:"updated_at"`
	SingleFileName      *string        `json:"single_file_name"`
}
