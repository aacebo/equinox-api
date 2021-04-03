package http

import "net/http"

var messages = map[int]string{
	http.StatusInternalServerError: "internal server error",
	http.StatusNotFound:            "not found",
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int) *Error {
	var e = new(Error)

	e.Code = code
	e.Message = messages[code]

	return e
}
