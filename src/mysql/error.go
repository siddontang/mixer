package mysql

import (
	"fmt"
)

type MySQLError struct {
	Code    uint16
	Message string
}

func (e *MySQLError) Error() string {
	return fmt.Sprintf("Error %d:%s", e.Code, e.Message)
}
