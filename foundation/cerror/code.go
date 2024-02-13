package cerror

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type StatusCoder interface {
	error
	StatusCode() int
}

type statusCoder struct {
	error
	code int
}

func (sc statusCoder) StatusCode() int {
	return sc.code
}

func WithStatusCode(err error, code int) error {
	if err == nil {
		err = errors.New(http.StatusText(code))
	}
	overrideCode := checkShouldOverride(err, code)

	return statusCoder{err, overrideCode}
}

func checkShouldOverride(err error, code int) int {
	var timeouter interface {
		error
		Timeout() bool
	}
	if errors.As(err, &timeouter) && timeouter.Timeout() {
		return http.StatusGatewayTimeout
	}
	var temper interface {
		error
		Temporary() bool
	}
	if errors.As(err, &temper) && temper.Temporary() {
		return http.StatusServiceUnavailable
	}
	if errors.Is(err, context.Canceled) {
		// HTTP 499 in Nginx means that the client closed the connection before the server answered the request.
		// usually caused by client side timeout.
		return 499
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	}
	return code
}

func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if sc := StatusCoder(nil); errors.As(err, &sc) {
		return sc.StatusCode()
	}

	return checkShouldOverride(err, http.StatusInternalServerError)
}
