package controller

import (
	"errors"
	"net/http"
)

type CtxKey struct{}

func GetField(r *http.Request, index int) (string, error) {

	fields, ok := r.Context().Value(CtxKey{}).([]string)

	if !ok {
		return "", errors.New("error: parameter is invalid")
	}

	if index > len(fields) {
		return "", errors.New("error: the parameter passed is bigger than fields values")
	}

	return fields[index], nil
}
