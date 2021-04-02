package http

import "net/http"

var Messages = map[int]string{
	http.StatusInternalServerError: "internal server error",
	http.StatusNotFound:            "not found",
}
