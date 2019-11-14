package main

import (
	"fmt"
)

func main() {
	var a []int     //声明int类型切片
	var b = []int{} //声明并初始化
	c := make([]int, 0)

	if a == nil {
		fmt.Println("a==nil")
	}
	fmt.Println(a, len(a), cap(a))

	if b == nil {
		fmt.Println("b==nil")
	}
	fmt.Println(b, len(b), cap(b))

	if c == nil {
		fmt.Println("c==nil")
	}
	fmt.Println(c, len(c), cap(c))

	//判断切片是否为空
	if len(a) == 0 { //不能用a == nil比较
		fmt.Println("切片为空")
	}
}
