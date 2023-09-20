package ddd

import (
	"github.com/google/uuid"
	"time"
)

type (
	EventOption interface {
		configureEvent(*event)
	}
	// EventPayload interface has no defined methods so that we can use it easily
	EventPayload interface{}
	Event        interface {
		IDer
		EventName() string
		Payload() EventPayload
		Metadata() Metadata
		OccurredAt() time.Time
	}
	event struct {
		Entity
		payload    EventPayload
		metadata   Metadata
		occurredAt time.Time
	}
)

var _ Event = (*event)(nil)

func NewEvent(name string, payload EventPayload, options ...EventOption) Event {
	return newEvent(name, payload, options...)
}

func newEvent(name string, payload EventPayload, options ...EventOption) event {
	evt := event{
		Entity:     NewEntity(uuid.New().String(), name),
		payload:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}
	for _, option := range options {
		option.configureEvent(&evt)
	}
	return evt
}

func (e event) EventName() string     { return e.EntityName() }
func (e event) Payload() EventPayload { return e.payload }
func (e event) Metadata() Metadata    { return e.metadata }
func (e event) OccurredAt() time.Time { return e.occurredAt }
