一、接口的定义与实现

语法:

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

- 接口名：使用`type`将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加`er`，如有写操作的接口叫`Writer`，有字符串功能的接口叫`Stringer`等。接口名最好要能突出该接口的类型含义。
- 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

实例:

```go
package main

import "fmt"

type Animal interface {
	say()
	move()
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (d dog) move() {
	fmt.Printf("小狗%s在跑\n", d.name)
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Printf("小猫%s在跑\n", c.name)
}

func main() {
	var animal1 Animal
	dogs := dog{"旺财"}
	animal1 = dogs
	animal1.say()
	animal1.move()
	fmt.Println(animal1)

	var animal2 Animal
	cats := cat{"哆啦A梦"}
	animal2 = cats
	animal2.say()
	animal2.move()
	fmt.Println(animal2)
}
```

注意事项:

1.Golang接口中没有不能包含任何变量

2.Golang的接口不需要显式实现，只要一个变量含有接口中所有的方法，那么这个变量就实现这个接口，所有Golang中没有implement关键字，且Golang中实现接口与接口名字无关

3.只要是自定义数据类型，就可以实现接口，而不仅仅是结构体类型

```go
package main

import "fmt"

type Number interface {
	show()
}

type ints int

func (i ints) show() {
	fmt.Println("i =", i)
}

func main() {
	var i ints = 10
	var N Number = i
	N.show()
}
```

4.接口是引用类型

二、值接收者和指针接收者实现接口的区别

值接受者接口:

```go
package main

import "fmt"

type Animal interface {
	move()
}

type dog struct {
	name string
}

func (d dog) move() {
	fmt.Printf("小狗%s在跑\n", d.name)
}

func main() {
	var a Animal

	dog1 := dog{"wangcai"} //旺财是dog类型
	a = dog1	//a可以接收dog类型
	dog1.move()
	fmt.Println(a)

	dog2 := &dog{"zhangtai"} //张泰是*dog类型
	a = dog2	//a可以接收*dog类型
	dog2.move()
	fmt.Println(a)
}
```

注:

使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针`张泰`内部会自动求值`*张泰`

指针接收者接口:

```go
package main

import "fmt"

type Animal interface {
	move()
}

type dog struct {
	name string
}

func (d *dog) move() {
	fmt.Printf("小狗%s在跑\n", d.name)
}

func main() {
	var a Animal

	//dog1 := dog{"wangcai"} //旺财是dog类型
	//a = dog1	//a不可以接收dog类型
	//dog1.move()
	//fmt.Println(a)

	dog2 := &dog{"zhangtai"} //张泰是*dog类型
	a = dog2	//a可以接收*dog类型
	dog2.move()
	fmt.Println(a)
}
```

注:

此时实现`move接口的是`*dog`类型，所以不能给`x`传入`dog`类型的wangcai，此时x只能存储`*dog`类型的值

三、接口嵌套

1.一个类型实现多个接口

```go
package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}
```

2.多个类型实现同一接口

```go
package main

import "fmt"

// Mover 接口
type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func main() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	x.move()
}
```

注:一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现

```go
// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}
```

3.接口嵌套

```go
package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

func main() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
```

注:接口嵌套时嵌套的接口中不能有相同的方法

```go
package main

type A interface {
	test1()
	test2()
}

type B interface {
	test1()
	test3()
}

type C interface {
	A
	B
}

func main() {
	
}
```

四、空接口

1.空接口类型的变量可以存储任意类型的变量

```go
func main() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}
```

2.使用空接口实现可以接收任意类型的函数参数

```go
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

3.空接口作为map的值

```go
// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
```

4.类型断言

语法:

```go
x.(T)
```

- x：表示类型为`interface{}`的变量
- T：表示断言`x`可能是的类型。

该语法返回两个参数，第一个参数是`x`转化为`T`类型后的变量，第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败。

实例1:

```go
func main() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
```

实例2:

```go
package main

import "fmt"

//Usb接口
type Usb interface {
	Start()
	Stop()
}

//手机
type Phone struct {}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p Phone) Call() {
	fmt.Println("手机打电话")
}

//相机
type Crema struct {}

func (c Crema) Start() {
	fmt.Println("相机开始工作")
}

func (c Crema) Stop() {
	fmt.Println("相机停止工作")
}

//计算机
type Computer struct {}

func (c Computer) working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok == true {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var computer Computer
	var usbArry [3]Usb
	usbArry[0] = Phone{}
	usbArry[1] = Phone{}
	usbArry[2] = Crema{}

	for _, v := range usbArry {
		computer.working(v)
		fmt.Println()
	}
	
	fmt.Println(usbArry)
}
```

如果要断言多次就需要写多个`if`判断，这个时候我们可以使用`switch`语句来实现:

```go
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```

五、接口实现对切片排序

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age int
}

type HeroSlice []Hero

func(hs HeroSlice) Len() int{
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i].Age
	hs[i].Age = hs[j].Age
	hs[j].Age = temp
    //hs[i], hs[j] = hs[j], hs[i]等价于上面三行代码

}

func main() {
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄:%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println("---------------排序前---------------")
	for _, v := range heros {
		fmt.Println(v)
	}
	fmt.Println("---------------排序后---------------")
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Println(v)
	}
}
```

六、接口是继承的一种补充

```go
package main

import "fmt"

type Flying interface {
	fly()
}

type Monkey struct {
	Name string
	Age int
}

func (m *Monkey) ClimbTree() {
	fmt.Println(m.Name, "会爬树")
}

type LittleMonkey struct {
	Monkey
}

func (lm *LittleMonkey) fly() {
	fmt.Println(lm.Name, "会飞翔")
}

func main() {
	Monkey1 := LittleMonkey{Monkey{
		Name: "悟空",
		Age:  6,
	}}
	Monkey1.ClimbTree()
	Monkey1.fly()
}
```

注:

1.继承的价值:解决代码的复用性和可维护性

2.接口的价值:设计好各种规范(方法)，让其自定义去实现这些方法

3.接口比继承更灵活

4.接口在一定程度上能让代码解耦