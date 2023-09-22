package domain

import "context"

type ProjectRepository interface {
	Save(ctx context.Context, project *Project) error
	Find(ctx context.Context, ProjectID string) (*Project, error)
}
