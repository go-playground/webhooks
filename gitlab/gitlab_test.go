package gitlab

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

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
	hook, err = New(Options.Secret("sampleToken!"))
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
			event:   PushEvents,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{},
		},
		{
			name:    "UnsubscribedEvent",
			event:   PushEvents,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gitlab-Event": []string{"noneexistant_event"},
			},
		},
		{
			name:    "BadBody",
			event:   PushEvents,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Gitlab-Event": []string{"Push Hook"},
				"X-Gitlab-Token": []string{"sampleToken!"},
			},
		},
		{
			name:    "TokenMismatch",
			event:   PushEvents,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gitlab-Event": []string{"Push Hook"},
				"X-Gitlab-Token": []string{"badsampleToken!!"},
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
			name:     "PushEvent",
			event:    PushEvents,
			typ:      PushEventPayload{},
			filename: "../testdata/gitlab/push-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Push Hook"},
			},
		},
		{
			name:     "TagEvent",
			event:    TagEvents,
			typ:      TagEventPayload{},
			filename: "../testdata/gitlab/tag-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Tag Push Hook"},
			},
		},
		{
			name:     "IssueEvent",
			event:    IssuesEvents,
			typ:      IssueEventPayload{},
			filename: "../testdata/gitlab/issue-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Issue Hook"},
			},
		},
		{
			name:     "ConfidentialIssueEvent",
			event:    ConfidentialIssuesEvents,
			typ:      ConfidentialIssueEventPayload{},
			filename: "../testdata/gitlab/confidential-issue-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Confidential Issue Hook"},
			},
		},
		{
			name:     "CommentCommitEvent",
			event:    CommentEvents,
			typ:      CommentEventPayload{},
			filename: "../testdata/gitlab/comment-commit-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Note Hook"},
			},
		},
		{
			name:     "CommentMergeRequestEvent",
			event:    CommentEvents,
			typ:      CommentEventPayload{},
			filename: "../testdata/gitlab/comment-merge-request-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Note Hook"},
			},
		},
		{
			name:     "CommentIssueEvent",
			event:    CommentEvents,
			typ:      CommentEventPayload{},
			filename: "../testdata/gitlab/comment-issue-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Note Hook"},
			},
		},
		{
			name:     "CommentSnippetEvent",
			event:    CommentEvents,
			typ:      CommentEventPayload{},
			filename: "../testdata/gitlab/comment-snippet-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Note Hook"},
			},
		},
		{
			name:     "MergeRequestEvent",
			event:    MergeRequestEvents,
			typ:      MergeRequestEventPayload{},
			filename: "../testdata/gitlab/merge-request-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Merge Request Hook"},
			},
		},
		{
			name:     "WikipageEvent",
			event:    WikiPageEvents,
			typ:      WikiPageEventPayload{},
			filename: "../testdata/gitlab/wikipage-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Wiki Page Hook"},
			},
		},
		{
			name:     "PipelineEvent",
			event:    PipelineEvents,
			typ:      PipelineEventPayload{},
			filename: "../testdata/gitlab/pipeline-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Pipeline Hook"},
			},
		},
		{
			name:     "BuildEvent",
			event:    BuildEvents,
			typ:      BuildEventPayload{},
			filename: "../testdata/gitlab/build-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Build Hook"},
			},
		},
		{
			name:     "JobEvent",
			event:    JobEvents,
			typ:      JobEventPayload{},
			filename: "../testdata/gitlab/job-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Job Hook"},
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
			req.Header.Set("X-Gitlab-Token", "sampleToken!")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}

func TestSystemHooks(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name     string
		event    Event
		typ      interface{}
		filename string
	}{
		{
			name:     "PushEvent",
			event:    PushEvents,
			typ:      PushEventPayload{},
			filename: "../testdata/gitlab/system-push-event.json",
		},
		{
			name:     "TagEvent",
			event:    TagEvents,
			typ:      TagEventPayload{},
			filename: "../testdata/gitlab/system-tag-event.json",
		},
		{
			name:     "MergeRequestEvent",
			event:    MergeRequestEvents,
			typ:      MergeRequestEventPayload{},
			filename: "../testdata/gitlab/system-merge-request-event.json",
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
				results, parseError = hook.Parse(r, SystemHookEvents, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
			assert.NoError(err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-Gitlab-Token", "sampleToken!")
			req.Header.Set("X-Gitlab-Event", "System Hook")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}
