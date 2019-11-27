一、调用方式不一样

函数调用方式:函数名(实参列表)

方法调用方式:变量.方法名(实参列表)

二、对于普通函数接收者为值类型时不能将指针类型的数据直接传递，反之亦然

```go
package main

import "fmt"

type Person struct {
	Name string
}

func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

func main() {
	p1 := Person{"tom"}
	test01(p1)
	//test01(&p1)

	p2 := Person{"tom"}
	test02(&p2)
}
```

三、对于方法(如struct的方法),接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以

```go
package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03 =",p.Name)
}

func (p *Person) test04() {
	p.Name = "marry"
	fmt.Println("test04 =",p.Name)
}

func main() {
	p3 := Person{"tom"}

	p3.test03() //jack
	fmt.Println("main.test03 =", p3.Name) //tom

	(&p3).test03() //形式上是传入地址，实际上是值拷贝,jack
	fmt.Println("main.test03 =", p3.Name) //tom

	p4 := Person{"wyl"}

	(&p4).test04() //marry
	fmt.Println("main.test04 =", p4.Name)

	p4.test04() //marry,等价于(&p4).test04()
	fmt.Println("main.test04 =", p4.Name) //marry
}
```

注:最终是值拷贝还是地址拷贝取决于绑定方法的类型

四、工厂模式

1.结构体为小写

model包

```go
package model

type student struct {
	Name string
	Score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}
```

main包

```go
package main

import (
	"Golang_Notes/Struct/Factory/model"
	"fmt"
)

func main() {
	var stu = model.NewStudent("wyl", 92.0)
	fmt.Println(*stu)
	fmt.Println(stu.Name, stu.Score)
	fmt.Println((*stu).Name, (*stu).Score)
}
```

2.结构体字段为小写

model包

```go
package model

type student struct {
	Name string
	score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
```

main包

```go
package main

import (
	"Golang_Notes/Struct/Factory/model"
	"fmt"
)

func main() {
	var stu = model.NewStudent("wyl", 92.0)
	fmt.Println(*stu)
	fmt.Println(stu.Name, stu.GetScore())
	fmt.Println((*stu).Name, (*stu).GetScore())
}
```

