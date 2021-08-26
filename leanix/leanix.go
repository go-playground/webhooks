// Webhook instance contains all methods needed to process events
package leanix

import "errors"

// Event defines a LeanIX hook event type
type Event string

const (
	FactSheetCreatedEvent  Event = "FACT_SHEET_CREATED"
	FactSheetUpdatedEvent  Event = "FACT_SHEET_UPDATED"
	FactSheetArchivedEvent Event = "FACT_SHEET_ARCHIVED"
	FactSheetDeletedEvent  Event = "FACT_SHEET_DELETED"
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
