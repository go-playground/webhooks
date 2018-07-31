package bitbucket

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
//
const (
	path = "/webhooks"
)

var hook *Webhook

func TestMain(m *testing.M) {

	// setup
	var err error
	hook, err = New(Options.UUID("MY_UUID"))
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

func TestUUIDMissingEvent(t *testing.T) {
	payload := "{}"
	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Event-Key", "noneexistant_event")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrMissingHookUUIDHeader)
}

func TestUUIDDoesNotMatchEvent(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "THIS_DOES_NOT_MATCH")
	req.Header.Set("X-Event-Key", "repo:push")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrUUIDVerificationFailed)
}

func TestBadNoEventHeader(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrMissingEventKeyHeader)
}

func TestBadBody(t *testing.T) {
	payload := ""

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:push")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrParsingPayload)
}

func TestUnsubscribedEvent(t *testing.T) {
	payload := "{}"

	var parseError error
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		_, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "noneexistant_event")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, ErrEventNotFound)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoPushEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:push")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoPushPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoForkEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:fork")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoForkPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoUpdatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoCommitCommentCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoCommitCommentCreatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoCommitStatusCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_status_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoCommitStatusCreatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, RepoCommitStatusUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "repo:commit_status_updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(RepoCommitStatusUpdatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssueCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssueCreatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssueUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssueUpdatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, IssueCommentCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "issue:comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(IssueCommentCreatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestCreatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestUpdatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestApprovedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:approved")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestApprovedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestUnapprovedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:unapproved")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestUnapprovedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestMergedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:fulfilled")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestMergedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestDeclinedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:rejected")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestDeclinedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestCommentCreatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:comment_created")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestCommentCreatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestCommentUpdatedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:comment_updated")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestCommentUpdatedPayload)
	Equal(t, ok, true)
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

	var parseError error
	var results interface{}
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		results, parseError = hook.Parse(r, PullRequestCommentDeletedEvent)
	})
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewBuffer([]byte(payload)))
	Equal(t, err, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hook-UUID", "MY_UUID")
	req.Header.Set("X-Event-Key", "pullrequest:comment_deleted")

	Equal(t, err, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, http.StatusOK)
	Equal(t, parseError, nil)
	_, ok := results.(PullRequestCommentDeletedPayload)
	Equal(t, ok, true)
}
