一、关键字与保留字

1.关键字(25个)

![001](D:\Golang_Notes\Golang变量与数据类型\Golang变量\001.png)

2.保留字(37个)

![002](D:\Golang_Notes\Golang变量与数据类型\Golang变量\002.png)

二、变量的声明

```go
var 变量名 变量类型
```



```go
//标准声明
var name string
var age int
var isOK bool

//批量声明
var (
		a int
		b float32
		c string
		d bool
	)

//短变量声明(在函数内部声明)
m := 10
n := 100
```

三、变量初始化

```go
var 变量名 变量类型 = 表达式
```



```go
//单变量初始化
var name string = "wyl"
var age int = 20

//一次性初始化
var names,ages = "wyls",21 //类型推导

//类型推导(可以在函数外部声明)
var job = "student"
var height = 175
```

四、匿名变量

```go
func foo(string,int){
    return 20,"wyl"
}
func main(){
    x,_ := foo()
    _,y := foo()
    fmt.Println(x)
    fmt.Println(y)
}
```

注:

1.在函数外的每个语句必须以关键字开始

2.匿名变量不占用命名空间，也不会分配内存，不存在重复声明

3._用于占位，表示忽略

五、常量

```go
//单独声明常量
const pi = 3.14

//一次性声明多个常量
const(
	n1 = 10
    n2 //不赋值默认为10
    n3 //同上
)
```



六、iota

iota声明跳过

```go
const(
	n1 = iota //0
    n2 //1
    _  //忽略
    n4 //3
)
```

iota声明中间插队

```go
const(
	n1 = iota //0
    n2 //1
    n3 = 100  //2
    n4 //3
)

const(
	n5 = iota //0
)
```

iota声明数量级

```go
const(
	iota = _
    KB = 1 << (10 * iota) //1<<10
    MB = 1 << (10 * iota) //1<<20
    GB = 1 << (10 * iota) //1<<30
    TB = 1 << (10 * iota) //1<<40
    PB = 1 << (10 * iota) //1<<50
)
```

多个iota声明在一行

```go
const(
	a,b = iota + 1,iota + 2 //1,2
    c,d						//3,4
    e.f						//5,6
)
```

注:

1.iota只能用于常量

2.iota在出现const关键字时将被重置为0，const中每新增一行将计数一次

