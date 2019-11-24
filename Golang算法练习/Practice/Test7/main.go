package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findArray(arry *[10]int, leftIndex int, rightIndex int, findValue int) {
	if leftIndex > rightIndex {
		fmt.Println("你要找的数不存在")
		return
	}
	middleIndex := (leftIndex + rightIndex) / 2
	if (*arry)[middleIndex] > findValue {
		findArray(arry, leftIndex, middleIndex - 1, findValue)
	} else if (*arry)[middleIndex] < findValue {
		findArray(arry, middleIndex + 1, rightIndex, findValue)
	} else {
		fmt.Printf("你要找的数%v的下标为%v\n", findValue, middleIndex)
	}
}

func main() {
	//随机生成10个整数(1-100)存入数组，进行冒泡排序，并用二分查找查询是否存在90这个数，如果存在便打印其下标
	var array [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}
	//打印随机数组
	fmt.Println(array)
	//二分排序并查找90
	findArray(&array, 0, len(array) - 1, 90)
}
