package main

import (
	"fmt"

	"github.com/nicob/errors/tipos"
)

// SampleType type
type SampleType string

// Sqrt error
func Sqrt(x SampleType) (SampleType, error) {
	if x == "A" {
		return "ese", tipos.InternalError{}
	}
	if x == "B" {
		return "aquel", ThirdPartyError{}
	}
	if x == "C" {
		return "otro", OtherError{}
	}
	return "", nil
}

func errors() {
	fmt.Println(Sqrt("A"))
	fmt.Println(Sqrt("B"))
	fmt.Println(Sqrt("C"))
	//for {
	//}
}
