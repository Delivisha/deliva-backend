package domain

import "context"

type DepartmentRepository interface {
	Save(ctx context.Context, department *Department) error
	Find(ctx context.Context, DeptID string) (*Department, error)
	Update(ctx context.Context, department *Department) error
}
