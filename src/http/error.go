package http

import (
	"github.com/aacebo/equinox-api/src/logger"
)

var log = logger.New("HttpError")

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, errors ...error) *Error {
	var e = new(Error)

	e.Code = code
	e.Message = Messages[code]

	if code >= 500 {
		log.Error(errors)
	}

	return e
}
