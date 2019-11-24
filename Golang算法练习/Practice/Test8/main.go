package main

import "fmt"

func main() {
	//一个数组有5个数，找出数组中最大的数和最小的数的坐标
	array := [5]int{1, 65, 4, 28, 79}
	fmt.Println(array)
	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - i - 1; j++ {
			if array[j] > array[j + 1] {
				var temp = array[j]
				array[j] = array[j + 1]
				array[j + 1] = temp
			}
		}
	}
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
}
