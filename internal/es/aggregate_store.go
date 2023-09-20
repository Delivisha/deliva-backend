package es

import (
	"context"
	"coolmarket/internal/ddd"
)

type EventSourcedAggregate interface {
	ddd.IDer
	AggregateName() string
	ddd.Eventer
	Versioner
	EventApplier
	EventCommitter
}

type AggregateStoreMiddleware func(store AggregateStore) AggregateStore

type AggregateStore interface {
	Load(ctx context.Context, aggregate EventSourcedAggregate) error
	Save(ctx context.Context, aggregate EventSourcedAggregate) error
}

func AggregateStoreWithMiddleware(store AggregateStore, mws ...AggregateStoreMiddleware) AggregateStore {
	s := store
	for i := len(mws) - 1; i >= 0; i-- {
		s = mws[i](s)
	}
	return s
}
