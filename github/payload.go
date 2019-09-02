package github

import "time"

// CheckRunPayload contains the information for GitHub's check_run hook event
type CheckRunPayload struct {
	Action   string `json:"action"`
	CheckRun struct {
		ID          int64     `json:"id"`
		NodeID      string    `json:"node_id"`
		Name        string    `json:"name"`
		HeadSHA     string    `json:"head_sha"`
		Status      string    `json:"status"`
		Conclusion  string    `json:"conclusion"`
		URL         string    `json:"url"`
		HtmlURL     string    `json:"html_url"`
		StarterAt   time.Time `json:"started_at"`
		CompletedAt time.Time `json:"completed_at"`
		Output      struct {
			Title            string `json:"title"`
			Summary          string `json:"summary"`
			Text             string `json:"text"`
			AnnotationsCount int64  `json:"annotations_count"`
			AnnotationsURL   string `json:"annotations_url"`
		}
		CheckSuite struct {
			ID           int64                `json:"id"`
			HeadBranch   string               `json:"head_branch"`
			HeadSHA      string               `json:"head_sha"`
			Status       string               `json:"status"`
			Conclusion   string               `json:"conclusion"`
			URL          string               `json:"url"`
			Before       string               `json:"before"`
			After        string               `json:"after"`
			PullRequests []PullRequestPayload `json:"pull_requests"`
			App          struct {
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
				Name        string `json:"name"`
				Description string `json:"description"`
				ExternalURL string `json:"external_url"`
				HtmlURL     string `json:"html_url"`
				CreatedAt   string `json:"created_at"`
				UpdatedAt   string `json:"updated_at"`
			} `json:"app"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"check_suite"`
		App struct {
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
			Name        string `json:"name"`
			Description string `json:"description"`
			ExternalURL string `json:"external_url"`
			HtmlURL     string `json:"html_url"`
			CreatedAt   string `json:"created_at"`
			UpdatedAt   string `json:"updated_at"`
		} `json:"app"`
		PullRequests []PullRequestPayload `json:"pull_requests"`
	} `json:"check_run"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// CheckSuitePayload contains the information for GitHub's check_suite hook event
type CheckSuitePayload struct {
	Action     string `json:"action"`
	CheckSuite struct {
		ID           int64                `json:"id"`
		NodeID       string               `json:"node_id"`
		HeadBranch   string               `json:"head_branch"`
		HeadSHA      string               `json:"head_sha"`
		Status       string               `json:"status"`
		Conclusion   string               `json:"conclusion"`
		URL          string               `json:"url"`
		Before       string               `json:"before"`
		After        string               `json:"after"`
		PullRequests []PullRequestPayload `json:"pull_requests"`
		App          struct {
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
			Name        string `json:"name"`
			Description string `json:"description"`
			ExternalURL string `json:"external_url"`
			HtmlURL     string `json:"html_url"`
			CreatedAt   string `json:"created_at"`
			UpdatedAt   string `json:"updated_at"`
		} `json:"app"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		LatestCheckRunsCount int64     `json:"latest_check_runs_count"`
		CheckRunsURL         string    `json:"check_runs_url"`
		HeadCommit           struct {
			ID        string    `json:"id"`
			TreeID    string    `json:"tree_id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
			Commiter struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"commiter"`
		} `json:"head_commit"`
	} `json:"check_suite"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// CommitCommentPayload contains the information for GitHub's commit_comment hook event
type CommitCommentPayload struct {
	Action  string `json:"action"`
	Comment struct {
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
		ID      int64  `json:"id"`
		NodeID  string `json:"node_id"`
		User    struct {
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
		Position          *int64    `json:"position"`
		Line              *int64    `json:"line"`
		Path              *string   `json:"path"`
		CommitID          string    `json:"commit_id"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Body              string    `json:"body"`
		AuthorAssociation string    `json:"author_association"`
	} `json:"comment"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// CreatePayload contains the information for GitHub's create hook event
type CreatePayload struct {
	Ref          string     `json:"ref"`
	RefType      string     `json:"ref_type"`
	MasterBranch string     `json:"master_branch"`
	Description  string     `json:"description"`
	PusherType   string     `json:"pusher_type"`
	Repository   Repository `json:"repository"`
	Sender       struct {
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
	} `json:"sender"`
}

// DeletePayload contains the information for GitHub's delete hook event
type DeletePayload struct {
	Ref        string     `json:"ref"`
	RefType    string     `json:"ref_type"`
	PusherType string     `json:"pusher_type"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// DeploymentPayload contains the information for GitHub's deployment hook
type DeploymentPayload struct {
	Deployment struct {
		URL         string   `json:"url"`
		ID          int64    `json:"id"`
		NodeID      string   `json:"node_id"`
		Sha         string   `json:"sha"`
		Ref         string   `json:"ref"`
		Task        string   `json:"task"`
		Payload     struct{} `json:"payload"`
		Environment string   `json:"environment"`
		Description *string  `json:"description"`
		Creator     struct {
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
		} `json:"creator"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		StatusesURL   string    `json:"statuses_url"`
		RepositoryURL string    `json:"repository_url"`
	} `json:"deployment"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// DeploymentStatusPayload contains the information for GitHub's deployment_status hook event
type DeploymentStatusPayload struct {
	DeploymentStatus struct {
		URL     string `json:"url"`
		ID      int64  `json:"id"`
		NodeID  string `json:"node_id"`
		State   string `json:"state"`
		Creator struct {
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
		} `json:"creator"`
		Description   *string   `json:"description"`
		TargetURL     *string   `json:"target_url"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		DeploymentURL string    `json:"deployment_url"`
		RepositoryURL string    `json:"repository_url"`
	} `json:"deployment_status"`
	Deployment struct {
		URL         string   `json:"url"`
		ID          int64    `json:"id"`
		NodeID      string   `json:"node_id"`
		Sha         string   `json:"sha"`
		Ref         string   `json:"ref"`
		Task        string   `json:"task"`
		Payload     struct{} `json:"payload"`
		Environment string   `json:"environment"`
		Description *string  `json:"description"`
		Creator     struct {
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
		} `json:"creator"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		StatusesURL   string    `json:"statuses_url"`
		RepositoryURL string    `json:"repository_url"`
	} `json:"deployment"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// ForkPayload contains the information for GitHub's fork hook event
type ForkPayload struct {
	Forkee struct {
		ID       int64  `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Owner    struct {
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
		SvnURL           string    `json:"svn_url"`
		Homepage         *string   `json:"homepage"`
		Size             int64     `json:"size"`
		StargazersCount  int64     `json:"stargazers_count"`
		WatchersCount    int64     `json:"watchers_count"`
		Language         *string   `json:"language"`
		HasIssues        bool      `json:"has_issues"`
		HasDownloads     bool      `json:"has_downloads"`
		HasWiki          bool      `json:"has_wiki"`
		HasPages         bool      `json:"has_pages"`
		ForksCount       int64     `json:"forks_count"`
		MirrorURL        *string   `json:"mirror_url"`
		OpenIssuesCount  int64     `json:"open_issues_count"`
		Forks            int64     `json:"forks"`
		OpenIssues       int64     `json:"open_issues"`
		Watchers         int64     `json:"watchers"`
		DefaultBranch    string    `json:"default_branch"`
		Public           bool      `json:"public"`
	} `json:"forkee"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
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
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// InstallationPayload contains the information for GitHub's installation and integration_installation hook events
type InstallationPayload struct {
	Action       string `json:"action"`
	Installation struct {
		ID      int64  `json:"id"`
		NodeID  string `json:"node_id"`
		Account struct {
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
		} `json:"account"`
		RepositorySelection string `json:"repository_selection"`
		AccessTokensURL     string `json:"access_tokens_url"`
		RepositoriesURL     string `json:"repositories_url"`
		HTMLURL             string `json:"html_url"`
		AppID               int    `json:"app_id"`
		TargetID            int    `json:"target_id"`
		TargetType          string `json:"target_type"`
		Permissions         struct {
			Issues             string `json:"issues"`
			Metadata           string `json:"metadata"`
			PullRequests       string `json:"pull_requests"`
			RepositoryProjects string `json:"repository_projects"`
		} `json:"permissions"`
		Events         []string `json:"events"`
		CreatedAt      int64    `json:"created_at"`
		UpdatedAt      int64    `json:"updated_at"`
		SingleFileName *string  `json:"single_file_name"`
	} `json:"installation"`
	Repositories []Repository `json:"repositories"`
	Sender       struct {
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
	} `json:"sender"`
}

// InstallationRepositoriesPayload contains the information for GitHub's installation_repositories hook events
type InstallationRepositoriesPayload struct {
	Action       string `json:"action"`
	Installation struct {
		ID      int64  `json:"id"`
		NodeID  string `json:"node_id"`
		Account struct {
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
		} `json:"account"`
		RepositorySelection string `json:"repository_selection"`
		AccessTokensURL     string `json:"access_tokens_url"`
		RepositoriesURL     string `json:"repositories_url"`
		HTMLURL             string `json:"html_url"`
		AppID               int    `json:"app_id"`
		TargetID            int    `json:"target_id"`
		TargetType          string `json:"target_type"`
		Permissions         struct {
			Issues              string `json:"issues"`
			Metadata            string `json:"metadata"`
			PullRequests        string `json:"pull_requests"`
			RepositoryProjects  string `json:"repository_projects"`
			VulnerabilityAlerts string `json:"vulnerability_alerts"`
			Statuses            string `json:"statuses"`
			Administration      string `json:"administration"`
			Deployments         string `json:"deployments"`
			Contents            string `json:"contents"`
		} `json:"permissions"`
		Events         []string `json:"events"`
		CreatedAt      int64    `json:"created_at"`
		UpdatedAt      int64    `json:"updated_at"`
		SingleFileName *string  `json:"single_file_name"`
	} `json:"installation"`
	RepositoriesAdded []struct {
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
	Sender struct {
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
	} `json:"sender"`
}

// IssueCommentPayload contains the information for GitHub's issue_comment hook event
type IssueCommentPayload struct {
	Action string `json:"action"`
	Issue  struct {
		URL         string `json:"url"`
		LabelsURL   string `json:"labels_url"`
		CommentsURL string `json:"comments_url"`
		EventsURL   string `json:"events_url"`
		HTMLURL     string `json:"html_url"`
		ID          int64  `json:"id"`
		NodeID      string `json:"node_id"`
		Number      int64  `json:"number"`
		Title       string `json:"title"`
		User        struct {
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
		Labels []struct {
			ID          int64  `json:"id"`
			NodeID      string `json:"node_id"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			Default     bool   `json:"default"`
		} `json:"labels"`
		State     string      `json:"state"`
		Locked    bool        `json:"locked"`
		Assignee  *Assignee   `json:"assignee"`
		Assignees []*Assignee `json:"assignees"`
		Milestone *Milestone  `json:"milestone"`
		Comments  int64       `json:"comments"`
		CreatedAt time.Time   `json:"created_at"`
		UpdatedAt time.Time   `json:"updated_at"`
		ClosedAt  *time.Time  `json:"closed_at"`
		Body      string      `json:"body"`
	} `json:"issue"`
	Comment struct {
		URL      string `json:"url"`
		HTMLURL  string `json:"html_url"`
		IssueURL string `json:"issue_url"`
		ID       int64  `json:"id"`
		NodeID   string `json:"node_id"`
		User     struct {
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
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Body              string    `json:"body"`
		AuthorAssociation string    `json:"author_association"`
	} `json:"comment"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// IssuesPayload contains the information for GitHub's issues hook event
type IssuesPayload struct {
	Action string `json:"action"`
	Issue  struct {
		URL         string `json:"url"`
		LabelsURL   string `json:"labels_url"`
		CommentsURL string `json:"comments_url"`
		EventsURL   string `json:"events_url"`
		HTMLURL     string `json:"html_url"`
		ID          int64  `json:"id"`
		NodeID      string `json:"node_id"`
		Number      int64  `json:"number"`
		Title       string `json:"title"`
		User        struct {
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
		Labels []struct {
			ID          int64  `json:"id"`
			NodeID      string `json:"node_id"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			Default     bool   `json:"default"`
		} `json:"labels"`
		State     string      `json:"state"`
		Locked    bool        `json:"locked"`
		Assignee  *Assignee   `json:"assignee"`
		Assignees []*Assignee `json:"assignees"`
		Milestone *Milestone  `json:"milestone"`
		Comments  int64       `json:"comments"`
		CreatedAt time.Time   `json:"created_at"`
		UpdatedAt time.Time   `json:"updated_at"`
		ClosedAt  *time.Time  `json:"closed_at"`
		Body      string      `json:"body"`
	} `json:"issue"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
	Assignee *Assignee `json:"assignee"`
	Label    *Label    `json:"label"`
}

// LabelPayload contains the information for GitHub's label hook event
type LabelPayload struct {
	Action string `json:"action"`
	Label  struct {
		ID          int64  `json:"id"`
		NodeID      string `json:"node_id"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Name        string `json:"name"`
		Color       string `json:"color"`
	} `json:"label"`
	Repository   Repository `json:"repository"`
	Organization struct {
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
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// MemberPayload contains the information for GitHub's member hook event
type MemberPayload struct {
	Action string `json:"action"`
	Member struct {
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
	} `json:"member"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// MembershipPayload contains the information for GitHub's membership hook event
type MembershipPayload struct {
	Action string `json:"action"`
	Scope  string `json:"scope"`
	Member struct {
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
	} `json:"member"`
	Sender struct {
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
	} `json:"sender"`
	Team struct {
		Name            string `json:"name"`
		ID              int64  `json:"id"`
		NodeID          string `json:"node_id"`
		Slug            string `json:"slug"`
		Permission      string `json:"permission"`
		URL             string `json:"url"`
		MembersURL      string `json:"members_url"`
		RepositoriesURL string `json:"repositories_url"`
	} `json:"team"`
	Organization struct {
		Login            string `json:"login"`
		ID               int64  `json:"id"`
		NodeID           string `json:"node_id"`
		URL              string `json:"url"`
		ReposURL         string `json:"repos_url"`
		EventsURL        string `json:"events_url"`
		MembersURL       string `json:"members_url"`
		PublicMembersURL string `json:"public_members_url"`
		AvatarURL        string `json:"avatar_url"`
	} `json:"organization"`
}

// MetaPayload contains the information for GitHub's meta hook event
type MetaPayload struct {
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
	Sender     struct {
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
	} `json:"sender"`
}

// MilestonePayload contains the information for GitHub's milestone hook event
type MilestonePayload struct {
	Action    string `json:"action"`
	Milestone struct {
		URL         string  `json:"url"`
		HTMLURL     string  `json:"html_url"`
		LabelsURL   string  `json:"labels_url"`
		ID          int64   `json:"id"`
		NodeID      string  `json:"node_id"`
		Number      int64   `json:"number"`
		Title       string  `json:"title"`
		Description *string `json:"description"`
		Creator     struct {
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
		} `json:"creator"`
		OpenIssues   int64      `json:"open_issues"`
		ClosedIssues int64      `json:"closed_issues"`
		State        string     `json:"state"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    time.Time  `json:"updated_at"`
		DueOn        *time.Time `json:"due_on"`
		ClosedAt     *time.Time `json:"closed_at"`
	} `json:"milestone"`
	Repository   Repository `json:"repository"`
	Organization struct {
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
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// OrganizationPayload contains the information for GitHub's organization hook event
type OrganizationPayload struct {
	Action     string `json:"action"`
	Invitation struct {
		ID     int64   `json:"id"`
		NodeID string  `json:"node_id"`
		Login  string  `json:"login"`
		Email  *string `json:"email"`
		Role   string  `json:"role"`
	} `json:"invitation"`
	Membership struct {
		URL             string `json:"url"`
		State           string `json:"state"`
		Role            string `json:"role"`
		OrganizationURL string `json:"organization_url"`
		User            struct {
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
	} `json:"membership"`
	Organization struct {
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
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// OrgBlockPayload contains the information for GitHub's org_block hook event
type OrgBlockPayload struct {
	Action      string `json:"action"`
	BlockedUser struct {
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
	} `json:"blocked_user"`
	Organization struct {
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
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// PageBuildPayload contains the information for GitHub's page_build hook event
type PageBuildPayload struct {
	ID     int64  `json:"id"`
	NodeID string `json:"node_id"`
	Build  struct {
		URL    string `json:"url"`
		Status string `json:"status"`
		Error  struct {
			Message *string `json:"message"`
		} `json:"error"`
		Pusher struct {
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
		} `json:"pusher"`
		Commit    string    `json:"commit"`
		Duration  int64     `json:"duration"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"build"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
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
	Sender     struct {
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
	} `json:"sender"`
}

// ProjectCardPayload contains the information for GitHub's project_payload hook event
type ProjectCardPayload struct {
	Action      string `json:"action"`
	ProjectCard struct {
		URL       string  `json:"url"`
		ColumnURL string  `json:"column_url"`
		ColumnID  int64   `json:"column_id"`
		ID        int64   `json:"id"`
		NodeID    string  `json:"node_id"`
		Note      *string `json:"note"`
		Creator   struct {
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
		} `json:"creator"`
		CreatedAt  int64  `json:"created_at"`
		UpdatedAt  int64  `json:"updated_at"`
		ContentURL string `json:"content_url"`
	} `json:"project_card"`
	Repository   Repository `json:"repository"`
	Organization struct {
		Login            string `json:"login"`
		ID               int64  `json:"id"`
		NodeID           string `json:"node_id"`
		URL              string `json:"url"`
		ReposURL         string `json:"repos_url"`
		EventsURL        string `json:"events_url"`
		MembersURL       string `json:"members_url"`
		PublicMembersURL string `json:"public_members_url"`
		AvatarURL        string `json:"avatar_url"`
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// ProjectColumnPayload contains the information for GitHub's project_column hook event
type ProjectColumnPayload struct {
	Action        string `json:"action"`
	ProjectColumn struct {
		URL        string `json:"url"`
		ProjectURL string `json:"project_url"`
		CardsURL   string `json:"cards_url"`
		ID         int64  `json:"id"`
		NodeID     string `json:"node_id"`
		Name       string `json:"name"`
		CreatedAt  int64  `json:"created_at"`
		UpdatedAt  int64  `json:"updated_at"`
	} `json:"project_column"`
	Repository   Repository `json:"repository"`
	Organization struct {
		Login            string `json:"login"`
		ID               int64  `json:"id"`
		NodeID           string `json:"node_id"`
		URL              string `json:"url"`
		ReposURL         string `json:"repos_url"`
		EventsURL        string `json:"events_url"`
		MembersURL       string `json:"members_url"`
		PublicMembersURL string `json:"public_members_url"`
		AvatarURL        string `json:"avatar_url"`
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// ProjectPayload contains the information for GitHub's project hook event
type ProjectPayload struct {
	Action  string `json:"action"`
	Project struct {
		OwnerURL   string `json:"owner_url"`
		URL        string `json:"url"`
		ColumnsURL string `json:"columns_url"`
		ID         int64  `json:"id"`
		NodeID     string `json:"node_id"`
		Name       string `json:"name"`
		Body       string `json:"body"`
		Number     int64  `json:"number"`
		State      string `json:"state"`
		Creator    struct {
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
		} `json:"creator"`
		CreatedAt int64 `json:"created_at"`
		UpdatedAt int64 `json:"updated_at"`
	} `json:"project"`
	Repository   Repository `json:"repository"`
	Organization struct {
		Login            string `json:"login"`
		ID               int64  `json:"id"`
		NodeID           string `json:"node_id"`
		URL              string `json:"url"`
		ReposURL         string `json:"repos_url"`
		EventsURL        string `json:"events_url"`
		MembersURL       string `json:"members_url"`
		PublicMembersURL string `json:"public_members_url"`
		AvatarURL        string `json:"avatar_url"`
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// PublicPayload contains the information for GitHub's public hook event
type PublicPayload struct {
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// PullRequestPayload contains the information for GitHub's pull_request hook event
type PullRequestPayload struct {
	Action      string `json:"action"`
	Number      int64  `json:"number"`
	PullRequest struct {
		URL      string `json:"url"`
		ID       int64  `json:"id"`
		NodeID   string `json:"node_id"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
		IssueURL string `json:"issue_url"`
		Number   int64  `json:"number"`
		State    string `json:"state"`
		Locked   bool   `json:"locked"`
		Title    string `json:"title"`
		User     struct {
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
		Body               string      `json:"body"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		ClosedAt           *time.Time  `json:"closed_at"`
		MergedAt           *time.Time  `json:"merged_at"`
		MergeCommitSha     *string     `json:"merge_commit_sha"`
		Assignee           *Assignee   `json:"assignee"`
		Assignees          []*Assignee `json:"assignees"`
		Milestone          *Milestone  `json:"milestone"`
		CommitsURL         string      `json:"commits_url"`
		ReviewCommentsURL  string      `json:"review_comments_url"`
		ReviewCommentURL   string      `json:"review_comment_url"`
		CommentsURL        string      `json:"comments_url"`
		StatusesURL        string      `json:"statuses_url"`
		RequestedReviewers []struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
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
		} `json:"requested_reviewers,omitempty"`
		Labels []struct {
			ID          int64  `json:"id"`
			NodeID      string `json:"node_id"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			Default     bool   `json:"default"`
		} `json:"labels"`
		Head struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
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
			Repo struct {
				ID       int64  `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
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
				SvnURL           string    `json:"svn_url"`
				Homepage         *string   `json:"homepage"`
				Size             int64     `json:"size"`
				StargazersCount  int64     `json:"stargazers_count"`
				WatchersCount    int64     `json:"watchers_count"`
				Language         *string   `json:"language"`
				HasIssues        bool      `json:"has_issues"`
				HasDownloads     bool      `json:"has_downloads"`
				HasWiki          bool      `json:"has_wiki"`
				HasPages         bool      `json:"has_pages"`
				ForksCount       int64     `json:"forks_count"`
				MirrorURL        *string   `json:"mirror_url"`
				OpenIssuesCount  int64     `json:"open_issues_count"`
				Forks            int64     `json:"forks"`
				OpenIssues       int64     `json:"open_issues"`
				Watchers         int64     `json:"watchers"`
				DefaultBranch    string    `json:"default_branch"`
			} `json:"repo"`
		} `json:"head"`
		Base struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
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
			Repo struct {
				ID       int64  `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
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
				SvnURL           string    `json:"svn_url"`
				Homepage         *string   `json:"homepage"`
				Size             int64     `json:"size"`
				StargazersCount  int64     `json:"stargazers_count"`
				WatchersCount    int64     `json:"watchers_count"`
				Language         *string   `json:"language"`
				HasIssues        bool      `json:"has_issues"`
				HasDownloads     bool      `json:"has_downloads"`
				HasWiki          bool      `json:"has_wiki"`
				HasPages         bool      `json:"has_pages"`
				ForksCount       int64     `json:"forks_count"`
				MirrorURL        *string   `json:"mirror_url"`
				OpenIssuesCount  int64     `json:"open_issues_count"`
				Forks            int64     `json:"forks"`
				OpenIssues       int64     `json:"open_issues"`
				Watchers         int64     `json:"watchers"`
				DefaultBranch    string    `json:"default_branch"`
			} `json:"repo"`
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
		Merged         bool      `json:"merged"`
		Mergeable      *bool     `json:"mergeable"`
		MergeableState string    `json:"mergeable_state"`
		MergedBy       *MergedBy `json:"merged_by"`
		Comments       int64     `json:"comments"`
		ReviewComments int64     `json:"review_comments"`
		Commits        int64     `json:"commits"`
		Additions      int64     `json:"additions"`
		Deletions      int64     `json:"deletions"`
		ChangedFiles   int64     `json:"changed_files"`
	} `json:"pull_request"`
	Label struct {
		ID          int64  `json:"id"`
		NodeID      string `json:"node_id"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Name        string `json:"name"`
		Color       string `json:"color"`
		Default     bool   `json:"default"`
	} `json:"label"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
	Assignee          *Assignee `json:"assignee"`
	RequestedReviewer *Assignee `json:"requested_reviewer"`
	RequestedTeam     struct {
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
	} `json:"requested_team"`
	Installation struct {
		ID int64 `json:"id"`
	} `json:"installation"`
}

// PullRequestReviewPayload contains the information for GitHub's pull_request_review hook event
type PullRequestReviewPayload struct {
	Action string `json:"action"`
	Review struct {
		ID     int64  `json:"id"`
		NodeID string `json:"node_id"`
		User   struct {
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
		Body           string    `json:"body"`
		SubmittedAt    time.Time `json:"submitted_at"`
		State          string    `json:"state"`
		HTMLURL        string    `json:"html_url"`
		PullRequestURL string    `json:"pull_request_url"`
		Links          struct {
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			PullRequest struct {
				Href string `json:"href"`
			} `json:"pull_request"`
		} `json:"_links"`
	} `json:"review"`
	PullRequest struct {
		URL      string `json:"url"`
		ID       int64  `json:"id"`
		NodeID   string `json:"node_id"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
		IssueURL string `json:"issue_url"`
		Number   int64  `json:"number"`
		State    string `json:"state"`
		Locked   bool   `json:"locked"`
		Title    string `json:"title"`
		User     struct {
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
		Body              string     `json:"body"`
		CreatedAt         time.Time  `json:"created_at"`
		UpdatedAt         time.Time  `json:"updated_at"`
		ClosedAt          *time.Time `json:"closed_at"`
		MergedAt          *time.Time `json:"merged_at"`
		MergeCommitSha    string     `json:"merge_commit_sha"`
		Assignee          *Assignee  `json:"assignee"`
		Assignees         []Assignee `json:"assignees"`
		Milestone         *Milestone `json:"milestone"`
		CommitsURL        string     `json:"commits_url"`
		ReviewCommentsURL string     `json:"review_comments_url"`
		ReviewCommentURL  string     `json:"review_comment_url"`
		CommentsURL       string     `json:"comments_url"`
		StatusesURL       string     `json:"statuses_url"`
		Head              struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
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
			Repo struct {
				ID       int64  `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
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
				Private          bool      `json:"private"`
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
				DeploymentsURL   string    `json:"deployments_url"`
				CreatedAt        time.Time `json:"created_at"`
				UpdatedAt        time.Time `json:"updated_at"`
				PushedAt         time.Time `json:"pushed_at"`
				GitURL           string    `json:"git_url"`
				SSHURL           string    `json:"ssh_url"`
				CloneURL         string    `json:"clone_url"`
				SvnURL           string    `json:"svn_url"`
				Homepage         *string   `json:"homepage"`
				Size             int64     `json:"size"`
				StargazersCount  int64     `json:"stargazers_count"`
				WatchersCount    int64     `json:"watchers_count"`
				Language         *string   `json:"language"`
				HasIssues        bool      `json:"has_issues"`
				HasDownloads     bool      `json:"has_downloads"`
				HasWiki          bool      `json:"has_wiki"`
				HasPages         bool      `json:"has_pages"`
				ForksCount       int64     `json:"forks_count"`
				MirrorURL        *string   `json:"mirror_url"`
				OpenIssuesCount  int64     `json:"open_issues_count"`
				Forks            int64     `json:"forks"`
				OpenIssues       int64     `json:"open_issues"`
				Watchers         int64     `json:"watchers"`
				DefaultBranch    string    `json:"default_branch"`
			} `json:"repo"`
		} `json:"head"`
		Base struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
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
			Repo struct {
				ID       int64  `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
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
				DeploymentsURL   string    `json:"deployments_url"`
				CreatedAt        time.Time `json:"created_at"`
				UpdatedAt        time.Time `json:"updated_at"`
				PushedAt         time.Time `json:"pushed_at"`
				GitURL           string    `json:"git_url"`
				SSHURL           string    `json:"ssh_url"`
				CloneURL         string    `json:"clone_url"`
				SvnURL           string    `json:"svn_url"`
				Homepage         *string   `json:"homepage"`
				Size             int64     `json:"size"`
				StargazersCount  int64     `json:"stargazers_count"`
				WatchersCount    int64     `json:"watchers_count"`
				Language         *string   `json:"language"`
				HasIssues        bool      `json:"has_issues"`
				HasDownloads     bool      `json:"has_downloads"`
				HasWiki          bool      `json:"has_wiki"`
				HasPages         bool      `json:"has_pages"`
				ForksCount       int64     `json:"forks_count"`
				MirrorURL        *string   `json:"mirror_url"`
				OpenIssuesCount  int64     `json:"open_issues_count"`
				Forks            int64     `json:"forks"`
				OpenIssues       int64     `json:"open_issues"`
				Watchers         int64     `json:"watchers"`
				DefaultBranch    string    `json:"default_branch"`
			} `json:"repo"`
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
	} `json:"pull_request"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
	Installation struct {
		ID int64 `json:"id"`
	} `json:"installation"`
}

// PullRequestReviewCommentPayload contains the information for GitHub's pull_request_review_comments hook event
type PullRequestReviewCommentPayload struct {
	Action  string `json:"action"`
	Comment struct {
		URL              string `json:"url"`
		ID               int64  `json:"id"`
		NodeID           string `json:"node_id"`
		DiffHunk         string `json:"diff_hunk"`
		Path             string `json:"path"`
		Position         int64  `json:"position"`
		OriginalPosition int64  `json:"original_position"`
		CommitID         string `json:"commit_id"`
		OriginalCommitID string `json:"original_commit_id"`
		User             struct {
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
		InReplyToID int64 `json:"in_reply_to_id"`
	} `json:"comment"`
	PullRequest struct {
		URL      string `json:"url"`
		ID       int64  `json:"id"`
		NodeID   string `json:"node_id"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
		IssueURL string `json:"issue_url"`
		Number   int64  `json:"number"`
		State    string `json:"state"`
		Locked   bool   `json:"locked"`
		Title    string `json:"title"`
		User     struct {
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
		Body              string      `json:"body"`
		CreatedAt         time.Time   `json:"created_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
		ClosedAt          *time.Time  `json:"closed_at"`
		MergedAt          *time.Time  `json:"merged_at"`
		MergeCommitSha    string      `json:"merge_commit_sha"`
		Assignee          *Assignee   `json:"assignee"`
		Assignees         []*Assignee `json:"assignees"`
		Milestone         *Milestone  `json:"milestone"`
		CommitsURL        string      `json:"commits_url"`
		ReviewCommentsURL string      `json:"review_comments_url"`
		ReviewCommentURL  string      `json:"review_comment_url"`
		CommentsURL       string      `json:"comments_url"`
		StatusesURL       string      `json:"statuses_url"`
		Head              struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
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
			Repo struct {
				ID       int64  `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
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
				SvnURL           string    `json:"svn_url"`
				Homepage         *string   `json:"homepage"`
				Size             int64     `json:"size"`
				StargazersCount  int64     `json:"stargazers_count"`
				WatchersCount    int64     `json:"watchers_count"`
				Language         *string   `json:"language"`
				HasIssues        bool      `json:"has_issues"`
				HasDownloads     bool      `json:"has_downloads"`
				HasWiki          bool      `json:"has_wiki"`
				HasPages         bool      `json:"has_pages"`
				ForksCount       int64     `json:"forks_count"`
				MirrorURL        *string   `json:"mirror_url"`
				OpenIssuesCount  int64     `json:"open_issues_count"`
				Forks            int64     `json:"forks"`
				OpenIssues       int64     `json:"open_issues"`
				Watchers         int64     `json:"watchers"`
				DefaultBranch    string    `json:"default_branch"`
			} `json:"repo"`
		} `json:"head"`
		Base struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
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
			Repo struct {
				ID       int64  `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
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
				SvnURL           string    `json:"svn_url"`
				Homepage         *string   `json:"homepage"`
				Size             int64     `json:"size"`
				StargazersCount  int64     `json:"stargazers_count"`
				WatchersCount    int64     `json:"watchers_count"`
				Language         *string   `json:"language"`
				HasIssues        bool      `json:"has_issues"`
				HasDownloads     bool      `json:"has_downloads"`
				HasWiki          bool      `json:"has_wiki"`
				HasPages         bool      `json:"has_pages"`
				ForksCount       int64     `json:"forks_count"`
				MirrorURL        *string   `json:"mirror_url"`
				OpenIssuesCount  int64     `json:"open_issues_count"`
				Forks            int64     `json:"forks"`
				OpenIssues       int64     `json:"open_issues"`
				Watchers         int64     `json:"watchers"`
				DefaultBranch    string    `json:"default_branch"`
			} `json:"repo"`
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
	} `json:"pull_request"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
	Installation struct {
		ID int64 `json:"id"`
	} `json:"installation"`
}

// PushPayload contains the information for GitHub's push hook event
type PushPayload struct {
	Ref     string  `json:"ref"`
	Before  string  `json:"before"`
	After   string  `json:"after"`
	Created bool    `json:"created"`
	Deleted bool    `json:"deleted"`
	Forced  bool    `json:"forced"`
	BaseRef *string `json:"base_ref"`
	Compare string  `json:"compare"`
	Commits []struct {
		Sha       string `json:"sha"`
		ID        string `json:"id"`
		NodeID    string `json:"node_id"`
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
	} `json:"commits"`
	HeadCommit struct {
		ID        string `json:"id"`
		NodeID    string `json:"node_id"`
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
	} `json:"head_commit"`
	Repository Repository `json:"repository"`
	Pusher     struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	Sender struct {
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
	} `json:"sender"`
	Installation struct {
		ID int `json:"id"`
	} `json:"installation"`
}

// ReleasePayload contains the information for GitHub's release hook event
type ReleasePayload struct {
	Action  string `json:"action"`
	Release struct {
		URL             string  `json:"url"`
		AssetsURL       string  `json:"assets_url"`
		UploadURL       string  `json:"upload_url"`
		HTMLURL         string  `json:"html_url"`
		ID              int64   `json:"id"`
		NodeID          string  `json:"node_id"`
		TagName         string  `json:"tag_name"`
		TargetCommitish string  `json:"target_commitish"`
		Name            *string `json:"name"`
		Draft           bool    `json:"draft"`
		Author          struct {
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
		} `json:"author"`
		Prerelease  bool      `json:"prerelease"`
		CreatedAt   time.Time `json:"created_at"`
		PublishedAt time.Time `json:"published_at"`
		Assets      []Asset   `json:"assets"`
		TarballURL  string    `json:"tarball_url"`
		ZipballURL  string    `json:"zipball_url"`
		Body        *string   `json:"body"`
	} `json:"release"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
	Installation struct {
		ID int `json:"id"`
	} `json:"installation"`
}

// RepositoryPayload contains the information for GitHub's repository hook event
type RepositoryPayload struct {
	Action       string     `json:"action"`
	Repository   Repository `json:"repository"`
	Organization struct {
		Login            string `json:"login"`
		ID               int64  `json:"id"`
		NodeID           string `json:"node_id"`
		URL              string `json:"url"`
		ReposURL         string `json:"repos_url"`
		EventsURL        string `json:"events_url"`
		MembersURL       string `json:"members_url"`
		PublicMembersURL string `json:"public_members_url"`
		AvatarURL        string `json:"avatar_url"`
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// RepositoryVulnerabilityAlertPayload contains the information for GitHub's repository_vulnerability_alert hook event.
type RepositoryVulnerabilityAlertPayload struct {
	Action string `json:"action"`
	Alert  struct {
		ID                  int64  `json:"id"`
		Summary             string `json:"summary"`
		AffectedRange       string `json:"affected_range"`
		AffectedPackageName string `json:"affected_package_name"`
		ExternalReference   string `json:"external_reference"`
		ExternalIdentifier  string `json:"external_identifier"`
		FixedIn             string `json:"fixed_in"`
		Dismisser           struct {
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
		} `json:"dismisser"`
	} `json:"alert"`
}

// SecurityAdvisoryPayload contains the information for GitHub's security_advisory hook event.
type SecurityAdvisoryPayload struct {
	Action           string `json:"action"`
	SecurityAdvisory struct {
		GHSAID      string `json:"ghsa_id"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Severity    string `json:"string"`
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
			}
			Severity               string `json:"severity"`
			VulnerableVersionRange string `json:"vulnerable_version_range"`
			FirstPatchedVersion    *struct {
				Identifier string `json:"identifier"`
			} `json:"first_patched_version"`
		} `json:"vulnerabilities"`
	} `json:"security_advisory"`
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
		} `json:"commit"`
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
		CommentsURL string `json:"comments_url"`
		Author      struct {
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
		} `json:"author"`
		Committer struct {
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
		} `json:"committer"`
		Parents []Parent `json:"parents"`
	} `json:"commit"`
	Branches []struct {
		Name   string `json:"name"`
		Commit struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"commit"`
	} `json:"branches"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// TeamPayload contains the information for GitHub's team hook event
type TeamPayload struct {
	Action string `json:"action"`
	Team   struct {
		Name            string `json:"name"`
		ID              int64  `json:"id"`
		NodeID          string `json:"node_id"`
		Slug            string `json:"slug"`
		Description     string `json:"description"`
		Privacy         string `json:"privacy"`
		URL             string `json:"url"`
		MembersURL      string `json:"members_url"`
		RepositoriesURL string `json:"repositories_url"`
		Permission      string `json:"permission"`
	} `json:"team"`
	Organization struct {
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
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// TeamAddPayload contains the information for GitHub's team_add hook event
type TeamAddPayload struct {
	Team struct {
		Name            string `json:"name"`
		ID              int64  `json:"id"`
		NodeID          string `json:"node_id"`
		Slug            string `json:"slug"`
		Description     string `json:"description"`
		Permission      string `json:"permission"`
		URL             string `json:"url"`
		MembersURL      string `json:"members_url"`
		RepositoriesURL string `json:"repositories_url"`
	} `json:"team"`
	Repository   Repository `json:"repository"`
	Organization struct {
		Login            string  `json:"login"`
		ID               int64   `json:"id"`
		NodeID           string  `json:"node_id"`
		URL              string  `json:"url"`
		ReposURL         string  `json:"repos_url"`
		EventsURL        string  `json:"events_url"`
		MembersURL       string  `json:"members_url"`
		PublicMembersURL string  `json:"public_members_url"`
		AvatarURL        string  `json:"avatar_url"`
		Description      *string `json:"description"`
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}

// WatchPayload contains the information for GitHub's watch hook event
type WatchPayload struct {
	Action     string     `json:"action"`
	Repository Repository `json:"repository"`
	Sender     struct {
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
	} `json:"sender"`
}

// Assignee contains GitHub's assignee information
type Assignee struct {
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

// Milestone contains GitHub's milestone information
type Milestone struct {
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	LabelsURL   string `json:"labels_url"`
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	Number      int64  `json:"number"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     struct {
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
	} `json:"creator"`
	OpenIssues   int64     `json:"open_issues"`
	ClosedIssues int64     `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosedAt     time.Time `json:"closed_at"`
	DueOn        time.Time `json:"due_on"`
}

// MergedBy contains GitHub's merged-by information
type MergedBy struct {
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
	Uploader           struct {
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
	} `json:"uploader"`
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

// Repository contains Repository information
type Repository struct {
	ID               int64     `json:"id"`
	NodeID           string    `json:"node_id"`
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
	SvnURL           string    `json:"svn_url"`
	Homepage         *string   `json:"homepage"`
	Size             int64     `json:"size"`
	StargazersCount  int64     `json:"stargazers_count"`
	WatchersCount    int64     `json:"watchers_count"`
	Language         *string   `json:"language"`
	HasIssues        bool      `json:"has_issues"`
	HasDownloads     bool      `json:"has_downloads"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	ForksCount       int64     `json:"forks_count"`
	MirrorURL        *string   `json:"mirror_url"`
	OpenIssuesCount  int64     `json:"open_issues_count"`
	Forks            int64     `json:"forks"`
	OpenIssues       int64     `json:"open_issues"`
	Watchers         int64     `json:"watchers"`
	DefaultBranch    string    `json:"default_branch"`
}

// Owner contains Repository Owner information
type Owner struct {
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
