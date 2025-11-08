package shared

/* import (
	"net/http"
) */

type PaginatedResponse struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"perPage"`
	TotalPages int         `json:"totalPages"`
	TotalItems int         `json:"totalItems"`
	Items      interface{} `json:"items"`
}

type SingleResponse interface{}

type ErrorData map[string]struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type ErrorResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    ErrorData `json:"data"`
}

const (
	BindFailed      = "bind_failed"
	Ok              = "ok"
	RecordUpdated   = "record_updated"
	RecordCreated   = "record_created"
	RecordDeleted   = "record_deleted"
	UnexpectedError = "unexpected_error"
	ValidationError = "validation_error"
	DuplicateError  = "duplicate_error"
	AuthError       = "auth_error"
)

type API struct{}

func NewAPI() *API {
	return &API{}
}

/* func (a *API) Ok(data interface{}) (int, interface{}) {
	return http.StatusOK, APIResponse{
		Message: Ok,
		Data:    data,
	}
}

func (a *API) Created(data interface{}) (int, APIResponse) {
	return http.StatusCreated, APIResponse{
		Message: RecordCreated,
		Data:    data,
	}
}

func (a *API) Updated(data interface{}) (int, APIResponse) {
	return http.StatusOK, APIResponse{
		Message: RecordUpdated,
		Data:    data,
	}
}

func (a *API) Deleted(data interface{}) (int, APIResponse) {
	return http.StatusOK, APIResponse{
		Message: RecordDeleted,
		Data:    data,
	}
}

func (a *API) BindFailed(err error) error {
	return nil
}

func (a *API) Error(who string, err error, msgs ...string) *Error {
	return nil
}
 */