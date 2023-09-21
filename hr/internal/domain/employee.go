package domain

import (
	"deliva/internal/ddd"
	"github.com/stackus/errors"
)

const EmployeeAggregate = "hr.EmployeeAggregate"

type Employee struct {
	ddd.Aggregate
	FirstName          string
	MiddleName         string
	LastName           string
	Gender             string
	DateOfBirth        string
	Email              string
	HomeAddress        string
	PhoneNumber        string
	Bank               string
	AccountNumber      string
	GrossSalary        int32
	NextOfKinName      string
	NextOfKinPhone     string
	ReferenceName      string
	ReferencePhone     string
	Department         string
	Country            string
	Suspended          bool
	SuspensionReason   string
	SuspensionDuration string
	Sacked             bool
	SackedReason       string
	password           string
}

var (
	ErrFirstNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the employee  first name cannot be blank")
)
