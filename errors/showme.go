package main

import (
	"errors"
	"fmt"
)

/*
// InternalError contains
type InternalError struct {
	res string
}

func (i *InternalError) Error() string {

	return i.res
}*/

// Sqrt error
func Sqrt(x string) error {
	if x == "C" {
		return &InternalError{"Error Interno"}

	}
	if x == "B" {
		return errors.New("errorThirdParty")
	}
	/*
		if x == "C" {
			return "otro", errors.New("errorOther")
		}*/
	return nil
}

func main() {
	fmt.Println(Sqrt("A"))
	fmt.Println(Sqrt("B"))
	fmt.Println(Sqrt("C"))
	fmt.Println(Sqrt("D"))
}
