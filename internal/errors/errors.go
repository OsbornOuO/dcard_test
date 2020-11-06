package errors

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// 自定義的 errors
var (
	ErrIPRateLimiting = &_error{Message: "Error", Status: http.StatusTooManyRequests}

	ErrPageNotFound = &_error{Message: "Page Not Found.", Status: http.StatusNotFound}

	ErrInternalError = &_error{Message: http.StatusText(http.StatusInternalServerError), Status: http.StatusBadRequest}
)

type _error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// HTTPError ...
type HTTPError struct {
	ErrorMessage string `json:"error_message"`
}

func (e *_error) Error() string {
	var b strings.Builder
	_, _ = b.WriteRune('[')
	_, _ = b.WriteString(strconv.Itoa(e.Status))
	_, _ = b.WriteRune(']')
	_, _ = b.WriteRune(' ')
	_, _ = b.WriteString(e.Message)
	return b.String()
}

// Is ...
func (e *_error) Is(target error) bool {
	causeErr := errors.Cause(target)
	tErr, ok := causeErr.(*_error)
	if !ok {
		return false
	}
	return e.Message == tErr.Message
}

// GetHTTPError ,,,
func GetHTTPError(c echo.Context, err *_error) HTTPError {
	msg := err.Message

	return HTTPError{
		ErrorMessage: msg,
	}
}

// NewWithMessage 抽換錯誤訊息
// 未定義的錯誤會被視為 ErrInternalError 類型
func NewWithMessage(err error, message string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	causeErr := errors.Cause(err)
	_err, ok := causeErr.(*_error)
	if !ok {
		return WithStack(&_error{
			Status:  ErrInternalError.Status,
			Message: ErrInternalError.Message,
		})
	}
	err = &_error{
		Status:  _err.Status,
		Message: message,
	}
	var msg string
	for i := 0; i < len(args); i++ {
		msg += "%+v"
	}
	return Wrapf(err, msg, args...)
}

// WithErrors 使用訂好的errors code 與訊息,如果未定義message 顯示對應的http status描述
func WithErrors(err error) error {
	if err == nil {
		return nil
	}
	causeErr := errors.Cause(err)
	_err, ok := causeErr.(*_error)
	if !ok {
		return WithStack(&_error{
			Status:  ErrInternalError.Status,
			Message: http.StatusText(ErrInternalError.Status),
		})
	}
	return WithStack(&_error{
		Status:  _err.Status,
		Message: _err.Message,
	})
}

// NewWithMessagef 抽換錯誤訊息
func NewWithMessagef(err error, format string, args ...interface{}) error {
	return NewWithMessage(err, fmt.Sprintf(format, args...))
}
