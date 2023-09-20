package hrpb

import (
	"deliva/internal/registry"
	"deliva/internal/registry/serdes"
)

const (
	EmployeeAggregateChannel   = "deliva.hr.events.Hr"
	DepartmentAggregateChannel = "deliva.hr.events.Hr"
	ProjectAggregateChannel    = "deliva.hr.events.Hr"

	DepartmentCreatedEvent   = "hrapi.DepartmentCreated"
	ProjectCreatedEvent      = "hrapi.ProjectCreated"
	EmployeeCreatedEvent     = "hrapi.EmployeeCreated"
	EmployeeSuspendedEvent   = "hrapi.EmployeeSuspended"
	EmployeeUnsuspendedEvent = "hrapi.EmployeeUnsuspended"
	EmployeeSackedEvent      = "hrapi.EmployeeSacked"

	CommandChannel = "deliva.hr.commands"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Department events
	if err := serde.Register(&DepartmentCreated{}); err != nil {
		return err
	}

	// Project events
	if err := serde.Register(&ProjectCreated{}); err != nil {
		return err
	}
	// Employee events
	if err := serde.Register(&EmployeeCreated{}); err != nil {
		return err
	}
	if err := serde.Register(&EmployeeSuspended{}); err != nil {
		return err
	}
	if err := serde.Register(&EmployeeUnsuspended{}); err != nil {
		return err
	}
	if err := serde.Register(&EmployeeSacked{}); err != nil {
		return err
	}
	return nil
}
func (*DepartmentCreated) Key() string   { return DepartmentCreatedEvent }
func (*ProjectCreated) Key() string      { return ProjectCreatedEvent }
func (*EmployeeCreated) Key() string     { return EmployeeCreatedEvent }
func (*EmployeeSuspended) Key() string   { return EmployeeSuspendedEvent }
func (*EmployeeUnsuspended) Key() string { return EmployeeUnsuspendedEvent }
func (*EmployeeSacked) Key() string      { return EmployeeSackedEvent }
