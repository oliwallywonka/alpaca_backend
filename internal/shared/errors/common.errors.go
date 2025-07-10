package commonerrors

import "errors"

var (
	NotFoundError     = errors.New("Record not found")
	DuplicateKeyError = errors.New("Record already exists")
)
