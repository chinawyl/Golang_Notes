package main

import "fmt"

func main() {
	//将1到10的奇数存储在数组并倒序打印
	var array [5]int
	var j  = 0
	for i := 0; i < 10; i++ {
		if i % 2 == 1 {
			array[j] = i
			j++
		}
	}
	fmt.Println(array)
	for i := 0; i < len(array) / 2; i++ {
		temps := array[len(array) - i - 1]
		array[len(array) - i - 1] = array[i]
		array[i] = temps
	}
	fmt.Println(array)
}
