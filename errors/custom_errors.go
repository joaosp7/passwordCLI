package customerrors

import "errors"

var (
	ErrEmptyInput = errors.New("input string cannot be empty")
	ErrDirectoryExists = errors.New("Directory already exists")
	ErrDirectoryInvalidString = errors.New("Directory string must be valid")
)