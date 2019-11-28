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

三、