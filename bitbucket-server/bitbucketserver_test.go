package bitbucketserver

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

const (
	path = "/webhooks"
)

var hook *Webhook

func TestMain(m *testing.M) {
	// setup
	var err error
	hook, err = New(Options.Secret("secret"))
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
			event:   RepositoryReferenceChangedEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{},
		},
		{
			name:    "BadSignatureLength",
			event:   RepositoryReferenceChangedEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Event-Key":     []string{"repo:refs_changed"},
				"X-Hub-Signature": []string{""},
			},
		},
		{
			name:    "BadSignatureMatch",
			event:   RepositoryReferenceChangedEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Event-Key":     []string{"repo:refs_changed"},
				"X-Hub-Signature": []string{"sha256=111"},
			},
		},
		{
			name:    "UnsubscribedEvent",
			event:   RepositoryReferenceChangedEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Event-Key":     []string{"nonexistent_event"},
				"X-Hub-Signature": []string{"sha256=111"},
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
		name        string
		event       Event
		payloadType interface{}
		filename    string
		headers     http.Header
	}{
		{
			name:        "Repository refs updated",
			event:       RepositoryReferenceChangedEvent,
			payloadType: RepositoryReferenceChangedPayload{},
			filename:    "../testdata/bitbucket-server/repo-refs-changed.json",
			headers: http.Header{
				"X-Event-Key":     []string{"repo:refs_changed"},
				"X-Hub-Signature": []string{"sha256=8a60f7487d167f55886df87d4077192035d76f76a8e0b3a48fd8ae8cad25f391"},
			},
		},
		{
			name:        "Repository modified",
			event:       RepositoryModifiedEvent,
			payloadType: RepositoryModifiedPayload{},
			filename:    "../testdata/bitbucket-server/repo-modified.json",
			headers: http.Header{
				"X-Event-Key":     []string{"repo:modified"},
				"X-Hub-Signature": []string{"sha256=1511ed69d7697ede1699b0217e17b7d0b492eeccc9a5649d5d30dd84f0e5a89a"},
			},
		},
		{
			name:        "Repository forked",
			event:       RepositoryForkedEvent,
			payloadType: RepositoryForkedPayload{},
			filename:    "../testdata/bitbucket-server/repo-forked.json",
			headers: http.Header{
				"X-Event-Key":     []string{"repo:forked"},
				"X-Hub-Signature": []string{"sha256=d34115023042f9e7ef7020200650e2f34da137e0217708475f9b749ad889a16d"},
			},
		},
		{
			name:        "Repository commit comment edited",
			event:       RepositoryCommentEditedEvent,
			payloadType: RepositoryCommentEditedPayload{},
			filename:    "../testdata/bitbucket-server/repo-comment-edited.json",
			headers: http.Header{
				"X-Event-Key":     []string{"repo:comment:edited"},
				"X-Hub-Signature": []string{"sha256=90a8f4898d8dd7a4ef99e33a7f1d86dd3645f45b2a5b59110493cc4b3062a712"},
			},
		},
		{
			name:        "Repository commit comment deleted",
			event:       RepositoryCommentDeletedEvent,
			payloadType: RepositoryCommentDeletedPayload{},
			filename:    "../testdata/bitbucket-server/repo-comment-deleted.json",
			headers: http.Header{
				"X-Event-Key":     []string{"repo:comment:deleted"},
				"X-Hub-Signature": []string{"sha256=e8b6d3d1581366c9f65949c93149b29ba33252f6afa807432f8f823fb08680e7"},
			},
		},
		{
			name:        "Repository commit comment added",
			event:       RepositoryCommentAddedEvent,
			payloadType: RepositoryCommentAddedPayload{},
			filename:    "../testdata/bitbucket-server/repo-comment-added.json",
			headers: http.Header{
				"X-Event-Key":     []string{"repo:comment:added"},
				"X-Hub-Signature": []string{"sha256=80b121d53ec48bb3f8bed9243ba53be62c8eb7d1ce0395ca87fef1938bf9620e"},
			},
		},
		{
			name:        "Pull request unapproved",
			event:       PullRequestReviewerUnapprovedEvent,
			payloadType: PullRequestReviewerUnapprovedPayload{},
			filename:    "../testdata/bitbucket-server/pr-reviewer-unapproved.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:reviewer:unapproved"},
				"X-Hub-Signature": []string{"sha256=9822024378738817dc85c0a41feb9fa4825058d28a9a1ee7065bfacd6a04c7c1"},
			},
		},
		{
			name:        "Pull request reviewer updated",
			event:       PullRequestReviewerUpdatedEvent,
			payloadType: PullRequestReviewerUpdatedPayload{},
			filename:    "../testdata/bitbucket-server/pr-reviewer-updated.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:reviewer:updated"},
				"X-Hub-Signature": []string{"sha256=07ca94a0a5c5913819a16ce0414f976023aeb0065fa9d80f990aad7f1d936be5"},
			},
		},
		{
			name:        "Pull request opened",
			event:       PullRequestOpenedEvent,
			payloadType: PullRequestOpenedPayload{},
			filename:    "../testdata/bitbucket-server/pr-opened.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:opened"},
				"X-Hub-Signature": []string{"sha256=b82c323978a741aa256c0a6bfa13a8f211e1795bd8ddb2641ced122769f7a7c6"},
			},
		},
		{
			name:        "Pull request modified",
			event:       PullRequestModifiedEvent,
			payloadType: PullRequestModifiedPayload{},
			filename:    "../testdata/bitbucket-server/pr-modified.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:modified"},
				"X-Hub-Signature": []string{"sha256=1e307462390ff6f0c59fcdb8eb4b2977058b5cbc502a24a0db385f5331136227"},
			},
		},
		{
			name:        "Pull request merged",
			event:       PullRequestMergedEvent,
			payloadType: PullRequestMergedPayload{},
			filename:    "../testdata/bitbucket-server/pr-merged.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:merged"},
				"X-Hub-Signature": []string{"sha256=adbee42ddd6a178b0c1160e89f666b53fb8c76495f782a4e3055e3fbee232704"},
			},
		},
		{
			name:        "Pull request marked needs work",
			event:       PullRequestReviewerNeedsWorkEvent,
			payloadType: PullRequestReviewerNeedsWorkPayload{},
			filename:    "../testdata/bitbucket-server/pr-reviewer-needs-work.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:reviewer:needs_work"},
				"X-Hub-Signature": []string{"sha256=3d10aadcede2131674654bb48c10fe904b0b2ed3d3b283bdc5c64dbc4856582d"},
			},
		},
		{
			name:        "Pull request deleted",
			event:       PullRequestDeletedEvent,
			payloadType: PullRequestDeletedPayload{},
			filename:    "../testdata/bitbucket-server/pr-deleted.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:deleted"},
				"X-Hub-Signature": []string{"sha256=657c5d9839c0e3c1c95e5ceceacb07f0e372328883dab6e25bb619ee8b19a359"},
			},
		},
		{
			name:        "Pull request declined",
			event:       PullRequestDeclinedEvent,
			payloadType: PullRequestDeclinedPayload{},
			filename:    "../testdata/bitbucket-server/pr-declined.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:declined"},
				"X-Hub-Signature": []string{"sha256=e323ab4d057f32475340d03c90aa9ec20cd3a96c15200d75e59f221c14053528"},
			},
		},
		{
			name:        "Pull request comment edited",
			event:       PullRequestCommentEditedEvent,
			payloadType: PullRequestCommentEditedPayload{},
			filename:    "../testdata/bitbucket-server/pr-comment-edited.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:comment:edited"},
				"X-Hub-Signature": []string{"sha256=66580d1b02904e470cb48c1333452ea0748aecc3a9806f5a0f949be3a8b0a5ec"},
			},
		},
		{
			name:        "Pull request comment deleted",
			event:       PullRequestCommentDeletedEvent,
			payloadType: PullRequestCommentDeletedPayload{},
			filename:    "../testdata/bitbucket-server/pr-comment-deleted.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:comment:deleted"},
				"X-Hub-Signature": []string{"sha256=7c9575a6e9b141e063ef34e5066ebe67c3a5b59241e633d1332d70aba468fd04"},
			},
		},
		{
			name:        "Pull request comment added",
			event:       PullRequestReviewerApprovedEvent,
			payloadType: PullRequestReviewerApprovedPayload{},
			filename:    "../testdata/bitbucket-server/pr-reviewer-approved.json",
			headers: http.Header{
				"X-Event-Key":     []string{"pr:reviewer:approved"},
				"X-Hub-Signature": []string{"sha256=a8b78d774dea02f234069f724ee6c6a3c5c13fc3a3b856dac0a33d8ed9ec1823"},
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
			assert.Equal(reflect.TypeOf(tc.payloadType), reflect.TypeOf(results))
		})
	}
}
