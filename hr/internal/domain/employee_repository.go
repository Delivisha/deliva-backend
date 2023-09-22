package domain

import "context"

type EmployeeRepository interface {
	Save(ctx context.Context, employee *Employee) error
	Find(ctx context.Context, employeeID string) (*Employee, error)
	Update(ctx context.Context, employee *Employee) error
}
