一、数组定义

```go
//var 数组变量名 [元素数量]T
var a [3]int //没初始化，结果为[0,0,0]
var b [4]int //没初始化，结果为[0,0,0,0]
a = b //不能这么做，因为二者是不同类型
```

注:数组一旦定义数据长度不能变

二、数组初始化

1.定义数组时使用初始值列表初始化

```go
var cityArray = [4]string{"北京", "上海", "深圳", "广州"}
fmt.Println(cityArray)
```

2.编译器推导数组长度

```go
var boolArray = [...]bool{true, false, true, true, false, false}
	fmt.Println(boolArray)
```

3.使用索引值初始化

```go
var languageArray = [...]string{1: "Golang", 3: "Java", 5: "Python", 8: "C"}
	fmt.Println(languageArray)
	fmt.Printf("%T\n", languageArray) //打印数组类型，结果为[9]string
```

三、数组遍历

```go
package main

import (
	"fmt"
)

func main() {
	var cityArray = [4]string{"北京", "上海", "深圳", "广州"}
	fmt.Println(cityArray)

	//for循环遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	//for range遍历
	for _, value := range cityArray {
		fmt.Println(value)
	}
}
```

四、二维数组

```go
package main

import (
	"fmt"
)

func main() {
	//二维数组的声明
	cityArray := [4][2]string{
		{"北京", "重庆"},
		{"上海", "杭州"},
		{"深圳", "西安"},
		{"广州", "成都"},
	}
	fmt.Println(cityArray)
	fmt.Println(cityArray[2][0])

	//二维数组的遍历
	for _, value1 := range cityArray {
		fmt.Println(value1)
		for _, value2 := range value1 {
			fmt.Println(value2)
		}
	}
}
```

注:编译器推导长度只能在外层写，内层不能写

```go
array = [...][2]string //正确
array = [...][...]string //错误
```

五、数组是值类型

```go
package main

import (
	"fmt"
)

func main() {
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
}

func f1(a [3][2]int) {
	a[0][0] = 100
}
/*
结果为
[[1 2] [3 4] [5 6]]
[[1 2] [3 4] [5 6]]
*/
```

注:值类型无法改变

六、练习

求一个数组两个数字和为8的数字的下标

```go
package main

import (
	"fmt"
)

func main() {
	var number = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(number); i++ {
		for j := i + 1; j < len(number); j++ {
			if number[i]+number[j] == 8 && number[i] != number[j] {
				fmt.Printf("i的下标是%v j的下标是%v\n", i, j)
			}
		}
	}
}

```

