package proxy

import (
	"fmt"
)

type MySQLError struct {
	Code    uint32
	Message string
}

func (e *MySQLError) Error() string {
	return fmt.Sprintf("%s (Error: %d)", e.Message, e.Code)
}

//default mysql error, must adapt errname message format
func NewDefaultMySQLError(errCode uint32, args ...interface{}) *MySQLError {
	e := new(MySQLError)
	e.Code = errCode
	if format, ok := ErrName[errCode]; ok {
		e.Message = fmt.Sprintf(format, args...)
	} else {
		e.Message = fmt.Sprint(args...)
	}

	return e
}

func NewMySQLError(errCode uint32, message string) *MySQLError {
	return &MySQLError{errCode, message}
}
