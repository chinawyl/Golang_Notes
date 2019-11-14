package main

import (
	"fmt"
)

func main() {
	//切片定义
	var a []string
	var b []int

	fmt.Println(a)
	fmt.Println(b)

	//切片定义及初始化
	var c = []bool{false, true}

	fmt.Println(c)

	//基于数组得到切片
	array := [5]int{5, 6, 7, 8, 9}
	array1 := array[1:4]
	fmt.Println(array1)
	fmt.Printf("%T\n", array1)

	//切片再次切片
	array2 := array1[0:len(array1)]
	fmt.Println(array2)
	fmt.Printf("%T\n", array2)

	//make函数构造切片
	d := make([]int, 5, 10) //第一个参数为切片的元素类型；第二个参数为切片中元素的数量；第三个参数为切片容量，不写默认和切片长度一样
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	//通过len()函数获取切片长度
	fmt.Println(len(d))

	//通过cap()函数获取切片容量
	fmt.Println(cap(d))
}
