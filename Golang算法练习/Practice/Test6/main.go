package main

import "fmt"

func main() {
	//在一个长度为10的数组中找出AA这个元素是否存在，如果存在多个AA，打印AA存在的个数
	var array = [10]string{"sas", "AA", "asas", "fff", "AA", "saasqq", "AA", "ssslll", "llll", "AA", }
	indexs := make([]int, len(array))
	sum := 0
	for index, value := range array {
		if value == "AA" {
			indexs = append(indexs, index)
			sum += 1
		}
	}
	if sum != 0 {
		fmt.Println("AA存在")
		fmt.Println("AA的下标为:")
		for i := 0; i < len(indexs); i++ {
			if indexs[i] != 0 {
				fmt.Printf("%v ", indexs[i])
			}
		}
		fmt.Println()
		fmt.Printf("AA存在的个数为%v", sum)
	} else {
		fmt.Println("AA不存在")
	}
}
