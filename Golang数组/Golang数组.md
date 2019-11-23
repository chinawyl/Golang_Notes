一、数组定义

```go
//var 数组变量名 [元素数量]T
var a [3]int //没初始化，结果为[0,0,0]
var b [4]int //没初始化，结果为[0,0,0,0]
a = b //不能这么做，因为二者是不同类型
```

注:

1.数组一旦定义数据长度不能变

2.数组创建后如果没有赋值会有默认值

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

注:数组内存地址是连续的，数组第一个元素地址就是数组的首地址

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
    for i := 0; i < len(cityArray); i++ {
        fmt.Println(i)
        for j := 0; j < len(cityArray[i]); j++{
            fmt.Println(j)
        }
    }
    
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

练习:输入3个班5个学生成绩并求每个班的平均分和所有班的平均分

```go
package main

import "fmt"

func main() {
	var Class_Grades [3][5]float64
	for i := 0; i < len(Class_Grades); i++ {
		for j := 0; j < len(Class_Grades[i]); j++ {
			fmt.Printf("请输入第%v个班第%v个同学的分数:", i+1, j+1)
			fmt.Scanln(&Class_Grades[i][j])
		}
	}

	number := 0
	total := 0.0
	for i := 0; i < len(Class_Grades); i++ {
		sum := 0.0
		for j := 0; j < len(Class_Grades[i]); j++ {
			sum += Class_Grades[i][j]
		}
		total += sum
		number = len(Class_Grades[i])
		fmt.Printf("第%v个班的总分为%v,平均分为%v\n", i+1, sum, sum / float64(number))
	}
	//换行要把,留在上面
	var totalavg = total / float64(len(Class_Grades) * number)
	totalavgs := fmt.Sprintf("%.2f", totalavg)
	fmt.Printf("%v个班的总分为%v,平均分为%v\n", len(Class_Grades), total, totalavgs)
}
```

五、数组是值类型

1.值类型无法改变

```go
package main

import (
	"fmt"
)

func f1(a [3][2]int) {
	a[0][0] = 100
}

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
/*
结果为
[[1 2] [3 4] [5 6]]
[[1 2] [3 4] [5 6]]
*/
```

2.如果想使用函数修改原来的数组，可以使用引用传递(指针方式)

```go
package main

import "fmt"

func test(arry *[4]int) {
	(*arry)[0] = 11
}

func main() {
	arry := [4]int{1, 2, 3, 4}
	test(&arry)
	fmt.Println("arry =", arry)
}
```

六、求数组[1, 3, 5, 7, 8]的和并且把两个元素之和等于8的下标求出

```go
package main

import (
	"fmt"
)

func main() {
	var arrayTest = [...]int{1, 3, 5, 7, 8}
	sum := 0
	for index, value := range arrayTest {
		sum += value
		for i := index + 1; i < len(arrayTest); i++ {
			if arrayTest[i]+value == 8 {
				fmt.Printf("(%v,%v)\n", index, i)
			}
		}
	}
	fmt.Println(sum)
}
```

七、给一个数组生成9个随机数并反转数组

```go
	var arryss [9]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arryss); i++ {
		arryss[i] = rand.Intn(100)
	}

	fmt.Println(arryss)

	temps := 0
	for i :=0; i < len(arryss) / 2; i++ {
		temps = arryss[len(arryss) - 1 - i]
		arryss[len(arryss) - 1 - i] = arryss[i]
		arryss[i] = temps
	}

	fmt.Println(arryss)
```