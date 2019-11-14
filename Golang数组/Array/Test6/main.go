package main

import (
	"fmt"
)

func main() {
	var arrayTest = [...]int{1, 3, 5, 7, 8}
	sum := 0
	for index, value := range arrayTest {
		sum += value
		for i := index + 1; i < len(arrayTest); i++ {
			if arrayTest[i]+value == 8 {
				fmt.Printf("(%v,%v)\n", index, i)
			}
		}
	}
	fmt.Println(sum)
}
