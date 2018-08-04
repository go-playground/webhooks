package gitlab

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

func TestBadNoEventHeader(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, PushEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrMissingGitLabEventHeader)
}

func TestUnsubscribedEvent(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, PushEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "noneexistant_event")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrEventNotFound)
}

func TestBadBody(t *testing.T) {
	payload := ""

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, PushEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Push Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrParsingPayload)
}

func TestTokenMismatch(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, PushEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Push Hook")
	req.Header.Set("X-Gitlab-Token", "badsampleToken!!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrGitLabTokenVerificationFailed)
}

func TestPushEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/push-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PushEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Push Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PushEventPayload)
	Equal(t, ok, true)
}

func TestTagEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/tag-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, TagEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Tag Push Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(TagEventPayload)
	Equal(t, ok, true)
}

func TestIssueEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/issue-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssuesEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Issue Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssueEventPayload)
	Equal(t, ok, true)
}

func TestConfidentialIssueEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/confidential-issue-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, ConfidentialIssuesEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Confidential Issue Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(ConfidentialIssueEventPayload)
	Equal(t, ok, true)
}

func TestCommentCommitEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/comment-commit-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, CommentEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Note Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(CommentEventPayload)
	Equal(t, ok, true)
}

func TestCommentMergeRequestEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/comment-merge-request-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, CommentEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Note Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(CommentEventPayload)
	Equal(t, ok, true)
}

func TestCommentIssueEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/comment-issue-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, CommentEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Note Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(CommentEventPayload)
	Equal(t, ok, true)
}

func TestCommentSnippetEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/comment-snippet-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, CommentEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Note Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(CommentEventPayload)
	Equal(t, ok, true)
}

func TestMergeRequestEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/merge-request-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, MergeRequestEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Merge Request Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(MergeRequestEventPayload)
	Equal(t, ok, true)
}

func TestWikipageEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/wikipage-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, WikiPageEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Wiki Page Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(WikiPageEventPayload)
	Equal(t, ok, true)
}

func TestPipelineEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/pipeline-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PipelineEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Pipeline Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PipelineEventPayload)
	Equal(t, ok, true)
}

func TestBuildEvent(t *testing.T) {

	payload, err := os.Open("../testdata/gitlab/build-event.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, BuildEvents)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", "Build Hook")
	req.Header.Set("X-Gitlab-Token", "sampleToken!")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(BuildEventPayload)
	Equal(t, ok, true)
}
