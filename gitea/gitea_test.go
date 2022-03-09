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
			name:    "BadSignatureLength",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Github-Event":  []string{"push"},
				"X-Gitea-Signature": []string{""},
			},
		},
		{
			name:    "BadSignatureMatch",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Github-Event":  []string{"push"},
				"X-Gitea-Signature": []string{"111"},
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
				"X-Gitea-Event":   []string{"create"},
				"X-Gitea-Signature": []string{"6f250ac7a090096574758e31bd31770eab63dfd0459404f0c18431f1c6b9024a"},
			},
		},
		{
			name:     "DeleteEvent",
			event:    DeleteEvent,
			typ:      DeletePayload{},
			filename: "../testdata/gitea/delete-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"delete"},
				"X-Gitea-Signature": []string{"84307b509e663cd897bc719b2a564e64fa4af8716fda389488f18369139e0fdd"},
			},
		},
		{
			name:     "ForkEvent",
			event:    ForkEvent,
			typ:      ForkPayload{},
			filename: "../testdata/gitea/fork-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"fork"},
				"X-Gitea-Signature": []string{"b7750f34adeaf333ac83a1fadcda4cbac097c8587e8dda297c3e7f059012215f"},
			},
		},
		{
			name:     "IssuesEvent",
			event:    IssuesEvent,
			typ:      IssuePayload{},
			filename: "../testdata/gitea/issues-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"issues"},
				"X-Gitea-Signature": []string{"98c44fd0ae42ca4208eac6f81e59b436837740abc8693bf828366b32d33b1cbc"},
			},
		},
		{
			name:     "IssueAssignEvent",
			event:    IssueAssignEvent,
			typ:      IssuePayload{},
			filename: "../testdata/gitea/issue-assign-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"issue_assign"},
				"X-Gitea-Signature": []string{"7d2adaf2fb3dc3769294c737ff48da003b7c3660b4f917b85c2c25dabd34a13c"},
			},
		},
		{
			name:     "IssueLabelEvent",
			event:    IssueLabelEvent,
			typ:      IssuePayload{},
			filename: "../testdata/gitea/issue-label-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"issue_label"},
				"X-Gitea-Signature": []string{"3611415860ed5904c87dd589a7c5fa7e87d2a72b0b2a92ea149ba9691ba8c785"},
			},
		},
		{
			name:     "IssueMilestoneEvent",
			event:    IssueMilestoneEvent,
			typ:      IssuePayload{},
			filename: "../testdata/gitea/issue-milestone-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"issue_milestone"},
				"X-Gitea-Signature": []string{"4b782b02035ca264c8e7782b8f2eb9d64e4a61344a1bc3a08fa85d7eed1e77b5"},
			},
		},
		{
			name:     "IssueCommentEvent",
			event:    IssueCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "../testdata/gitea/issue-comment-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"issue_comment"},
				"X-Gitea-Signature": []string{"690180a8c853460cba88f9b09911a531d449e9f63ed6b1ff0a0def3b972ca744"},
			},
		},
		{
			name:     "PushEvent",
			event:    PushEvent,
			typ:      PushPayload{},
			filename: "../testdata/gitea/push-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"push"},
				"X-Gitea-Signature": []string{"60fe446c74fa0cb9474f98cc557db79e10c7aaf22cf324ad65239600b9e4d915"},
			},
		},
		{
			name:     "PullRequestEvent",
			event:    PullRequestEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"pull_request"},
				"X-Gitea-Signature": []string{"65c18a212efc7bde0f336acaec87f596fe20e80b2a0e7e51a790dd38393ff771"},
			},
		},
		{
			name:     "PullRequestAssignEvent",
			event:    PullRequestAssignEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-assign-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"pull_request_assign"},
				"X-Gitea-Signature": []string{"6e96f0515898d427d87fc022cef60e4a02695739e8eae05f8cccd79b2ce4809a"},
			},
		},
		{
			name:     "PullRequestLabelEvent",
			event:    PullRequestLabelEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-label-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"pull_request_label"},
				"X-Gitea-Signature": []string{"c52fa035b8d9ac4d94449349b16bac5892bc59faa72be19ff39f33b6bc24315a"},
			},
		},
		{
			name:     "PullRequestMilestoneEvent",
			event:    PullRequestMilestoneEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-milestone-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"pull_request_milestone"},
				"X-Gitea-Signature": []string{"38b2cf88e15b15795371517cb4121a92c7db1116f83eba57c13c192a7e0730dc"},
			},
		},
		{
			name:     "PullRequestCommentEvent",
			event:    PullRequestCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "../testdata/gitea/pull-request-comment-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"pull_request_comment"},
				"X-Gitea-Signature": []string{"8b38bf221adbaef2ce01cfa810c6d9cb977414fec306895973aea61e10e8d5a8"},
			},
		},
		{
			name:     "PullRequestReviewEvent",
			event:    PullRequestReviewEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/gitea/pull-request-review-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"pull_request_review"},
				"X-Gitea-Signature": []string{"5d5c315cc199807a23da81b788dcf7874299a223ba81fc77d76ab248fdab4d1c"},
			},
		},
		{
			name:     "RepositoryEvent",
			event:    RepositoryEvent,
			typ:      RepositoryPayload{},
			filename: "../testdata/gitea/repository-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"repository"},
				"X-Gitea-Signature": []string{"52ca39cfb873254a9cbbda10f8a878f9df16216da6ae92267fc81352ac685d97"},
			},
		},
		{
			name:     "ReleaseEvent",
			event:    ReleaseEvent,
			typ:      ReleasePayload{},
			filename: "../testdata/gitea/release-event.json",
			headers: http.Header{
				"X-Gitea-Event":   []string{"release"},
				"X-Gitea-Signature": []string{"847fcef001c2e59dadac3fa5fa01ca26c9985a5faa10e48a9868a8ad98e9dd18"},
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
