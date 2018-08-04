package bitbucket

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"io"

	"reflect"

	"github.com/stretchr/testify/require"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//
//
const (
	path = "/webhooks"
)

var hook *Webhook

func TestMain(m *testing.M) {

	// setup
	var err error
	hook, err = New(Options.UUID("MY_UUID"))
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())

	// teardown
}

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	return httptest.NewServer(mux)
}

func TestBadRequests(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name    string
		event   Event
		payload io.Reader
		headers http.Header
	}{
		{
			name:    "UUIDMissingEvent",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Event-Key": []string{"noneexistant_event"},
			},
		},
		{
			name:    "UUIDDoesNotMatchEvent",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Hook-UUID": []string{"THIS_DOES_NOT_MATCH"},
				"X-Event-Key": []string{"repo:push"},
			},
		},
		{
			name:    "BadNoEventHeader",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
			},
		},
		{
			name:    "BadBody",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:push"},
			},
		},
		{
			name:    "UnsubscribedEvent",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"noneexistant_event"},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var parseError error
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				_, parseError = hook.Parse(r, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, tc.payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.Error(parseError)
		})
	}
}

func TestWebhooks(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name     string
		event    Event
		typ      interface{}
		filename string
		headers  http.Header
	}{
		{
			name:     "RepoPush",
			event:    RepoPushEvent,
			typ:      RepoPushPayload{},
			filename: "../testdata/bitbucket/repo-push.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:push"},
			},
		},
		{
			name:     "RepoFork",
			event:    RepoForkEvent,
			typ:      RepoForkPayload{},
			filename: "../testdata/bitbucket/repo-fork.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:fork"},
			},
		},
		{
			name:     "RepoUpdated",
			event:    RepoUpdatedEvent,
			typ:      RepoUpdatedPayload{},
			filename: "../testdata/bitbucket/repo-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:updated"},
			},
		},
		{
			name:     "RepoCommitCommentCreated",
			event:    RepoCommitCommentCreatedEvent,
			typ:      RepoCommitCommentCreatedPayload{},
			filename: "../testdata/bitbucket/commit-comment-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:commit_comment_created"},
			},
		},
		{
			name:     "RepoCommitStatusCreated",
			event:    RepoCommitStatusCreatedEvent,
			typ:      RepoCommitStatusCreatedPayload{},
			filename: "../testdata/bitbucket/repo-commit-status-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:commit_status_created"},
			},
		},
		{
			name:     "RepoCommitStatusUpdated",
			event:    RepoCommitStatusUpdatedEvent,
			typ:      RepoCommitStatusUpdatedPayload{},
			filename: "../testdata/bitbucket/repo-commit-status-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:commit_status_updated"},
			},
		},
		{
			name:     "IssueCreated",
			event:    IssueCreatedEvent,
			typ:      IssueCreatedPayload{},
			filename: "../testdata/bitbucket/issue-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"issue:created"},
			},
		},
		{
			name:     "IssueUpdated",
			event:    IssueUpdatedEvent,
			typ:      IssueUpdatedPayload{},
			filename: "../testdata/bitbucket/issue-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"issue:updated"},
			},
		},
		{
			name:     "IssueUpdated",
			event:    IssueUpdatedEvent,
			typ:      IssueUpdatedPayload{},
			filename: "../testdata/bitbucket/issue-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"issue:updated"},
			},
		},
		{
			name:     "IssueCommentCreated",
			event:    IssueCommentCreatedEvent,
			typ:      IssueCommentCreatedPayload{},
			filename: "../testdata/bitbucket/issue-comment-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"issue:comment_created"},
			},
		},
		{
			name:     "PullRequestCreated",
			event:    PullRequestCreatedEvent,
			typ:      PullRequestCreatedPayload{},
			filename: "../testdata/bitbucket/pull-request-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:created"},
			},
		},
		{
			name:     "PullRequestUpdated",
			event:    PullRequestUpdatedEvent,
			typ:      PullRequestUpdatedPayload{},
			filename: "../testdata/bitbucket/pull-request-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:updated"},
			},
		},
		{
			name:     "PullRequestApproved",
			event:    PullRequestApprovedEvent,
			typ:      PullRequestApprovedPayload{},
			filename: "../testdata/bitbucket/pull-request-approved.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:approved"},
			},
		},
		{
			name:     "PullRequestApprovalRemoved",
			event:    PullRequestUnapprovedEvent,
			typ:      PullRequestUnapprovedPayload{},
			filename: "../testdata/bitbucket/pull-request-approval-removed.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:unapproved"},
			},
		},
		{
			name:     "PullRequestMerged",
			event:    PullRequestMergedEvent,
			typ:      PullRequestMergedPayload{},
			filename: "../testdata/bitbucket/pull-request-merged.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:fulfilled"},
			},
		},
		{
			name:     "PullRequestDeclined",
			event:    PullRequestDeclinedEvent,
			typ:      PullRequestDeclinedPayload{},
			filename: "../testdata/bitbucket/pull-request-declined.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:rejected"},
			},
		},
		{
			name:     "PullRequestCommentCreated",
			event:    PullRequestCommentCreatedEvent,
			typ:      PullRequestCommentCreatedPayload{},
			filename: "../testdata/bitbucket/pull-request-comment-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:comment_created"},
			},
		},
		{
			name:     "PullRequestCommentUpdated",
			event:    PullRequestCommentUpdatedEvent,
			typ:      PullRequestCommentUpdatedPayload{},
			filename: "../testdata/bitbucket/pull-request-comment-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:comment_updated"},
			},
		},
		{
			name:     "PullRequestCommentDeleted",
			event:    PullRequestCommentDeletedEvent,
			typ:      PullRequestCommentDeletedPayload{},
			filename: "../testdata/bitbucket/pull-request-comment-deleted.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:comment_deleted"},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			payload, err := os.Open(tc.filename)
			assert.NoError(err)
			defer func() {
				_ = payload.Close()
			}()

			var parseError error
			var results interface{}
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				results, parseError = hook.Parse(r, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}
