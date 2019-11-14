package main

import (
	"fmt"
	"sort"
)

func main() {
	//append(a[:index], a[index+1:]...)
	a := []string{"北京", "上海", "深圳", "广州"}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)

	//排序
	var b = []int{8, 7, 13, 9, 2}
	sort.Ints(b)
	fmt.Println(b)
}
