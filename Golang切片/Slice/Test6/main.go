package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 5)
	copy(b, a) //a数组赋值给b数组
	fmt.Println(a)
	fmt.Println(b)

	b[0] = 100
	fmt.Println(a) //[1, 2, 3, 4, 5]
	fmt.Println(b) //[100, 2, 3, 4, 5]
}
