package gitea

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"io"

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

const (
	path = "/webhooks"
)

var hook *Webhook

func TestMain(m *testing.M) {

	// setup
	var err error
	hook, err = New(Options.Secret("IsWishesWereHorsesWedAllBeEatingSteak!"))
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
			name:    "BadNoEventHeader",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{},
		},
		{
			name:    "UnsubscribedEvent",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gitea-Event": []string{"noneexistant_event"},
			},
		},
		{
			name:    "BadBody",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Gitea-Event": []string{"push"},
			},
		},
		{
			name:    "TokenMismatch",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gitea-Event": []string{"push"},
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
			name:     "CreateEvent",
			event:    CreateEvent,
			typ:      CreatePayload{},
			filename: "../testdata/gitea/create-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"create"},
			},
		},
		{
			name:     "DeleteEvent",
			event:    DeleteEvent,
			typ:      DeletePayload{},
			filename: "../testdata/gitea/delete-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"delete"},
			},
		},
		{
			name:     "ForkEvent",
			event:    ForkEvent,
			typ:      ForkPayload{},
			filename: "../testdata/gitea/fork-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"fork"},
			},
		},
		{
			name:     "IssuesEvent",
			event:    IssuesEvent,
			typ:      IssuePayload{},
			filename: "../testdata/gitea/issues-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"issues"},
			},
		},
		{
			name:     "IssueAssignEvent",
			event:    IssueAssignEvent,
			typ:      IssuePayload{},
			filename: "../testdata/gitea/issue-assign-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"issue_assign"},
			},
		},
		{
			name:     "IssueLabelEvent",
			event:    IssueLabelEvent,
			typ:      IssuePayload{},
			filename: "../testdata/gitea/issue-label-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"issue_label"},
			},
		},
		{
			name:     "IssueMilestoneEvent",
			event:    IssueMilestoneEvent,
			typ:      IssuePayload{},
			filename: "../testdata/gitea/issue-milestone-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"issue_milestone"},
			},
		},
		{
			name:     "IssueCommentEvent",
			event:    IssueCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "../testdata/gitea/issue-comment-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"issue_comment"},
			},
		},
		{
			name:     "PushEvent",
			event:    PushEvent,
			typ:      PushPayload{},
			filename: "../testdata/gitea/push-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"push"},
			},
		},
		{
			name:     "PullRequestEvent",
			event:    PullRequestEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"pull_request"},
			},
		},
		{
			name:     "PullRequestAssignEvent",
			event:    PullRequestAssignEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-assign-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"pull_request_assign"},
			},
		},
		{
			name:     "PullRequestLabelEvent",
			event:    PullRequestLabelEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-label-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"pull_request_label"},
			},
		},
		{
			name:     "PullRequestMilestoneEvent",
			event:    PullRequestMilestoneEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-milestone-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"pull_request_milestone"},
			},
		},
		{
			name:     "PullRequestCommentEvent",
			event:    PullRequestCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "../testdata/gitea/pull-request-comment-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"pull_request_comment"},
			},
		},
		{
			name:     "PullRequestReviewEvent",
			event:    PullRequestReviewEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-review-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"pull_request_review"},
			},
		},
		{
			name:     "RepositoryEvent",
			event:    RepositoryEvent,
			typ:      RepositoryPayload{},
			filename: "../testdata/gitea/repository-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"repository"},
			},
		},
		{
			name:     "ReleaseEvent",
			event:    ReleaseEvent,
			typ:      ReleasePayload{},
			filename: "../testdata/gitea/release-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"release"},
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
