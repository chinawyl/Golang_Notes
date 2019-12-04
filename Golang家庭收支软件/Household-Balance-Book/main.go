package main

import "fmt"

func main() {
	//声明一个变量，保存接受用户输入的选项
	key := ""
	//声明一个变量，控制是否退出for
	loup := true
	//定义一个收支情况变量
	flags := false
	//定义账号的余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//收支的详情
	detalis := "收支\t账户金额\t收支金额\t说\t明"
	//显式这个主菜单
	for {
		fmt.Println("\n------------------家庭收支记账软件------------------")
		fmt.Println("					1.收支明细")
		fmt.Println("					2.登记收入")
		fmt.Println("					3.登记支出")
		fmt.Println("					4.退出软件")
		fmt.Println("请选择(1-4): ")

		fmt.Scanln(&key)

		switch key {
			case "1":
				fmt.Println("------------------当前收支明细------------------")
				if flags {
					fmt.Println(detalis)
				} else {
					fmt.Println("当前没有收支明细...请来一笔吧")
				}
			case "2":
				fmt.Println("------------------当前收入明细------------------")
				fmt.Println("本次收入金额")
				fmt.Scanln(&money)
				balance += money
				fmt.Println("本次收入说明")
				fmt.Scanln(&note)
				detalis += fmt.Sprintf("\n收入\t%v\t%v\t\t%v", balance, money, note)
				flags = true
			case "3":
				fmt.Println("------------------当前支出明细------------------")
				fmt.Println("本次支出金额")
				fmt.Scanln(&money)
				if money > balance {
					fmt.Println("余额不足，无法支出")
					break
				}
				balance -= money
				fmt.Println("本次支出说明")
				fmt.Scanln(&note)
				detalis += fmt.Sprintf("\n支出\t%v\t%v\t\t%v", balance, -money, note)
				flags = true
			case "4":
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
					loup = false
				}
			default:
				fmt.Println("请输入正确的选项..")
		}
		if !loup {
			break
		}
	}
	fmt.Println("你退出家庭记账软件的使用...")
}