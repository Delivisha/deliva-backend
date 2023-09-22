package domain

import (
	"deliva/internal/ddd"
	"github.com/stackus/errors"
)

const ProjectAggregate = "hr.ProjectAggregate"

type Project struct {
	ddd.Aggregate
	Name     string
	Duration string
	Lead     string
	Budget   int32
	Resource string
}

var (
	ErrProjectNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "project name cannot be blank")
)

func NewProject(id string) *Project {
	return &Project{
		Aggregate: ddd.NewAggregate(id, ProjectAggregate),
	}
}

func CreateProject(id, name, duration, lead, resource string, budget int32) (*Project, error) {
	if name == "" {
		return nil, ErrProjectNameCannotBeBlank
	}
	project := NewProject(id)
	project.Name = name
	project.Duration = duration
	project.Lead = lead
	project.Resource = resource
	project.Budget = budget

	project.AddEvent(ProjectCreatedEvent, &ProjectCreated{
		Project: project,
	})

	return project, nil

}
