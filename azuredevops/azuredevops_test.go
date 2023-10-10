package azuredevops

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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

const (
	virtualDir = "/webhooks"
)

var hook *Webhook

func TestMain(m *testing.M) {

	// setup
	var err error
	hook, err = New()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
	// teardown
}

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(virtualDir, handler)
	return httptest.NewServer(mux)
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
			name:     "build.complete",
			event:    BuildCompleteEventType,
			typ:      BuildCompleteEvent{},
			filename: "../testdata/azuredevops/build.complete.json",
		},
		{
			name:     "git.pullrequest.created",
			event:    GitPullRequestCreatedEventType,
			typ:      GitPullRequestEvent{},
			filename: "../testdata/azuredevops/git.pullrequest.created.json",
		},
		{
			name:     "git.pullrequest.merged",
			event:    GitPullRequestMergedEventType,
			typ:      GitPullRequestEvent{},
			filename: "../testdata/azuredevops/git.pullrequest.merged.json",
		},
		{
			name:     "git.pullrequest.updated",
			event:    GitPullRequestUpdatedEventType,
			typ:      GitPullRequestEvent{},
			filename: "../testdata/azuredevops/git.pullrequest.updated.json",
		},
		{
			name:     "git.push",
			event:    GitPushEventType,
			typ:      GitPushEvent{},
			filename: "../testdata/azuredevops/git.push.json",
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
			req, err := http.NewRequest(http.MethodPost, server.URL+virtualDir, payload)
			assert.NoError(err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}
