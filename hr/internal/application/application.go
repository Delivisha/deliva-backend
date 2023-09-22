package application

import (
	"context"
	"deliva/hr/internal/domain"
	"deliva/internal/ddd"
)

type (
	CreateEmployee struct {
		ID                 string
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
		DeptID             string
		Country            string
		Suspended          bool
		SuspensionReason   string
		SuspensionDuration string
		Sacked             bool
		SackedReason       string
		password           string
	}
	CreateDepartment struct {
		ID                string
		Name              string
		HeadOfDepartment  string
		NumberOfEmployees int
		Budget            int32
	}
	CreateProject struct {
		ID       string
		Name     string
		Duration string
		Budget   int32
		Resource string
	}
	SuspendEmployee struct {
		ID       string
		Reason   string
		Duration string
	}
	UnsuspendEmployee struct {
		ID string
	}
	UpdateNumberOfEmployees struct {
		ID       string
		Employee int
	}
	UpdateBudget struct {
		ID     string
		Budget int32
	}
	SackEmployee struct {
		ID     string
		Reason string
	}
	GetDepartment struct {
		ID string
	}

	App interface {
		CreateDepartment(ctx context.Context, createDept CreateDepartment) error
		GetDepartment(ctx context.Context, getDept GetDepartment) (*domain.Department, error)
		CreateEmployee(ctx context.Context, createEmp CreateEmployee) error
		SuspendEmployee(ctx context.Context, suspend SuspendEmployee) error
		UnsuspendEmployee(ctx context.Context, unsuspend UnsuspendEmployee) error
		CreateProject(ctx context.Context, createPro CreateProject) error
		UpdateNumberOfEmployees(ctx context.Context, update UpdateNumberOfEmployees) error
		UpdateBudget(ctx context.Context, budget UpdateBudget) error
		SackEmployee(ctx context.Context, sack SackEmployee) error
	}
	Application struct {
		departments     domain.DepartmentRepository
		employees       domain.EmployeeRepository
		projects        domain.ProjectRepository
		domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
	}
)

var _ App = (*Application)(nil)

func New(departments domain.DepartmentRepository, employees domain.EmployeeRepository, projects domain.ProjectRepository, domainPublisher ddd.EventPublisher[ddd.AggregateEvent]) *Application {
	return &Application{
		departments:     departments,
		employees:       employees,
		projects:        projects,
		domainPublisher: domainPublisher,
	}
}

func (a Application) CreateDepartment(ctx context.Context, createDept CreateDepartment) error {
	department, err := domain.CreateDepartment(createDept.ID, createDept.Name, createDept.HeadOfDepartment)
	if err != nil {
		return err
	}
	if err = a.departments.Save(ctx, department); err != nil {
		return err
	}
	// Publish domain event
	if err = a.domainPublisher.Publish(ctx, department.Events()...); err != nil {
		return err
	}
	return nil
}
func (a Application) GetDepartment(ctx context.Context, getDept GetDepartment) (*domain.Department, error) {
	return a.departments.Find(ctx, getDept.ID)
}
func (a Application) CreateEmployee(ctx context.Context, createEmp CreateEmployee) error {
	employee, err := domain.CreateEmployee(
		createEmp.ID,
		createEmp.FirstName,
		createEmp.MiddleName,
		createEmp.LastName,
		createEmp.Gender,
		createEmp.DateOfBirth,
		createEmp.Email,
		createEmp.HomeAddress,
		createEmp.PhoneNumber,
		createEmp.Bank,
		createEmp.AccountNumber,
		createEmp.GrossSalary,
		createEmp.NextOfKinName,
		createEmp.NextOfKinPhone,
		createEmp.ReferenceName,
		createEmp.ReferencePhone,
		createEmp.DeptID,
		createEmp.Country,
		createEmp.password,
	)
	if err != nil {
		return err
	}
	if err = a.employees.Save(ctx, employee); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, employee.Events()...); err != nil {
		return err
	}
	return nil
}
func (a Application) SuspendEmployee(ctx context.Context, suspend SuspendEmployee) error {
	employee, err := a.employees.Find(ctx, suspend.ID)
	if err != nil {
		return err
	}
	if err = employee.SuspendEmployee(); err != nil {
		return err
	}
	if err = a.employees.Update(ctx, employee); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, employee.Events()...); err != nil {
		return err
	}
	return nil

}
func (a Application) UnsuspendEmployee(ctx context.Context, unsuspend UnsuspendEmployee) error {
	employee, err := a.employees.Find(ctx, unsuspend.ID)
	if err != nil {
		return err
	}
	if err = employee.UnsuspendEmployee(); err != nil {
		return err
	}
	if err = a.employees.Update(ctx, employee); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, employee.Events()...); err != nil {
		return err
	}
	return nil
}
func (a Application) CreateProject(ctx context.Context, createPro CreateProject) error {
	project, err := domain.CreateProject(createPro.ID, createPro.Name, createPro.Duration, createPro.Resource, createPro.Budget)
	if err != nil {
		return err
	}
	if err = a.projects.Save(ctx, project); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, project.Events()...); err != nil {
		return err
	}
	return nil
}
func (a Application) UpdateNumberOfEmployees(ctx context.Context, update UpdateNumberOfEmployees) error {
	department, err := a.departments.Find(ctx, update.ID)
	if err != nil {
		return err
	}
	if err = department.UpdateNumberOfEmployees(update.Employee); err != nil {
		return err
	}
	if err = a.departments.Update(ctx, department); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, department.Events()...); err != nil {
		return err
	}
	return nil
}
func (a Application) UpdateBudget(ctx context.Context, budget UpdateBudget) error {
	department, err := a.departments.Find(ctx, budget.ID)
	if err != nil {
		return err
	}
	if err = department.UpdateBudget(budget.Budget); err != nil {
		return err
	}
	if err = a.departments.Update(ctx, department); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, department.Events()...); err != nil {
		return err
	}
	return nil
}
func (a Application) SackEmployee(ctx context.Context, sack SackEmployee) error {
	employee, err := a.employees.Find(ctx, sack.ID)
	if err != nil {
		return err
	}
	if err = employee.SackEmployee(); err != nil {
		return err
	}
	if err = a.employees.Update(ctx, employee); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, employee.Events()...); err != nil {
		return err
	}
	return nil
}
