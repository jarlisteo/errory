package errory

import (
	"strings"
)

type Err struct {
	Code    int      `json:"code"`
	Message string   `json:"error"`
	Errors  []string `json:"errors"`
}

const (
	NotFound   = 404
	BadRequest = 400
	Internal   = 500
)

func (e Err) Error() string {
	return e.Message
}

func (e *Err) GenerateErrors() {
	e.Errors = strings.Split(e.Message, ",")
}

func GenerateMessage(code int) string {
	switch code {
	case NotFound:
		return "Resource Not found"
	case Internal:
		return "Server Error"
	default:
		return "Unexpected Error"
	}
}

func New(code int, previousErr error, messages ...string) Err {
	message := GenerateMessage(code)
	message = manageMultipleErrors(message, messages)

	if previousErr != nil && code != Internal {
		message = previousErr.Error() + "," + message
	}

	err := Err{
		Code:    code,
		Message: message,
	}

	err.GenerateErrors()
	return err
}

func manageMultipleErrors(message string, messages []string) string {
	if len(messages) > 0 {
		message = ""
		for k, mess := range messages {
			if k > 0 {
				message += ","
			}
			message = message + mess
		}
	}
	return message
}

func GetErrorCode(e error) int {
	err := e.(Err)
	return err.Code
}
