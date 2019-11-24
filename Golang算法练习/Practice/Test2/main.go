package main

import "fmt"

func main() {
	//给一个已经排好序的数组(升序)添加一个元素，且再次排序(升序)并打印
	array := [5]int{1, 12, 31, 65, 148}
	fmt.Println("没有添加元素的数组:", array)
	var arrays[6]int
	for i := 0; i < len(arrays); i++ {
		if i < len(arrays) - 1 {
			arrays[i] = array[i]
		} else {
			var number int
			fmt.Println("请输入你要添加的数(整数):")
			fmt.Scanln(&number)
			arrays[i] = number
		}
	}
	fmt.Println("添加了元素的数组:", arrays)
	for i := 0; i < len(arrays) - 1; i++ {
		for j:= 0; j < len(arrays) - i -1; j++ {
			if arrays[j] > arrays[j + 1] {
				temps := array[j]
				arrays[j] = arrays[j + 1]
				arrays[j + 1] = temps
			}
		}
	}
	fmt.Println("添加元素后重新排序的数组", arrays)
}
