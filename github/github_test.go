package github

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

func TestBadNoEventHeader(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, CreateEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrMissingGithubEventHeader)
}

func TestUnsubscribedEvent(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, CreateEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "noneexistant_event")

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
		_, parseError = hook.Parse(r, CommitCommentEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "commit_comment")
	req.Header.Set("X-Hub-Signature", "sha1=156404ad5f721c53151147f3d3d302329f95a3ab")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrParsingPayload)
}

func TestBadSignatureLength(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, CommitCommentEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "commit_comment")
	req.Header.Set("X-Hub-Signature", "")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrMissingHubSignatureHeader)
}

func TestBadSignatureMatch(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, CommitCommentEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "commit_comment")
	req.Header.Set("X-Hub-Signature", "sha1=111")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrHMACVerificationFailed)
}

func TestCommitCommentEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/commit-comment.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, CommitCommentEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "commit_comment")
	req.Header.Set("X-Hub-Signature", "sha1=156404ad5f721c53151147f3d3d302329f95a3ab")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(CommitCommentPayload)
	Equal(t, ok, true)
}

func TestCreateEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/create.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, CreateEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "create")
	req.Header.Set("X-Hub-Signature", "sha1=77ff16ca116034bbeed77ebfce83b36572a9cbaf")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(CreatePayload)
	Equal(t, ok, true)
}

func TestDeleteEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/delete.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, DeleteEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "delete")
	req.Header.Set("X-Hub-Signature", "sha1=4ddef04fd05b504c7041e294fca3ad1804bc7be1")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(DeletePayload)
	Equal(t, ok, true)
}

func TestDeploymentEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/deployment.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, DeploymentEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "deployment")
	req.Header.Set("X-Hub-Signature", "sha1=bb47dc63ceb764a6b1f14fe123e299e5b814c67c")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(DeploymentPayload)
	Equal(t, ok, true)
}

func TestDeploymentStatusEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/deployment-status.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, DeploymentStatusEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "deployment_status")
	req.Header.Set("X-Hub-Signature", "sha1=1b2ce08e0c3487fdf22bed12c63dc734cf6dc8a4")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(DeploymentStatusPayload)
	Equal(t, ok, true)
}

func TestForkEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/fork.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, ForkEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "fork")
	req.Header.Set("X-Hub-Signature", "sha1=cec5f8fb7c383514c622d3eb9e121891dfcca848")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(ForkPayload)
	Equal(t, ok, true)
}

func TestGollumEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/gollum.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, GollumEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "gollum")
	req.Header.Set("X-Hub-Signature", "sha1=a375a6dc8ceac7231ee022211f8eb85e2a84a5b9")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(GollumPayload)
	Equal(t, ok, true)
}

func TestInstallationEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/installation.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, InstallationEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "installation")
	req.Header.Set("X-Hub-Signature", "sha1=2058cf6cc28570710afbc638e669f5c67305a2db")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(InstallationPayload)
	Equal(t, ok, true)
}

func TestIntegrationInstallationEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/integration-installation.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IntegrationInstallationEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "integration_installation")
	req.Header.Set("X-Hub-Signature", "sha1=bb2769f05f1a11af3a1edf8f9fac11bae7402a1e")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(InstallationPayload)
	Equal(t, ok, true)
}

func TestIssueCommentEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/issue-comment.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssueCommentEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "issue_comment")
	req.Header.Set("X-Hub-Signature", "sha1=e724c9f811fcf5f511aac32e4251b08ab1a0fd87")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssueCommentPayload)
	Equal(t, ok, true)
}

func TestIssuesEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/issues.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssuesEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "issues")
	req.Header.Set("X-Hub-Signature", "sha1=dfc9a3428f3df86e4ecd78e34b41c55bba5d0b21")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssuesPayload)
	Equal(t, ok, true)
}

func TestLabelEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/label.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, LabelEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "label")
	req.Header.Set("X-Hub-Signature", "sha1=efc13e7ad816235222e4a6b3f96d3fd1e162dbd4")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(LabelPayload)
	Equal(t, ok, true)
}

func TestMemberEvent(t *testing.T) {
	payload, err := os.Open("../testdata/github/member.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, MemberEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "member")
	req.Header.Set("X-Hub-Signature", "sha1=597e7d6627a6636d4c3283e36631983fbd57bdd0")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(MemberPayload)
	Equal(t, ok, true)
}

func TestMembershipEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/membership.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, MembershipEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "membership")
	req.Header.Set("X-Hub-Signature", "sha1=16928c947b3707b0efcf8ceb074a5d5dedc9c76e")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(MembershipPayload)
	Equal(t, ok, true)
}

func TestMilestoneEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/milestone.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, MilestoneEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "milestone")
	req.Header.Set("X-Hub-Signature", "sha1=8b63f58ea58e6a59dcfc5ecbaea0d1741a6bf9ec")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(MilestonePayload)
	Equal(t, ok, true)
}

func TestOrganizationEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/organization.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, OrganizationEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "organization")
	req.Header.Set("X-Hub-Signature", "sha1=7e5ad88557be0a05fb89e86c7893d987386aa0d5")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(OrganizationPayload)
	Equal(t, ok, true)
}

func TestOrgBlockEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/org-block.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, OrgBlockEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "org_block")
	req.Header.Set("X-Hub-Signature", "sha1=21fe61da3f014c011edb60b0b9dfc9aa7059a24b")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(OrgBlockPayload)
	Equal(t, ok, true)
}

func TestPageBuildEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/page-build.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PageBuildEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "page_build")
	req.Header.Set("X-Hub-Signature", "sha1=b3abad8f9c1b3fc0b01c4eb107447800bb5000f9")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PageBuildPayload)
	Equal(t, ok, true)
}

func TestPingEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/ping.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PingEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "ping")
	req.Header.Set("X-Hub-Signature", "sha1=f80e1cfc04245b65228b86612119ab5c894133c2")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PingPayload)
	Equal(t, ok, true)
}

func TestProjectCardEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/project-card.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, ProjectCardEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "project_card")
	req.Header.Set("X-Hub-Signature", "sha1=495dec0d6449d16b71f2ddcd37d595cb9b04b1d8")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(ProjectCardPayload)
	Equal(t, ok, true)
}

func TestProjectColumnEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/project-column.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, ProjectColumnEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "project_column")
	req.Header.Set("X-Hub-Signature", "sha1=7d5dd49d9863e982a4f577170717ea8350a69db0")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(ProjectColumnPayload)
	Equal(t, ok, true)
}

func TestProjectEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/project.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, ProjectEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "project")
	req.Header.Set("X-Hub-Signature", "sha1=7295ab4f205434208f1b86edf2b55adae34c6c92")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(ProjectPayload)
	Equal(t, ok, true)
}

func TestPublicEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/public.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PublicEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "public")
	req.Header.Set("X-Hub-Signature", "sha1=73edb2a8c69c1ac35efb797ede3dc2cde618c10c")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PublicPayload)
	Equal(t, ok, true)
}

func TestPullRequestEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/pull-request.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "pull_request")
	req.Header.Set("X-Hub-Signature", "sha1=35712c8d2bc197b7d07621dcf20d2fb44620508f")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestPayload)
	Equal(t, ok, true)
}

func TestPullRequestReviewEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/pull-request-review.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestReviewEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "pull_request_review")
	req.Header.Set("X-Hub-Signature", "sha1=55345ce92be7849f97d39b9426b95261d4bd4465")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestReviewPayload)
	Equal(t, ok, true)
}

func TestPullRequestReviewCommentEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/pull-request-review-comment.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestReviewCommentEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "pull_request_review_comment")
	req.Header.Set("X-Hub-Signature", "sha1=a9ece15dbcbb85fa5f00a0bf409494af2cbc5b60")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestReviewCommentPayload)
	Equal(t, ok, true)
}

func TestPushEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/push.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "push")
	req.Header.Set("X-Hub-Signature", "sha1=0534736f52c2fc5896ef1bd5a043127b20d233ba")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PushPayload)
	Equal(t, ok, true)
}

func TestReleaseEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/release.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, ReleaseEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "release")
	req.Header.Set("X-Hub-Signature", "sha1=e62bb4c51bc7dde195b9525971c2e3aecb394390")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(ReleasePayload)
	Equal(t, ok, true)
}

func TestRepositoryEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/repository.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepositoryEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "repository")
	req.Header.Set("X-Hub-Signature", "sha1=df442a8af41edd2d42ccdd997938d1d111b0f94e")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepositoryPayload)
	Equal(t, ok, true)
}

func TestStatusEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/status.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, StatusEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "status")
	req.Header.Set("X-Hub-Signature", "sha1=3caa5f062a2deb7cce1482314bb9b4c99bf0ab45")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(StatusPayload)
	Equal(t, ok, true)
}

func TestTeamEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/team.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, TeamEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "team")
	req.Header.Set("X-Hub-Signature", "sha1=ff5b5d58faec10bd40fc96834148df408e7a4608")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(TeamPayload)
	Equal(t, ok, true)
}

func TestTeamAddEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/team-add.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, TeamAddEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "team_add")
	req.Header.Set("X-Hub-Signature", "sha1=5f3953476e270b79cc6763780346110da880609a")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(TeamAddPayload)
	Equal(t, ok, true)
}

func TestWatchEvent(t *testing.T) {

	payload, err := os.Open("../testdata/github/watch.json")
	Equal(t, err, nil)

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, WatchEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "watch")
	req.Header.Set("X-Hub-Signature", "sha1=a317bcfe69ccb8bece74c20c7378e5413c4772f1")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(WatchPayload)
	Equal(t, ok, true)
}
