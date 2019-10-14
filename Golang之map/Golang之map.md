一、map的定义与初始化

1.语法

```go
map定义
map[KeyType]ValueType //KeyType:表示键的类型，ValueType:表示键对应的值的类型

map初始化
make(map[KeyType]ValueType, [cap]) //cap表示map的容量，该参数虽然不是必须的，但最好声明清楚
```

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
    
    //声明的同时初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T", b)
}
```

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

