一、抽象

```go
package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

func (account *Account) SaveMoney(pwd string, money float64) {
	if account.Pwd != pwd {
		fmt.Println("输入密码错误")
		return
	}
	if money < 0 {
		fmt.Println("存款金额输入有误")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account) GetMoney(pwd string, money float64) {
	if account.Pwd != pwd {
		fmt.Println("输入密码错误")
		return
	}
	if money < 0 {
		fmt.Println("取款金额输入有误")
		return
	}
	if money > account.Balance {
		fmt.Println("取款金额大于账户余额")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

func (account *Account) Selectoney(pwd string) {
	if account.Pwd != pwd {
		fmt.Println("输入密码错误")
		return
	}
	fmt.Printf("你的账号是%v 你的余额是%v\n", account.AccountNo, account.Balance)
}
func main() {
	a := Account{
		AccountNo: "zgyh123456",
		Pwd:       "123456",
		Balance:   100,
	}
	a.Selectoney("123456")
	a.SaveMoney("123456", 200)
	a.Selectoney("123456")
	a.GetMoney("123456", 120)
	a.Selectoney("123456")
	a.GetMoney("123456", 500)
}
```

二、封装

golang没有特别强调封装，这点不像java，golang面向对象特性做了简化

Model包

```go
package Model

import "fmt"

type person struct {
	Name string
	age int
	sal float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄超出取值范围")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水超出取值范围")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
```

Main包

```go
package main

import (
	"Golang_Notes/Object/Encapsulation/Model"
	"fmt"
)

func main() {
	p := Model.NewPerson("wyl")
	fmt.Println(p)
	p.SetAge(20)
	p.SetSal(8000)
	fmt.Println(p)
}
```

三、继承

```go
package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float64
}

func (stu *Student) SelctScore() {
	fmt.Printf("%v的成绩为%v\n", stu.Name, stu.Score)
}

func (stu *Student) SetScore(score float64) {
	stu.Score = score
	fmt.Printf("%v的成绩修改成功\n", stu.Name)
}

func (stu *Student) GetScore() float64 {
	return stu.Score
}

type XXS struct {
	Student
}

func (xxs *XXS) Testing1() {
	fmt.Println("小学生" + xxs.Name + "正在考试")
}

type DXS struct {
	Student
}

func (dxs *XXS) Testing2() {
	fmt.Println("大学生" + dxs.Name + "正在考试")
}

func main() {
	Studnet1 := &Student{
		Name:  "zt",
		Age:   8,
		Score: 60,
	}
	Studnet1.SelctScore()
	Studnet1.SetScore(66)
	Studnet1.SelctScore()

	Student2 := &Student{
		Name:  "wyl",
		Age:   20,
		Score: 90,
	}
	Student2.SelctScore()
	Student2.SetScore(96)
	Student2.SelctScore()
}
```

注:

1.结构体可以使用嵌套匿名结构体所有的字段和方法，即:首字母大小写的字段和方法都可以使用

2.匿名结构体访问字段可以简写，匿名结构图访问时先访问本结构体，如果没有就去嵌套结构体里面找

3.当结构体和匿名结构体有相同字段或方法时，编译器就近访问，但如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来访问，即具体访问

4.如果嵌入的两个或多个匿名结构体有相同的字段或方法时(结构体本身没有相同的字段或方法时)，在访问时必须指定匿名结构体名字，否则编译会报错

5.如果一个结构图嵌套了一个有名结构体，这种关系为组合，访问组合的结构体的字段或方法时必须带上结构体名字

6.结构体允许嵌套基本数据类型，比如int，但不允许有多个int

7.多重继承就是嵌套了多个结构体