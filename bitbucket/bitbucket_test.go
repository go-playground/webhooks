package bitbucket

import (
	"bytes"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	. "gopkg.in/go-playground/assert.v1"
	"gopkg.in/go-playground/webhooks.v3"
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
//
const (
	port = 3009
	path = "/webhooks"
)

// HandlePayload handles GitHub event(s)
func HandlePayload(payload interface{}, header webhooks.Header) {

}

var hook *Webhook

func TestMain(m *testing.M) {

	// setup
	hook = New(&Config{UUID: "MY_UUID"})
	hook.RegisterEvents(
		HandlePayload,
		RepoPushEvent,
		RepoForkEvent,
		RepoUpdatedEvent,
		RepoCommitCommentCreatedEvent,
		RepoCommitStatusCreatedEvent,
		RepoCommitStatusUpdatedEvent,
		IssueCreatedEvent,
		IssueUpdatedEvent,
		IssueCommentCreatedEvent,
		PullRequestCreatedEvent,
		PullRequestUpdatedEvent,
		PullRequestApprovedEvent,
		PullRequestUnapprovedEvent,
		PullRequestMergedEvent,
		PullRequestDeclinedEvent,
		PullRequestCommentCreatedEvent,
		PullRequestCommentUpdatedEvent,
		PullRequestCommentDeletedEvent,
	)

	go webhooks.Run(hook, "127.0.0.1:"+strconv.Itoa(port), path)
	time.Sleep(time.Millisecond * 500)

	os.Exit(m.Run())

	// teardown
}

func TestProvider(t *testing.T) {
	Equal(t, hook.Provider(), webhooks.Bitbucket)
}

func TestUUIDMissingEvent(t *testing.T) {
	payload := "{}"

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Event-Key", "noneexistant_event")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusBadRequest)
}

func TestUUIDDoesNotMatchEvent(t *testing.T) {
	payload := "{}"

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "THIS_DOES_NOT_MATCH")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusForbidden)
}

func TestBadNoEventHeader(t *testing.T) {
	payload := "{}"

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusBadRequest)
}

func TestUnsubscribedEvent(t *testing.T) {
	payload := "{}"

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "noneexistant_event")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestBadBody(t *testing.T) {
	payload := ""

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:push")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusInternalServerError)
}

func TestRepoPush(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "push":{
    "changes":[
      {
        "new":{
          "type":"branch",
          "name":"name-of-branch",
          "target":{
            "type":"commit",
            "hash":"709d658dc5b6d6afcd46049c2f332ee3f515a67d",
            "author":{
              "username":"emmap1",
              "display_name":"Emma",
              "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
              "links":{
                "self":{
                  "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
                },
                "html":{
                  "href":"https://api.bitbucket.org/emmap1"
                },
                "avatar":{
                  "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
                }
              }
            },
            "message":"new commit message\n",
            "date":"2015-06-09T03:34:49+00:00",
            "parents":[
              {
                "type":"commit",
                "hash":"1e65c05c1d5171631d92438a13901ca7dae9618c",
                "links":{
                  "self":{
                    "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/commit/8cbbd65829c7ad834a97841e0defc965718036a0"
                  },
                  "html":{
                    "href":"https://bitbucket.org/user_name/repo_name/commits/8cbbd65829c7ad834a97841e0defc965718036a0"
                  }
                }
              }
            ],
            "links":{
              "self":{
                "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/commit/c4b2b7914156a878aa7c9da452a09fb50c2091f2"
              },
              "html":{
                "href":"https://bitbucket.org/user_name/repo_name/commits/c4b2b7914156a878aa7c9da452a09fb50c2091f2"
              }
            }
          },
          "links":{
            "self":{
              "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/refs/branches/master"
            },
            "commits":{
              "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/commits/master"
            },
            "html":{
              "href":"https://bitbucket.org/user_name/repo_name/branch/master"
            }
          }
        },
        "old":{
          "type":"branch",
          "name":"name-of-branch",
          "target":{
            "type":"commit",
            "hash":"1e65c05c1d5171631d92438a13901ca7dae9618c",
            "author":{
              "username":"emmap1",
              "display_name":"Emma",
              "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
              "links":{
                "self":{
                  "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
                },
                "html":{
                  "href":"https://api.bitbucket.org/emmap1"
                },
                "avatar":{
                  "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
                }
              }
            },
            "message":"old commit message\n",
            "date":"2015-06-08T21:34:56+00:00",
            "parents":[
              {
                "type":"commit",
                "hash":"e0d0c2041e09746be5ce4b55067d5a8e3098c843",
                "links":{
                  "self":{
                    "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/commit/9c4a3452da3bc4f37af5a6bb9c784246f44406f7"
                  },
                  "html":{
                    "href":"https://bitbucket.org/user_name/repo_name/commits/9c4a3452da3bc4f37af5a6bb9c784246f44406f7"
                  }
                }
              }
            ],
            "links":{
              "self":{
                "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/commit/b99ea6dad8f416e57c5ca78c1ccef590600d841b"
              },
              "html":{
                "href":"https://bitbucket.org/user_name/repo_name/commits/b99ea6dad8f416e57c5ca78c1ccef590600d841b"
              }
            }
          },
          "links":{
            "self":{
              "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/refs/branches/master"
            },
            "commits":{
              "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/commits/master"
            },
            "html":{
              "href":"https://bitbucket.org/user_name/repo_name/branch/master"
            }
          }
        },
        "links":{
          "html":{
            "href":"https://bitbucket.org/user_name/repo_name/branches/compare/c4b2b7914156a878aa7c9da452a09fb50c2091f2..b99ea6dad8f416e57c5ca78c1ccef590600d841b"
          },
          "diff":{
            "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/diff/c4b2b7914156a878aa7c9da452a09fb50c2091f2..b99ea6dad8f416e57c5ca78c1ccef590600d841b"
          },
          "commits":{
            "href":"https://api.bitbucket.org/2.0/repositories/user_name/repo_name/commits?include=c4b2b7914156a878aa7c9da452a09fb50c2091f2&exclude=b99ea6dad8f416e57c5ca78c1ccef590600d841b"
          }
        },
        "created":false,
        "forced":false,
        "closed":false,
        "commits":[
          {
            "hash":"03f4a7270240708834de475bcf21532d6134777e",
            "type":"commit",
            "message":"commit message\n",
            "author":{
              "username":"emmap1",
              "display_name":"Emma",
              "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
              "links":{
                "self":{
                  "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
                },
                "html":{
                  "href":"https://api.bitbucket.org/emmap1"
                },
                "avatar":{
                  "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
                }
              }
            },
            "links":{
              "self":{
                "href":"https://api.bitbucket.org/2.0/repositories/user/repo/commit/03f4a7270240708834de475bcf21532d6134777e"
              },
              "html":{
                "href":"https://bitbucket.org/user/repo/commits/03f4a7270240708834de475bcf21532d6134777e"
              }
            }
          }
        ],
        "truncated":false
      }
    ]
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:push")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestRepoFork(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "fork":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:fork")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestRepoUpdated(t *testing.T) {

	payload := `{
	"actor": {
		"type": "user",
		"username": "emmap1",
		"display_name": "Emma",
		"uuid": "{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
		"links": {
			"self": {
				"href": "https://api.bitbucket.org/api/2.0/users/emmap1"
			},
			"html": {
				"href": "https://api.bitbucket.org/emmap1"
			},
			"avatar": {
				"href": "https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
			}
		}
	},
	"repository": {
		"type": "repository",
		"links": {
			"self": {
				"href": "https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
			},
			"html": {
				"href": "https://api.bitbucket.org/bitbucket/bitbucket"
			},
			"avatar": {
				"href": "https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
			}
		},
		"uuid": "{673a6070-3421-46c9-9d48-90745f7bfe8e}",
		"project": {
			"type": "project",
			"project": "Untitled project",
			"uuid": "{3b7898dc-6891-4225-ae60-24613bb83080}",
			"links": {
				"html": {
					"href": "https://bitbucket.org/account/user/teamawesome/projects/proj"
				},
				"avatar": {
					"href": "https://bitbucket.org/account/user/teamawesome/projects/proj/avatar/32"
				}
			},
			"key": "proj"
		},
		"full_name": "team_name/repo_name",
		"name": "repo_name",
		"website": "https://mywebsite.com/",
		"owner": {
			"type": "user",
			"username": "emmap1",
			"display_name": "Emma",
			"uuid": "{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
			"links": {
				"self": {
					"href": "https://api.bitbucket.org/api/2.0/users/emmap1"
				},
				"html": {
					"href": "https://api.bitbucket.org/emmap1"
				},
				"avatar": {
					"href": "https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
				}
			}
		},
		"scm": "git",
		"is_private": true
	},
	"changes": {
		"name": {
			"new": "repository",
			"old": "repository_name"
		},
		"website": {
			"new": "http://www.example.com/",
			"old": ""
		},
		"language": {
			"new": "java",
			"old": ""
		},
		"links": {
			"new": {
				"avatar": {
					"href": "https://bitbucket.org/teamawesome/repository/avatar/32/"
				},
				"self": {
					"href": "https://api.bitbucket.org/2.0/repositories/teamawesome/repository"
				},
				"html": {
					"href": "https://bitbucket.org/teamawesome/repository"
				}
			},
			"old": {
				"avatar": {
					"href": "https://bitbucket.org/teamawesome/repository_name/avatar/32/"
				},
				"self": {
					"href": "https://api.bitbucket.org/2.0/repositories/teamawesome/repository_name"
				},
				"html": {
					"href": "https://bitbucket.org/teamawesome/repository_name"
				}
			}
		},
		"description": {
			"new": "This is a better description.",
			"old": "This is a description."
		},
		"full_name": {
			"new": "teamawesome/repository",
			"old": "teamawesome/repository_name"
		}
	}
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestRepoCommitCommentCreated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "comment":{
    "id":17,
    "parent":{
      "id":16
    },
    "content":{
      "raw":"Comment text",
      "html":"<p>Comment text</p>",
      "markup":"markdown"
    },
    "inline":{
      "path":"path/to/file",
      "from":null,
      "to":10
    },
    "created_on":"2015-04-06T16:52:29.982346+00:00",
    "updated_on":"2015-04-06T16:52:29.983730+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/comments/comment_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/comment_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "commit":{
    "hash":"d3022fc0ca3d65c7f6654eea129d6bf0cf0ee08e"
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestRepoCommitStatusCreated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "commit_status":{
    "name":"Unit Tests (Python)",
    "description":"Build started",
    "state":"INPROGRESS",
    "key":"mybuildtool",
    "url":"https://my-build-tool.com/builds/MY-PROJECT/BUILD-777",
    "type":"build",
    "created_on":"2015-11-19T20:37:35.547563+00:00",
    "updated_on":"2015-11-19T20:37:35.547563+00:00",
    "links":{
      "commit":{
        "href":"http://api.bitbucket.org/2.0/repositories/tk/test/commit/9fec847784abb10b2fa567ee63b85bd238955d0e"
      },
      "self":{
        "href":"http://api.bitbucket.org/2.0/repositories/tk/test/commit/9fec847784abb10b2fa567ee63b85bd238955d0e/statuses/build/mybuildtool"
      }
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_status_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestRepoCommitStatusUpdated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "commit_status":{
    "name":"Unit Tests (Python)",
    "description":"All tests passed",
    "state":"SUCCESSFUL",
    "key":"mybuildtool",
    "url":"https://my-build-tool.com/builds/MY-PROJECT/BUILD-792",
    "type":"build",
    "created_on":"2015-11-19T20:37:35.547563+00:00",
    "updated_on":"2015-11-20T08:01:16.433108+00:00",
    "links":{
      "commit":{
        "href":"http://api.bitbucket.org/2.0/repositories/tk/test/commit/9fec847784abb10b2fa567ee63b85bd238955d0e"
      },
      "self":{
        "href":"http://api.bitbucket.org/2.0/repositories/tk/test/commit/9fec847784abb10b2fa567ee63b85bd238955d0e/statuses/build/mybuildtool"
      }
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_status_updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestIssueCreated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "issue":{
    "id":1,
    "component":"component",
    "title":"Issue title",
    "content":{
      "raw":"Issue description",
      "html":"<p>Issue description</p>",
      "markup":"markdown"
    },
    "priority":"trivial|minor|major|critical|blocker",
    "state":"new|open|on hold|resolved|duplicate|invalid|wontfix|closed",
    "type":"bug|enhancement|proposal|task",
    "milestone":{
      "name":"milestone 1"
    },
    "version":{
      "name":"version 1"
    },
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.179678+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/issues/issue_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/issue_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestIssueUpdated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "issue":{
    "id":1,
    "component":"component",
    "title":"Issue title",
    "content":{
      "raw":"Issue description",
      "html":"<p>Issue description</p>",
      "markup":"markdown"
    },
    "priority":"trivial|minor|major|critical|blocker",
    "state":"new|open|on hold|resolved|duplicate|invalid|wontfix|closed",
    "type":"bug|enhancement|proposal|task",
    "milestone":{
      "name":"milestone 1"
    },
    "version":{
      "name":"version 1"
    },
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.179678+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/issues/issue_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/issue_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "comment":{
    "id":17,
    "parent":{
      "id":16
    },
    "content":{
      "raw":"Comment text",
      "html":"<p>Comment text</p>",
      "markup":"markdown"
    },
    "inline":{
      "path":"path/to/file",
      "from":null,
      "to":10
    },
    "created_on":"2015-04-06T16:52:29.982346+00:00",
    "updated_on":"2015-04-06T16:52:29.983730+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/comments/comment_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/comment_id"
      }
    }
  },
  "changes":{
    "status":{
      "old":"open",
      "new":"on hold"
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestIssueCommentCreated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "issue":{
    "id":1,
    "component":"component",
    "title":"Issue title",
    "content":{
      "raw":"Issue description",
      "html":"<p>Issue description</p>",
      "markup":"markdown"
    },
    "priority":"trivial|minor|major|critical|blocker",
    "state":"new|open|on hold|resolved|duplicate|invalid|wontfix|closed",
    "type":"bug|enhancement|proposal|task",
    "milestone":{
      "name":"milestone 1"
    },
    "version":{
      "name":"version 1"
    },
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.179678+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/issues/issue_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/issue_id"
      }
    }
  },
  "comment":{
    "id":17,
    "parent":{
      "id":16
    },
    "content":{
      "raw":"Comment text",
      "html":"<p>Comment text</p>",
      "markup":"markdown"
    },
    "inline":{
      "path":"path/to/file",
      "from":null,
      "to":10
    },
    "created_on":"2015-04-06T16:52:29.982346+00:00",
    "updated_on":"2015-04-06T16:52:29.983730+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/comments/comment_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/comment_id"
      }
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestCreated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestUpdated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestApproved(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "approval":{
    "date":"2015-04-06T16:34:59.195330+00:00",
    "user":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:approved")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestApprovalRemoved(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "approval":{
    "date":"2015-04-06T16:34:59.195330+00:00",
    "user":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:unapproved")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestMerged(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:fulfilled")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestDeclined(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:rejected")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestCommentCreated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "comment":{
    "id":17,
    "parent":{
      "id":16
    },
    "content":{
      "raw":"Comment text",
      "html":"<p>Comment text</p>",
      "markup":"markdown"
    },
    "inline":{
      "path":"path/to/file",
      "from":null,
      "to":10
    },
    "created_on":"2015-04-06T16:52:29.982346+00:00",
    "updated_on":"2015-04-06T16:52:29.983730+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/comments/comment_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/comment_id"
      }
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestCommentUpdated(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "comment":{
    "id":17,
    "parent":{
      "id":16
    },
    "content":{
      "raw":"Comment text",
      "html":"<p>Comment text</p>",
      "markup":"markdown"
    },
    "inline":{
      "path":"path/to/file",
      "from":null,
      "to":10
    },
    "created_on":"2015-04-06T16:52:29.982346+00:00",
    "updated_on":"2015-04-06T16:52:29.983730+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/comments/comment_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/comment_id"
      }
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:comment_updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}

func TestPullRequestCommentDeleted(t *testing.T) {

	payload := `{
  "actor":{
    "username":"emmap1",
    "display_name":"Emma",
    "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
      },
      "html":{
        "href":"https://api.bitbucket.org/emmap1"
      },
      "avatar":{
        "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
      }
    }
  },
  "repository":{
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
      },
      "html":{
        "href":"https://api.bitbucket.org/bitbucket/bitbucket"
      },
      "avatar":{
        "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
      }
    },
    "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
    "full_name":"team_name/repo_name",
    "name":"repo_name",
    "scm":"git",
    "is_private":true
  },
  "pullrequest":{
    "id":1,
    "title":"Title of pull request",
    "description":"Description of pull request",
    "state":"OPEN|MERGED|DECLINED",
    "author":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "source":{
      "branch":{
        "name":"branch2"
      },
      "commit":{
        "hash":"d3022fc0ca3d"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "destination":{
      "branch":{
        "name":"master"
      },
      "commit":{
        "hash":"ce5965ddd289"
      },
      "repository":{
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/repositories/bitbucket/bitbucket"
          },
          "html":{
            "href":"https://api.bitbucket.org/bitbucket/bitbucket"
          },
          "avatar":{
            "href":"https://api-staging-assetroot.s3.amazonaws.com/c/photos/2014/Aug/01/bitbucket-logo-2629490769-3_avatar.png"
          }
        },
        "uuid":"{673a6070-3421-46c9-9d48-90745f7bfe8e}",
        "full_name":"team_name/repo_name",
        "name":"repo_name",
        "scm":"git",
        "is_private":true
      }
    },
    "merge_commit":{
      "hash":"764413d85e29"
    },
    "participants":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "reviewers":[
      {
        "username":"emmap1",
        "display_name":"Emma",
        "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
        "links":{
          "self":{
            "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
          },
          "html":{
            "href":"https://api.bitbucket.org/emmap1"
          },
          "avatar":{
            "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
          }
        }
      }
    ],
    "close_source_branch":true,
    "closed_by":{
      "username":"emmap1",
      "display_name":"Emma",
      "uuid":"{a54f16da-24e9-4d7f-a3a7-b1ba2cd98aa3}",
      "links":{
        "self":{
          "href":"https://api.bitbucket.org/api/2.0/users/emmap1"
        },
        "html":{
          "href":"https://api.bitbucket.org/emmap1"
        },
        "avatar":{
          "href":"https://bitbucket-api-assetroot.s3.amazonaws.com/c/photos/2015/Feb/26/3613917261-0-emmap1-avatar_avatar.png"
        }
      }
    },
    "reason":"reason for declining the PR (if applicable)",
    "created_on":"2015-04-06T15:23:38.179678+00:00",
    "updated_on":"2015-04-06T15:23:38.205705+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/pullrequests/pullrequest_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/pullrequest_id"
      }
    }
  },
  "comment":{
    "id":17,
    "parent":{
      "id":16
    },
    "content":{
      "raw":"Comment text",
      "html":"<p>Comment text</p>",
      "markup":"markdown"
    },
    "inline":{
      "path":"path/to/file",
      "from":null,
      "to":10
    },
    "created_on":"2015-04-06T16:52:29.982346+00:00",
    "updated_on":"2015-04-06T16:52:29.983730+00:00",
    "links":{
      "self":{
        "href":"https://api.bitbucket.org/api/2.0/comments/comment_id"
      },
      "html":{
        "href":"https://api.bitbucket.org/comment_id"
      }
    }
  }
}
`

	req, err := http.NewRequest("POST", "http://127.0.0.1:3009/webhooks", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pull_request:comment_deleted")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)

	defer resp.Body.Close()

	Equal(t, resp.StatusCode, http.StatusOK)
}
