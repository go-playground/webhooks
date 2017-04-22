package bitbucket

import "time"

// RepoPushPayload is the Bitbucket repo:push payload
type RepoPushPayload struct {
	Actor      Owner      `json:"actor"`
	Repository Repository `json:"repository"`
	Push       struct {
		Changes []struct {
			New struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Target struct {
					Type    string    `json:"type"`
					Hash    string    `json:"hash"`
					Author  Owner     `json:"author"`
					Message string    `json:"message"`
					Date    time.Time `json:"date"`
					Parents []struct {
						Type  string `json:"type"`
						Hash  string `json:"hash"`
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
						} `json:"links"`
					} `json:"parents"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
				} `json:"target"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"new"`
			Old struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Target struct {
					Type    string    `json:"type"`
					Hash    string    `json:"hash"`
					Author  Owner     `json:"author"`
					Message string    `json:"message"`
					Date    time.Time `json:"date"`
					Parents []struct {
						Type  string `json:"type"`
						Hash  string `json:"hash"`
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
						} `json:"links"`
					} `json:"parents"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
				} `json:"target"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"old"`
			Links struct {
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Diff struct {
					Href string `json:"href"`
				} `json:"diff"`
				Commits struct {
					Href string `json:"href"`
				} `json:"commits"`
			} `json:"links"`
			Created bool `json:"created"`
			Forced  bool `json:"forced"`
			Closed  bool `json:"closed"`
			Commits []struct {
				Hash    string `json:"hash"`
				Type    string `json:"type"`
				Message string `json:"message"`
				Author  Owner  `json:"author"`
				Links   struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"commits"`
			Truncated bool `json:"truncated"`
		} `json:"changes"`
	} `json:"push"`
}

// RepoForkPayload is the Bitbucket repo:fork payload
type RepoForkPayload struct {
	Actor      Owner      `json:"actor"`
	Repository Repository `json:"repository"`
	Fork       Repository `json:"fork"`
}

// RepoUpdatedPayload is the Bitbucket repo:updated payload
type RepoUpdatedPayload struct {
	Actor      Owner      `json:"actor"`
	Repository Repository `json:"repository"`
	Changes    struct {
		Name struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"name"`
		Website struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"website"`
		Language struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"language"`
		Links struct {
			New struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"new"`
			Old struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"old"`
		} `json:"links"`
		Description struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"description"`
		FullName struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"full_name"`
	} `json:"changes"`
}

// RepoCommitCommentCreatedPayload is the Bitbucket repo:commit_comment_created payload
type RepoCommitCommentCreatedPayload struct {
	Actor      Owner      `json:"actor"`
	Comment    Comment    `json:"comment"`
	Repository Repository `json:"repository"`
	Commit     struct {
		Hash string `json:"hash"`
	} `json:"commit"`
}

// RepoCommitStatusCreatedPayload is the Bitbucket repo:commit_status_created payload
type RepoCommitStatusCreatedPayload struct {
	Actor        Owner      `json:"actor"`
	Repository   Repository `json:"repository"`
	CommitStatus struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		State       string    `json:"state"`
		Key         string    `json:"key"`
		URL         string    `json:"url"`
		Type        string    `json:"type"`
		CreatedOn   time.Time `json:"created_on"`
		UpdatedOn   time.Time `json:"updated_on"`
		Links       struct {
			Commit struct {
				Href string `json:"href"`
			} `json:"commit"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"commit_status"`
}

// RepoCommitStatusUpdatedPayload is the Bitbucket repo:commit_status_updated payload
type RepoCommitStatusUpdatedPayload struct {
	Actor        Owner      `json:"actor"`
	Repository   Repository `json:"repository"`
	CommitStatus struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		State       string    `json:"state"`
		Key         string    `json:"key"`
		URL         string    `json:"url"`
		Type        string    `json:"type"`
		CreatedOn   time.Time `json:"created_on"`
		UpdatedOn   time.Time `json:"updated_on"`
		Links       struct {
			Commit struct {
				Href string `json:"href"`
			} `json:"commit"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"commit_status"`
}

// IssueCreatedPayload is the Bitbucket issue:created payload
type IssueCreatedPayload struct {
	Actor      Owner      `json:"actor"`
	Issue      Issue      `json:"issue"`
	Repository Repository `json:"repository"`
}

// IssueUpdatedPayload is the Bitbucket issue:updated payload
type IssueUpdatedPayload struct {
	Actor      Owner      `json:"actor"`
	Issue      Issue      `json:"issue"`
	Repository Repository `json:"repository"`
	Comment    Comment    `json:"comment"`
	Changes    struct {
		Status struct {
			Old string `json:"old"`
			New string `json:"new"`
		} `json:"status"`
	} `json:"changes"`
}

// IssueCommentCreatedPayload is the Bitbucket pullrequest:created payload
type IssueCommentCreatedPayload struct {
	Actor      Owner      `json:"actor"`
	Repository Repository `json:"repository"`
	Issue      Issue      `json:"issue"`
	Comment    Comment    `json:"comment"`
}

// PullRequestCreatedPayload is the Bitbucket pullrequest:created payload
type PullRequestCreatedPayload struct {
	Actor       Owner       `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
}

// PullRequestUpdatedPayload is the Bitbucket pullrequest:updated payload
type PullRequestUpdatedPayload struct {
	Actor       Owner       `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
}

// PullRequestApprovedPayload is the Bitbucket pullrequest:approved payload
type PullRequestApprovedPayload struct {
	Actor       Owner       `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
	Approval    struct {
		Date time.Time `json:"date"`
		User Owner     `json:"user"`
	} `json:"approval"`
}

// PullRequestUnapprovedPayload is the Bitbucket pullrequest:unapproved payload
type PullRequestUnapprovedPayload struct {
	Actor       Owner       `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
	Approval    struct {
		Date time.Time `json:"date"`
		User Owner     `json:"user"`
	} `json:"approval"`
}

// PullRequestMergedPayload is the Bitbucket pullrequest:fulfilled payload
type PullRequestMergedPayload struct {
	Actor       Owner       `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
}

// PullRequestDeclinedPayload is the Bitbucket pullrequest:rejected payload
type PullRequestDeclinedPayload struct {
	Actor       Owner       `json:"actor"`
	PullRequest PullRequest `json:"pullrequest"`
	Repository  Repository  `json:"repository"`
}

// PullRequestCommentCreatedPayload is the Bitbucket pullrequest:comment_updated payload
type PullRequestCommentCreatedPayload struct {
	Actor       Owner       `json:"actor"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pullrequest"`
	Comment     Comment     `json:"comment"`
}

// PullRequestCommentUpdatedPayload is the Bitbucket pullrequest:comment_created payload
type PullRequestCommentUpdatedPayload struct {
	Actor       Owner       `json:"actor"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pullrequest"`
	Comment     Comment     `json:"comment"`
}

// PullRequestCommentDeletedPayload is the Bitbucket pullrequest:comment_deleted payload
type PullRequestCommentDeletedPayload struct {
	Actor       Owner       `json:"actor"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pullrequest"`
	Comment     Comment     `json:"comment"`
}

// Owner is the common Bitbucket Owner Sub Entity
type Owner struct {
	Type        string `json:"type"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	UUID        string `json:"uuid"`
	Links       struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
	} `json:"links"`
}

// Repository is the common Bitbucket Repository Sub Entity
type Repository struct {
	Type  string `json:"type"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
	} `json:"links"`
	UUID      string  `json:"uuid"`
	Project   Project `json:"project"`
	FullName  string  `json:"full_name"`
	Name      string  `json:"name"`
	Website   string  `json:"website"`
	Owner     Owner   `json:"owner"`
	Scm       string  `json:"scm"`
	IsPrivate bool    `json:"is_private"`
}

// Project is the common Bitbucket Project Sub Entity
type Project struct {
	Type    string `json:"type"`
	Project string `json:"project"`
	UUID    string `json:"uuid"`
	Links   struct {
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
	} `json:"links"`
	Key string `json:"key"`
}

// Issue is the common Bitbucket Issue Sub Entity
type Issue struct {
	ID        int64  `json:"id"`
	Component string `json:"component"`
	Title     string `json:"title"`
	Content   struct {
		Raw    string `json:"raw"`
		HTML   string `json:"html"`
		Markup string `json:"markup"`
	} `json:"content"`
	Priority  string `json:"priority"`
	State     string `json:"state"`
	Type      string `json:"type"`
	Milestone struct {
		Name string `json:"name"`
	} `json:"milestone"`
	Version struct {
		Name string `json:"name"`
	} `json:"version"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
}

// Comment is the common Bitbucket Comment Sub Entity
type Comment struct {
	ID     int64 `json:"id"`
	Parent struct {
		ID int64 `json:"id"`
	} `json:"parent"`
	Content struct {
		Raw    string `json:"raw"`
		HTML   string `json:"html"`
		Markup string `json:"markup"`
	} `json:"content"`
	Inline struct {
		Path string `json:"path"`
		From *int64 `json:"from"`
		To   int64  `json:"to"`
	} `json:"inline"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
}

// PullRequest is the common Bitbucket Pull Request Sub Entity
type PullRequest struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string `json:"state"`
	Author      Owner  `json:"author"`
	Source      struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
		Commit struct {
			Hash string `json:"hash"`
		} `json:"commit"`
		Repository Repository `json:"repository"`
	} `json:"source"`
	Destination struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
		Commit struct {
			Hash string `json:"hash"`
		} `json:"commit"`
		Repository Repository `json:"repository"`
	} `json:"destination"`
	MergeCommit struct {
		Hash string `json:"hash"`
	} `json:"merge_commit"`
	Participants      []Owner   `json:"participants"`
	Reviewers         []Owner   `json:"reviewers"`
	CloseSourceBranch bool      `json:"close_source_branch"`
	ClosedBy          Owner     `json:"closed_by"`
	Reason            string    `json:"reason"`
	CreatedOn         time.Time `json:"created_on"`
	UpdatedOn         time.Time `json:"updated_on"`
	Links             struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
}
