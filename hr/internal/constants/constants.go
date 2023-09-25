package constants

const ServiceName = "hr"

const (
	HrServiceName = "Hr"
)

const (
	RegistryKey                 = "registry"
	DomainDispatcherKey         = "domainDispatcher"
	DatabaseTransactionKey      = "tx"
	MessagePublisherKey         = "messagePublisher"
	MessageSubscriberKey        = "messageSubscriber"
	EventPublisherKey           = "eventPublisher"
	CommandPublisherKey         = "commandPublisher"
	ReplyPublisherKey           = "replyPublisher"
	SagaStoreKey                = "sagaStore"
	InboxStoreKey               = "inboxStore"
	ApplicationKey              = "app"
	DomainEventHandlersKey      = "domainEventHandlers"
	IntegrationEventHandlersKey = "integrationEventHandlers"
	CommandHandlersKey          = "commandHandlers"
	ReplyHandlersKey            = "replyHandlers"

	DepartmentsRepoKey = "departmentsRepo"
	EmployeesRepoKey   = "employeesRepo"
	ProjectsRepoKey    = "projectsRepo"
)

// Repository Table Names
const (
	OutboxTableName    = ServiceName + ".outbox"
	InboxTableName     = ServiceName + ".inbox"
	EventsTableName    = ServiceName + ".events"
	SnapshotsTableName = ServiceName + ".snapshots"
	SagasTableName     = ServiceName + ".sagas"

	DepartmentsTableName = ServiceName + ".departments"
	EmployeesTableName   = ServiceName + ".employees"
	ProjectsTableName    = ServiceName + ".projects"
)

const (
	DepartmentsCreatedCount = "departments_created_count"
	EmployeesCreatedCount   = "employees_created_count"
	ProjectsCreatedCount    = "projects_created_count"
)
