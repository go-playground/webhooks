package github

import (
	"bytes"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/joeybloggs/webhooks"
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
	port = 3009
	path = "/webhooks"
)

// HandlePayload handles GitHub event(s)
func HandlePayload(payload interface{}) {

}

var hook *Webhook

func TestMain(m *testing.M) {

	// setup
	hook = New(&Config{Secret: "IsWishesWereHorsesWedAllBeEatingSteak!"})
	hook.RegisterEvents(HandlePayload, CommitCommentEvent, CreateEvent, CreateEvent, DeleteEvent, DeploymentEvent, DeploymentStatusEvent, ForkEvent, GollumEvent, IssueCommentEvent, IssuesEvent, MemberEvent, MembershipEvent, PageBuildEvent, PublicEvent, PullRequestReviewCommentEvent, PullRequestEvent, PushEvent, RepositoryEvent, ReleaseEvent, StatusEvent, TeamAddEvent, WatchEvent)

	go webhooks.Run(hook, ":"+strconv.Itoa(port), path)

	os.Exit(m.Run())

	// teardown
}

func TestProvider(t *testing.T) {
	Equal(t, hook.Provider(), webhooks.GitHub)
}

func TestBadNoEventHeader(t *testing.T) {
	payload := "{}"

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusBadRequest)
}

func TestUnsubscribedEvent(t *testing.T) {
	payload := "{}"

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "noneexistant_event")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestBadBody(t *testing.T) {
	payload := ""

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "commit_comment")
	req.Header.Set("X-Hub-Signature", "sha1=156404ad5f721c53151147f3d3d302329f95a3ab")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusInternalServerError)
}

func TestBadSignatureLength(t *testing.T) {
	payload := "{}"

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "commit_comment")
	req.Header.Set("X-Hub-Signature", "")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusForbidden)
}

func TestBadSignatureMatch(t *testing.T) {
	payload := "{}"

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "commit_comment")
	req.Header.Set("X-Hub-Signature", "sha1=111")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusForbidden)
}

func TestCommitCommentEvent(t *testing.T) {

	payload := `{
  "action": "created",
  "comment": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/comments/11056394",
    "html_url": "https://github.com/baxterthehacker/public-repo/commit/9049f1265b7d61be4a8904a9a27120d2064dab3b#commitcomment-11056394",
    "id": 11056394,
    "user": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "position": null,
    "line": null,
    "path": null,
    "commit_id": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
    "created_at": "2015-05-05T23:40:29Z",
    "updated_at": "2015-05-05T23:40:29Z",
    "body": "This is a really good change! :+1:"
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:12Z",
    "pushed_at": "2015-05-05T23:40:27Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "commit_comment")
	req.Header.Set("X-Hub-Signature", "sha1=156404ad5f721c53151147f3d3d302329f95a3ab")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestCreateEvent(t *testing.T) {

	payload := `{
  "ref": "0.0.1",
  "ref_type": "tag",
  "master_branch": "master",
  "description": "",
  "pusher_type": "user",
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:38Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "create")
	req.Header.Set("X-Hub-Signature", "sha1=77ff16ca116034bbeed77ebfce83b36572a9cbaf")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestDeleteEvent(t *testing.T) {

	payload := `{
  "ref": "simple-tag",
  "ref_type": "tag",
  "pusher_type": "user",
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:40Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "delete")
	req.Header.Set("X-Hub-Signature", "sha1=4ddef04fd05b504c7041e294fca3ad1804bc7be1")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestDeploymentEvent(t *testing.T) {

	payload := `{
  "deployment": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/deployments/710692",
    "id": 710692,
    "sha": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
    "ref": "master",
    "task": "deploy",
    "payload": {
    },
    "environment": "production",
    "description": null,
    "creator": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "created_at": "2015-05-05T23:40:38Z",
    "updated_at": "2015-05-05T23:40:38Z",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/deployments/710692/statuses",
    "repository_url": "https://api.github.com/repos/baxterthehacker/public-repo"
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:38Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "deployment")
	req.Header.Set("X-Hub-Signature", "sha1=bb47dc63ceb764a6b1f14fe123e299e5b814c67c")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestDeploymentStatusEvent(t *testing.T) {

	payload := `{
  "deployment": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/deployments/710692",
    "id": 710692,
    "sha": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
    "ref": "master",
    "task": "deploy",
    "payload": {
    },
    "environment": "production",
    "description": null,
    "creator": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "created_at": "2015-05-05T23:40:38Z",
    "updated_at": "2015-05-05T23:40:38Z",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/deployments/710692/statuses",
    "repository_url": "https://api.github.com/repos/baxterthehacker/public-repo"
  },
  "deployment_status": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/deployments/710692/statuses/1115122",
    "id": 1115122,
    "state": "success",
    "creator": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "description": null,
    "target_url": null,
    "created_at": "2015-05-05T23:40:39Z",
    "updated_at": "2015-05-05T23:40:39Z",
    "deployment_url": "https://api.github.com/repos/baxterthehacker/public-repo/deployments/710692",
    "repository_url": "https://api.github.com/repos/baxterthehacker/public-repo"
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:38Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "deployment_status")
	req.Header.Set("X-Hub-Signature", "sha1=8dc0bd0be97440e282e1b4c9ec8445a8d095dc28")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestForkEvent(t *testing.T) {

	payload := `{
  "forkee": {
    "id": 35129393,
    "name": "public-repo",
    "full_name": "baxterandthehackers/public-repo",
    "owner": {
      "login": "baxterandthehackers",
      "id": 7649605,
      "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterandthehackers",
      "html_url": "https://github.com/baxterandthehackers",
      "followers_url": "https://api.github.com/users/baxterandthehackers/followers",
      "following_url": "https://api.github.com/users/baxterandthehackers/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterandthehackers/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterandthehackers/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterandthehackers/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterandthehackers/orgs",
      "repos_url": "https://api.github.com/users/baxterandthehackers/repos",
      "events_url": "https://api.github.com/users/baxterandthehackers/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterandthehackers/received_events",
      "type": "Organization",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterandthehackers/public-repo",
    "description": "",
    "fork": true,
    "url": "https://api.github.com/repos/baxterandthehackers/public-repo",
    "forks_url": "https://api.github.com/repos/baxterandthehackers/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterandthehackers/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterandthehackers/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterandthehackers/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterandthehackers/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterandthehackers/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterandthehackers/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterandthehackers/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterandthehackers/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterandthehackers/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterandthehackers/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterandthehackers/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterandthehackers/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterandthehackers/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterandthehackers/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterandthehackers/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterandthehackers/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterandthehackers/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterandthehackers/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterandthehackers/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterandthehackers/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterandthehackers/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterandthehackers/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterandthehackers/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterandthehackers/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterandthehackers/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterandthehackers/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:30Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:27Z",
    "git_url": "git://github.com/baxterandthehackers/public-repo.git",
    "ssh_url": "git@github.com:baxterandthehackers/public-repo.git",
    "clone_url": "https://github.com/baxterandthehackers/public-repo.git",
    "svn_url": "https://github.com/baxterandthehackers/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": false,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 0,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master",
    "public": true
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:27Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 1,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 1,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterandthehackers",
    "id": 7649605,
    "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterandthehackers",
    "html_url": "https://github.com/baxterandthehackers",
    "followers_url": "https://api.github.com/users/baxterandthehackers/followers",
    "following_url": "https://api.github.com/users/baxterandthehackers/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterandthehackers/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterandthehackers/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterandthehackers/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterandthehackers/orgs",
    "repos_url": "https://api.github.com/users/baxterandthehackers/repos",
    "events_url": "https://api.github.com/users/baxterandthehackers/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterandthehackers/received_events",
    "type": "Organization",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "fork")
	req.Header.Set("X-Hub-Signature", "sha1=cec5f8fb7c383514c622d3eb9e121891dfcca848")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestGollumEvent(t *testing.T) {

	payload := `{
  "pages": [
    {
      "page_name": "Home",
      "title": "Home",
      "summary": null,
      "action": "created",
      "sha": "91ea1bd42aa2ba166b86e8aefe049e9837214e67",
      "html_url": "https://github.com/baxterthehacker/public-repo/wiki/Home"
    }
  ],
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:12Z",
    "pushed_at": "2015-05-05T23:40:17Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 0,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "jasonrudolph",
    "id": 2988,
    "avatar_url": "https://avatars.githubusercontent.com/u/2988?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/jasonrudolph",
    "html_url": "https://github.com/jasonrudolph",
    "followers_url": "https://api.github.com/users/jasonrudolph/followers",
    "following_url": "https://api.github.com/users/jasonrudolph/following{/other_user}",
    "gists_url": "https://api.github.com/users/jasonrudolph/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/jasonrudolph/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/jasonrudolph/subscriptions",
    "organizations_url": "https://api.github.com/users/jasonrudolph/orgs",
    "repos_url": "https://api.github.com/users/jasonrudolph/repos",
    "events_url": "https://api.github.com/users/jasonrudolph/events{/privacy}",
    "received_events_url": "https://api.github.com/users/jasonrudolph/received_events",
    "type": "User",
    "site_admin": true
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "gollum")
	req.Header.Set("X-Hub-Signature", "sha1=a375a6dc8ceac7231ee022211f8eb85e2a84a5b9")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestIssueCommentEvent(t *testing.T) {

	payload := `{
  "action": "created",
  "issue": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/labels{/name}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/comments",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/events",
    "html_url": "https://github.com/baxterthehacker/public-repo/issues/2",
    "id": 73464126,
    "number": 2,
    "title": "Spelling error in the README file",
    "user": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [
      {
        "url": "https://api.github.com/repos/baxterthehacker/public-repo/labels/bug",
        "name": "bug",
        "color": "fc2929"
      }
    ],
    "state": "open",
    "locked": false,
    "assignee": null,
    "milestone": null,
    "comments": 1,
    "created_at": "2015-05-05T23:40:28Z",
    "updated_at": "2015-05-05T23:40:28Z",
    "closed_at": null,
    "body": "It looks like you accidently spelled 'commit' with two 't's."
  },
  "comment": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments/99262140",
    "html_url": "https://github.com/baxterthehacker/public-repo/issues/2#issuecomment-99262140",
    "issue_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2",
    "id": 99262140,
    "user": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "created_at": "2015-05-05T23:40:28Z",
    "updated_at": "2015-05-05T23:40:28Z",
    "body": "You are totally right! I'll get this fixed right away."
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:12Z",
    "pushed_at": "2015-05-05T23:40:27Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "issue_comment")
	req.Header.Set("X-Hub-Signature", "sha1=e724c9f811fcf5f511aac32e4251b08ab1a0fd87")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestIssuesEvent(t *testing.T) {

	payload := `{
  "action": "opened",
  "issue": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/labels{/name}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/comments",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/events",
    "html_url": "https://github.com/baxterthehacker/public-repo/issues/2",
    "id": 73464126,
    "number": 2,
    "title": "Spelling error in the README file",
    "user": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [
      {
        "url": "https://api.github.com/repos/baxterthehacker/public-repo/labels/bug",
        "name": "bug",
        "color": "fc2929"
      }
    ],
    "state": "open",
    "locked": false,
    "assignee": null,
    "milestone": null,
    "comments": 0,
    "created_at": "2015-05-05T23:40:28Z",
    "updated_at": "2015-05-05T23:40:28Z",
    "closed_at": null,
    "body": "It looks like you accidently spelled 'commit' with two 't's."
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:12Z",
    "pushed_at": "2015-05-05T23:40:27Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "issues")
	req.Header.Set("X-Hub-Signature", "sha1=266736f9446195ffefd3d0cfcd1e096ab129ccad")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestMemberEvent(t *testing.T) {

	payload := `{
  "action": "added",
  "member": {
    "login": "octocat",
    "id": 583231,
    "avatar_url": "https://avatars.githubusercontent.com/u/583231?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:40Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "member")
	req.Header.Set("X-Hub-Signature", "sha1=597e7d6627a6636d4c3283e36631983fbd57bdd0")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestMembershipEvent(t *testing.T) {

	payload := `{
  "action": "added",
  "scope": "team",
  "member": {
    "login": "kdaigle",
    "id": 2501,
    "avatar_url": "https://avatars.githubusercontent.com/u/2501?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/kdaigle",
    "html_url": "https://github.com/kdaigle",
    "followers_url": "https://api.github.com/users/kdaigle/followers",
    "following_url": "https://api.github.com/users/kdaigle/following{/other_user}",
    "gists_url": "https://api.github.com/users/kdaigle/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/kdaigle/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/kdaigle/subscriptions",
    "organizations_url": "https://api.github.com/users/kdaigle/orgs",
    "repos_url": "https://api.github.com/users/kdaigle/repos",
    "events_url": "https://api.github.com/users/kdaigle/events{/privacy}",
    "received_events_url": "https://api.github.com/users/kdaigle/received_events",
    "type": "User",
    "site_admin": true
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=2",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  },
  "team": {
    "name": "Contractors",
    "id": 123456,
    "slug": "contractors",
    "permission": "admin",
    "url": "https://api.github.com/teams/123456",
    "members_url": "https://api.github.com/teams/123456/members{/member}",
    "repositories_url": "https://api.github.com/teams/123456/repos"
  },
  "organization": {
    "login": "baxterandthehackers",
    "id": 7649605,
    "url": "https://api.github.com/orgs/baxterandthehackers",
    "repos_url": "https://api.github.com/orgs/baxterandthehackers/repos",
    "events_url": "https://api.github.com/orgs/baxterandthehackers/events",
    "members_url": "https://api.github.com/orgs/baxterandthehackers/members{/member}",
    "public_members_url": "https://api.github.com/orgs/baxterandthehackers/public_members{/member}",
    "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=2"
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "membership")
	req.Header.Set("X-Hub-Signature", "sha1=16928c947b3707b0efcf8ceb074a5d5dedc9c76e")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPageBuildEvent(t *testing.T) {

	payload := `{
  "id": 15995382,
  "build": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/pages/builds/15995382",
    "status": "built",
    "error": {
      "message": null
    },
    "pusher": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "commit": "053b99542c83021d6b202d1a1f5ecd5ef7084e55",
    "duration": 3790,
    "created_at": "2015-05-05T23:40:13Z",
    "updated_at": "2015-05-05T23:40:17Z"
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:12Z",
    "pushed_at": "2015-05-05T23:40:17Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 0,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "page_build")
	req.Header.Set("X-Hub-Signature", "sha1=b3abad8f9c1b3fc0b01c4eb107447800bb5000f9")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPublicEvent(t *testing.T) {

	payload := `{
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:41Z",
    "pushed_at": "2015-05-05T23:40:40Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "public")
	req.Header.Set("X-Hub-Signature", "sha1=73edb2a8c69c1ac35efb797ede3dc2cde618c10c")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestReviewCommentEvent(t *testing.T) {

	payload := `{
  "action": "created",
  "comment": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/comments/29724692",
    "id": 29724692,
    "diff_hunk": "@@ -1 +1 @@\n-# public-repo",
    "path": "README.md",
    "position": 1,
    "original_position": 1,
    "commit_id": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
    "original_commit_id": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
    "user": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "body": "Maybe you should use more emojji on this line.",
    "created_at": "2015-05-05T23:40:27Z",
    "updated_at": "2015-05-05T23:40:27Z",
    "html_url": "https://github.com/baxterthehacker/public-repo/pull/1#discussion_r29724692",
    "pull_request_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1",
    "_links": {
      "self": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/comments/29724692"
      },
      "html": {
        "href": "https://github.com/baxterthehacker/public-repo/pull/1#discussion_r29724692"
      },
      "pull_request": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1"
      }
    }
  },
  "pull_request": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1",
    "id": 34778301,
    "html_url": "https://github.com/baxterthehacker/public-repo/pull/1",
    "diff_url": "https://github.com/baxterthehacker/public-repo/pull/1.diff",
    "patch_url": "https://github.com/baxterthehacker/public-repo/pull/1.patch",
    "issue_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/1",
    "number": 1,
    "state": "open",
    "locked": false,
    "title": "Update the README with new information",
    "user": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "body": "This is a pretty simple change that we need to pull into master.",
    "created_at": "2015-05-05T23:40:27Z",
    "updated_at": "2015-05-05T23:40:27Z",
    "closed_at": null,
    "merged_at": null,
    "merge_commit_sha": "18721552ba489fb84e12958c1b5694b5475f7991",
    "assignee": null,
    "milestone": null,
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1/commits",
    "review_comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1/comments",
    "review_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/comments{/number}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/1/comments",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
    "head": {
      "label": "baxterthehacker:changes",
      "ref": "changes",
      "sha": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
      "user": {
        "login": "baxterthehacker",
        "id": 6752317,
        "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/baxterthehacker",
        "html_url": "https://github.com/baxterthehacker",
        "followers_url": "https://api.github.com/users/baxterthehacker/followers",
        "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
        "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
        "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
        "repos_url": "https://api.github.com/users/baxterthehacker/repos",
        "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
        "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 35129377,
        "name": "public-repo",
        "full_name": "baxterthehacker/public-repo",
        "owner": {
          "login": "baxterthehacker",
          "id": 6752317,
          "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
          "gravatar_id": "",
          "url": "https://api.github.com/users/baxterthehacker",
          "html_url": "https://github.com/baxterthehacker",
          "followers_url": "https://api.github.com/users/baxterthehacker/followers",
          "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
          "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
          "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
          "repos_url": "https://api.github.com/users/baxterthehacker/repos",
          "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
          "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
          "type": "User",
          "site_admin": false
        },
        "private": false,
        "html_url": "https://github.com/baxterthehacker/public-repo",
        "description": "",
        "fork": false,
        "url": "https://api.github.com/repos/baxterthehacker/public-repo",
        "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
        "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
        "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
        "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
        "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
        "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
        "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
        "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
        "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
        "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
        "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
        "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
        "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
        "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
        "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
        "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
        "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
        "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
        "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
        "created_at": "2015-05-05T23:40:12Z",
        "updated_at": "2015-05-05T23:40:12Z",
        "pushed_at": "2015-05-05T23:40:27Z",
        "git_url": "git://github.com/baxterthehacker/public-repo.git",
        "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
        "clone_url": "https://github.com/baxterthehacker/public-repo.git",
        "svn_url": "https://github.com/baxterthehacker/public-repo",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": true,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 1,
        "forks": 0,
        "open_issues": 1,
        "watchers": 0,
        "default_branch": "master"
      }
    },
    "base": {
      "label": "baxterthehacker:master",
      "ref": "master",
      "sha": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
      "user": {
        "login": "baxterthehacker",
        "id": 6752317,
        "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/baxterthehacker",
        "html_url": "https://github.com/baxterthehacker",
        "followers_url": "https://api.github.com/users/baxterthehacker/followers",
        "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
        "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
        "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
        "repos_url": "https://api.github.com/users/baxterthehacker/repos",
        "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
        "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 35129377,
        "name": "public-repo",
        "full_name": "baxterthehacker/public-repo",
        "owner": {
          "login": "baxterthehacker",
          "id": 6752317,
          "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
          "gravatar_id": "",
          "url": "https://api.github.com/users/baxterthehacker",
          "html_url": "https://github.com/baxterthehacker",
          "followers_url": "https://api.github.com/users/baxterthehacker/followers",
          "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
          "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
          "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
          "repos_url": "https://api.github.com/users/baxterthehacker/repos",
          "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
          "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
          "type": "User",
          "site_admin": false
        },
        "private": false,
        "html_url": "https://github.com/baxterthehacker/public-repo",
        "description": "",
        "fork": false,
        "url": "https://api.github.com/repos/baxterthehacker/public-repo",
        "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
        "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
        "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
        "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
        "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
        "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
        "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
        "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
        "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
        "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
        "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
        "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
        "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
        "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
        "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
        "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
        "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
        "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
        "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
        "created_at": "2015-05-05T23:40:12Z",
        "updated_at": "2015-05-05T23:40:12Z",
        "pushed_at": "2015-05-05T23:40:27Z",
        "git_url": "git://github.com/baxterthehacker/public-repo.git",
        "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
        "clone_url": "https://github.com/baxterthehacker/public-repo.git",
        "svn_url": "https://github.com/baxterthehacker/public-repo",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": true,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 1,
        "forks": 0,
        "open_issues": 1,
        "watchers": 0,
        "default_branch": "master"
      }
    },
    "_links": {
      "self": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1"
      },
      "html": {
        "href": "https://github.com/baxterthehacker/public-repo/pull/1"
      },
      "issue": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/issues/1"
      },
      "comments": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/issues/1/comments"
      },
      "review_comments": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1/comments"
      },
      "review_comment": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/comments{/number}"
      },
      "commits": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1/commits"
      },
      "statuses": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c"
      }
    }
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:12Z",
    "pushed_at": "2015-05-05T23:40:27Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 1,
    "forks": 0,
    "open_issues": 1,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "pull_request_review_comment")
	req.Header.Set("X-Hub-Signature", "sha1=a9ece15dbcbb85fa5f00a0bf409494af2cbc5b60")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestEvent(t *testing.T) {

	payload := `{
  "action": "opened",
  "number": 1,
  "pull_request": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1",
    "id": 34778301,
    "html_url": "https://github.com/baxterthehacker/public-repo/pull/1",
    "diff_url": "https://github.com/baxterthehacker/public-repo/pull/1.diff",
    "patch_url": "https://github.com/baxterthehacker/public-repo/pull/1.patch",
    "issue_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/1",
    "number": 1,
    "state": "open",
    "locked": false,
    "title": "Update the README with new information",
    "user": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "body": "This is a pretty simple change that we need to pull into master.",
    "created_at": "2015-05-05T23:40:27Z",
    "updated_at": "2015-05-05T23:40:27Z",
    "closed_at": null,
    "merged_at": null,
    "merge_commit_sha": null,
    "assignee": null,
    "milestone": null,
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1/commits",
    "review_comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1/comments",
    "review_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/comments{/number}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/1/comments",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
    "head": {
      "label": "baxterthehacker:changes",
      "ref": "changes",
      "sha": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
      "user": {
        "login": "baxterthehacker",
        "id": 6752317,
        "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/baxterthehacker",
        "html_url": "https://github.com/baxterthehacker",
        "followers_url": "https://api.github.com/users/baxterthehacker/followers",
        "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
        "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
        "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
        "repos_url": "https://api.github.com/users/baxterthehacker/repos",
        "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
        "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 35129377,
        "name": "public-repo",
        "full_name": "baxterthehacker/public-repo",
        "owner": {
          "login": "baxterthehacker",
          "id": 6752317,
          "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
          "gravatar_id": "",
          "url": "https://api.github.com/users/baxterthehacker",
          "html_url": "https://github.com/baxterthehacker",
          "followers_url": "https://api.github.com/users/baxterthehacker/followers",
          "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
          "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
          "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
          "repos_url": "https://api.github.com/users/baxterthehacker/repos",
          "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
          "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
          "type": "User",
          "site_admin": false
        },
        "private": false,
        "html_url": "https://github.com/baxterthehacker/public-repo",
        "description": "",
        "fork": false,
        "url": "https://api.github.com/repos/baxterthehacker/public-repo",
        "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
        "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
        "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
        "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
        "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
        "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
        "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
        "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
        "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
        "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
        "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
        "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
        "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
        "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
        "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
        "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
        "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
        "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
        "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
        "created_at": "2015-05-05T23:40:12Z",
        "updated_at": "2015-05-05T23:40:12Z",
        "pushed_at": "2015-05-05T23:40:26Z",
        "git_url": "git://github.com/baxterthehacker/public-repo.git",
        "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
        "clone_url": "https://github.com/baxterthehacker/public-repo.git",
        "svn_url": "https://github.com/baxterthehacker/public-repo",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": true,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 1,
        "forks": 0,
        "open_issues": 1,
        "watchers": 0,
        "default_branch": "master"
      }
    },
    "base": {
      "label": "baxterthehacker:master",
      "ref": "master",
      "sha": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
      "user": {
        "login": "baxterthehacker",
        "id": 6752317,
        "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/baxterthehacker",
        "html_url": "https://github.com/baxterthehacker",
        "followers_url": "https://api.github.com/users/baxterthehacker/followers",
        "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
        "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
        "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
        "repos_url": "https://api.github.com/users/baxterthehacker/repos",
        "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
        "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 35129377,
        "name": "public-repo",
        "full_name": "baxterthehacker/public-repo",
        "owner": {
          "login": "baxterthehacker",
          "id": 6752317,
          "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
          "gravatar_id": "",
          "url": "https://api.github.com/users/baxterthehacker",
          "html_url": "https://github.com/baxterthehacker",
          "followers_url": "https://api.github.com/users/baxterthehacker/followers",
          "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
          "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
          "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
          "repos_url": "https://api.github.com/users/baxterthehacker/repos",
          "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
          "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
          "type": "User",
          "site_admin": false
        },
        "private": false,
        "html_url": "https://github.com/baxterthehacker/public-repo",
        "description": "",
        "fork": false,
        "url": "https://api.github.com/repos/baxterthehacker/public-repo",
        "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
        "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
        "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
        "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
        "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
        "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
        "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
        "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
        "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
        "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
        "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
        "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
        "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
        "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
        "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
        "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
        "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
        "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
        "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
        "created_at": "2015-05-05T23:40:12Z",
        "updated_at": "2015-05-05T23:40:12Z",
        "pushed_at": "2015-05-05T23:40:26Z",
        "git_url": "git://github.com/baxterthehacker/public-repo.git",
        "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
        "clone_url": "https://github.com/baxterthehacker/public-repo.git",
        "svn_url": "https://github.com/baxterthehacker/public-repo",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": true,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 1,
        "forks": 0,
        "open_issues": 1,
        "watchers": 0,
        "default_branch": "master"
      }
    },
    "_links": {
      "self": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1"
      },
      "html": {
        "href": "https://github.com/baxterthehacker/public-repo/pull/1"
      },
      "issue": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/issues/1"
      },
      "comments": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/issues/1/comments"
      },
      "review_comments": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1/comments"
      },
      "review_comment": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/comments{/number}"
      },
      "commits": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/pulls/1/commits"
      },
      "statuses": {
        "href": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c"
      }
    },
    "merged": false,
    "mergeable": null,
    "mergeable_state": "unknown",
    "merged_by": null,
    "comments": 0,
    "review_comments": 0,
    "commits": 1,
    "additions": 1,
    "deletions": 1,
    "changed_files": 1
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:12Z",
    "pushed_at": "2015-05-05T23:40:26Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 1,
    "forks": 0,
    "open_issues": 1,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "pull_request")
	req.Header.Set("X-Hub-Signature", "sha1=5b342f365078abd366111158b17a8edf5b41ef2a")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPushEvent(t *testing.T) {

	payload := `{
  "ref": "refs/heads/changes",
  "before": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
  "after": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
  "created": false,
  "deleted": false,
  "forced": false,
  "base_ref": null,
  "compare": "https://github.com/baxterthehacker/public-repo/compare/9049f1265b7d...0d1a26e67d8f",
  "commits": [
    {
      "id": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
      "distinct": true,
      "message": "Update README.md",
      "timestamp": "2015-05-05T19:40:15-04:00",
      "url": "https://github.com/baxterthehacker/public-repo/commit/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
      "author": {
        "name": "baxterthehacker",
        "email": "baxterthehacker@users.noreply.github.com",
        "username": "baxterthehacker"
      },
      "committer": {
        "name": "baxterthehacker",
        "email": "baxterthehacker@users.noreply.github.com",
        "username": "baxterthehacker"
      },
      "added": [

      ],
      "removed": [

      ],
      "modified": [
        "README.md"
      ]
    }
  ],
  "head_commit": {
    "id": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
    "distinct": true,
    "message": "Update README.md",
    "timestamp": "2015-05-05T19:40:15-04:00",
    "url": "https://github.com/baxterthehacker/public-repo/commit/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
    "author": {
      "name": "baxterthehacker",
      "email": "baxterthehacker@users.noreply.github.com",
      "username": "baxterthehacker"
    },
    "committer": {
      "name": "baxterthehacker",
      "email": "baxterthehacker@users.noreply.github.com",
      "username": "baxterthehacker"
    },
    "added": [

    ],
    "removed": [

    ],
    "modified": [
      "README.md"
    ]
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "name": "baxterthehacker",
      "email": "baxterthehacker@users.noreply.github.com"
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://github.com/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": 1430869212,
    "updated_at": "2015-05-05T23:40:12Z",
    "pushed_at": 1430869217,
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 0,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master",
    "stargazers": 0,
    "master_branch": "master"
  },
  "pusher": {
    "name": "baxterthehacker",
    "email": "baxterthehacker@users.noreply.github.com"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "push")
	req.Header.Set("X-Hub-Signature", "sha1=d683a72295b08a42a55bf6fbf2598dc7603e0b98")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestRepositoryEvent(t *testing.T) {

	payload := `{
  "action": "created",
  "repository": {
    "id": 27496774,
    "name": "new-repository",
    "full_name": "baxterandthehackers/new-repository",
    "owner": {
      "login": "baxterandthehackers",
      "id": 7649605,
      "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterandthehackers",
      "html_url": "https://github.com/baxterandthehackers",
      "followers_url": "https://api.github.com/users/baxterandthehackers/followers",
      "following_url": "https://api.github.com/users/baxterandthehackers/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterandthehackers/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterandthehackers/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterandthehackers/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterandthehackers/orgs",
      "repos_url": "https://api.github.com/users/baxterandthehackers/repos",
      "events_url": "https://api.github.com/users/baxterandthehackers/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterandthehackers/received_events",
      "type": "Organization",
      "site_admin": false
    },
    "private": true,
    "html_url": "https://github.com/baxterandthehackers/new-repository",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterandthehackers/new-repository",
    "forks_url": "https://api.github.com/repos/baxterandthehackers/new-repository/forks",
    "keys_url": "https://api.github.com/repos/baxterandthehackers/new-repository/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterandthehackers/new-repository/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterandthehackers/new-repository/teams",
    "hooks_url": "https://api.github.com/repos/baxterandthehackers/new-repository/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterandthehackers/new-repository/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterandthehackers/new-repository/events",
    "assignees_url": "https://api.github.com/repos/baxterandthehackers/new-repository/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterandthehackers/new-repository/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterandthehackers/new-repository/tags",
    "blobs_url": "https://api.github.com/repos/baxterandthehackers/new-repository/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterandthehackers/new-repository/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterandthehackers/new-repository/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterandthehackers/new-repository/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterandthehackers/new-repository/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterandthehackers/new-repository/languages",
    "stargazers_url": "https://api.github.com/repos/baxterandthehackers/new-repository/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterandthehackers/new-repository/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterandthehackers/new-repository/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterandthehackers/new-repository/subscription",
    "commits_url": "https://api.github.com/repos/baxterandthehackers/new-repository/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterandthehackers/new-repository/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterandthehackers/new-repository/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterandthehackers/new-repository/issues/comments/{number}",
    "contents_url": "https://api.github.com/repos/baxterandthehackers/new-repository/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterandthehackers/new-repository/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterandthehackers/new-repository/merges",
    "archive_url": "https://api.github.com/repos/baxterandthehackers/new-repository/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterandthehackers/new-repository/downloads",
    "issues_url": "https://api.github.com/repos/baxterandthehackers/new-repository/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterandthehackers/new-repository/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterandthehackers/new-repository/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterandthehackers/new-repository/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterandthehackers/new-repository/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterandthehackers/new-repository/releases{/id}",
    "created_at": "2014-12-03T16:39:25Z",
    "updated_at": "2014-12-03T16:39:25Z",
    "pushed_at": "2014-12-03T16:39:25Z",
    "git_url": "git://github.com/baxterandthehackers/new-repository.git",
    "ssh_url": "git@github.com:baxterandthehackers/new-repository.git",
    "clone_url": "https://github.com/baxterandthehackers/new-repository.git",
    "svn_url": "https://github.com/baxterandthehackers/new-repository",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 0,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master"
  },
  "organization": {
    "login": "baxterandthehackers",
    "id": 7649605,
    "url": "https://api.github.com/orgs/baxterandthehackers",
    "repos_url": "https://api.github.com/orgs/baxterandthehackers/repos",
    "events_url": "https://api.github.com/orgs/baxterandthehackers/events",
    "members_url": "https://api.github.com/orgs/baxterandthehackers/members{/member}",
    "public_members_url": "https://api.github.com/orgs/baxterandthehackers/public_members{/member}",
    "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=2"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=2",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "repository")
	req.Header.Set("X-Hub-Signature", "sha1=df442a8af41edd2d42ccdd997938d1d111b0f94e")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestReleaseEvent(t *testing.T) {

	payload := `{
  "action": "published",
  "release": {
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/releases/1261438",
    "assets_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases/1261438/assets",
    "upload_url": "https://uploads.github.com/repos/baxterthehacker/public-repo/releases/1261438/assets{?name}",
    "html_url": "https://github.com/baxterthehacker/public-repo/releases/tag/0.0.1",
    "id": 1261438,
    "tag_name": "0.0.1",
    "target_commitish": "master",
    "name": null,
    "draft": false,
    "author": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "prerelease": false,
    "created_at": "2015-05-05T23:40:12Z",
    "published_at": "2015-05-05T23:40:38Z",
    "assets": [

    ],
    "tarball_url": "https://api.github.com/repos/baxterthehacker/public-repo/tarball/0.0.1",
    "zipball_url": "https://api.github.com/repos/baxterthehacker/public-repo/zipball/0.0.1",
    "body": null
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:38Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "release")
	req.Header.Set("X-Hub-Signature", "sha1=e62bb4c51bc7dde195b9525971c2e3aecb394390")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestStatusEvent(t *testing.T) {

	payload := `{
  "id": 214015194,
  "sha": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
  "name": "baxterthehacker/public-repo",
  "target_url": null,
  "context": "default",
  "description": null,
  "state": "success",
  "commit": {
    "sha": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
    "commit": {
      "author": {
        "name": "baxterthehacker",
        "email": "baxterthehacker@users.noreply.github.com",
        "date": "2015-05-05T23:40:12Z"
      },
      "committer": {
        "name": "baxterthehacker",
        "email": "baxterthehacker@users.noreply.github.com",
        "date": "2015-05-05T23:40:12Z"
      },
      "message": "Initial commit",
      "tree": {
        "sha": "02b49ad0ba4f1acd9f06531b21e16a4ac5d341d0",
        "url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees/02b49ad0ba4f1acd9f06531b21e16a4ac5d341d0"
      },
      "url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits/9049f1265b7d61be4a8904a9a27120d2064dab3b",
      "comment_count": 1
    },
    "url": "https://api.github.com/repos/baxterthehacker/public-repo/commits/9049f1265b7d61be4a8904a9a27120d2064dab3b",
    "html_url": "https://github.com/baxterthehacker/public-repo/commit/9049f1265b7d61be4a8904a9a27120d2064dab3b",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits/9049f1265b7d61be4a8904a9a27120d2064dab3b/comments",
    "author": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "committer": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "parents": [

    ]
  },
  "branches": [
    {
      "name": "master",
      "commit": {
        "sha": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
        "url": "https://api.github.com/repos/baxterthehacker/public-repo/commits/9049f1265b7d61be4a8904a9a27120d2064dab3b"
      }
    },
    {
      "name": "changes",
      "commit": {
        "sha": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
        "url": "https://api.github.com/repos/baxterthehacker/public-repo/commits/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c"
      }
    },
    {
      "name": "gh-pages",
      "commit": {
        "sha": "b11bb7545ac14abafc6191a0481b0d961e7793c6",
        "url": "https://api.github.com/repos/baxterthehacker/public-repo/commits/b11bb7545ac14abafc6191a0481b0d961e7793c6"
      }
    }
  ],
  "created_at": "2015-05-05T23:40:39Z",
  "updated_at": "2015-05-05T23:40:39Z",
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:39Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "status")
	req.Header.Set("X-Hub-Signature", "sha1=3caa5f062a2deb7cce1482314bb9b4c99bf0ab45")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestTeamAddEvent(t *testing.T) {

	payload := `{
  "team": {
    "name": "github",
    "id": 836012,
    "slug": "github",
    "description": "",
    "permission": "pull",
    "url": "https://api.github.com/teams/836012",
    "members_url": "https://api.github.com/teams/836012/members{/member}",
    "repositories_url": "https://api.github.com/teams/836012/repos"
  },
  "repository": {
    "id": 35129393,
    "name": "public-repo",
    "full_name": "baxterandthehackers/public-repo",
    "owner": {
      "login": "baxterandthehackers",
      "id": 7649605,
      "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterandthehackers",
      "html_url": "https://github.com/baxterandthehackers",
      "followers_url": "https://api.github.com/users/baxterandthehackers/followers",
      "following_url": "https://api.github.com/users/baxterandthehackers/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterandthehackers/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterandthehackers/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterandthehackers/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterandthehackers/orgs",
      "repos_url": "https://api.github.com/users/baxterandthehackers/repos",
      "events_url": "https://api.github.com/users/baxterandthehackers/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterandthehackers/received_events",
      "type": "Organization",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterandthehackers/public-repo",
    "description": "",
    "fork": true,
    "url": "https://api.github.com/repos/baxterandthehackers/public-repo",
    "forks_url": "https://api.github.com/repos/baxterandthehackers/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterandthehackers/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterandthehackers/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterandthehackers/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterandthehackers/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterandthehackers/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterandthehackers/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterandthehackers/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterandthehackers/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterandthehackers/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterandthehackers/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterandthehackers/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterandthehackers/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterandthehackers/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterandthehackers/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterandthehackers/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterandthehackers/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterandthehackers/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterandthehackers/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterandthehackers/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterandthehackers/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterandthehackers/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterandthehackers/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterandthehackers/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterandthehackers/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterandthehackers/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterandthehackers/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:30Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:27Z",
    "git_url": "git://github.com/baxterandthehackers/public-repo.git",
    "ssh_url": "git@github.com:baxterandthehackers/public-repo.git",
    "clone_url": "https://github.com/baxterandthehackers/public-repo.git",
    "svn_url": "https://github.com/baxterandthehackers/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": false,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 0,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master"
  },
  "organization": {
    "login": "baxterandthehackers",
    "id": 7649605,
    "url": "https://api.github.com/orgs/baxterandthehackers",
    "repos_url": "https://api.github.com/orgs/baxterandthehackers/repos",
    "events_url": "https://api.github.com/orgs/baxterandthehackers/events",
    "members_url": "https://api.github.com/orgs/baxterandthehackers/members{/member}",
    "public_members_url": "https://api.github.com/orgs/baxterandthehackers/public_members{/member}",
    "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=3",
    "description": null
  },
  "sender": {
    "login": "baxterandthehackers",
    "id": 7649605,
    "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterandthehackers",
    "html_url": "https://github.com/baxterandthehackers",
    "followers_url": "https://api.github.com/users/baxterandthehackers/followers",
    "following_url": "https://api.github.com/users/baxterandthehackers/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterandthehackers/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterandthehackers/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterandthehackers/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterandthehackers/orgs",
    "repos_url": "https://api.github.com/users/baxterandthehackers/repos",
    "events_url": "https://api.github.com/users/baxterandthehackers/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterandthehackers/received_events",
    "type": "Organization",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "team_add")
	req.Header.Set("X-Hub-Signature", "sha1=5f3953476e270b79cc6763780346110da880609a")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestWatchEvent(t *testing.T) {

	payload := `{
  "action": "started",
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/baxterthehacker",
      "html_url": "https://github.com/baxterthehacker",
      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/baxterthehacker/public-repo",
    "description": "",
    "fork": false,
    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
    "created_at": "2015-05-05T23:40:12Z",
    "updated_at": "2015-05-05T23:40:30Z",
    "pushed_at": "2015-05-05T23:40:27Z",
    "git_url": "git://github.com/baxterthehacker/public-repo.git",
    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
    "svn_url": "https://github.com/baxterthehacker/public-repo",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 2,
    "forks": 0,
    "open_issues": 2,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/baxterthehacker",
    "html_url": "https://github.com/baxterthehacker",
    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

	req, err := http.NewRequest("POST", "http://localhost:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Github-Event", "watch")
	req.Header.Set("X-Hub-Signature", "sha1=a317bcfe69ccb8bece74c20c7378e5413c4772f1")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}
