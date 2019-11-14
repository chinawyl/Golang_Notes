package main

import (
	"fmt"
)

func main() {
	//二维数组的声明
	cityArray := [4][2]string{
		{"北京", "重庆"},
		{"上海", "杭州"},
		{"深圳", "西安"},
		{"广州", "成都"},
	}
	fmt.Println(cityArray)
	fmt.Println(cityArray[2][0])

	//二维数组的遍历
	for _, value1 := range cityArray {
		fmt.Println(value1)
		for _, value2 := range value1 {
			fmt.Println(value2)
		}
	}
}
