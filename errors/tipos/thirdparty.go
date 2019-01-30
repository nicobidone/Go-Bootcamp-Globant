package tipos

import "fmt"

// ThirdPartyError struct
type ThirdPartyError struct {
}

func (i ThirdPartyError) Error() string {

	return fmt.Sprint("Error de un tercero")
}
