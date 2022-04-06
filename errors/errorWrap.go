package errors

import "fmt"

type BusinessError struct {
	Code int
	Msg  string
}

func (b *BusinessError) Error() string {
	return fmt.Sprintf("msg: %s", b.Msg)
}

func NotFoundError(msg string) *BusinessError {
	return &BusinessError{Code: 404, Msg: msg}
}

func BadRequestError(msg string) *BusinessError {
	return &BusinessError{Code: 400, Msg: msg}
}

func InternalServerError(msg string) *BusinessError {
	return &BusinessError{Code: 500, Msg: msg}
}

func Unauthorized(msg string) *BusinessError {
	return &BusinessError{Code: 403, Msg: msg}
}
