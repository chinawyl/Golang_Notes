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

注:结果显示a的类型是`main.NewInt`，表示main包下定义的`NewInt`类型。b的类型是`int`。`MyInt`类型只会在代码中存在，编译完成时并不会有`MyInt`类型

二、结构体定义与实例化

1.定义结构体与实例化

```go
package main

import (
	"fmt"
)

//定义结构体
//type 类型名 struct {
//    字段名 字段类型
//    字段名 字段类型
//    …
//}
type person struct {
	name, city string
	age        int8
}

func main() {
	//结构体实例化
    //var 结构体实例 结构体类型 
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

2.指针类型结构体

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

1).`p3.name = "七米"`其实在底层是`(*p3).name = "七米"`，这是Go语言帮我们实现的语法糖

2).没有初始化的结构体，其成员变量都是对应其类型的零值

3.使用键值对初始化

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

4.使用值的列表初始化

```go
p8 := &person{
	"沙河娜扎",
	"北京",
	28,
}
fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
```

注:

1).必须初始化结构体的所有字段。

2).初始值的填充顺序必须与字段在结构体中的声明顺序一致。

3).该方式不能和键值初始化方式混用。

5.结构体占用一块连续内存

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

6.构造函数

```go
type person struct {
	name, city string
	age        int8
}

//直接复制拷贝，开销较大
func newPerson(name, city string, age int8) person {
	return person{
		name: name,
		city: city,
		age:  age,
	}
}

//结构体是值类型，开销大，换成指针
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	p9 := newPerson("张三", "重庆", 90)
	fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"重庆", age:90}
}
```

三、方法和接收者

1.方法和接收者

1).方法的定义格式:

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

注:

- 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是`self`、`this`之类的命名。例如，`Person`类型的接收者变量应该命名为 `p`，`Connector`类型的接收者变量应该命名为`c`等。
- 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
- 方法名、参数列表、返回参数：具体格式与函数定义相同。

2).示例

```go
package main

import (
	"fmt"
)

//Person 结构体
type Person struct {
	name string
	age  int
}

//NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream 接收方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {
	p1 := NewPerson("wyl", 25)
	p1.Dream()
}
```

2.指针类型和值类型的接收者

1).指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。

2).当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

```go
package main

import (
	"fmt"
)

//Person 结构体
type Person struct {
	name string
	age  int
}

//NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream 接收方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("wyl", 25)
	p1.Dream()
    
    p2 := NewPerson("wyl", 25)
	fmt.Println(p2.age) // 25
	p2.SetAge(30)
	fmt.Println(p2.age) // 30
    
    p3 := NewPerson("wyl", 35)
	fmt.Println(p3.age) // 35
	p3.SetAge2(40) // (*p3).SetAge2(40)
	fmt.Println(p3.age) // 40
}
```

注:什么时候应该使用指针类型接收者

a.需要修改接收者中的值

b.接收者是拷贝代价比较大的大对象

c.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

3.任意类型添加方法

在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的`int`类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

```go
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
func main() {
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}
```

