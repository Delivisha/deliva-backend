package domain

import (
	"deliva/internal/ddd"
	"github.com/stackus/errors"
)

const EmployeeAggregate = "hr.EmployeeAggregate"

type Employee struct {
	ddd.Aggregate
	FirstName     string
	MiddleName    string
	LastName      string
	DateOfBirth   string
	HomeAddress   string
	PhoneNumber   string
	Bank          string
	AccountNumber string
	GrossSalary   int32
	Department    string
	Country       string
	Suspended     bool
	Sacked        bool
}

var (
	ErrFirstNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the employee  first name cannot be blank")
)
