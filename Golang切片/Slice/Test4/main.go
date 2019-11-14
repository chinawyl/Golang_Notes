package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}

	//普通遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println()

	//for range 遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
