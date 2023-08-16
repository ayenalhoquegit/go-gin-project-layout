package errorPkg

import (
	"errors"
	"net/http"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/pkg/constant"
)

type HTTPErrorPkg struct {
	Code int
	Err  error
}

func NewError(errorMsg string) error{
	return errors.New(errorMsg)
}

func HandleError (err error) *HTTPErrorPkg{
	httpErr := &HTTPErrorPkg{Code: http.StatusBadRequest, Err: err}
	if err.Error() == "sql: no rows in result set" {
		httpErr.Code = http.StatusNotFound
		httpErr.Err = NewError(constant.NotFound)
	}
	return httpErr
}