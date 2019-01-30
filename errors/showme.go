package main

import (
	"fmt"

	"github.com/nicob/errors/tipos"
)

/*// InternalError contains
type InternalError struct {
	res string
}

func (i *InternalError) Error() string {

	return i.res
}*/

func whatType(v interface{}) {
	fmt.Printf("type of %v: %T\n", v, v)
}

// Sqrt error
func Sqrt(x string) error {
	if x == "C" {
		return tipos.NewInternalError()
	}
	if x == "B" {
		return tipos.NewThirdPartyError()
	}
	if x == "C" {
		return tipos.NewOtherError()
	}
	return nil
}

func main() {

	whatType(Sqrt("A"))
	whatType(Sqrt("C"))
	fmt.Println(Sqrt("C"))
	fmt.Println(Sqrt("D"))
}
