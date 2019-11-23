一、内部排序法

1.交换式排序法

​	1).冒泡排序法

```go
package main

import "fmt"

func main() {
	arry := [5]int{24,69,80,17,23}
	for i := 0; i < len(arry) - 1; i++ {
		for j := 0; j < len(arry) - i - 1; j++ {
			if arry[j] > arry[j + 1] {
				var temp = arry[j]
				arry[j] = arry[j + 1]
				arry[j + 1] = temp
			}
		}
	}
	fmt.Println(arry)
}
```

​	2).快速排序法

```go

```

2.选择式排序法

​	1).顺序查找

```go
package main

import "fmt"

func main() {
	//顺序查找,方式一
	arry1 := [4]string{"红", "黄", "蓝", "绿"}
	var colors = ""
	fmt.Println("请输入颜色:")
	fmt.Scanln(&colors)
	for i := 0; i < len(arry1); i++ {
		if arry1[i] == colors {
			fmt.Printf("你找到的颜色%v存在,下标为%v\n",arry1[i], i)
			break
		} else if i == len(arry1) - 1 {
			fmt.Printf("没有找到%v\n", colors)
		}
	}

	//顺序查找,方式二
	var languages = ""
	index := -1
	fmt.Println("请输入语言")
	fmt.Scanln(&languages)
	arry2 := [4]string{"java", "c", "golang", "python"}
	for i := 0; i < len(arry2); i++ {
		if arry2[i] == languages {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v，下标为%v\n", languages, index)
	} else {
		fmt.Printf("没有找到%v\n", languages)
	}
}
```

​	2).二分查找

```go
package main

import "fmt"

func findArray(arry *[6]int, leftIndex int, rightIndex int, findValue int) {
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
	var arry = [6]int{1, 2, 8, 20, 60, 88}
	findArray(&arry, 0, len(arry) - 1, 20)
}
```

