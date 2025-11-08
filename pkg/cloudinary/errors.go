package cloudinaryfx

import "errors"

var (
	ConnectionError = errors.New("cloudinary connection error")
	UploadError     = errors.New("cloudinary upload error")
	DeleteError     = errors.New("cloudinary delete error")
)
