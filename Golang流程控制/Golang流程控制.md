一、单分支

```go
func main(){
    if age := 20; age > 18{ //Golang可以在if中声明一个局部变量
        fmt.Println("可以看片了")
    }
}
```

二、双分支

```go
func main(){
    var a = 1
    var b = 3
    if (a + b) > 1 && (a + b) % 2 == 0 {
        fmt.Println("符合所以条件")
    } else if (a + b) > 1 || (a + b) % 2 == 0 {
        fmt.Println("符合一个条件")
    } else {
        fmt.Println("不符合任何条件")
    }
}
```

注:

1.判断条件不用()包起来，虽然不会报错，编译也能运行

2.不论执行语句有几行，都要用{}包起来

3.if的{注意空格，else和else if 注意前后空格

4.if的条件判断式不能有赋值语句

三、switch

1.case后是一个表达式

```go
package main

import (
	"fmt"
)

func main() {
	var month int
	fmt.Println("请输入月份:")
	fmt.Scanf("%d", &month)
	switch month {
	case 1:
		fmt.Println("一月")
	case 2:
		fmt.Println("二月")
	case 3:
		fmt.Println("三月")
	case 4:
		fmt.Println("四月")
	case 5:
		fmt.Println("五月")
	case 6:
		fmt.Println("六月")
	case 7:
		fmt.Println("七月")
	case 8:
		fmt.Println("八月")
	case 9:
		fmt.Println("九月")
	case 10:
		fmt.Println("十月")
	case 11:
		fmt.Println("十一月")
	case 12:
		fmt.Println("十二月")
	default:
		fmt.Println("输入有误")
	}
}
```

2.case后表达式的数据类型要与switch数据类型保持一致

```go
//会报错
var n1 int32 = 10
var n2 int64 = 20
switch n1 {
    case n2:
    	fmt.Println("OK")
}
```

3.case后可以带多个表达式，使用逗号间隔

```go
var n1 int32 = 10
var n2 int32 = 20
switch n1 {
    case n2, 6, 5:
    	fmt.Println("OK")
}
```

4.case表达式如果后面是常量，则不能重复

```go
var n1 int32 = 10
var n2 int32 = 20
var n3 int32 = 5
switch n1 {
    case n2, 5:
    	fmt.Println("OK1")
    case 5: //这里重复，会报错
    	fmt.Println("OK2")
    case n3: //可以用变量欺骗编译器，不会报错
    	fmt.Println("OK3")
}
```

5.switch可以不带表达式，类似于if-else

```go
var age int = 10
switch {
    case age == 10:
    	fmt.Println("age == 10")
    case age == 20:
    	fmt.Println("age == 20")
    default:
    	fmt.Println("age == 默认")
}
```

6.switch后可以直接声明一个变量，分号结束，不推荐

```go
switch age := 20 {
    case age == 10:
    	fmt.Println("age == 10")
    case age == 20:
    	fmt.Println("age == 20")
    default:
    	fmt.Println("age == 默认")
}
```

7.switch可以用fallthrough穿透

```go
switch age := 20 {
    case age == 10:
    	fmt.Println("age == 10")
    	fallthrough //紧跟后面的case语句或default语句无论条件是否满足都会被执行
    case age == 20:
    	fmt.Println("age == 20")
    default:
    	fmt.Println("age == 默认")
}
```

8.判断interface中变量实际指向的变量类型

```go
var x interface{}
var y = 10.0
x = y
switch i := x.(type) {
    case ni1:
    	fmt.Printf("x的类型~ :%T", i)
    case int:
    	fmt.Printf("x是int型")
    case float64:
    	fmt.Printf("x是float64型")
    case func(int) float64:
    	fmt.Printf("x是func(int)型")
    case bool, string:
    	fmt.Printf("x是bool或string型")
    default:
    	fmt.Printf("x是未知型")
}
```

四、for循环

1.for循环的写法

```go
package main

import (
	"fmt"
)

func main() {
	//标准写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//省略初始语句，但必须保留;
	var j int = 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	//省略初始语句和结束语句
	var k int = 10
	for k > 0 {
		fmt.Println(k)
		k--
	}
    
    //无限循环
    for {
        fmt.Println("Hello")
    }
}

```

2.for循环遍历字符串

```go
package main

import (
	"fmt"
)

func main() {
	//传统方式:按字节遍历，而汉字是三个字节，中文会乱码
	var str string = "我爱你，中国"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	//for range方式:按字符遍历，中文不会乱码
	var str1 string = "我爱你，中国"
	for index, value := range str1 {
		fmt.Printf("index=%d value=%c \n", index, value)
	}

	//解决汉字乱码
	var str2 string = "我爱你，中国"
	str2s := []rune(str2) //把str2转换成[]range即切片
	for i := 0; i < len(str2s); i++ {
		fmt.Printf("%c \n", str2s[i])
	}

}

```

五、while和do while

注:golang中没有while和do while，可以用for循环代替

1.代替while

```go
//使用while方式输出10句Hello World
var i int = 1
for {
    if i > 10 { //循环条件
        break //跳出并结束for循环
    }
    fmt.Println("Hello World", i)
    i++
}
fmt.Println("i =", i)
```

2.代替do...while

```go
//使用do...while方式输出10句Hello World
var j int = 1
for {
    fmt.Println("Hello world", j)
    j++
    if j > 10 {
        break
    }
}
```

六、经典练习

1.打印空心金字塔

```go
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("请输入你想打印的层数:")
	fmt.Scanln(&number)
	for i := 1; i <= number; i++ {
		for k := 1; k <= number-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i*2-1; j++ {
			if j == 1 || j == i*2-1 || i == number {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

```

2.打印九九乘法表

```go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Println()
	}
}

```

七、break

```go
//标签指明终止层数
lable2:
for i := 1; i <= 3; i++ {
    // label1
	for j := 1; j <= 10; j++ {
        if j == 2 {
            // break label1
            break label2
        }
        fmt.Println("j =", j)
	}
```

注:不使用标签终止就是就近原则

八、continue

```go
//标签指明终止层数
lable2:
for i := 1; i <= 3; i++ {
    // label1
	for j := 1; j <= 10; j++ {
        if j == 2 {
            // continue label1
           continue label2
        }
        fmt.Println("j =", j)
	}
```

注:不使用标签终止就是就近原则

九、goto

```go
n := 30
fmt.Println("OK1")
if n > 20 {
    goto lable1
}
fmt.Println("OK2")
fmt.Println("OK3")
fmt.Println("OK4")
lable1:
fmt.Println("OK5")
fmt.Println("OK6")
//打印结果为:OK1 OK2 OK3
```

注:

1.goto语句可以无条件转移到程序中指定的行

2.一般不要用goto语句，容易造成程序流程混乱

十、return(跳出所在方法或函数)

1.如果return是在普通的函数，即表示跳出函数，后面代码不执行

2.如果return是在main函数，表示终止main函数，也就是终止程序
