package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机生成10个0-100的整数并存储到数组中
	var array [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}
	fmt.Println("初始数组:", array)

	//反转初始数组
	for i := 0; i < len(array) / 2 ; i++ {
		temps := array[len(array) - i - 1]
		array[len(array) - i - 1] = array[i]
		array[i] = temps
	}
	fmt.Println("反转数组:", array)

	//求数组平均值
	sum := 0.0
	for _, value := range array {
		sum += float64(value)
	}
	fmt.Println("数组的平均值为:", sum / float64(len(array)))

	//求数组最大值和最小值的下标
	maxNumber := 0
	minNumber := 100
	for i := 0; i < len(array); i++  {
		if array[i] > maxNumber {
			maxNumber = array[i]
		}
		if array[i] < minNumber {
			minNumber = array[i]
		}
	}
	maxIndex := 0
	minIndex := 0
	for i := 0; i < len(array) ; i++  {
		if array[i] == maxNumber {
			maxIndex = i
		} else if array[i] == minNumber {
			minIndex = i
		}
	}

	fmt.Printf("数组最大值的下标为:%v\n", maxIndex)
	fmt.Printf("数组最小值的下标为:%v\n", minIndex)

	//查找数组里面是否有55
	isOk := false
	for i := 0; i < len(array) ; i++  {
		if array[i] == 55 {
			isOk = true
			break
		} else {
			isOk = false
		}
	}
	if isOk {
		fmt.Println("数组里面存在55\n")
	} else {
		fmt.Println("数组里面不存在55\n")
	}
}
