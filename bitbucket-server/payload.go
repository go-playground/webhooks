package bitbucketserver

import (
	"fmt"
	"strings"
	"time"
)

type DiagnosticsPingPayload struct{}

type RepositoryReferenceChangedPayload struct {
	Date       Date               `json:"date"`
	EventKey   Event              `json:"eventKey"`
	Actor      User               `json:"actor"`
	Repository Repository         `json:"repository"`
	Changes    []RepositoryChange `json:"changes"`
}

type RepositoryModifiedPayload struct {
	Date     Date       `json:"date"`
	EventKey Event      `json:"eventKey"`
	Actor    User       `json:"actor"`
	Old      Repository `json:"old"`
	New      Repository `json:"new"`
}

type RepositoryForkedPayload struct {
	Date       Date       `json:"date"`
	EventKey   Event      `json:"eventKey"`
	Actor      User       `json:"actor"`
	Repository Repository `json:"repository"`
}

type RepositoryCommentAddedPayload struct {
	Date       Date       `json:"date"`
	EventKey   Event      `json:"eventKey"`
	Actor      User       `json:"actor"`
	Comment    Comment    `json:"comment"`
	Repository Repository `json:"repository"`
	Commit     string     `json:"commit"`
}

type RepositoryCommentEditedPayload struct {
	Date            Date       `json:"date"`
	EventKey        Event      `json:"eventKey"`
	Actor           User       `json:"actor"`
	Comment         Comment    `json:"comment"`
	PreviousComment string     `json:"previousComment"`
	Repository      Repository `json:"repository"`
	Commit          string     `json:"commit"`
}

type RepositoryCommentDeletedPayload struct {
	Date       Date       `json:"date"`
	EventKey   Event      `json:"eventKey"`
	Actor      User       `json:"actor"`
	Comment    Comment    `json:"comment"`
	Repository Repository `json:"repository"`
	Commit     string     `json:"commit"`
}

type PullRequestOpenedPayload struct {
	Date        Date        `json:"date"`
	EventKey    Event       `json:"eventKey"`
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullRequest"`
}

type PullRequestFromReferenceUpdatedPayload struct {
	Date             Date        `json:"date"`
	EventKey         Event       `json:"eventKey"`
	Actor            User        `json:"actor"`
	PullRequest      PullRequest `json:"pullRequest"`
	PreviousFromHash string      `json:"previousFromHash"`
}

type PullRequestModifiedPayload struct {
	Date                Date                   `json:"date"`
	EventKey            Event                  `json:"eventKey"`
	Actor               User                   `json:"actor"`
	PullRequest         PullRequest            `json:"pullRequest"`
	PreviousTitle       string                 `json:"previousTitle"`
	PreviousDescription string                 `json:"previousDescription"`
	PreviousTarget      map[string]interface{} `json:"previousTarget"`
}

type PullRequestMergedPayload struct {
	Date        Date        `json:"date"`
	EventKey    Event       `json:"eventKey"`
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullRequest"`
}

type PullRequestDeclinedPayload struct {
	Date        Date        `json:"date"`
	EventKey    Event       `json:"eventKey"`
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullRequest"`
}

type PullRequestDeletedPayload struct {
	Date        Date        `json:"date"`
	EventKey    Event       `json:"eventKey"`
	Actor       User        `json:"actor"`
	PullRequest PullRequest `json:"pullRequest"`
}

type PullRequestReviewerUpdatedPayload struct {
	Date             Date        `json:"date"`
	EventKey         Event       `json:"eventKey"`
	Actor            User        `json:"actor"`
	PullRequest      PullRequest `json:"pullRequest"`
	RemovedReviewers []User      `json:"removedReviewers"`
	AddedReviewers   []User      `json:"addedReviewers"`
}

type PullRequestReviewerApprovedPayload struct {
	Date           Date                   `json:"date"`
	EventKey       Event                  `json:"eventKey"`
	Actor          User                   `json:"actor"`
	PullRequest    PullRequest            `json:"pullRequest"`
	Participant    PullRequestParticipant `json:"participant"`
	PreviousStatus string                 `json:"previousStatus"`
}

type PullRequestReviewerUnapprovedPayload struct {
	Date           Date                   `json:"date"`
	EventKey       Event                  `json:"eventKey"`
	Actor          User                   `json:"actor"`
	PullRequest    PullRequest            `json:"pullRequest"`
	Participant    PullRequestParticipant `json:"participant"`
	PreviousStatus string                 `json:"previousStatus"`
}

type PullRequestReviewerNeedsWorkPayload struct {
	Date           Date                   `json:"date"`
	EventKey       Event                  `json:"eventKey"`
	Actor          User                   `json:"actor"`
	PullRequest    PullRequest            `json:"pullRequest"`
	Participant    PullRequestParticipant `json:"participant"`
	PreviousStatus string                 `json:"previousStatus"`
}

type PullRequestCommentAddedPayload struct {
	Date            Date        `json:"date"`
	EventKey        Event       `json:"eventKey"`
	Actor           User        `json:"actor"`
	PullRequest     PullRequest `json:"pullRequest"`
	Comment         Comment     `json:"comment"`
	CommentParentId uint64      `json:"commentParentId,omitempty"`
}

type PullRequestCommentEditedPayload struct {
	Date            Date        `json:"date"`
	EventKey        Event       `json:"eventKey"`
	Actor           User        `json:"actor"`
	PullRequest     PullRequest `json:"pullRequest"`
	Comment         Comment     `json:"comment"`
	CommentParentId string      `json:"commentParentId,omitempty"`
	PreviousComment string      `json:"previousComment"`
}

type PullRequestCommentDeletedPayload struct {
	Date            Date        `json:"date"`
	EventKey        Event       `json:"eventKey"`
	Actor           User        `json:"actor"`
	PullRequest     PullRequest `json:"pullRequest"`
	Comment         Comment     `json:"comment"`
	CommentParentId uint64      `json:"commentParentId,omitempty"`
}

// -----------------------

type User struct {
	ID           uint64                 `json:"id"`
	Name         string                 `json:"name"`
	EmailAddress string                 `json:"emailAddress"`
	DisplayName  string                 `json:"displayName"`
	Active       bool                   `json:"active"`
	Slug         string                 `json:"slug"`
	Type         string                 `json:"type"`
	Links        map[string]interface{} `json:"links"`
}

type Repository struct {
	ID            uint64                 `json:"id"`
	Slug          string                 `json:"slug"`
	Name          string                 `json:"name"`
	ScmId         string                 `json:"scmId"`
	State         string                 `json:"state"`
	StatusMessage string                 `json:"statusMessage"`
	Forkable      bool                   `json:"forkable"`
	Origin        *Repository            `json:"origin,omitempty"`
	Project       Project                `json:"project"`
	Public        bool                   `json:"public"`
	Links         map[string]interface{} `json:"links"`
}

type Project struct {
	ID     uint64                 `json:"id"`
	Key    string                 `json:"key"`
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Public *bool                  `json:"public,omitempty"`
	Owner  User                   `json:"owner"`
	Links  map[string]interface{} `json:"links"`
}

type PullRequest struct {
	ID           uint64                   `json:"id"`
	Version      uint64                   `json:"version"`
	Title        string                   `json:"title"`
	Description  string                   `json:"description,omitempty"`
	State        string                   `json:"state"`
	Open         bool                     `json:"open"`
	Closed       bool                     `json:"closed"`
	CreatedDate  uint64                   `json:"createdDate"`
	UpdatedDate  uint64                   `json:"updatedDate,omitempty"`
	ClosedDate   uint64                   `json:"closedDate,omitempty"`
	FromRef      RepositoryReference      `json:"fromRef"`
	ToRef        RepositoryReference      `json:"toRef"`
	Locked       bool                     `json:"locked"`
	Author       PullRequestParticipant   `json:"author"`
	Reviewers    []PullRequestParticipant `json:"reviewers"`
	Participants []PullRequestParticipant `json:"participants"`
	Properties   map[string]interface{}   `json:"properties,omitempty"`
	Links        map[string]interface{}   `json:"links"`
}

type RepositoryChange struct {
	Reference   RepositoryReference `json:"ref"`
	ReferenceId string              `json:"refId"`
	FromHash    string              `json:"fromHash"`
	ToHash      string              `json:"toHash"`
	Type        string              `json:"type"`
}

type RepositoryReference struct {
	ID           string     `json:"id"`
	DisplayId    string     `json:"displayId"`
	Type         string     `json:"type,omitempty"`
	LatestCommit string     `json:"latestCommit,omitempty"`
	Repository   Repository `json:"repository,omitempty"`
}

type Comment struct {
	ID                  uint64                   `json:"id"`
	Properties          map[string]interface{}   `json:"properties,omitempty"`
	Version             uint64                   `json:"version"`
	Text                string                   `json:"text"`
	Author              User                     `json:"author"`
	CreatedDate         uint64                   `json:"createdDate"`
	UpdatedDate         uint64                   `json:"updatedDate"`
	Comments            []map[string]interface{} `json:"comments"`
	Tasks               []map[string]interface{} `json:"tasks"`
	PermittedOperations map[string]interface{}   `json:"permittedOperations,omitempty"`
}

type PullRequestParticipant struct {
	User               User   `json:"user"`
	LastReviewedCommit string `json:"lastReviewedCommit,omitempty"`
	Role               string `json:"role"`
	Approved           bool   `json:"approved"`
	Status             string `json:"status"`
}

type Date time.Time

func (b *Date) UnmarshalJSON(p []byte) error {
	t, err := time.Parse("2006-01-02T15:04:05Z0700", strings.Replace(string(p), "\"", "", -1))
	if err != nil {
		return err
	}
	*b = Date(t)
	return nil
}

func (b Date) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(b).Format("2006-01-02T15:04:05Z0700"))
	return []byte(stamp), nil
}
