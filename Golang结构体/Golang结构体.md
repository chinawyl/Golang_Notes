一、自定义类型和类别名

```go
package main

import (
	"fmt"
)

//自定义类型

//newInt是基于int类型的自定义类型
type newInt int

//类型别名

//myInt类型是int类型别名
type myInt = int

func main() {
	var a newInt
	var b myInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.newInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}
```

二、结构体定义与实例化

结构体实例化与匿名函数

```go
package main

import (
	"fmt"
)

//定义结构体
type person struct {
	name, city string
	age        int8
}

func main() {
	//结构体实例化
	var p person
	p.name = "wyl"
	p.city = "重庆"
	p.age = 20
	fmt.Printf("p=%v\n", p)  //p={wyl 重庆 20}
	fmt.Printf("p=%#v\n", p) //p=main.person{name:"wyl", city:"重庆", age:20}

	//匿名结构体
	var user struct {
		name string
		age  int
	}
	user.name = "cxk"
	user.age = 18
	fmt.Printf("user=%#v\n", user) //user=struct { name string; age int }{name:"cxk", age:18}
}
```

