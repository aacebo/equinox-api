package errors

import "net/http"

var HttpErrorMessages = map[int]string{
	http.StatusInternalServerError: "internal server error",
}
