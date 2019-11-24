一、map的定义与初始化

1.语法

```go
map定义
map[KeyType]ValueType //KeyType:表示键的类型，ValueType:表示键对应的值的类型

map初始化
make(map[KeyType]ValueType, [cap]) //cap表示map的容量，该参数虽然不是必须的，但最好声明清楚
```

注:map只声明是不会分配内存的

2.举例

```go
package main

import (
	"fmt"
)

func main() {
	//只声明map类型，但不初始化，a就是初始值nil
	var a map[string]int
	fmt.Println(a == nil) //true

	//map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil) //false

	//map中添加键值对
	a["王艳丽"] = 100
	a["李云龙"] = 200
	fmt.Println(a)
	fmt.Printf("type of a:%T\n", a)
    
    //声明的同时make
    var c = make(map[int]string)
    c[1] = "aaa"
    c[2] = "bbb"
    c[3] = "ccc"
    
    //声明的同时初始化
	b := map[int]bool{
		1: true,
		2: false, //不要忘记,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T", b)
}
```

注:

1).map的key不能重复，key重复会以最后的key-value存储

2).map存储时为无序

3).map增加key-value时如果达到容量，可以继续添加，不会发生panic错误

二、判断键是否存在

1.语法

```go
value, ok := map[key]
```

2.

```go
package main

import (
	"fmt"
)

func main() {
	//判断键是否不存在
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
```

三、map的遍历

1.无序遍历与删除键值对

```go
package main

import (
	"fmt"
)

func main() {
    //map是无序的，与添加键值对顺序无关
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
    
    //只遍历map中的key
    for k := range scoreMap {
		fmt.Println(k)
	}
    
    //只遍历map中的value
    for _, v := range scoreMap {
		fmt.Println(v)
	}  
    
    //使用delete()函数删除键值对
    //语法delete(map, key)
    delete(scoreMap, "小明")//将小明:100从map中删除
	for k,v := range scoreMap{
		fmt.Println(k, v)
	}
}
```

2.按顺序遍历

```go
package main

import (
	"fmt"
    "math/rand"
	"sort"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```

四、元素为map类型的切片

```go
package main

import (
	"fmt"
)

func main() {
	var mapSlice = make([]map[string]string, 3) //对切片初始化
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	mapSlice[0] = make(map[string]string, 10) //对切片中的map元素进行初始化
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
```

注:如果想要动态增加map切片可以使用append函数

```go
package main

import "fmt"

func main() {
	var message = make([]map[string]string, 2)
	message[0] = make(map[string]string, 3)
	message[0]["name"] = "wyl"
	message[0]["age"] = "18"
	message[0]["address"] = "重庆"

	message[1] = make(map[string]string, 3)
	message[1]["name"] = "sxb"
	message[1]["age"] = "22"
	message[1]["address"] = "四川"
	fmt.Println(message)
	newmessage := map[string]string{
		"name" : "ty",
		"age" : "22",
		"address" : "上海",
	}
	message = append(message, newmessage)
	fmt.Println(message)
}
```

五、值为切片类型的map

```go
package main

import (
	"fmt"
)

func main() {
	var sliceMap = make(map[string][]string, 3) //map初始化
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2) //切片初始化
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
```

六、map的增删改查

```go
package main

import "fmt"

func main() {
	city := make(map[int]string, 10)
	city[1] = "重庆"
	city[2] = "上海"
	city[3] = "成都"

	//增加
	city[4] = "东京"

	//修改
	city[3] = "北京"

	//删除
	delete(city, 1) //即使键不存在也不会报错

	//map清空
	city = make(map[int]string) //重新make一个新空间，原来的map当做垃圾回收

	//查询
	val, ok := city[1]
	if ok {
		fmt.Printf("有1这个key，值为%v", val)
	} else {
		fmt.Printf("没有1这个key，值不存在")
	}
}
```

七、练习:统计一个字符串中每个单词出现的次数，比如：”how do you do”中how=1 do=2 you=1

```go
package main

import(
	"fmt"
	"strings"
)

func main() {
	var s = "How do you do"
	var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}

	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
```

