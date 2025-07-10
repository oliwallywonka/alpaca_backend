package activityerrors

import "errors"

var (
	NotFoundError     = errors.New("activity not found")
	DuplicateKeyError = errors.New("activity already exists")
)
