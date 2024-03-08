package util

import (
	"errors"
	"net/http"
)

var (
	ErrDuplicateEntry                = errors.New("duplicate entry")
	ErrEmptyResult                   = errors.New("empty result")
	ErrCategoryNotFound              = errors.New("category not found")
	ErrDatabase                      = errors.New("database error")
	ErrUnableToDecodeJSON            = errors.New("unable to decode json, input format is wrong")
	ErrInputJSONMustOnlyHaveOneValue = errors.New("input json must only have one value")
)

var CustomErrorType = map[error]int{
	ErrDatabase: http.StatusInternalServerError,
}
