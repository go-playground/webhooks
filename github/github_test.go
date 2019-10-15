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
				"X-Hub-Signature": []string{"sha1=283d86ddd75a3a508102e968e901f815553fc82c"},
			},
		},
		{
			name:     "CheckSuiteEvent",
			event:    CheckSuiteEvent,
			typ:      CheckSuitePayload{},
			filename: "../testdata/github/check-suite.json",
			headers: http.Header{
				"X-Github-Event":  []string{"check_suite"},
				"X-Hub-Signature": []string{"sha1=38b7799610fa366e98b6c2bd09af31ece07d7063"},
			},
		},
		{
			name:     "CommitCommentEvent",
			event:    CommitCommentEvent,
			typ:      CommitCommentPayload{},
			filename: "../testdata/github/commit-comment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"commit_comment"},
				"X-Hub-Signature": []string{"sha1=b11cee473f9fd27648145dfeebea04164e9412b1"},
			},
		},
		{
			name:     "CreateEvent",
			event:    CreateEvent,
			typ:      CreatePayload{},
			filename: "../testdata/github/create.json",
			headers: http.Header{
				"X-Github-Event":  []string{"create"},
				"X-Hub-Signature": []string{"sha1=7f1f9400733d3db721ee88747da6d2944cadeef0"},
			},
		},
		{
			name:     "DeleteEvent",
			event:    DeleteEvent,
			typ:      DeletePayload{},
			filename: "../testdata/github/delete.json",
			headers: http.Header{
				"X-Github-Event":  []string{"delete"},
				"X-Hub-Signature": []string{"sha1=2eae3b0e47b867e7b92cfeddf9e32b1c287a4390"},
			},
		},
		{
			name:     "DeploymentEvent",
			event:    DeploymentEvent,
			typ:      DeploymentPayload{},
			filename: "../testdata/github/deployment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"deployment"},
				"X-Hub-Signature": []string{"sha1=4fc1d4b8b57b3c449d4641a48db4bc1524c6d085"},
			},
		},
		{
			name:     "DeploymentStatusEvent",
			event:    DeploymentStatusEvent,
			typ:      DeploymentStatusPayload{},
			filename: "../testdata/github/deployment-status.json",
			headers: http.Header{
				"X-Github-Event":  []string{"deployment_status"},
				"X-Hub-Signature": []string{"sha1=ced7fbc1885edb99b5b1ca0c3f857e6ac9816142"},
			},
		},
		{
			name:     "ForkEvent",
			event:    ForkEvent,
			typ:      ForkPayload{},
			filename: "../testdata/github/fork.json",
			headers: http.Header{
				"X-Github-Event":  []string{"fork"},
				"X-Hub-Signature": []string{"sha1=917d7252153839963b78c6ce984c66e2b23623b6"},
			},
		},
		{
			name:     "GollumEvent",
			event:    GollumEvent,
			typ:      GollumPayload{},
			filename: "../testdata/github/gollum.json",
			headers: http.Header{
				"X-Github-Event":  []string{"gollum"},
				"X-Hub-Signature": []string{"sha1=b6e1d3428880139208eca930248b949d1ddfc938"},
			},
		},
		{
			name:     "InstallationEvent",
			event:    InstallationEvent,
			typ:      InstallationPayload{},
			filename: "../testdata/github/installation.json",
			headers: http.Header{
				"X-Github-Event":  []string{"installation"},
				"X-Hub-Signature": []string{"sha1=74017b4ee1d0ed6fd2525dbd7321e996fe1443d3"},
			},
		},
		{
			name:     "InstallationRepositoriesEvent",
			event:    InstallationRepositoriesEvent,
			typ:      InstallationRepositoriesPayload{},
			filename: "../testdata/github/installation-repositories.json",
			headers: http.Header{
				"X-Github-Event":  []string{"installation_repositories"},
				"X-Hub-Signature": []string{"sha1=9ef54a43a7320412c76074649adef6f2743065d5"},
			},
		},
		// TODO IntegrationInstallationEvent is not listed in github docs
		// {
		// 	name:     "IntegrationInstallationEvent",
		// 	event:    IntegrationInstallationEvent,
		// 	typ:      InstallationPayload{},
		// 	filename: "../testdata/github/integration-installation.json",
		// 	headers: http.Header{
		// 		"X-Github-Event":  []string{"integration_installation"},
		// 		"X-Hub-Signature": []string{"sha1=bb2769f05f1a11af3a1edf8f9fac11bae7402a1e"},
		// 	},
		// },

		// TODO IntegrationInstallationRepositoriesEvent is not listed in github docs
		// {
		// 	name:     "IntegrationInstallationRepositoriesEvent",
		// 	event:    IntegrationInstallationRepositoriesEvent,
		// 	typ:      InstallationRepositoriesPayload{},
		// 	filename: "../testdata/github/integration-installation-repositories.json",
		// 	headers: http.Header{
		// 		"X-Github-Event":  []string{"integration_installation_repositories"},
		// 		"X-Hub-Signature": []string{"sha1=2f00a982574188342c2894eb9d1b1e93434687fb"},
		// 	},
		// },
		{
			name:     "IssueCommentEvent",
			event:    IssueCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "../testdata/github/issue-comment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"issue_comment"},
				"X-Hub-Signature": []string{"sha1=93121a2300c181e748857c02fdd42fcda7b0becb"},
			},
		},
		{
			name:     "IssuesEvent",
			event:    IssuesEvent,
			typ:      IssuesPayload{},
			filename: "../testdata/github/issues.json",
			headers: http.Header{
				"X-Github-Event":  []string{"issues"},
				"X-Hub-Signature": []string{"sha1=e9652a499e41cb07994e30d89ef86e658e7f85b6"},
			},
		},
		{
			name:     "LabelEvent",
			event:    LabelEvent,
			typ:      LabelPayload{},
			filename: "../testdata/github/label.json",
			headers: http.Header{
				"X-Github-Event":  []string{"label"},
				"X-Hub-Signature": []string{"sha1=35ee1ad349fe974ac19210e7c0ca932119ddb5d5"},
			},
		},
		{
			name:     "MemberEvent",
			event:    MemberEvent,
			typ:      MemberPayload{},
			filename: "../testdata/github/member.json",
			headers: http.Header{
				"X-Github-Event":  []string{"member"},
				"X-Hub-Signature": []string{"sha1=073e497440a70d36722e0bbf003559d2b2910b66"},
			},
		},
		{
			name:     "MembershipEvent",
			event:    MembershipEvent,
			typ:      MembershipPayload{},
			filename: "../testdata/github/membership.json",
			headers: http.Header{
				"X-Github-Event":  []string{"membership"},
				"X-Hub-Signature": []string{"sha1=35abc79a5c458805022baf75cf6b04f4a5200d94"},
			},
		},
		{
			name:     "MilestoneEvent",
			event:    MilestoneEvent,
			typ:      MilestonePayload{},
			filename: "../testdata/github/milestone.json",
			headers: http.Header{
				"X-Github-Event":  []string{"milestone"},
				"X-Hub-Signature": []string{"sha1=a3fe81287e8d43745fc1882f4d2ef856f9bc8a16"},
			},
		},
		{
			name:     "OrganizationEvent",
			event:    OrganizationEvent,
			typ:      OrganizationPayload{},
			filename: "../testdata/github/organization.json",
			headers: http.Header{
				"X-Github-Event":  []string{"organization"},
				"X-Hub-Signature": []string{"sha1=f05df4b6721ed45d310a99742e6fecfa606bcb0e"},
			},
		},
		{
			name:     "OrgBlockEvent",
			event:    OrgBlockEvent,
			typ:      OrgBlockPayload{},
			filename: "../testdata/github/org-block.json",
			headers: http.Header{
				"X-Github-Event":  []string{"org_block"},
				"X-Hub-Signature": []string{"sha1=98df61755a92678d8e75e6fd5eb961d788d7af93"},
			},
		},
		{
			name:     "PageBuildEvent",
			event:    PageBuildEvent,
			typ:      PageBuildPayload{},
			filename: "../testdata/github/page-build.json",
			headers: http.Header{
				"X-Github-Event":  []string{"page_build"},
				"X-Hub-Signature": []string{"sha1=1c2badf38c4657b2080d0c159d7a499b2ca11513"},
			},
		},
		// TODO PingEvent is not listed in github docs
		// {
		// 	name:     "PingEvent",
		// 	event:    PingEvent,
		// 	typ:      PingPayload{},
		// 	filename: "../testdata/github/ping.json",
		// 	headers: http.Header{
		// 		"X-Github-Event":  []string{"ping"},
		// 		"X-Hub-Signature": []string{"sha1=f80e1cfc04245b65228b86612119ab5c894133c2"},
		// 	},
		// },
		{
			name:     "ProjectCardEvent",
			event:    ProjectCardEvent,
			typ:      ProjectCardPayload{},
			filename: "../testdata/github/project-card.json",
			headers: http.Header{
				"X-Github-Event":  []string{"project_card"},
				"X-Hub-Signature": []string{"sha1=2a2765bf29885b3db578d98c248d76c208fced75"},
			},
		},
		{
			name:     "ProjectColumnEvent",
			event:    ProjectColumnEvent,
			typ:      ProjectColumnPayload{},
			filename: "../testdata/github/project-column.json",
			headers: http.Header{
				"X-Github-Event":  []string{"project_column"},
				"X-Hub-Signature": []string{"sha1=31624226b9ea5d3825e6b3fcd0f99ddf091cf39c"},
			},
		},
		{
			name:     "ProjectEvent",
			event:    ProjectEvent,
			typ:      ProjectPayload{},
			filename: "../testdata/github/project.json",
			headers: http.Header{
				"X-Github-Event":  []string{"project"},
				"X-Hub-Signature": []string{"sha1=ebac78c119e31b1f591b9fe3a692e5d670831380"},
			},
		},
		{
			name:     "PublicEvent",
			event:    PublicEvent,
			typ:      PublicPayload{},
			filename: "../testdata/github/public.json",
			headers: http.Header{
				"X-Github-Event":  []string{"public"},
				"X-Hub-Signature": []string{"sha1=174dd61f161dad3d0709d5c314ecad00cdcef73b"},
			},
		},
		{
			name:     "PullRequestEvent",
			event:    PullRequestEvent,
			typ:      PullRequestPayload{},
			filename: "../testdata/github/pull-request.json",
			headers: http.Header{
				"X-Github-Event":  []string{"pull_request"},
				"X-Hub-Signature": []string{"sha1=c488ac6544c2f4526811899e22fb62cea4fcda03"},
			},
		},
		{
			name:     "PullRequestReviewEvent",
			event:    PullRequestReviewEvent,
			typ:      PullRequestReviewPayload{},
			filename: "../testdata/github/pull-request-review.json",
			headers: http.Header{
				"X-Github-Event":  []string{"pull_request_review"},
				"X-Hub-Signature": []string{"sha1=ad813b10b1ca686b1ed7e34ed4820fd23151c45b"},
			},
		},
		{
			name:     "PullRequestReviewCommentEvent",
			event:    PullRequestReviewCommentEvent,
			typ:      PullRequestReviewCommentPayload{},
			filename: "../testdata/github/pull-request-review-comment.json",
			headers: http.Header{
				"X-Github-Event":  []string{"pull_request_review_comment"},
				"X-Hub-Signature": []string{"sha1=390a08a10a1388852dc27d060e1a38ffa198b4c6"},
			},
		},
		{
			name:     "PushEvent",
			event:    PushEvent,
			typ:      PushPayload{},
			filename: "../testdata/github/push.json",
			headers: http.Header{
				"X-Github-Event":  []string{"push"},
				"X-Hub-Signature": []string{"sha1=08da466a59780b773016d8c82dd003bad1c4aa5c"},
			},
		},
		{
			name:     "ReleaseEvent",
			event:    ReleaseEvent,
			typ:      ReleasePayload{},
			filename: "../testdata/github/release.json",
			headers: http.Header{
				"X-Github-Event":  []string{"release"},
				"X-Hub-Signature": []string{"sha1=74a656e593c579d59ed0442acb3323a502855402"},
			},
		},
		{
			name:     "RepositoryEvent",
			event:    RepositoryEvent,
			typ:      RepositoryPayload{},
			filename: "../testdata/github/repository.json",
			headers: http.Header{
				"X-Github-Event":  []string{"repository"},
				"X-Hub-Signature": []string{"sha1=c0b344f0434f9bf9c136ba08725873a76636e6f7"},
			},
		},
		{
			name:     "RepositoryVulnerabilityAlertEvent",
			event:    RepositoryVulnerabilityAlertEvent,
			typ:      RepositoryVulnerabilityAlertPayload{},
			filename: "../testdata/github/repository-vulnerability-alert.json",
			headers: http.Header{
				"X-Github-Event":  []string{"repository_vulnerability_alert"},
				"X-Hub-Signature": []string{"sha1=158f1c382a98c44b65e63bfe408abc790eae5909"},
			},
		},
		{
			name:     "SecurityAdvisoryEvent",
			event:    SecurityAdvisoryEvent,
			typ:      SecurityAdvisoryPayload{},
			filename: "../testdata/github/security-advisory.json",
			headers: http.Header{
				"X-Github-Event":  []string{"security_advisory"},
				"X-Hub-Signature": []string{"sha1=e05faeac4f88a114a78d95ff2f8cadd81b100381"},
			},
		},
		{
			name:     "StatusEvent",
			event:    StatusEvent,
			typ:      StatusPayload{},
			filename: "../testdata/github/status.json",
			headers: http.Header{
				"X-Github-Event":  []string{"status"},
				"X-Hub-Signature": []string{"sha1=76bd9fb549bd1f9468d9faff8019c0a377e32d55"},
			},
		},
		{
			name:     "TeamEvent",
			event:    TeamEvent,
			typ:      TeamPayload{},
			filename: "../testdata/github/team.json",
			headers: http.Header{
				"X-Github-Event":  []string{"team"},
				"X-Hub-Signature": []string{"sha1=42463baff0eb057ba02bb446a5a1e916f10b3fd5"},
			},
		},
		{
			name:     "TeamAddEvent",
			event:    TeamAddEvent,
			typ:      TeamAddPayload{},
			filename: "../testdata/github/team-add.json",
			headers: http.Header{
				"X-Github-Event":  []string{"team_add"},
				"X-Hub-Signature": []string{"sha1=d4f592285c079a618703d83fab925c1e28498716"},
			},
		},
		{
			name:     "WatchEvent",
			event:    WatchEvent,
			typ:      WatchPayload{},
			filename: "../testdata/github/watch.json",
			headers: http.Header{
				"X-Github-Event":  []string{"watch"},
				"X-Hub-Signature": []string{"sha1=2693ef71ae69be979bbcd4b8b95fcbe130ab393d"},
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
