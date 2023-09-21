package domain

import (
	"deliva/internal/ddd"
	"github.com/stackus/errors"
)

type Department struct {
	ddd.Aggregate
	name              string
	headOfDepartment  string
	numberOfEmployees string
	budget            int32
}

var (
	ErrNameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "department name cannot be blank")
)
