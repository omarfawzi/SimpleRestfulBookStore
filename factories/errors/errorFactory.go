package errors

import (
	"database/sql"
	"errors"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func Make(err error) *Error {
	if errors.Is(err, sql.ErrNoRows) {
		return &Error{
			Code:    http.StatusNotFound,
			Message: "No records found",
		}
	}

	return &Error{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}
