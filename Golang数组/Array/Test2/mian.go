package main

import (
	"fmt"
)

func main() {
	var cityArray = [4]string{"北京", "上海", "深圳", "广州"}
	fmt.Println(cityArray)

	//for循环遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	//for range遍历
	for _, value := range cityArray {
		fmt.Println(value)
	}
}
