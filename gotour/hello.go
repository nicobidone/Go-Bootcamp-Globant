package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	fmt.Println((strings.Fields(s))[2])
	fmt.Println(len(strings.Fields(s)))
	return map[string]int{"x": 1}

}
