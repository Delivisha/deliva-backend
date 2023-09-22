package postgres

import (
	"context"
	"deliva/hr/internal/domain"
	"deliva/internal/postgres"
	"fmt"
)

type DepartmentRepository struct {
	tableName string
	db        postgres.DB
}

var _ domain.DepartmentRepository = (*DepartmentRepository)(nil)

func NewDepartmentRepository(tableName string, db postgres.DB) DepartmentRepository {
	return DepartmentRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r DepartmentRepository) Save(ctx context.Context, department *domain.Department) error {
	const query = "INSERT INTO %s (id, name, head_of_department, number_of_employees, budget) VALUES($1, $2, $3, $4, $5, $6)"

	_, err := r.db.ExecContext(ctx, r.table(query), department.ID(), department.Name, department.HeadOfDepartment, department.NumberOfEmployees, department.Budget)
	return err
}
func (r DepartmentRepository) Find(ctx context.Context, deptID string) (*domain.Department, error) {
	const query = "SELECT * FROM %s WHERE id = $1 LIMIT 1"

	department := domain.NewDepartment(deptID)
	err := r.db.QueryRowContext(ctx, r.table(query), deptID).Scan(&department.Name, &department.HeadOfDepartment, &department.NumberOfEmployees, &department.Budget)
	return department, err
}
func (r DepartmentRepository) Update(ctx context.Context, department *domain.Department) error {
	const query = "UPDATE %s SET  head_of_department = $3, number_of_employees = $4, budget = $5 WHERE id = $1"
	_, err := r.db.ExecContext(ctx, r.table(query), department.ID(), department.Name, department.HeadOfDepartment, department.NumberOfEmployees, department.Budget)
	return err
}

func (r DepartmentRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
