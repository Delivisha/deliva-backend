package hr

import (
	"context"
	"database/sql"
	"deliva/hr/hrpb"
	"deliva/hr/internal/constants"
	"deliva/hr/internal/handlers"
	"deliva/hr/internal/postgres"
	"deliva/hr/internal/rest"
	"deliva/internal/am"
	"deliva/internal/amotel"
	"deliva/internal/amprom"
	"deliva/internal/ddd"
	"deliva/internal/di"
	"deliva/internal/jetstream"
	pg "deliva/internal/postgres"
	"deliva/internal/postgresotel"
	"deliva/internal/registry"
	"deliva/internal/system"
	"deliva/internal/tm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono system.System) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	container := di.New()

	// setup Driven adapters
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		reg := registry.New()
		if err := hrpb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	stream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})
	container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return svc.DB().Begin()
	})
	container.AddScoped(constants.DepartmentsRepoKey, func(c di.Container) (any, error) {
		return postgres.NewDepartmentRepository(
			constants.DepartmentsTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
		), nil
	})
	sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	container.AddScoped(constants.EmployeesRepoKey, func(c di.Container) (any, error) {
		return postgres.NewEmployeeRepository(
			constants.EmployeesTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
		), nil
	})
	container.AddScoped(constants.ProjectsRepoKey, func(c di.Container) (any, error) {
		return postgres.NewProjectRepository(
			constants.ProjectsTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
		), nil
	})
	container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		outboxStore := pg.NewOutboxStore(constants.OutboxTableName, tx)
		return am.NewEventPublisher(
			stream,
			amotel.OtelMessageContextInjector(),
			sentCounter,
			tm.OutboxPublisher(outboxStore),
		), nil
	})
	container.AddScoped(constants.EventPublisherKey, func(c di.Container) (any, error) {
		return am.NewEventPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})
	container.AddScoped(constants.ReplyPublisherKey, func(c di.Container) (any, error) {
		return am.NewReplyPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})
	container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		return pg.NewInboxStore(constants.InboxTableName, tx), nil
	})

	// Prometheus counters
	departmentCreated := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.DepartmentsCreatedCount,
	})
	employeesCreated := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.EmployeesCreatedCount,
	})
	projectCreated := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.ProjectsCreatedCount,
	})

	// setup application

	// setup Driver adapters
	if err = rest.RegisterGateway(ctx, svc.Mux(), svc.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(svc.Mux()); err != nil {
		return err
	}
	handlers.RegisterDomainEvents()

}

func startOutboxProcessor(ctx context.Context, outboxProcessor tm.OutboxProcessor, logger zerolog.Logger) {

}
