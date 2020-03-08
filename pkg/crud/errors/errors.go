package errors

import "fmt"

type apiError struct {
	text string
	err error
}

func ApiError(text string, err error) *apiError {
	return &apiError{text: text, err: err}
}
func Erroring(text string) string {
	return text
}

func (receiver *apiError) Error() string {
	return fmt.Sprintf("error: %v", receiver.err.Error())
}

func (receiver *apiError) Unwrap() error {
	return receiver.err
}
