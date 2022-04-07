package httpapi

import (
	"errors"
	"github.com/mdapathy/imageuploader/pkg/domain/model"
	"net/http"
)

var (
	ErrUnhealthy          = NewError(http.StatusInternalServerError, "system.unhealthy")
	ErrInvalidRequestBody = NewError(http.StatusBadRequest, "record.invalid_request_body")

)

type Error struct {
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
}

func (e Error) Error() string {
	return e.Messages[0]
}

func (e Error) Status() int {
	return e.Code
}

func NewError(code int, messages ...string) Error {
	return Error{
		Code:     code,
		Messages: messages,
	}
}

func handleErr(err error) error {
	var domainErr model.Error

	if errors.As(err, &domainErr) {
		return NewError(http.StatusBadRequest, domainErr.Error())
	}

	return err
}
