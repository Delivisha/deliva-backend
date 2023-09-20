package ddd

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type (
	ReplyHandler[T Reply] interface {
		HandleReply(ctx context.Context, reply T) error
	}
	ReplyHandlerFunc[T Reply] func(ctx context.Context, reply T) error
	ReplyOption               interface {
		configureReply(*reply)
	}
	ReplyPayload any

	Reply interface {
		ID() string
		ReplyName() string
		Payload() ReplyPayload
		Metadata() Metadata
		OccurredAt() time.Time
	}

	reply struct {
		Entity
		paylaod    ReplyPayload
		metadata   Metadata
		occurredAt time.Time
	}
)

var _ Reply = (*reply)(nil)

func NewReply(name string, payload ReplyPayload, options ...ReplyOption) Reply {
	return newReply(name, payload, options...)
}

func newReply(name string, payload ReplyPayload, options ...ReplyOption) reply {
	rep := reply{
		Entity:     NewEntity(uuid.New().String(), name),
		paylaod:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}
	for _, option := range options {
		option.configureReply(&rep)
	}
	return rep
}

func (e reply) ReplyName() string     { return e.ReplyName() }
func (e reply) Payload() ReplyPayload { return e.paylaod }
func (e reply) Metadata() Metadata    { return e.metadata }
func (e reply) OccurredAt() time.Time { return e.occurredAt }

func (f ReplyHandlerFunc[T]) HandleReply(ctx context.Context, reply T) error {
	return f(ctx, reply)
}
