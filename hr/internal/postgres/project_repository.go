package postgres

import (
	"context"
	"deliva/hr/internal/domain"
	"deliva/internal/postgres"
	"fmt"
)

type ProjectRepository struct {
	tableName string
	db        postgres.DB
}

func NewProjectRepository(tableName string, db postgres.DB) ProjectRepository {
	return ProjectRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r ProjectRepository) Save(ctx context.Context, project *domain.Project) error {
	const query = "INSERT INTO %s (id, dept_id, name, duration, employee_id, budget, resource) VALUES($1, $2, $3, $4, $5, $6, $7)"

	_, err := r.db.ExecContext(ctx, r.table(query), project.ID(), project.Name, project.Duration, project.Lead, project.Budget, project.Resource)
	return err
}
func (r ProjectRepository) Find(ctx context.Context, projectID string) (*domain.Project, error) {
	const query = "SELECT name, duration, budget FROM %s WHERE id = $1 LIMIT 1"

	project := domain.NewProject(projectID)
	err := r.db.QueryRowContext(ctx, r.table(query), projectID).Scan(&project.Name, &project.Duration, &project.Budget)

	return project, err

}
func (r ProjectRepository) Update() {

}

func (r ProjectRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
