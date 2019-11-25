package main

import "fmt"

func main() {
	//输入8个整数，找出其中大于平均值的数的个数，小于平均值的数的个数
	var array [8]int
	var number int
	var moreavg = 0
	var lessavg = 0
	var equalavg = 0
	var sum = 0.0
	for i := 0; i < len(array); i++ {
		fmt.Println("请输入整数:")
		fmt.Scanln(&number)
		array[i] = number
	}
	fmt.Println(array)
	for _, value := range array {
		sum += float64(value)
	}
	avg := sum / float64(5)
	fmt.Printf("平均数为%v", avg)
	for _, value := range array {
		if float64(value) < avg {
			lessavg++
		} else if float64(value) > avg {
			moreavg++
		} else {
			equalavg++
		}
	}
	fmt.Printf("大于平均值的数的个数:%v\n", moreavg)
	fmt.Printf("小于平均值的数的个数:%v\n", lessavg)
	fmt.Printf("等于平均值的数的个数:%v\n", equalavg)
}
