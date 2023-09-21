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

func NewEmployee(id string) *Employee {
	return &Employee{
		Aggregate: ddd.NewAggregate(id, EmployeeAggregate),
	}
}

func CreateEmployee(id, firstName, middleName, lastName, gender, dateOfBirth, email, homeAddress, phoneNumber, bank, accountNumber string, grossSalary int32, nextOfKinName, nextOfKinPhone, referenceName, referencePhone, deptId, country, password string) (*Employee, error) {
	if firstName == "" {
		return nil, ErrFirstNameCannotBeBlank
	}
	employee := NewEmployee(id)
	employee.FirstName = firstName
	employee.MiddleName = middleName
	employee.LastName = lastName
	employee.Email = email
	employee.Gender = gender
	employee.DateOfBirth = dateOfBirth
	employee.HomeAddress = homeAddress
	employee.PhoneNumber = phoneNumber
	employee.Bank = bank
	employee.AccountNumber = accountNumber
	employee.GrossSalary = grossSalary
	employee.NextOfKinName = nextOfKinName
	employee.NextOfKinPhone = nextOfKinPhone
	employee.ReferenceName = referenceName
	employee.ReferencePhone = referencePhone
	employee.Department = deptId
	employee.Suspended = false
	employee.Sacked = false
	employee.Country = country
	employee.password = password

	employee.AddEvent(EmployeeCreatedEvent, &EmployeeCreated{
		Employee: employee,
	})

	return employee, nil
}
func (Employee) Key() string {
	return EmployeeAggregate
}

func (e *Employee) SuspendEmployee() error {
	e.Suspended = true

	e.AddEvent(EmployeeSuspendedEvent, &EmployeeSuspended{
		Employee: e,
	})
	return nil
}

func (e *Employee) UnsuspendEmployee() error {
	e.Suspended = false

	e.AddEvent(EmployeeUnsuspendedEvent, &EmployeeUnsuspended{
		Employee: e,
	})
	return nil
}

//
//func (e *Employee) ChangeBankAccount() error  {
//
//}
