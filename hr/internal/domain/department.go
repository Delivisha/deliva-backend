package domain

import (
	"deliva/internal/ddd"
	"github.com/stackus/errors"
)

const (
	DepartmentAggregate = "hr.DepartmentAggregate"
)

type Department struct {
	ddd.Aggregate
	name              string
	headOfDepartment  string
	numberOfEmployees int
	budget            int32
}

var (
	ErrNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "department name cannot be blank")
)

func NewDepartment(id string) *Department {
	return &Department{
		Aggregate: ddd.NewAggregate(id, DepartmentAggregate),
	}
}

func CreateDepartment(id, name, headOfDepartment string) (*Department, error) {
	if name == "" {
		return nil, ErrNameCannotBeBlank
	}
	department := NewDepartment(id)
	department.name = name
	department.headOfDepartment = headOfDepartment

	department.AddEvent(DepartmentCreatedEvent, &DepartmentCreated{
		Department: department,
	})

	return department, nil
}

func (d *Department) UpdateNumberOfEmployees(employee int) error {
	d.numberOfEmployees += employee
	d.AddEvent(NumberOfEmployeesUpdatedEvent, &NumberOfEmployeesUpdated{
		Department: d,
	})
	return nil
}

func (d *Department) UpdateBudget(budget int32) error {
	d.budget += budget
	d.AddEvent(BudgetUpdatedEvent, &BudgetUpdated{
		Department: d,
	})
	return nil
}
