package pepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// parse errors
var (
	ErrInvalidHTTPMethod      = errors.New("invalid http method")
	ErrMissingTimestampHeader = errors.New("missing timestamp header")
	ErrMissingSignatureHeader = errors.New("missing signature header")
	ErrMissingVersionHeader   = errors.New("missing version header")
	ErrParsingPayload         = errors.New("error parsing payload")
)

// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// Secret registers the GitHub secret
func (WebhookOptions) Secret(secret string) Option {
	return func(hook *Webhook) error {
		hook.secret = secret
		return nil
	}
}

// Webhook instance contains all methods needed to process events
type Webhook struct {
	secret string
}

// New creates and returns a WebHook instance denoted by the Provider type
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("Error applying Option")
		}
	}
	return hook, nil
}

// Parse verifies and parses the topics specified and returns the payload object or an error
func (hook Webhook) Parse(r *http.Request) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	// If we have a secret set, we should check the signature
	if len(hook.secret) > 0 {
		timestamp := r.Header.Get("pepo-timestamp")
		if len(timestamp) == 0 {
			return nil, ErrMissingTimestampHeader
		}
		signature := r.Header.Get("pepo-signature")
		if len(signature) == 0 {
			return nil, ErrMissingSignatureHeader
		}
		version := r.Header.Get("pepo-version")
		if len(version) == 0 {
			return nil, ErrMissingVersionHeader
		}
		fmt.Println(timestamp)
		fmt.Println(signature)
		fmt.Println(version)
	}

	var pl EventPayload
	err = json.Unmarshal([]byte(payload), &pl)
	return pl, err
}
