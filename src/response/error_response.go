package response

import "net/http"

var messages = map[int]string{
	http.StatusInternalServerError: "internal server error",
	http.StatusNotFound:            "not found",
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int) *ErrorResponse {
	var e = new(ErrorResponse)

	e.Code = code
	e.Message = messages[code]

	return e
}
