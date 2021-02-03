package domain

import "time"

// EventCategory are the available categories for an domain.
type EventCategory string

const (
	// EventInfo will be used to describe some information on an domain.
	EventInfo EventCategory = "info"

	// EventError will be used to describe an error on an domain.
	EventError EventCategory = "error"
)

// Event is an domain that has occurred on a task.
type Event struct {
	Category   EventCategory     `json:"category"`
	Message    string            `json:"message"`
	CreatedAt  time.Time         `json:"created_at"`
	ErrorTrail []string          `json:"error_trail"`
	Data       map[string]string `json:"data"`
}

// Logger handles events
type Logger interface {
	AddEvent(event Event)
	Events() []Event
}
