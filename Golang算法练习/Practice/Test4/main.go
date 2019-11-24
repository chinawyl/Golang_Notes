package main

import "fmt"

func main() {
	//定义一个四行四列的二维数组，逐个从键盘输入值，并将1,2行和3，4行的值交换
	var arrayTwo [4][4]int
	for i := 0; i < len(arrayTwo); i++  {
		for j := 0; j < len(arrayTwo[i]); j++ {
			fmt.Printf("请输入第%v行第%v列的数:", i+1, j+1)
			var number int
			fmt.Scanln(&number)
			arrayTwo[i][j] = number
		}
	}
	//打印输入完成的二维数组
	fmt.Println("输入完成的二维数组")
	for i := 0; i < len(arrayTwo); i++  {
		for j := 0; j < len(arrayTwo[i]); j++ {
			fmt.Printf("%v ", arrayTwo[i][j])
		}
		fmt.Println()
	}
	//打印交换后的二维数组

}
