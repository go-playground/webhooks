package pepo

import (
	"bytes"
	"encoding/json"
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
		payload io.Reader
		headers http.Header
	}{
		{
			name:    "ErrMissingTimestampHeader",
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"pepo-signature": []string{"sha1=229f4920493b455398168cd86dc6b366064bdf3f"},
				"pepo-version":   []string{"1"},
			},
		},
		{
			name:    "ErrMissingSignatureHeader",
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"pepo-timestamp": []string{"1584103295"},
				"pepo-version":   []string{"1"},
			},
		},
		{
			name:    "ErrMissingVersionHeader",
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"pepo-timestamp": []string{"1584103295"},
				"pepo-signature": []string{"sha1=229f4920493b455398168cd86dc6b366064bdf3f"},
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
				_, parseError = hook.Parse(r)
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
		typ      interface{}
		filename string
		headers  http.Header
	}{
		{
			name:     "VideoUpdate",
			typ:      EventPayload{},
			filename: "../testdata/pepo/video-update.json",
			headers: http.Header{
				"pepo-signature": []string{"sha1=229f4920493b455398168cd86dc6b366064bdf3f"},
				"pepo-timestamp": []string{"1584103295"},
				"pepo-version":   []string{"1"},
			},
		},
		{
			name:     "VideoContribution",
			typ:      EventPayload{},
			filename: "../testdata/pepo/video-contribution.json",
			headers: http.Header{
				"pepo-signature": []string{"sha1=229f4920493b455398168cd86dc6b366064bdf3f"},
				"pepo-timestamp": []string{"1584103295"},
				"pepo-version":   []string{"1"},
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
			var result interface{}
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				result, parseError = hook.Parse(r)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)

			t.Log(result)

			resultsMarshalled, _ := json.Marshal(result)
			t.Log(string(resultsMarshalled))

			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(result))
		})
	}
}
