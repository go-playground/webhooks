package bitbucket

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "gopkg.in/go-playground/assert.v1"
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

func TestUUIDMissingEvent(t *testing.T) {
	payload := "{}"
	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Event-Key", "noneexistant_event")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrMissingHookUUIDHeader)
}

func TestUUIDDoesNotMatchEvent(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "THIS_DOES_NOT_MATCH")
	req.Header.Set("X-Event-Key", "repo:push")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrUUIDVerificationFailed)
}

func TestBadNoEventHeader(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrMissingEventKeyHeader)
}

func TestBadBody(t *testing.T) {
	payload := ""

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:push")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrParsingPayload)
}

func TestUnsubscribedEvent(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "noneexistant_event")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrEventNotFound)
}

func TestRepoPush(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/repo-push.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:push")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoPushPayload)
	Equal(t, ok, true)
}

func TestRepoFork(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/repo-fork.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoForkEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:fork")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoForkPayload)
	Equal(t, ok, true)
}

func TestRepoUpdated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/repo-updated.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoUpdatedPayload)
	Equal(t, ok, true)
}

func TestRepoCommitCommentCreated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/commit-comment-created.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoCommitCommentCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoCommitCommentCreatedPayload)
	Equal(t, ok, true)
}

func TestRepoCommitStatusCreated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/repo-commit-status-created.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoCommitStatusCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_status_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoCommitStatusCreatedPayload)
	Equal(t, ok, true)
}

func TestRepoCommitStatusUpdated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/repo-commit-status-updated.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoCommitStatusUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_status_updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoCommitStatusUpdatedPayload)
	Equal(t, ok, true)
}

func TestIssueCreated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/issue-created.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssueCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssueCreatedPayload)
	Equal(t, ok, true)
}

func TestIssueUpdated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/issue-updated.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssueUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssueUpdatedPayload)
	Equal(t, ok, true)
}

func TestIssueCommentCreated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/issue-comment-created.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssueCommentCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssueCommentCreatedPayload)
	Equal(t, ok, true)
}

func TestPullRequestCreated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-created.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestCreatedPayload)
	Equal(t, ok, true)
}

func TestPullRequestUpdated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-updated.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestUpdatedPayload)
	Equal(t, ok, true)
}

func TestPullRequestApproved(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-approved.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestApprovedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:approved")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestApprovedPayload)
	Equal(t, ok, true)
}

func TestPullRequestApprovalRemoved(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-approval-removed.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestUnapprovedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:unapproved")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestUnapprovedPayload)
	Equal(t, ok, true)
}

func TestPullRequestMerged(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-merged.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestMergedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:fulfilled")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestMergedPayload)
	Equal(t, ok, true)
}

func TestPullRequestDeclined(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-declined.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestDeclinedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:rejected")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestDeclinedPayload)
	Equal(t, ok, true)
}

func TestPullRequestCommentCreated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-comment-created.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestCommentCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestCommentCreatedPayload)
	Equal(t, ok, true)
}

func TestPullRequestCommentUpdated(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-comment-updated.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestCommentUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:comment_updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestCommentUpdatedPayload)
	Equal(t, ok, true)
}

func TestPullRequestCommentDeleted(t *testing.T) {

	payload, err := os.Open("../testdata/bitbucket/pull-request-comment-deleted.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestCommentDeletedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:comment_deleted")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestCommentDeletedPayload)
	Equal(t, ok, true)
}
