一、数据类型概要

![001](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\001.png)

注:

1.uint代表无符号

2.字符型只能存字母，字母占一个字节，汉字占三个字节

二、整数类型

1.有符号

![002](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\002.png)

2.无符号

![003](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\003.png)

3.特殊

![004](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\004.png)

注:rune处理中文

4.查看数据类型(Printf)和字节数(Sizeof，注意导包)

![005](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\005.png)

二、浮点数类型

1.基本类型

![006](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\006.png)

2.浮点类型精度表示

![007](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\007.png)

3.浮点型使用细节

![008](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\008.png)

注:没有OS即没有float，e后面跟负数相当于除

三、字符型

![009](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\009.png)

注:

1.Golang没有字符串类型，一般用byte，使用UTF-8编码

2.如果小于255用byte，否则用int

四、布尔类型

只能取true或false，不能用0或1代替

五、字符串类型

1.在Golang中字符串不可变

2.使用反引号``，输出大批代码，不用转义

3.字符串拼接过长时可以换行，但要把+留在上面

六、基本数据类型默认值

![010](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\010.png)

七、数据类型转换

注:Golang基本数据类型转换必须显示转换，不能自动转换

1.基本数据类型相互转换

![011](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\011.png)

![012](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\012.png)

![013](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\013.png)

2.基本数据类型转string类型

方式一

![014](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\014.png)方式二

![015](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\015.png)

![016](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\016.png)

注意导入strconv包

3.string转基本数据类型

![017](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\017.png)

![018](D:\Golang_Notes\Golang变量与数据类型\Golang数据类型\018.png)

注:如果将"hello"转换整数，结果为0，请注意转换有效性