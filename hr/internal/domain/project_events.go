package domain

const (
	ProjectCreatedEvent = "hr.ProjectCreated"
)

type ProjectCreated struct {
	Project *Project
}

func (ProjectCreated) Key() string {
	return ProjectCreatedEvent
}
