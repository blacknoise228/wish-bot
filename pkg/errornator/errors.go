package errornator

import (
	"errors"
	"fmt"
	"runtime"
)

type Errornate struct {
	Code uint
	File string
	Line int
	Err  error
}

func (e *Errornate) Error() string {
	return fmt.Sprintf("%s:%d: %v", e.File, e.Line, e.Err)
}
func (e *Errornate) Unwrap() error {
	return e.Err
}

func ErrBadRequest(msg string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	return &Errornate{
		Code: 400,
		File: file,
		Line: line,
		Err:  errors.New("bad request: " + msg),
	}
}

func ErrNotFound(msg string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	return &Errornate{
		Code: 404,
		File: file,
		Line: line,
		Err:  errors.New("not found: " + msg),
	}
}

func ErrForbidden(msg string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	return &Errornate{
		Code: 403,
		File: file,
		Line: line,
		Err:  errors.New("forbidden: " + msg),
	}
}

func ErrInternalServerError(msg string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	return &Errornate{
		Code: 500,
		File: file,
		Line: line,
		Err:  errors.New("internal server error: " + msg),
	}
}

func ErrUnauthorized(msg string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	return &Errornate{
		Code: 401,
		File: file,
		Line: line,
		Err:  errors.New("unauthorized: " + msg),
	}
}

func CustomError(message string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	return &Errornate{
		Code: 999,
		File: file,
		Line: line,
		Err:  errors.New(message),
	}
}

func ErrConflict(msg string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	return &Errornate{
		Code: 409,
		File: file,
		Line: line,
		Err:  errors.New("unique constraint violation: " + msg),
	}
}
