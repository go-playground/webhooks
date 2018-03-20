/*
 * Copyright (c) 2018, 奶爸<1@5.nu>
 * All rights reserved.
 */

package gogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/naiba/webhooks"
	client "github.com/gogits/go-gogs-client"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Webhook instance contains all methods needed to process events
type Webhook struct {
	provider   webhooks.Provider
	secret     string
	eventFuncs map[Event]webhooks.ProcessPayloadFunc
}

// Config defines the configuration to create a new Gogs Webhook instance
type Config struct {
	Secret string
}

// Event defines a Gogs hook event type
type Event string

// Gogs hook types
const (
	CreateEvent       Event = "create"
	DeleteEvent       Event = "delete"
	ForkEvent         Event = "fork"
	PushEvent         Event = "push"
	IssuesEvent       Event = "issues"
	IssueCommentEvent Event = "issue_comment"
	PullRequestEvent  Event = "pull_request"
	ReleaseEvent      Event = "release"
)

// New creates and returns a WebHook instance denoted by the Provider type
func New(config *Config) *Webhook {
	return &Webhook{
		provider:   webhooks.Gogs,
		secret:     config.Secret,
		eventFuncs: map[Event]webhooks.ProcessPayloadFunc{},
	}
}

// Provider returns the current hooks provider ID
func (hook Webhook) Provider() webhooks.Provider {
	return hook.provider
}

// RegisterEvents registers the function to call when the specified event(s) are encountered
func (hook Webhook) RegisterEvents(fn webhooks.ProcessPayloadFunc, events ...Event) {

	for _, event := range events {
		hook.eventFuncs[event] = fn
	}
}

// ParsePayload parses and verifies the payload and fires off the mapped function, if it exists.
func (hook Webhook) ParsePayload(w http.ResponseWriter, r *http.Request) {
	webhooks.DefaultLog.Info("Parsing Payload...")

	event := r.Header.Get("X-Gogs-Event")
	if len(event) == 0 {
		webhooks.DefaultLog.Error("Missing X-Gogs-Event Header")
		http.Error(w, "400 Bad Request - Missing X-Gogs-Event Header", http.StatusBadRequest)
		return
	}
	webhooks.DefaultLog.Debug(fmt.Sprintf("X-Gogs-Event:%s", event))

	gogsEvent := Event(event)

	fn, ok := hook.eventFuncs[gogsEvent]
	// if no event registered
	if !ok {
		webhooks.DefaultLog.Info(fmt.Sprintf("Webhook Event %s not registered, it is recommended to setup only events in gogs that will be registered in the webhook to avoid unnecessary traffic and reduce potential attack vectors.", event))
		return
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		webhooks.DefaultLog.Error("Issue reading Payload")
		http.Error(w, "Issue reading Payload", http.StatusInternalServerError)
		return
	}
	webhooks.DefaultLog.Debug(fmt.Sprintf("Payload:%s", string(payload)))

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		webhooks.DefaultLog.Info("Checking secret")
		signature := r.Header.Get("X-Gogs-Signature")
		if len(signature) == 0 {
			webhooks.DefaultLog.Error("Missing X-Gogs-Signature required for HMAC verification")
			http.Error(w, "403 Forbidden - Missing X-Gogs-Signature required for HMAC verification", http.StatusForbidden)
			return
		}
		webhooks.DefaultLog.Debug(fmt.Sprintf("X-Gogs-Signature:%s", signature))

		mac := hmac.New(sha256.New, []byte(hook.secret))
		mac.Write(payload)

		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
			webhooks.DefaultLog.Error("HMAC verification failed")
			webhooks.DefaultLog.Debug("LocalHMAC:" + expectedMAC)
			webhooks.DefaultLog.Debug("RemoteHMAC:" + signature)
			webhooks.DefaultLog.Debug("Secret:" + hook.secret)
			webhooks.DefaultLog.Debug(string(payload))
			http.Error(w, "403 Forbidden - HMAC verification failed", http.StatusForbidden)
			return
		}
	}

	// Make headers available to ProcessPayloadFunc as a webhooks type
	hd := webhooks.Header(r.Header)

	switch gogsEvent {
	case CreateEvent:
		var pe client.CreatePayload
		json.Unmarshal([]byte(payload), &pe)
		hook.runProcessPayloadFunc(fn, pe, hd)

	case ReleaseEvent:
		var re client.ReleasePayload
		json.Unmarshal([]byte(payload), &re)
		hook.runProcessPayloadFunc(fn, re, hd)

	case PushEvent:
		var pe client.PushPayload
		json.Unmarshal([]byte(payload), &pe)
		hook.runProcessPayloadFunc(fn, pe, hd)

	case DeleteEvent:
		var de client.DeletePayload
		json.Unmarshal([]byte(payload), &de)
		hook.runProcessPayloadFunc(fn, de, hd)

	case ForkEvent:
		var fe client.ForkPayload
		json.Unmarshal([]byte(payload), &fe)
		hook.runProcessPayloadFunc(fn, fe, hd)

	case IssuesEvent:
		var ie client.IssuesPayload
		json.Unmarshal([]byte(payload), &ie)
		hook.runProcessPayloadFunc(fn, ie, hd)

	case IssueCommentEvent:
		var ice client.IssueCommentPayload
		json.Unmarshal([]byte(payload), &ice)
		hook.runProcessPayloadFunc(fn, ice, hd)

	case PullRequestEvent:
		var pre client.PullRequestPayload
		json.Unmarshal([]byte(payload), &pre)
		hook.runProcessPayloadFunc(fn, pre, hd)
	}
}

func (hook Webhook) runProcessPayloadFunc(fn webhooks.ProcessPayloadFunc, results interface{}, header webhooks.Header) {
	go func(fn webhooks.ProcessPayloadFunc, results interface{}, header webhooks.Header) {
		fn(results, header)
	}(fn, results, header)
}
