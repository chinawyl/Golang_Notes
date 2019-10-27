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

指针类型结构体

```go
package main

import (
	"fmt"
)

type person struct {
	name string
	city string
	age  int8
}

func main() {
	//创建指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

	//指针类结构体实例化
	p2.name = "wyl"
	p2.age = 18
	p2.city = "重庆"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"wyl", city:"重庆", age:18}

	//取结构体的地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}

	//没有初始化的结构体
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
}
```

注:

1.`p3.name = "七米"`其实在底层是`(*p3).name = "七米"`，这是Go语言帮我们实现的语法糖

2.没有初始化的结构体，其成员变量都是对应其类型的零值

使用键值对初始化

```go
package main

import (
	"fmt"
)

type person struct {
	name string
	city string
	age  int8
}

func main() {
	//使用键值对对结构体进行初始化
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}

	//结构体指针进行键值对初始化
	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}

}
```

注:当某些字段没有初始值的时候，该字段可以不写，此时，没有指定初始值的字段的值就是该字段类型的零值

使用值的列表初始化

```go
p8 := &person{
	"沙河娜扎",
	"北京",
	28,
}
fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
```

注:

1. 必须初始化结构体的所有字段。
2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
3. 该方式不能和键值初始化方式混用。

结构体占用一块连续内存

```go
type test struct {
	a int8
	b int8
	c int8
	d int8
}
n := test{
	1, 2, 3, 4,
}
fmt.Printf("n.a %p\n", &n.a) //n.a 0xc0000a0060
fmt.Printf("n.b %p\n", &n.b) //n.b 0xc0000a0061
fmt.Printf("n.c %p\n", &n.c) //n.b 0xc0000a0062
fmt.Printf("n.d %p\n", &n.d) //n.b 0xc0000a0063
```

构造函数

```go
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

p9 := newPerson("张三", "沙河", 90)
fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}

//结构体是值类型，开销大，换成指针
p10 := &newPerson("张三", "沙河", 90)
```

