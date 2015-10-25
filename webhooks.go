package webhooks

// Provider defines the type of webhook
type Provider int

// webhooks available providers
const (
	GitHub Provider = iota
)

// Webhook interface defines a webhook to recieve events
type Webhook interface {
	Provider() Provider
}

// Config interface defines the config to setup a webhook instance
type Config interface {
	UnderlyingProvider() Provider
}

// New creates and returns a WebHook instance denoted by the Provider type
func New(config Config) Webhook {

	switch config.UnderlyingProvider() {
	case GitHub:
		c := config.(*GitHubConfig)
		return &GitHubWebhook{
			provider: c.Provider,
		}
	default:
		panic("Invalid config type")
	}
}
