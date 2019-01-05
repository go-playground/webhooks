package googlecalendar

// recieve the Google Calendar Push Notifications
// https://developers.google.com/calendar/v3/push

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	parseheader "github.com/bnfinet/go-httpheader"
)

// parse errors
var (
	ErrInvalidHTTPMethod = errors.New("invalid HTTP Method")
	ErrParsingPayload    = errors.New("error parsing payload")
)

// Event defines a Docker hook event type
type Event string

// GoogleCalendar hook types
const (
	SyncEvent      Event = "sync"
	ExistsEvent    Event = "exists"
	NotExistsEvent Event = "not_exists"
)

// GoogleCalendarPayload a google calendar notice
// https://developers.google.com/calendar/v3/push
type GoogleCalendarPayload struct {
	ChannelID         string    `header:"X-Goog-Channel-ID"`
	ChannelToken      string    `header:"X-Goog-Channel-Token,omitempty"`
	ChannelExpiration time.Time `header:"X-Goog-Channel-Expiration,omitempty"`
	ResourceID        string    `header:"X-Goog-Resource-ID"`
	ResourceURI       string    `header:"X-Goog-Resource-URI"`
	ResourceState     string    `header:"X-Goog-Resource-State"`
	MessageNumber     int       `header:"X-Goog-Message-Number"`
}

// Webhook instance contains all methods needed to process events
type Webhook struct {
	secret string
}

// New creates and returns a WebHook instance
func New() (*Webhook, error) {
	hook := new(Webhook)
	return hook, nil
}

// Parse verifies and parses the events specified and returns the payload object or an error
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	headers := r.Header
	if len(headers) == 0 {
		return nil, ErrParsingPayload
	}

	gc := &GoogleCalendarPayload{}
	err := parseheader.ParseHeader(headers, gc)
	if err != nil {
		fmt.Println(err)
		return nil, ErrParsingPayload
	}
	return gc, err
}
