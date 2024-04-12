package main

import (
	"fmt"
	"strconv"
)

type Test struct {
	Title *string
}

func main() {
	a := "10"

	// t := Test{Title: &a}
	i, err := strconv.Atoi(a)

	if err != nil {
		return
	}

	fmt.Println(i)
}