package customerrors

import "errors"

var (
	ErrEmptyInput = errors.New("input string cannot be empty")
)