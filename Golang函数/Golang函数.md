一、函数定义与调用

1.语法

```go
func 函数名(参数)(返回值){
    函数体
}
```

- 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名。
- 参数：参数由参数变量和参数变量的类型组成，多个参数之间使用`,`分隔。
- 返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用`()`包裹，并用`,`分隔。
- 函数体：实现指定功能的代码块。

2.实例

```go
package main

import (
	"fmt"
)

//无参函数
func sayHello() {
	fmt.Println("Hello")
}

//有参函数
func calculate(a, b float64) float64 { //函数的参数中如果相邻变量的类型相同，则可以省略类型
	ret := a + b
	return ret
}

//简写有参函数
func calculate2(a, b float64) (ret float64) {
	ret = a + b //已经声明返回值为ret和他的类型，不用:=
	return      //已经声明返回值为ret和他的类型，不用写返回值
}

//可变参数函数
func calculate3(a ...int) int {
	fmt.Println(a) //打印的是一个切片
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

//固定参数出现，可变产生放后面
func calculate4(a int, b ...int) int {
	fmt.Println(a, b)
	ret := a
	for _, v := range b {
		ret += v
	}
	return ret
}

//多返回值函数
func calculate5(a, b int) (ret1, ret2 int) {
	ret1 = a + b
	ret2 = a - b
	return
}

func main() {
	//调用无参函数
	sayHello()

	//调用有参函数
	var result float64 = calculate(10.0, 20.1)
	fmt.Println(result)

	//调用可变参数函数
	sum1 := calculate3()                //[]
	sum2 := calculate3(10)              //[10]
	sum3 := calculate3(10, 20)          //[10 20]
	sum4 := calculate3(10, 20, 30)      //[10 20 30]
	fmt.Println(sum1, sum2, sum3, sum4) //0 10 30 60

	//调用有固定参数的可变参数函数
	ret1 := calculate4(0)               //0 []
	ret2 := calculate4(10)              //10 []
	ret3 := calculate4(10, 20)          //10 [20]
	ret4 := calculate4(10, 20, 30)      //10 [20 30]
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60

	//调用多返回值函数
	x, y := calculate5(100, 200)
	fmt.Println(x, y)
}
```

注:go语言没有默认参数

3.defer(延迟执行)

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("start ...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end ...")
}
/*
start ...
end ...
3
2
1
*/
```

练习题

```go
//defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

//A 1 2 3
//B 10 2 12
//BB 10 12 22
//AA 1 3 4
```

Go语言中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行

二、变量作用域

```go
package main

import (
	"fmt"
)

var num = 10 //全局变量

func pln() {
	//先在自己函数中找，找到了就用自己的
	//没有找到就往外层找全局变量
	num := 100 //局部变量
	fmt.Println("变量num =", num)
}

func main() {
	pln() //变量num = 100

	abc := pln              //函数可以作为变量
	fmt.Printf("%T\n", abc) //返回的类型是func()
	abc()                   //直接调用
}
```

注:for循环的i也是局部变量

三、函数作为参数或返回值

```go
package main

import (
	"fmt"
)

//函数作为参数
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	r1 := cal(100, 200, add)
	fmt.Println(r1)
	r2 := cal(100, 200, sub)
	fmt.Println(r2)
    
    r3 := do("s")
    fmt.Println(r3)
}
```

四、匿名函数

```go
package main

import (
	"fmt"
)

func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
```

五、闭包

1.初级案例

```go
//定义一个函数，它的返回值是函数
func a(name string) func() {
	return func() {
		fmt.Println("Hello ", name)
	}
}

func main() {
	//闭包=函数 + 外层变量的引用
	r := a("wyl") //r次数相当于一个闭包
	r()           //相当于执行a函数内部的匿名函数
}
```

```go
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}
```

```go
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}
```

2.中级案例

```go
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
```

3.高级案例

```go
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
```

4.常用内置函数

![001](D:\Golang_Notes\Golang函数与包\001.png)

**注意：**

1. `recover()`必须搭配`defer`使用。
2. `defer`一定要在可能引发`panic`的语句之前定义。
