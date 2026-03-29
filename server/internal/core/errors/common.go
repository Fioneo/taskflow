package core_errors

import "errors"

var (
	ErrNotFould        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
)
