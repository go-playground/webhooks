package googlecalendar

import (
	"fmt"
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
	path = "/webhooks"
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
	mux.HandleFunc(path, handler)
	return httptest.NewServer(mux)
}

var basicHeader = map[string][]string{
	"X-Goog-Channel-ID":         {"channel-ID-value"},
	"X-Goog-Channel-Token":      {"channel-token-value"},
	"X-Goog-Channel-Expiration": {"Tue, 19 Nov 2013 01:13:52 GMT"},
	"X-Goog-Resource-ID":        {"identifier-for-the-watched-resource"},
	"X-Goog-Resource-URI":       {"version-specific-URI-of-the-watched-resource"},
	"X-Goog-Resource-State":     {"sync"},
	"X-Goog-Message-Number":     {"1"},
}

func TestWebhooks(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name    string
		event   Event
		typ     interface{}
		headers http.Header
	}{
		{
			name:    "SyncEvent",
			event:   SyncEvent,
			typ:     &GoogleCalendarPayload{},
			headers: basicHeader,
		},
		{
			name:    "ExistsEvent",
			event:   ExistsEvent,
			typ:     &GoogleCalendarPayload{},
			headers: basicHeader,
		},
		{
			name:    "NotExistsEvent",
			event:   NotExistsEvent,
			typ:     &GoogleCalendarPayload{},
			headers: basicHeader,
		},
	}

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var parseError error
			var results interface{}
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				tt.headers["X-Goog-Resource-State"] = []string{string(tt.event)}
				r.Header = tt.headers
				results, parseError = hook.Parse(r, tc.event)
				if parseError != nil {
					fmt.Println(parseError)
				}
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, nil)
			assert.NoError(err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)

			gc := results.(*GoogleCalendarPayload)
			assert.Equal(gc.ResourceState, string(tt.event))
			assert.Equal(gc.MessageNumber, 1)

			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}
