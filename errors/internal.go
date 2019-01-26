package errors

import "fmt"

// InternalError contains
type InternalError struct {
}

func (i InternalError) Error() string {

	return fmt.Sprint("Error interno")
}
