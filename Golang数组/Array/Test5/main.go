package main

import (
	"fmt"
)

func main() {
	var number = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(number); i++ {
		for j := i + 1; j < len(number); j++ {
			if number[i]+number[j] == 8 && number[i] != number[j] {
				fmt.Printf("i的下标是%v j的下标是%v\n", i, j)
			}
		}
	}
}
