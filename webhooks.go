package webhooks

import (
	"errors"
	"fmt"
	"net/http"
)

// Provider defines the type of webhook
type Provider int

func (p Provider) String() string {
	switch p {
	case GitHub:
		return "GitHub"
	default:
		return "Unknown"
	}
}

// webhooks available providers
const (
	GitHub Provider = iota
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

	return run(s)
}

// RunTLS runs a server with TLS configuration.
func RunTLS(hook Webhook, addr string, path string, certFile string, keyFile string) error {
	srv := &server{
		hook: hook,
		path: path,
	}

	s := &http.Server{Addr: addr, Handler: srv}

	return run(s, certFile, keyFile)
}

// RunServer runs a custom server.
func RunServer(s *http.Server, hook Webhook, addr string, path string) error {

	srv := &server{
		hook: hook,
		path: path,
	}

	s.Handler = srv

	return run(s)
}

// RunTLSServer runs a custom server with TLS configuration.
// NOTE: http.Server Handler will be overridden by this library, just set it to nil
func RunTLSServer(s *http.Server, hook Webhook, addr string, path string, certFile string, keyFile string) error {

	srv := &server{
		hook: hook,
		path: path,
	}

	s.Handler = srv

	return run(s, certFile, keyFile)
}

func run(s *http.Server, files ...string) error {
	if len(files) == 0 {
		return s.ListenAndServe()
	} else if len(files) == 2 {
		return s.ListenAndServeTLS(files[0], files[1])
	}

	return errors.New("invalid server configuration")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	fmt.Println("GOT HERE!")

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
