package webhooks

import "net/http"

// Provider defines the type of webhook
type Provider int

func (p Provider) String() string {
	switch p {
	case GitHub:
		return "GitHub"
	case Bitbucket:
		return "Bitbucket"
	default:
		return "Unknown"
	}
}

// webhooks available providers
const (
	GitHub Provider = iota
	Bitbucket
)

// Webhook interface defines a webhook to recieve events
type Webhook interface {
	Provider() Provider
	ParsePayload(w http.ResponseWriter, r *http.Request)
}

type server struct {
	hook Webhook
	path string
}

// ProcessPayloadFunc is a common function for payload return values
type ProcessPayloadFunc func(payload interface{})

// Run runs a server
func Run(hook Webhook, addr string, path string) error {
	srv := &server{
		hook: hook,
		path: path,
	}

	s := &http.Server{Addr: addr, Handler: srv}

	return s.ListenAndServe()
}

// RunServer runs a custom server.
func RunServer(s *http.Server, hook Webhook, path string) error {

	srv := &server{
		hook: hook,
		path: path,
	}

	s.Handler = srv

	return s.ListenAndServe()
}

// RunTLSServer runs a custom server with TLS configuration.
// NOTE: http.Server Handler will be overridden by this library, just set it to nil.
// Setting the Certificates can be done in the http.Server.TLSConfig.Certificates
// see example here: https://gopkg.in/go-playground/webhooks.v1/blob/master/webhooks_test.go#L178
func RunTLSServer(s *http.Server, hook Webhook, path string) error {

	srv := &server{
		hook: hook,
		path: path,
	}

	s.Handler = srv

	return s.ListenAndServeTLS("", "")
}

// ServeHTTP is the Handler for every posted WebHook Event
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != s.path {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	s.hook.ParsePayload(w, r)
}
