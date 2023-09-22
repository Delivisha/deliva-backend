package postgres

import (
	"deliva/internal/postgres"
	"fmt"
)

type EmployeeRepository struct {
	tableName string
	db        postgres.DB
}

func NewEmployeeRepository(tableName string, db postgres.DB) EmployeeRepository {
	return EmployeeRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r EmployeeRepository) Save() {

}

func (r EmployeeRepository) Find() {

}

func (r EmployeeRepository) Update() {

}

func (r EmployeeRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
