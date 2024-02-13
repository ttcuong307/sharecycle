package cerror

import (
	"errors"
	"fmt"
	"net/http"
)

type UserMessenger interface {
	error
	UserMessage() string
}

type messenger struct {
	error
	msg string
}

func (msgr messenger) Unwrap() error {
	return msgr.error
}

func (msgr messenger) UserMessage() string {
	return msgr.msg
}

// adds a UserMessenger to err's error chain.
// This func will wrap a nil error.
func WithUserMessage(err error, msg string) error {
	if err == nil {
		err = errors.New(msg)
	}
	return messenger{err, msg}
}

// calls fmt.Sprintf before calling WithUserMessage.
func WithUserMessagef(err error, format string, v ...interface{}) error {
	return WithUserMessage(err, fmt.Sprintf(format, v...))
}

// UserMessage returns the user message associated with an error.
// If no message is found, it checks StatusCode and returns that message.
// Default status is 500, so the default message will be "Internal Server Error".
func UserMessage(err error) string {
	if err == nil {
		return ""
	}
	if um := UserMessenger(nil); errors.As(err, &um) {
		return um.UserMessage()
	}
	return http.StatusText(StatusCode(err))
}
