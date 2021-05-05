package response

import "net/http"

var messages = map[int]string{
	http.StatusInternalServerError: "internal server error",
	http.StatusNotFound:            "not found",
	http.StatusBadRequest:          "bad request",
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int) *ErrorResponse {
	var self = new(ErrorResponse)

	self.Code = code
	self.Message = messages[code]

	return self
}
