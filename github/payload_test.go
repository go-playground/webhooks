package github

import (
	"encoding/json"
	"io/ioutil"
	pth "path"
	"testing"

	jd "github.com/josephburnett/jd/lib"
	"github.com/stretchr/testify/require"
)

func ParsePayload(payload string, pl interface{}) (string, error) {
	err := json.Unmarshal([]byte(payload), pl)
	if err != nil {
		return "", err
	}
	serialized, err := json.MarshalIndent(pl, "", "  ")
	if err != nil {
		return "", err
	}
	return string(serialized), nil
}

func TestPayloads(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name     string
		filename string
		typ      interface{}
	}{
		{
			name:     "CheckRunPayload",
			filename: "check-run.json",
			typ:      &CheckRunPayload{},
		},
		{
			name:     "CheckSuitePayload",
			filename: "check-suite.json",
			typ:      &CheckSuitePayload{},
		},
		{
			name:     "CommitCommentPayload",
			filename: "commit-comment.json",
			typ:      &CommitCommentPayload{},
		},
		{
			name:     "CreatePayload",
			filename: "create.json",
			typ:      &CreatePayload{},
		},
		{
			name:     "DeletePayload",
			filename: "delete.json",
			typ:      &DeletePayload{},
		},
		{
			name:     "DeploymentStatusPayload",
			filename: "deployment-status.json",
			typ:      &DeploymentStatusPayload{},
		},
		{
			name:     "DeploymentPayload",
			filename: "deployment.json",
			typ:      &DeploymentPayload{},
		},
		{
			name:     "ForkPayload",
			filename: "fork.json",
			typ:      &ForkPayload{},
		},
		{
			name:     "GollumPayload",
			filename: "gollum.json",
			typ:      &GollumPayload{},
		},
		{
			name:     "InstallationPayload",
			filename: "installation.json",
			typ:      &InstallationPayload{},
		},
		{
			name:     "InstallationRepositoriesPayload",
			filename: "installation-repositories.json",
			typ:      &InstallationRepositoriesPayload{},
		},
		{
			name:     "IssueCommentPayload",
			filename: "issue-comment.json",
			typ:      &IssueCommentPayload{},
		},
		{
			name:     "IssuesPayload",
			filename: "issues.json",
			typ:      &IssuesPayload{},
		},
		{
			name:     "LabelPayload",
			filename: "label.json",
			typ:      &LabelPayload{},
		},
		{
			name:     "MemberPayload",
			filename: "member.json",
			typ:      &MemberPayload{},
		},
		{
			name:     "MembershipPayload",
			filename: "membership.json",
			typ:      &MembershipPayload{},
		},
		{
			name:     "MilestonePayload",
			filename: "milestone.json",
			typ:      &MilestonePayload{},
		},
		// {
		// 	name:     "OrgBlockPayload",
		// 	filename: "org-block.json",
		// 	typ:      &OrgBlockPayload{},
		// },
		// {
		// 	name:     "OrganizationPayload",
		// 	filename: "organization.json",
		// 	typ:      &OrganizationPayload{},
		// },
		// {
		// 	name:     "PageBuildPayload",
		// 	filename: "page-build.json",
		// 	typ:      &PageBuildPayload{},
		// },
		// {
		// 	name:     "PingPayload",
		// 	filename: "ping.json",
		// 	typ:      &PingPayload{},
		// },
		// {
		// 	name:     "ProjectCardPayload",
		// 	filename: "project-card.json",
		// 	typ:      &ProjectCardPayload{},
		// },
		// {
		// 	name:     "ProjectColumnPayload",
		// 	filename: "project-column.json",
		// 	typ:      &ProjectColumnPayload{},
		// },
		// {
		// 	name:     "ProjectPayload",
		// 	filename: "project.json",
		// 	typ:      &ProjectPayload{},
		// },
		// {
		// 	name:     "PullRequestReviewCommentPayload",
		// 	filename: "pull-request-review-comment.json",
		// 	typ:      &PullRequestReviewCommentPayload{},
		// },
		// {
		// 	name:     "PullRequestReviewPayload",
		// 	filename: "pull-request-review.json",
		// 	typ:      &PullRequestReviewPayload{},
		// },
		// {
		// 	name:     "PullRequestPayload",
		// 	filename: "pull-request.json",
		// 	typ:      &PullRequestPayload{},
		// },
		// {
		// 	name:     "PushPayload",
		// 	filename: "push.json",
		// 	typ:      &PushPayload{},
		// },
		// {
		// 	name:     "ReleasePayload",
		// 	filename: "release.json",
		// 	typ:      &ReleasePayload{},
		// },
		// {
		// 	name:     "RepositoryVulnerabilityAlertPayload",
		// 	filename: "repository-vulnerability-alert.json",
		// 	typ:      &RepositoryVulnerabilityAlertPayload{},
		// },
		// {
		// 	name:     "RepositoryPayload",
		// 	filename: "repository.json",
		// 	typ:      &RepositoryPayload{},
		// },
		// {
		// 	name:     "SecurityAdvisoryPayload",
		// 	filename: "security-advisory.json",
		// 	typ:      &SecurityAdvisoryPayload{},
		// },
		// {
		// 	name:     "StatusPayload",
		// 	filename: "status.json",
		// 	typ:      &StatusPayload{},
		// },
		// {
		// 	name:     "TeamAddPayload",
		// 	filename: "team-add.json",
		// 	typ:      &TeamAddPayload{},
		// },
		// {
		// 	name:     "TeamPayload",
		// 	filename: "team.json",
		// 	typ:      &TeamPayload{},
		// },
		// {
		// 	name:     "WatchPayload",
		// 	filename: "watch.json",
		// 	typ:      &WatchPayload{},
		// },
	}
	for _, tt := range tests {
		tc := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			payload, err := ioutil.ReadFile(pth.Join("../tmp", tc.filename))
			assert.NoError(err)
			parsedPayload, err := ParsePayload(string(payload), tc.typ)
			assert.NoError(err)
			a, _ := jd.ReadJsonString(string(payload))
			b, _ := jd.ReadJsonString(parsedPayload)
			diff := a.Diff(b).Render()
			assert.Equal("", diff)
		})
	}
}
