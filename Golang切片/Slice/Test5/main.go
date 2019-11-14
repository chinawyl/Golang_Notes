package main

import (
	"fmt"
)

func main() {
	var a []int //没有申请内存
	//a[0] = 100 //编译通过，运行错误

	//添加单一元素
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("a:%v len:%d cap:%d ptr:%p\n", a, len(a), cap(a), a)
	}

	//添加多个元素
	var b []int
	b = append(b, 1, 2, 3, 4, 5)
	fmt.Println(b)

	//切片元素堆加
	c := []int{12, 13, 14, 15}
	b = append(b, c...)
	fmt.Println(b)
}
