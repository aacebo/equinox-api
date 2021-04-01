package http

import (
	"github.com/aacebo/equinox-api/src/log"
)

var _log = log.New("HttpError")

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, err error) *Error {
	var e = new(Error)

	e.Code = code
	e.Message = Messages[code]

	if code >= 500 {
		_log.Error(err)
	}

	return e
}
