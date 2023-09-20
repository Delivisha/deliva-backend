package es

import (
	"context"
	"deliva/internal/ddd"
	"deliva/internal/registry"
	"fmt"
)

type AggregateRepository[T EventSourcedAggregate] struct {
	AggregateName string
	registry      registry.Registry
	store         AggregateStore
}

func NewAggregateRepository[T EventSourcedAggregate](aggregateName string, registry registry.Registry, store AggregateStore) AggregateRepository[T] {
	return AggregateRepository[T]{
		AggregateName: aggregateName,
		registry:      registry,
		store:         store,
	}
}

func (r AggregateRepository[T]) Find(ctx context.Context, aggregateID string) (agg T, err error) {
	var v any
	v, err = r.registry.Build(
		r.AggregateName,
		ddd.SetID(aggregateID),
		ddd.SetName(r.AggregateName),
	)
	if err != nil {
		return agg, err
	}
	var ok bool
	if agg, ok = v.(T); !ok {
		return agg, fmt.Errorf("%T is not expected type %T", v, agg)
	}
	if err = r.store.Load(ctx, agg); err != nil {
		return agg, err
	}
	return agg, err

}

func (r AggregateRepository[T]) Save(ctx context.Context, aggregate T) error {
	if aggregate.Version() == aggregate.PendingVersion() {
		return nil
	}
	for _, event := range aggregate.Events() {
		if err := aggregate.ApplyEvent(event); err != nil {
			return err
		}
	}
	err := r.store.Save(ctx, aggregate)
	if err != nil {
		return err
	}
	aggregate.CommitEvents()
	return nil
}

//func (r AggregateRepository[T]) Update(ctx context.Context) {
//
//}
