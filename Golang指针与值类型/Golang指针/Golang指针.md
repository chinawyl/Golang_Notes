一、指针及其内存布局

```go
package main

import (
	"fmt"
)

func main() {
	//普通变量
	var i int = 10

	fmt.Printf("i = %v\n", i)     //打印普通变量值
	fmt.Printf("i的地址 = %v\n", &i) //打印普通变量地址

	//指针变量
	var ptr *int = &i

	fmt.Printf("ptr=%v\n", ptr)      //打印ptr值，是个地址
	fmt.Printf("ptr的地址=%v\n", &ptr)  //打印ptr的地址
	fmt.Printf("ptr指向的值=%v\n", *ptr) //打印ptr指向的变量的值，即i=10

	*ptr = 11             //i的值会改变
	fmt.Printf("i=%v", i) //i=11
}
```

![001](001.png)

二、指针使用细节

指针变量于执向变量数据类型需要一致

```go
var a int = 300
var ptr *float32 = &a //编译会报错
```

注:值类型都有对应指针类型，形式为*数据类型

三、指针传值实例

```go
package main

import (
	"fmt"
)

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modify1(a)
	fmt.Println(a) //1
	modify2(&a)
	fmt.Println(a) //100
}
```

四、new与make区别

1.实例

```go
func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["wyl"] = 100
	fmt.Println(b)
}
```

注:

1)`var a *int`只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值

2)`var b map[string]int`只是声明变量b是一个map类型的变量，需要使用make函数进行初始化操作之后，才能对其进行键值对赋值

2.new函数

使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值

```go
func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}	
```

3.make函数

make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了

```go
func main() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["wyl"] = 100
	fmt.Println(b)
}
```

4.new函数与make函数区别

1)二者都是用来做内存分配的。

2)make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；

3)而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。