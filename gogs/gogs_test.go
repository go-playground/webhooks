package gogs

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	client "github.com/gogits/go-gogs-client"
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

	//teardown

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
			event:   CreateEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{},
		},
		{
			name:    "UnsubscribedEvent",
			event:   CreateEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gogs-Event": []string{"noneexistant_event"},
			},
		},
		{
			name:    "BadBody",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Gogs-Event":     []string{"push"},
				"X-Gogs-Signature": []string{"0dacdb7c00bc1cdc0c24038b7b244cf65146b33da862a769e2608a063529fffc"},
			},
		},
		{
			name:    "BadSignatureLength",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gogs-Event":     []string{"push"},
				"X-Gogs-Signature": []string{""},
			},
		},
		{
			name:    "BadSignatureMatch",
			event:   PushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Github-Event":   []string{"push"},
				"X-Gogs-Signature": []string{"111"},
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
			typ:      client.CreatePayload{},
			filename: "../testdata/gogs/create-event.json",
			headers: http.Header{
				"X-Gogs-Delivery":  []string{"f6266f16-1bf3-46a5-9ea4-602e06ead473"},
				"X-Gogs-Event":     []string{"create"},
				"X-Gogs-Signature": []string{"d3c99a5711de365c4ff897be0692d9016de2cee6e3e1d6b3fb49154f9efb354d"},
			},
		},
		{
			name:     "DeleteEvent",
			event:    DeleteEvent,
			typ:      client.DeletePayload{},
			filename: "../testdata/gogs/delete-event.json",
			headers: http.Header{
				"X-Gogs-Delivery":  []string{"f6266f16-1bf3-46a5-9ea4-602e06ead473"},
				"X-Gogs-Event":     []string{"delete"},
				"X-Gogs-Signature": []string{"6208970efc3c5283df65f726d69ae8331491d7857f812dfaeaedf11a68d8da79"},
			},
		},
		{
			name:     "ForkEvent",
			event:    ForkEvent,
			typ:      client.ForkPayload{},
			filename: "../testdata/gogs/fork-event.json",
			headers: http.Header{
				"X-Gogs-Delivery":  []string{"f6266f16-1bf3-46a5-9ea4-602e06ead473"},
				"X-Gogs-Event":     []string{"fork"},
				"X-Gogs-Signature": []string{"bd5fda972bd27745d9e108c3d92a926a47baf8b715a422f73ff9acdfa6a86402"},
			},
		},
		{
			name:     "PushEvent",
			event:    PushEvent,
			typ:      client.PushPayload{},
			filename: "../testdata/gogs/push-event.json",
			headers: http.Header{
				"X-Gogs-Delivery":  []string{"f6266f16-1bf3-46a5-9ea4-602e06ead473"},
				"X-Gogs-Event":     []string{"push"},
				"X-Gogs-Signature": []string{"83d4163fb936904aeb9ffd6ce22cf86e4b36273b2e1c63a57f5a6ddc371ce3ba"},
			},
		},
		{
			name:     "IssuesEvent",
			event:    IssuesEvent,
			typ:      client.IssuesPayload{},
			filename: "../testdata/gogs/issues-event.json",
			headers: http.Header{
				"X-Gogs-Delivery":  []string{"f6266f16-1bf3-46a5-9ea4-602e06ead473"},
				"X-Gogs-Event":     []string{"issues"},
				"X-Gogs-Signature": []string{"bb61f632278cc601f37bc0a611e2a06bba3b53453f3e58092dc20ffa0cec6916"},
			},
		},
		{
			name:     "IssueCommentEvent",
			event:    IssueCommentEvent,
			typ:      client.IssueCommentPayload{},
			filename: "../testdata/gogs/issue-comment-event.json",
			headers: http.Header{
				"X-Gogs-Delivery":  []string{"f6266f16-1bf3-46a5-9ea4-602e06ead473"},
				"X-Gogs-Event":     []string{"issue_comment"},
				"X-Gogs-Signature": []string{"1d45142f03ed44a06d6630e57aa8c8c0ad7ed90a57cc0ee9fec8fb5434efd92e"},
			},
		},
		{
			name:     "PullRequestEvent",
			event:    PullRequestEvent,
			typ:      client.PullRequestPayload{},
			filename: "../testdata/gogs/pull-request-event.json",
			headers: http.Header{
				"X-Gogs-Delivery":  []string{"f6266f16-1bf3-46a5-9ea4-602e06ead473"},
				"X-Gogs-Event":     []string{"pull_request"},
				"X-Gogs-Signature": []string{"69105ac4cada9f0b099be34722b26d3854be20988c26bd9e135fec74a830d00a"},
			},
		},
		{
			name:     "ReleaseEvent",
			event:    ReleaseEvent,
			typ:      client.ReleasePayload{},
			filename: "../testdata/gogs/release-event.json",
			headers: http.Header{
				"X-Gogs-Delivery":  []string{"f6266f16-1bf3-46a5-9ea4-602e06ead473"},
				"X-Gogs-Event":     []string{"release"},
				"X-Gogs-Signature": []string{"93bfcfb879642dbc28136266b1f7d56d053fa08332aa1f4d884a66ced2b5f10a"},
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
