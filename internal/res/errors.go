package res

import (
	"fmt"
)

type HttpError struct {
	underlying error
}

func (e HttpError) Unwrap() error {
	return e.underlying
}

func (e HttpError) Error() string {
	return fmt.Sprintf("HTTP error: %s", e.underlying.Error())
}

type ParsingError struct {
	underlying error
}

func (e ParsingError) Unwrap() error {
	return e.underlying
}

func (e ParsingError) Error() string {
	return fmt.Sprintf("Parsing error: %s", e.underlying.Error())
}

type UnauthorizedError struct{}

func (e UnauthorizedError) Error() string {
	return "unauthorized, invalid or missing token"
}

type UnexpectedResponseError struct {
	StatusCode int
	Body       []byte
}

func (e *UnexpectedResponseError) Error() string {
	return "unexpected response"
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "not found"
}
