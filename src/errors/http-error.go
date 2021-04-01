package errors

import (
	"github.com/aacebo/equinox-api/src/log"
)

var _log = log.New("HttpError")

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, err error) *HttpError {
	var e = new(HttpError)

	e.Code = code
	e.Message = HttpErrorMessages[code]

	if code >= 500 {
		_log.Error(err)
	}

	return e
}
