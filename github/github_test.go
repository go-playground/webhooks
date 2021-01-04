package github

import (
	"bytes"
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
				"X-Github-Event": []string{"noneexistant_event"},
			},
		},
		{
			name:    "BadBody",
			event:   CommitCommentEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Github-Event":  []string{"commit_comment"},
				"X-Hub-Signature": []string{"sha1=156404ad5f721c53151147f3d3d302329f95a3ab"},
			},
		},
		{
			name:    "BadSignatureLength",
			event:   CommitCommentEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Github-Event":  []string{"commit_comment"},
				"X-Hub-Signature": []string{""},
			},
		},
		{
			name:    "BadSignatureMatch",
			event:   CommitCommentEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Github-Event":  []string{"commit_comment"},
				"X-Hub-Signature": []string{"sha1=111"},
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
			name:     "CheckRunEvent",
			event:    CheckRunEvent,
			typ:      CheckRunPayload{},
			filename: "../testdata/github/check-run.json",
			headers: http.Header{
				"X-Github-Event":  []string{"check_run"},
				"X-Hub-Signature": []string{"sha1=229f4920493b455398168cd86dc6b366064bdf3f"},
			},
		},
		{
			name:     "CheckSuiteEvent",
			event:    CheckSuiteEvent,
			typ:      CheckSuitePayload{},
			filename: "../testdata/github/check-suite.json",
			headers: http.Header{
				"X-Github-Event":  []string{"check_suite"},
				"X-Hub-Signature": []string{"sha1=250ad5a340f8d91e67dc5682342f3190fd2006a1"},
			},
		},
		{
			name:     "CommitCommentEvent",
			event:    CommitCommentEvent,
			typ:      CommitCommentPayload{},
			filename: "../testdata/github/commit-comment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"commit_comment"},
				"X-Hub-Signature": []string{"sha1=156404ad5f721c53151147f3d3d302329f95a3ab"},
			},
		},
		{
			name:     "CreateEvent",
			event:    CreateEvent,
			typ:      CreatePayload{},
			filename: "../testdata/github/create.json",
			headers: http.Header{
				"X-Github-Event":  []string{"create"},
				"X-Hub-Signature": []string{"sha1=77ff16ca116034bbeed77ebfce83b36572a9cbaf"},
			},
		},
		{
			name:     "DeleteEvent",
			event:    DeleteEvent,
			typ:      DeletePayload{},
			filename: "../testdata/github/delete.json",
			headers: http.Header{
				"X-Github-Event":  []string{"delete"},
				"X-Hub-Signature": []string{"sha1=4ddef04fd05b504c7041e294fca3ad1804bc7be1"},
			},
		},
		{
			name:     "DeploymentEvent",
			event:    DeploymentEvent,
			typ:      DeploymentPayload{},
			filename: "../testdata/github/deployment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"deployment"},
				"X-Hub-Signature": []string{"sha1=bb47dc63ceb764a6b1f14fe123e299e5b814c67c"},
			},
		},
		{
			name:     "DeploymentStatusEvent",
			event:    DeploymentStatusEvent,
			typ:      DeploymentStatusPayload{},
			filename: "../testdata/github/deployment-status.json",
			headers: http.Header{
				"X-Github-Event":  []string{"deployment_status"},
				"X-Hub-Signature": []string{"sha1=1b2ce08e0c3487fdf22bed12c63dc734cf6dc8a4"},
			},
		},
		{
			name:     "ForkEvent",
			event:    ForkEvent,
			typ:      ForkPayload{},
			filename: "../testdata/github/fork.json",
			headers: http.Header{
				"X-Github-Event":  []string{"fork"},
				"X-Hub-Signature": []string{"sha1=cec5f8fb7c383514c622d3eb9e121891dfcca848"},
			},
		},
		{
			name:     "GollumEvent",
			event:    GollumEvent,
			typ:      GollumPayload{},
			filename: "../testdata/github/gollum.json",
			headers: http.Header{
				"X-Github-Event":  []string{"gollum"},
				"X-Hub-Signature": []string{"sha1=a375a6dc8ceac7231ee022211f8eb85e2a84a5b9"},
			},
		},
		{
			name:     "InstallationEvent",
			event:    InstallationEvent,
			typ:      InstallationPayload{},
			filename: "../testdata/github/installation.json",
			headers: http.Header{
				"X-Github-Event":  []string{"installation"},
				"X-Hub-Signature": []string{"sha1=2058cf6cc28570710afbc638e669f5c67305a2db"},
			},
		},
		{
			name:     "InstallationRepositoriesEvent",
			event:    InstallationRepositoriesEvent,
			typ:      InstallationRepositoriesPayload{},
			filename: "../testdata/github/installation-repositories.json",
			headers: http.Header{
				"X-Github-Event":  []string{"installation_repositories"},
				"X-Hub-Signature": []string{"sha1=c587fbd9dd169db8ae592b3bcc80b08e2e6f4f45"},
			},
		},
		{
			name:     "IntegrationInstallationEvent",
			event:    IntegrationInstallationEvent,
			typ:      InstallationPayload{},
			filename: "../testdata/github/integration-installation.json",
			headers: http.Header{
				"X-Github-Event":  []string{"integration_installation"},
				"X-Hub-Signature": []string{"sha1=bb2769f05f1a11af3a1edf8f9fac11bae7402a1e"},
			},
		},
		{
			name:     "IntegrationInstallationRepositoriesEvent",
			event:    IntegrationInstallationRepositoriesEvent,
			typ:      InstallationRepositoriesPayload{},
			filename: "../testdata/github/integration-installation-repositories.json",
			headers: http.Header{
				"X-Github-Event":  []string{"integration_installation_repositories"},
				"X-Hub-Signature": []string{"sha1=2f00a982574188342c2894eb9d1b1e93434687fb"},
			},
		},
		{
			name:     "IssueCommentEvent",
			event:    IssueCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "../testdata/github/issue-comment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"issue_comment"},
				"X-Hub-Signature": []string{"sha1=e724c9f811fcf5f511aac32e4251b08ab1a0fd87"},
			},
		},
		{
			name:     "PullRequestIssueCommentEvent",
			event:    IssueCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "../testdata/github/pull-request-issue-comment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"issue_comment"},
				"X-Hub-Signature": []string{"sha1=6c969b99ef881b5c98b2dbfc66a34465fcf0e7d4"},
			},
		},
		{
			name:     "IssuesEvent",
			event:    IssuesEvent,
			typ:      IssuesPayload{},
			filename: "../testdata/github/issues.json",
			headers: http.Header{
				"X-Github-Event":  []string{"issues"},
				"X-Hub-Signature": []string{"sha1=dfc9a3428f3df86e4ecd78e34b41c55bba5d0b21"},
			},
		},
		{
			name:     "LabelEvent",
			event:    LabelEvent,
			typ:      LabelPayload{},
			filename: "../testdata/github/label.json",
			headers: http.Header{
				"X-Github-Event":  []string{"label"},
				"X-Hub-Signature": []string{"sha1=efc13e7ad816235222e4a6b3f96d3fd1e162dbd4"},
			},
		},
		{
			name:     "MemberEvent",
			event:    MemberEvent,
			typ:      MemberPayload{},
			filename: "../testdata/github/member.json",
			headers: http.Header{
				"X-Github-Event":  []string{"member"},
				"X-Hub-Signature": []string{"sha1=597e7d6627a6636d4c3283e36631983fbd57bdd0"},
			},
		},
		{
			name:     "MembershipEvent",
			event:    MembershipEvent,
			typ:      MembershipPayload{},
			filename: "../testdata/github/membership.json",
			headers: http.Header{
				"X-Github-Event":  []string{"membership"},
				"X-Hub-Signature": []string{"sha1=16928c947b3707b0efcf8ceb074a5d5dedc9c76e"},
			},
		},
		{
			name:     "MilestoneEvent",
			event:    MilestoneEvent,
			typ:      MilestonePayload{},
			filename: "../testdata/github/milestone.json",
			headers: http.Header{
				"X-Github-Event":  []string{"milestone"},
				"X-Hub-Signature": []string{"sha1=8b63f58ea58e6a59dcfc5ecbaea0d1741a6bf9ec"},
			},
		},
		{
			name:     "OrganizationEvent",
			event:    OrganizationEvent,
			typ:      OrganizationPayload{},
			filename: "../testdata/github/organization.json",
			headers: http.Header{
				"X-Github-Event":  []string{"organization"},
				"X-Hub-Signature": []string{"sha1=7e5ad88557be0a05fb89e86c7893d987386aa0d5"},
			},
		},
		{
			name:     "OrgBlockEvent",
			event:    OrgBlockEvent,
			typ:      OrgBlockPayload{},
			filename: "../testdata/github/org-block.json",
			headers: http.Header{
				"X-Github-Event":  []string{"org_block"},
				"X-Hub-Signature": []string{"sha1=21fe61da3f014c011edb60b0b9dfc9aa7059a24b"},
			},
		},
		{
			name:     "PageBuildEvent",
			event:    PageBuildEvent,
			typ:      PageBuildPayload{},
			filename: "../testdata/github/page-build.json",
			headers: http.Header{
				"X-Github-Event":  []string{"page_build"},
				"X-Hub-Signature": []string{"sha1=b3abad8f9c1b3fc0b01c4eb107447800bb5000f9"},
			},
		},
		{
			name:     "PingEvent",
			event:    PingEvent,
			typ:      PingPayload{},
			filename: "../testdata/github/ping.json",
			headers: http.Header{
				"X-Github-Event":  []string{"ping"},
				"X-Hub-Signature": []string{"sha1=f80e1cfc04245b65228b86612119ab5c894133c2"},
			},
		},
		{
			name:     "ProjectCardEvent",
			event:    ProjectCardEvent,
			typ:      ProjectCardPayload{},
			filename: "../testdata/github/project-card.json",
			headers: http.Header{
				"X-Github-Event":  []string{"project_card"},
				"X-Hub-Signature": []string{"sha1=f5ed1572b04f0e03c8d5f5e3f7fa63737bef76d7"},
			},
		},
		{
			name:     "ProjectColumnEvent",
			event:    ProjectColumnEvent,
			typ:      ProjectColumnPayload{},
			filename: "../testdata/github/project-column.json",
			headers: http.Header{
				"X-Github-Event":  []string{"project_column"},
				"X-Hub-Signature": []string{"sha1=7d5dd49d9863e982a4f577170717ea8350a69db0"},
			},
		},
		{
			name:     "ProjectEvent",
			event:    ProjectEvent,
			typ:      ProjectPayload{},
			filename: "../testdata/github/project.json",
			headers: http.Header{
				"X-Github-Event":  []string{"project"},
				"X-Hub-Signature": []string{"sha1=7295ab4f205434208f1b86edf2b55adae34c6c92"},
			},
		},
		{
			name:     "PublicEvent",
			event:    PublicEvent,
			typ:      PublicPayload{},
			filename: "../testdata/github/public.json",
			headers: http.Header{
				"X-Github-Event":  []string{"public"},
				"X-Hub-Signature": []string{"sha1=73edb2a8c69c1ac35efb797ede3dc2cde618c10c"},
			},
		},
		{
			name:     "PullRequestEvent",
			event:    PullRequestEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/github/pull-request.json",
			headers: http.Header{
				"X-Github-Event":  []string{"pull_request"},
				"X-Hub-Signature": []string{"sha1=88972f972db301178aa13dafaf112d26416a15e6"},
			},
		},
		{
			name:     "PullRequestReviewEvent",
			event:    PullRequestReviewEvent,
			typ:      PullRequestReviewPayload{},
			filename: "../testdata/github/pull-request-review.json",
			headers: http.Header{
				"X-Github-Event":  []string{"pull_request_review"},
				"X-Hub-Signature": []string{"sha1=55345ce92be7849f97d39b9426b95261d4bd4465"},
			},
		},
		{
			name:     "PullRequestReviewCommentEvent",
			event:    PullRequestReviewCommentEvent,
			typ:      PullRequestReviewCommentPayload{},
			filename: "../testdata/github/pull-request-review-comment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"pull_request_review_comment"},
				"X-Hub-Signature": []string{"sha1=a9ece15dbcbb85fa5f00a0bf409494af2cbc5b60"},
			},
		},
		{
			name:     "PushEvent",
			event:    PushEvent,
			typ:      PushPayload{},
			filename: "../testdata/github/push.json",
			headers: http.Header{
				"X-Github-Event":  []string{"push"},
				"X-Hub-Signature": []string{"sha1=0534736f52c2fc5896ef1bd5a043127b20d233ba"},
			},
		},
		{
			name:     "ReleaseEvent",
			event:    ReleaseEvent,
			typ:      ReleasePayload{},
			filename: "../testdata/github/release.json",
			headers: http.Header{
				"X-Github-Event":  []string{"release"},
				"X-Hub-Signature": []string{"sha1=e62bb4c51bc7dde195b9525971c2e3aecb394390"},
			},
		},
		{
			name:     "RepositoryEvent",
			event:    RepositoryEvent,
			typ:      RepositoryPayload{},
			filename: "../testdata/github/repository.json",
			headers: http.Header{
				"X-Github-Event":  []string{"repository"},
				"X-Hub-Signature": []string{"sha1=df442a8af41edd2d42ccdd997938d1d111b0f94e"},
			},
		},
		{
			name:     "RepositoryVulnerabilityAlertEvent",
			event:    RepositoryVulnerabilityAlertEvent,
			typ:      RepositoryVulnerabilityAlertPayload{},
			filename: "../testdata/github/repository-vulnerability-alert.json",
			headers: http.Header{
				"X-Github-Event":  []string{"repository_vulnerability_alert"},
				"X-Hub-Signature": []string{"sha1=c42c0649e7e06413bcd756763edbab48dff400db"},
			},
		},
		{
			name:     "SecurityAdvisoryEvent",
			event:    SecurityAdvisoryEvent,
			typ:      SecurityAdvisoryPayload{},
			filename: "../testdata/github/security-advisory.json",
			headers: http.Header{
				"X-Github-Event":  []string{"security_advisory"},
				"X-Hub-Signature": []string{"sha1=6a71f24fa69f55469843a91dc3a5c3e29714a565"},
			},
		},
		{
			name:     "StatusEvent",
			event:    StatusEvent,
			typ:      StatusPayload{},
			filename: "../testdata/github/status.json",
			headers: http.Header{
				"X-Github-Event":  []string{"status"},
				"X-Hub-Signature": []string{"sha1=3caa5f062a2deb7cce1482314bb9b4c99bf0ab45"},
			},
		},
		{
			name:     "TeamEvent",
			event:    TeamEvent,
			typ:      TeamPayload{},
			filename: "../testdata/github/team.json",
			headers: http.Header{
				"X-Github-Event":  []string{"team"},
				"X-Hub-Signature": []string{"sha1=ff5b5d58faec10bd40fc96834148df408e7a4608"},
			},
		},
		{
			name:     "TeamAddEvent",
			event:    TeamAddEvent,
			typ:      TeamAddPayload{},
			filename: "../testdata/github/team-add.json",
			headers: http.Header{
				"X-Github-Event":  []string{"team_add"},
				"X-Hub-Signature": []string{"sha1=5f3953476e270b79cc6763780346110da880609a"},
			},
		},
		{
			name:     "WatchEvent",
			event:    WatchEvent,
			typ:      WatchPayload{},
			filename: "../testdata/github/watch.json",
			headers: http.Header{
				"X-Github-Event":  []string{"watch"},
				"X-Hub-Signature": []string{"sha1=a317bcfe69ccb8bece74c20c7378e5413c4772f1"},
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
