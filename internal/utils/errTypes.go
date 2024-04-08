package utils

import (
	"github.com/pkg/errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrBadRequest      = errors.New("bad request")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNoAuthorization = errors.New("no authorization")
	ErrAccessDenied    = errors.New("access denied")
)
