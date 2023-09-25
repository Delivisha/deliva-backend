package handlers

import (
	"context"
	"deliva/hr/hrpb"
	"deliva/hr/internal/domain"
	"deliva/internal/am"
	"deliva/internal/ddd"
	"deliva/internal/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type domainHandlers[T ddd.AggregateEvent] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainHandlers(publisher am.EventPublisher) ddd.EventHandler[ddd.AggregateEvent] {
	return &domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func RegisterDomainEvents(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handlers ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handlers,
		domain.DepartmentCreatedEvent,
		domain.EmployeeCreatedEvent,
		domain.ProjectCreatedEvent,
		domain.EmployeeSuspendedEvent,
		domain.EmployeeUnsuspendedEvent,
		domain.EmployeeSackedEvent,
		domain.EmployeeBankAccountChangedEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent("Encountered an error handling domain event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled domain event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())
	span.AddEvent("Handling domain event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))
	switch event.EventName() {
	case domain.EmployeeCreatedEvent:
		return h.onEmployeeCreated(ctx, event)
	case domain.DepartmentCreatedEvent:
		return h.onDepartmentCreated(ctx, event)
	case domain.EmployeeSackedEvent:
		return h.onEmployeeSacked(ctx, event)
	case domain.EmployeeSuspendedEvent:
		return h.onEmployeeSuspended(ctx, event)
	case domain.EmployeeUnsuspendedEvent:
		return h.onEmployeeUnsuspended(ctx, event)
		//case domain.EmployeeBankAccountChangedEvent:
		//	return h.onEmployeeChangedBankAccount(ctx, event)

	}
	return nil
}

func (h domainHandlers[T]) onDepartmentCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.DepartmentCreated)
	return h.publisher.Publish(ctx, hrpb.DepartmentAggregateChannel,
		ddd.NewEvent(hrpb.DepartmentCreatedEvent, &hrpb.DepartmentCreated{
			Id:                payload.Department.ID(),
			Name:              payload.Department.Name,
			HeadOfDepartment:  payload.Department.HeadOfDepartment,
			NumberOfEmployees: int32(payload.Department.NumberOfEmployees),
			Budget:            payload.Department.Budget,
		}))
}

func (h domainHandlers[T]) onEmployeeCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.EmployeeCreated)
	return h.publisher.Publish(ctx, hrpb.EmployeeAggregateChannel,
		ddd.NewEvent(hrpb.EmployeeCreatedEvent, &hrpb.EmployeeCreated{
			Id:             payload.Employee.ID(),
			FirstName:      payload.Employee.FirstName,
			MiddleName:     payload.Employee.MiddleName,
			LastName:       payload.Employee.LastName,
			Email:          payload.Employee.Email,
			Gender:         payload.Employee.Gender,
			DateOfBirth:    payload.Employee.DateOfBirth,
			HomeAddress:    payload.Employee.HomeAddress,
			PhoneNumber:    payload.Employee.PhoneNumber,
			Bank:           payload.Employee.Bank,
			AccountNumber:  payload.Employee.AccountNumber,
			GrossSalary:    payload.Employee.GrossSalary,
			NextOfKinName:  payload.Employee.NextOfKinName,
			NextOfKinPhone: payload.Employee.NextOfKinPhone,
			ReferenceName:  payload.Employee.ReferenceName,
			ReferencePhone: payload.Employee.ReferencePhone,
			DeptId:         payload.Employee.Department,
			Suspended:      payload.Employee.Suspended,
			Sacked:         payload.Employee.Sacked,
			Country:        payload.Employee.Country,
			Password:       payload.Employee.Password,
		}))
}

func (h domainHandlers[T]) onEmployeeSacked(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.EmployeeSacked)
	return h.publisher.Publish(ctx, hrpb.EmployeeAggregateChannel,
		ddd.NewEvent(hrpb.ProjectCreatedEvent, &hrpb.EmployeeSacked{
			Id:     payload.Employee.ID(),
			Reason: payload.Employee.SackedReason,
		}))
}

func (h domainHandlers[T]) onEmployeeSuspended(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.EmployeeSuspended)
	return h.publisher.Publish(ctx, hrpb.EmployeeAggregateChannel,
		ddd.NewEvent(hrpb.EmployeeSuspendedEvent, &hrpb.EmployeeSuspended{
			Id:     payload.Employee.ID(),
			Reason: payload.Employee.SuspensionReason,
		}))
}

func (h domainHandlers[T]) onEmployeeUnsuspended(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.EmployeeUnsuspended)
	return h.publisher.Publish(ctx, hrpb.EmployeeAggregateChannel,
		ddd.NewEvent(hrpb.EmployeeUnsuspendedEvent, &hrpb.EmployeeUnsuspended{
			Id: payload.Employee.ID(),
		}))
}

//func (h domainHandlers[T]) onEmployeeChangedBankAccount(ctx context.Context, event ddd.AggregateEvent) error {
//	payload := event.Payload().(*domain.EmployeeBankAccountChanged)
//	return h.publisher.Publish(ctx, hrpb.EmployeeAggregateChannel,
//		ddd.NewEvent(hrpb.EmployeeC))
//}
