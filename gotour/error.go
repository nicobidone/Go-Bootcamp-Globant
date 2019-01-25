package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//return fmt.Sprint(float64(e), " pifio")
	return fmt.Sprint("pifio con el signo")
}

func Sqrt(x ErrNegativeSqrt) (ErrNegativeSqrt, error) {
	if x < 0.0 {
		return 0, nil
	}
	z := 1.0
	for i := 0.0; z-i > 0.0000000000000000001; {
		i = z
		z -= (z*z - float64(x)) / (2 * z)
		fmt.Println("z:", z)
	}
	fmt.Println()
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2)) // Porque se imprime dos veces?
	for {

	}
}
