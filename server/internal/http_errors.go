package internal

import (
	"net/http"
	"time"
)

type httpErrorDto struct {
	Path   string    `json:"path"`
	Code   int       `json:"code"`
	Reason string    `json:"reason"`
	Time   time.Time `json:"time"`
}

func createHTTPError(path string, code int, reason string) httpErrorDto {
	return httpErrorDto{
		path,
		code, reason, time.Now(),
	}
}

func CreateHTTPNotFoundError(path string, reason string) httpErrorDto {
	return createHTTPError(path, http.StatusNotFound, reason)
}

func CreateHTTPBadRequestError(path string, reason string) httpErrorDto {
	return createHTTPError(path, http.StatusBadRequest, reason)
}

func CreateHTTPServerError(path string, reason string) httpErrorDto {
	return createHTTPError(path, http.StatusInternalServerError, reason)
}
