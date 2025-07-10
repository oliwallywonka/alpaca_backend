package hotelerrors

import "errors"

var (
	NotFoundError = errors.New("hotel not found")
	DuplicateKeyError = errors.New("hotel already exists")
)