package gitea_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"code.gitea.io/gitea/modules/structs"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/webhooks.v5/gitea"
)

const (
	path = "/webhooks"
)

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	return httptest.NewServer(mux)
}

func TestWebhooks(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name     string
		event    gitea.Event
		typ      interface{}
		filename string
		headers  http.Header
	}{
		{
			name:     "PushEvent",
			event:    gitea.PushEvent,
			typ:      structs.PushPayload{},
			filename: "../testdata/gitea/push-event.json",
			headers: http.Header{
				"X-Gitea-Event": []string{"push"},
			},
		},
	}

	hook, _ := gitea.New(gitea.Options.Secret("mytoken"))

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
			req.Header.Set("X-Gitea-Token", "mytoken")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}

func TestBadRequests(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name           string
		event          gitea.Event
		payload        io.Reader
		headers        http.Header
		method         string
		expcectedError error
	}{
		{
			name:           "ErrEventNotSpecifiedToParse",
			event:          "",
			method:         http.MethodPost,
			payload:        bytes.NewBuffer([]byte("{}")),
			headers:        http.Header{},
			expcectedError: gitea.ErrEventNotSpecifiedToParse,
		},
		{
			name:           "ErrInvalidHTTPMethod",
			event:          gitea.PushEvent,
			method:         http.MethodGet,
			payload:        bytes.NewBuffer([]byte("{}")),
			headers:        http.Header{},
			expcectedError: gitea.ErrInvalidHTTPMethod,
		},
		{
			name:           "ErrMissingGiteaEventHeader",
			event:          gitea.PushEvent,
			method:         http.MethodPost,
			payload:        bytes.NewBuffer([]byte("{}")),
			headers:        http.Header{},
			expcectedError: gitea.ErrMissingGiteaEventHeader,
		},
		{
			name:    "ErrEventNotFound",
			event:   gitea.PushEvent,
			method:  http.MethodPost,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gitea-Event": []string{"create"},
			},
			expcectedError: gitea.ErrEventNotFound,
		},
		{
			name:    "ErrParsingPayload",
			event:   gitea.PushEvent,
			method:  http.MethodPost,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Gitea-Event": []string{"push"},
			},
			expcectedError: gitea.ErrParsingPayload,
		},
		{
			name:    "ErrSecretNotMatch",
			event:   gitea.PushEvent,
			method:  http.MethodPost,
			payload: bytes.NewBuffer([]byte("{\"secret\":\"test\"}")),
			headers: http.Header{
				"X-Gitea-Event": []string{"push"},
			},
			expcectedError: gitea.ErrSecretNotMatch,
		},
	}

	hook, _ := gitea.New(gitea.Options.Secret("mytoken"))

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var parseError error
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				if tc.event != "" {
					_, parseError = hook.Parse(r, tc.event)
				} else {
					_, parseError = hook.Parse(r)
				}
			})
			defer server.Close()
			req, err := http.NewRequest(tc.method, server.URL+path, tc.payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.Equal(tc.expcectedError, parseError)
		})
	}
}
