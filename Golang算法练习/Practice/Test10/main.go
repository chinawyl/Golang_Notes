package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//(1).随机生成8个整数，作为评委打分的分数，去掉一个最高分和最低分，剩下的分数除以剩下评委个数为最终得分，找出打分最高和最低评委的下标
	//(2).找出最佳评委(与最终得分差值最小)，找出最差评委(与最终得分差值最大)

	//(1)
	rand.Seed(time.Now().UnixNano())
	var array [8]int
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}
	fmt.Println("评委的打分情况:", array)
	var arrays [8]int = array
	avg := 0.0
	sum := 0.0
	for i := 0; i < len(arrays) - 1; i++ {
		for j := 0; j < len(arrays) - i - 1; j++ {
			if arrays[j] > arrays[j+1] {
				temps := arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = temps
			}
		}
	}
	for _, value := range arrays {
		sum += float64(value)
	}
	avg = sum / float64(len(arrays))
	fmt.Sprintf("%.1f", avg)
	fmt.Printf("最终得分为:%v\n", avg)
	var maxscore = arrays[len(arrays) - 1]
	var minscore = arrays[0]
	var maxscoreindex int
	var minscoreindex int
	for index, value := range array {
		if value == maxscore {
			maxscoreindex = index
		} else if value == minscore {
			minscoreindex = index
		}
	}
	fmt.Printf("打分最高的评委下标是:%v\n", maxscoreindex)
	fmt.Printf("打分最低的评委下标是:%v\n", minscoreindex)

	//(2)
	var subarray[8]float64
	for index, value := range array {
		var sub = float64(value) - avg
		if sub >= 0 {
			sub = sub
			subarray[index] = sub
		} else {
			sub = - sub
			subarray[index] = sub
		}
	}
	fmt.Println("评委打分的差值:", subarray)

	var subarrays[8]float64 = subarray
	for i := 0; i < len(subarrays) - 1; i++ {
		for j := 0; j < len(subarrays) - i - 1; j++ {
			if subarrays[j] > subarrays[j+1] {
				temps := subarrays[j]
				subarrays[j] = subarrays[j+1]
				subarrays[j+1] = temps
			}
		}
	}
	var bestvalue float64 = subarrays[0]
	var badvalue float64 = subarrays[len(subarrays)-1]

	var bastindex int
	var badindex int
	for index, value := range subarray {
		if value == bestvalue {
			bastindex = index
		} else if value == badvalue {
			badindex = index
		}
	}
	fmt.Printf("最佳评委的下标为:%v\n", bastindex)
	fmt.Printf("最差评委的下标为:%v\n", badindex)
}
