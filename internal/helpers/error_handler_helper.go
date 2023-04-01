package helpers

import (
	res "gotodo/internal/domain/models/response"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	InternalServerError(w, r, err)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := res.DefaultServiceResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "INTERNAL SERVER ERROR",
		IsSuccess:  false,
		Data:       err,
	}

	WriteToResponseBody(w, response)
}
