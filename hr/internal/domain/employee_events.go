package domain

const (
	EmployeeCreatedEvent            = "hr.EmployeeCreated"
	EmployeeSuspendedEvent          = "hr.EmployeeSuspended"
	EmployeeUnsuspendedEvent        = "hr.EmployeeUnsuspended"
	EmployeeBankAccountChangedEvent = "hr.EmployeeBankAccountChanged"
	EmployeeSackedEvent             = "hr.EmployeeSacked"
)

type EmployeeCreated struct {
	Employee *Employee
}

func (EmployeeCreated) Key() string {
	return EmployeeCreatedEvent
}

type EmployeeSuspended struct {
	Employee *Employee
}

func (EmployeeSuspended) Key() string {
	return EmployeeSuspendedEvent
}

type EmployeeUnsuspended struct {
	Employee *Employee
}

func (EmployeeUnsuspended) Key() string {
	return EmployeeUnsuspendedEvent
}

type EmployeeBankAccountChanged struct {
	Employee *Employee
}

func (EmployeeBankAccountChanged) Key() string {
	return EmployeeBankAccountChangedEvent
}

type EmployeeSacked struct {
	Employee *Employee
}

func (EmployeeSacked) Key() string {
	return EmployeeSackedEvent
}
