package utils

import "fmt"

type FamilyAccount struct {
	//保存接受用户输入的选项
	key string
	//控制是否退出for
	loup bool
	//定义一个收支情况
	flags bool
	//定义账号的余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//收支的详情
	detalis string
}

//编写工厂模式的构造方法
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loup:    true,
		flags:   false,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		detalis: "收支\t账户金额\t收支金额\t说\t明",
	}
}

//收支明细方法
func (this *FamilyAccount) ShowMenu() {
	fmt.Println("------------------当前收支明细------------------")
	if this.flags {
		fmt.Println(this.detalis)
	} else {
		fmt.Println("当前没有收支明细...请来一笔吧")
	}
}

//登记收入方法
func (this *FamilyAccount) income() {
	fmt.Println("------------------当前收入明细------------------")
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.detalis += fmt.Sprintf("\n收入\t%v\t%v\t\t%v", this.balance, this.money, this.note)
	this.flags = true
}

//登记支出方法
func (this *FamilyAccount) pay() {
	fmt.Println("------------------当前支出明细------------------")
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足，无法支出")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.detalis += fmt.Sprintf("\n支出\t%v\t%v\t\t%v", this.balance, -this.money, this.note)
	this.flags = true
}

//退出软件方法
func (this *FamilyAccount) exit() {
	fmt.Println("请问你确定要退出吗，请输入y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入y/n")
	}
	if choice == "y"{
		this.loup = false
	}
}

//显式主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n------------------家庭收支记账软件------------------")
		fmt.Println("					1.收支明细")
		fmt.Println("					2.登记收入")
		fmt.Println("					3.登记支出")
		fmt.Println("					4.退出软件")
		fmt.Println("请选择(1-4): ")
		fmt.Scanln(&this.key)

		switch this.key {
			case "1":
				this.ShowMenu()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("请输入正确的选项..")
			}
			if !this.loup {
				break
			}
	}
	fmt.Println("你退出家庭记账软件的使用...")
}