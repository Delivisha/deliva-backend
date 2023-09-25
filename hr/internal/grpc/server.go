package grpc

import (
	"context"
	"deliva/hr/hrpb"
	"deliva/hr/internal/application"
	"deliva/hr/internal/domain"
	"deliva/internal/errorsotel"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	hrpb.UnimplementedHrServiceServer
}

var _ hrpb.HrServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	hrpb.RegisterHrServiceServer(registrar, server{
		app: app,
	})
	return nil
}

func (s server) CreateDepartment(ctx context.Context, req *hrpb.CreateDepartmentRequest) (resp *hrpb.CreateDepartmentResponse, err error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()
	span.SetAttributes(attribute.String("DeptID", id))

	err = s.app.CreateDepartment(ctx, application.CreateDepartment{
		ID:                id,
		Name:              req.Department.GetName(),
		HeadOfDepartment:  req.Department.GetHeadOfDepartment(),
		NumberOfEmployees: int(req.Department.GetNumberOfEmployees()),
		Budget:            req.Department.GetBudget(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}
	return &hrpb.CreateDepartmentResponse{Id: id}, err
}

func (s server) CreateEmployee(ctx context.Context, req *hrpb.CreateEmployeeRequest) (resp *hrpb.CreateEmployeeResponse, err error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()
	span.SetAttributes(attribute.String("EmpId", id))

	err = s.app.CreateEmployee(ctx, application.CreateEmployee{
		ID:             id,
		FirstName:      req.Employee.GetFirstName(),
		MiddleName:     req.Employee.GetMiddleName(),
		LastName:       req.Employee.GetLastName(),
		Email:          req.Employee.GetEmail(),
		Gender:         req.Employee.GetGender(),
		DateOfBirth:    req.Employee.GetDateOfBirth(),
		HomeAddress:    req.Employee.GetHomeAddress(),
		PhoneNumber:    req.Employee.GetPhoneNumber(),
		Bank:           req.Employee.GetBank(),
		AccountNumber:  req.Employee.GetAccountNumber(),
		GrossSalary:    req.Employee.GetGrossSalary(),
		NextOfKinName:  req.Employee.GetNextOfKinName(),
		NextOfKinPhone: req.Employee.GetNextOfKinPhone(),
		ReferencePhone: req.Employee.GetReferencePhone(),
		ReferenceName:  req.Employee.GetReferenceName(),
		DeptID:         req.Employee.GetDeptId(),
		Suspended:      req.Employee.GetSuspended(),
		Sacked:         req.Employee.GetSacked(),
		Country:        req.Employee.GetCountry(),
		Password:       req.Employee.GetPassword(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}
	return &hrpb.CreateEmployeeResponse{Id: id}, err

}

func (s server) SuspendEmployee(ctx context.Context, req *hrpb.SuspendEmployeeRequest) (resp *hrpb.SuspendEmployeeResponse, err error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("EmpID", req.GetId()),
	)
	err = s.app.SuspendEmployee(ctx, application.SuspendEmployee{
		ID:       req.GetId(),
		Reason:   req.GetReason(),
		Duration: req.GetDuration(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}
	return &hrpb.SuspendEmployeeResponse{Success: true}, err
}

func (s server) UnsuspendEmployee(ctx context.Context, req *hrpb.UnsuspendEmployeeRequest) (resp *hrpb.UnsuspendEmployeeResponse, err error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("EmpID", req.GetId()),
	)
	err = s.app.UnsuspendEmployee(ctx, application.UnsuspendEmployee{
		ID: req.GetId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}
	return &hrpb.UnsuspendEmployeeResponse{Success: true}, err
}

func (s server) SackEmployee(ctx context.Context, req *hrpb.SackEmployeeRequest) (resp *hrpb.SackEmployeeResponse, err error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("EmpID", req.GetId()),
	)
	err = s.app.SackEmployee(ctx, application.SackEmployee{
		ID:     req.GetId(),
		Reason: req.GetReason(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}
	return &hrpb.SackEmployeeResponse{Success: true}, err
}

func (s server) GetEmployee(ctx context.Context, req *hrpb.GetEmployeeRequest) (resp *hrpb.GetEmployeeResponse, err error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("EmpID", req.GetId()),
	)
	employee, err := s.app.GetEmployee(ctx, application.GetEmployee{
		ID: req.GetId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}
	return &hrpb.GetEmployeeResponse{
		Employee: s.EmployeeFromDomain(employee),
	}, nil
}

func (s server) EmployeeFromDomain(employee *domain.Employee) *hrpb.Employee {
	return &hrpb.Employee{
		Id:                 employee.ID(),
		FirstName:          employee.FirstName,
		MiddleName:         employee.MiddleName,
		LastName:           employee.LastName,
		Gender:             employee.Gender,
		DateOfBirth:        employee.DateOfBirth,
		Email:              employee.Email,
		HomeAddress:        employee.HomeAddress,
		PhoneNumber:        employee.PhoneNumber,
		Bank:               employee.Bank,
		AccountNumber:      employee.AccountNumber,
		GrossSalary:        employee.GrossSalary,
		NextOfKinName:      employee.NextOfKinName,
		NextOfKinPhone:     employee.NextOfKinPhone,
		ReferencePhone:     employee.ReferencePhone,
		ReferenceName:      employee.ReferenceName,
		DeptId:             employee.Department,
		Country:            employee.Country,
		Sacked:             employee.Sacked,
		Suspended:          employee.Suspended,
		SuspensionReason:   employee.SuspensionReason,
		SuspensionDuration: employee.SuspensionDuration,
		Password:           employee.Password,
	}
}
