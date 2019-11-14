package main

import (
	"fmt"
)

func main() {
	a := make([]int, 3)
	b := a

	b[0] = 100

	fmt.Println(a)
	fmt.Println(b)
}
