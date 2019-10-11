一、算术运算符

![001](D:\Golang_Notes\Golang运算符与进制转换\Golang运算符\001.png)

注:

1.golang中/两边都是整数，即使结果为浮点数，小数点后面取不到

```go
fmt.Println(10 / 4) //2

var result float32 = 10 / 4
fmt.Println(result) //2，即使变量类型为浮点型

fmt.Println(10.0 / 4) //2.5
```

2.golang中没有++i和--i即前置自增和自减

3.i++必须单独使用，i--类似

```go
var i int = 10
i++
fmt.Println(i)

//编译无法通过
//var j int = 10
//if j++ >10{
//    println("OK")
//}
```

4.取模公式:a % b = a - (a / b) * b

```go
fmt.Println(10 % 3) //1
fmt.Println(-10 % 3) //-1
fmt.Println(10 % -3) //1
fmt.Println(-10 % -3) //-1
```

二、关系运算符

![002](D:\Golang_Notes\Golang运算符与进制转换\Golang运算符\002.png)

三、逻辑运算符

![003](D:\Golang_Notes\Golang运算符与进制转换\Golang运算符\003.png)

注:

a.&&也叫短路与，&&的第一个条件为false，则第二个条件不会判断，最终结果为false

b.||也叫短路或，||的第一个条件为true，则第二个条件不会判断最最最结果为true

四、位运算符

![004](D:\Golang_Notes\Golang运算符与进制转换\Golang运算符\004.png)

注:位运算都是补码运算

五、赋值运算符

![005](D:\Golang_Notes\Golang运算符与进制转换\Golang运算符\005.png)

练习:不用中间变量把a,b值互换

```go
var a int = 10
var b int = 20

a = a + b //a = 10 + 20 = 30

b = a - b //b = 30 - 20 = 10
a = a - b //a = 30 - 10 = 20
fmt.Printf("a=%v b=%v",a ,b) //a=20,b=10
```

六、其他运算符

![006](D:\Golang_Notes\Golang运算符与进制转换\Golang运算符\006.png)

注:golang中没有三元运算符

七、运算符优先级

![007](D:\Golang_Notes\Golang运算符与进制转换\Golang运算符\007.png)