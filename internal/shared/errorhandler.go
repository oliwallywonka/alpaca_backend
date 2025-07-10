package shared

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HttpErrorHandler(err error, c echo.Context) {
	// custom error
	e, ok := err.(*Error)
	if ok {
		_ = c.JSON(getResponseError(e))
		return
	}

	// check echo error
	if echoErr, ok := err.(*echo.HTTPError); ok {
		msg, ok := echoErr.Message.(string)

		if !ok {
			msg = "Upps! server error1"
		}

		_ = c.JSON(echoErr.Code, APIResponse{
			Message: UnexpectedError,
			Error:   []string{msg},
		})
		return
	}

	// if the handler not returns a "model.Error" then it returns a generic error JSON response
	_ = c.JSON(http.StatusInternalServerError, APIResponse{
		Message: UnexpectedError,
		Error:   []string{"Upps! server error2"},
	})
}

func getResponseError(err *Error) (int, APIResponse) {
	outputStatus := 0
	outputResponse := APIResponse{}

	if !err.HasCode() {
		err.Code = UnexpectedError
	}

	if !err.HasStatusHTTP() {
		err.StatusHTTP = http.StatusInternalServerError
	}

	outputStatus = err.StatusHTTP
	outputResponse.Message = err.Code
	outputResponse.Error = err.APIMessage

	return outputStatus, outputResponse
}
