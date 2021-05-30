package errors

import (
	"net/http"
)

type DataNotFoundError string

func (e DataNotFoundError) Error() string {
	return string(e)
}

func (e DataNotFoundError) StatusCode() int {
	return http.StatusNotFound
}

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

func (e ValidationError) StatusCode() int {
	return http.StatusBadRequest
}
