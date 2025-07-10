package shared

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PaginatedResponse struct {
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Items interface{} `json:"items"`
}

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   []string    `json:"error"`
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

func (a *API) Ok(data interface{}) (int, APIResponse) {
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
	e := NewError()
	e.Err = err
	e.Code = BindFailed
	e.StatusHTTP = http.StatusBadRequest
	e.Who = "c.Bind()"

	log.Warnf("%s", e.Error())
	return &e
}

func (a *API) Error(c echo.Context, who string, err error, msgs ...string) *Error {
	if msgs == nil {
		msgs = []string{"!Upssss! something went wrong"}
	}

	if len(msgs) == 0 {
		msgs = []string{"!Upssss! something went wrong"}
	}

	e := NewError()
	e.Err = err
	e.APIMessage = msgs
	e.Code = UnexpectedError
	e.StatusHTTP = http.StatusInternalServerError
	e.Who = who

	userID, ok := c.Get("userID").(uuid.UUID)

	if !ok {
		log.Errorf("cannot get/parse uuid from userID")
	}
	e.UserID = userID.String()

	log.Errorf("%s", e.Error())
	return &e
}
