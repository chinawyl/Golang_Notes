package main

import (
	"fmt"
)

func main() {
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
}

func f1(a [3][2]int) {
	a[0][0] = 100
}
