package main

import (
	"fmt"
)

func main() {
	//map是无序的，与添加键值对顺序无关
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//只遍历map中的key
	for k := range scoreMap {
		fmt.Println(k)
	}

	//只遍历map中的value
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	//使用delete()函数删除键值对
	//语法delete(map, key)
	delete(scoreMap, "小明") //将小明:100从map中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
