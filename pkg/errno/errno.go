package errno

import (
	"fmt"
)

type Errno struct {
	Code    int
	Message string
}

// Get the message
func (err *Errno) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

func (err *Err) Add(message string) *Err {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) *Err {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func Decode(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
		return InternalServerError.Code, err.Error()
	}
}

func IsErrUserFieldMissed(err error) bool {
	code, _ := Decode(err)
	return code == UserFieldMissed.Code
}
