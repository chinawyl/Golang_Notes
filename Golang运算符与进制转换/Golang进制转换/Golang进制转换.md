

一、golang中进制输出

```go
//二进制输出
var i int = 5
fmt.Printf("%b \n", i)

//八进制输出
var j int = 047
fmt.Println("j =", j)

//十六进制输出
var k int = 0x11
fmt.Println("k =", k)
```

注:

a.golang中二进制无法直接输出，要用%b

b.八进制以数字0开头，十六进制以0x或0X开头

二、原码，反码，补码(原码一般为8位)

1.正数转换

原码、反码和补码是一样的，即看到符号位（第一位）是0，就可以照着写出其他两种码

2.负数转换

1) 原码转换为反码：符号位不变，数值位分别“按位取反” 
2) 反码转换为原码也是一样：符号位不变，数值位分别“按位取反” 
3) 原码转换为补码：符号位不变，数值位按位取反,末位再加1
4) 补码转换为原码：符号位不变，数值位按位取反,末位再加1

即补码的补码等于原码

三、移位运算

例:

1>>2

0000 0001 —— 0000 0000

1<<2

0000 0001 —— 0000 0100

注:

a.<<，符号位不变，低位补0

b.>>，低位溢出，符号位不变，并用符号位补溢出的高位