package bitbucket

import "gopkg.in/go-playground/webhooks.v1"

// Webhook instance contains all methods needed to process events
type Webhook struct {
	provider   webhooks.Provider
	secret     string
	eventFuncs map[Event]webhooks.ProcessPayloadFunc
}

// Config defines the configuration to create a new GitHubWebhook instance
type Config struct {
	Secret string
}

// Event defines a GitHub hook event type
type Event string
