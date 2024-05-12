package utils

import (
	"errors"
	"io"
	"net/http"
)

func ReadBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, errors.New("no body in request")
	}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	return bodyBytes, nil
}
