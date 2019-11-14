package main

import (
	"fmt"
)

func main() {
	//只声明map类型，但不初始化，a就是初始值nil
	var a map[string]int
	fmt.Println(a == nil) //true

	//map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil) //false

	//map中添加键值对
	a["王艳丽"] = 100
	a["李云龙"] = 200
	fmt.Println(a)
	fmt.Printf("type of a:%T\n", a)

	//声明的同时初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T", b)
}
