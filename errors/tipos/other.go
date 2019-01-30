package tipos

import "fmt"

// OtherError struct
type OtherError struct {
}

func (i OtherError) Error() string {

	return fmt.Sprint("Error de otro tipo")
}
