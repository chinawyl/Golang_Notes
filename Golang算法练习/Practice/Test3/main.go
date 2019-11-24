package main

import "fmt"

func changeArrayTwo(arrayTwo *[3][4]int) {
	(*arrayTwo)[0][0] = 0
	(*arrayTwo)[0][3] = 0
	(*arrayTwo)[2][0] = 0
	(*arrayTwo)[2][3] = 0
}

func main() {
	//定义一个三行四列的二维数组，逐个从键盘输入值，并将四角的值清空
	var arrayTwo [3][4]int
	for i := 0; i < len(arrayTwo); i++  {
		for j := 0; j < len(arrayTwo[i]); j++ {
			fmt.Printf("请输入第%v行第%v列的数:", i+1, j+1)
			var number int
			fmt.Scanln(&number)
			arrayTwo[i][j] = number
		}
	}
	//打印输入完成的二维数组
	fmt.Println("四周未清0的二维数组")
	for i := 0; i < len(arrayTwo); i++  {
		for j := 0; j < len(arrayTwo[i]); j++ {
			fmt.Printf("%v ", arrayTwo[i][j])
		}
		fmt.Println()
	}
	//将二维数组四周清0
	fmt.Println()
	fmt.Println("四周清0后的二维数组")
	changeArrayTwo(&arrayTwo)
	for i := 0; i < len(arrayTwo); i++  {
		for j := 0; j < len(arrayTwo[i]); j++ {
			fmt.Printf("%v ", arrayTwo[i][j])
		}
		fmt.Println()
	}
}
