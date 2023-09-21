package domain

import "github.com/stackus/errors"

type Project struct {
	name     string
	duration string
	budget   int32
	resource string
}

var (
	ErrProjectNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "project name cannot be blank")
)
