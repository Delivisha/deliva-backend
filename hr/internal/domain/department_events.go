package domain

const (
	DepartmentCreatedEvent        = "hr.DepartmentCreated"
	NumberOfEmployeesUpdatedEvent = "hr.NumberOfEmployeesUpdated"
	BudgetUpdatedEvent            = "hr.BudgetUpdated"
)

type DepartmentCreated struct {
	Department *Department
}

func (DepartmentCreated) Key() string {
	return DepartmentCreatedEvent
}

type NumberOfEmployeesUpdated struct {
	Department *Department
}

func (NumberOfEmployeesUpdated) Key() string {
	return NumberOfEmployeesUpdatedEvent
}

type BudgetUpdated struct {
	Department *Department
}

func (BudgetUpdated) Key() string {
	return BudgetUpdatedEvent
}
