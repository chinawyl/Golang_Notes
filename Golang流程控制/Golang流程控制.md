一、单分支

![001](001.png)

![002](002.png)

二、双分支

![003](003.png)

注:

1.执行条件不用()包起来，虽然不会保存，编译也能运行

2.不论执行语句有几行，都要用{}包起来

3.else不用换号

4.if的条件判断式不能有赋值语句

示例:

![004](004.png)

注意导入math包
<<<<<<< HEAD

三、switch

![005](005.png)

![006](006.png)

![007](007.png)

1)case后是一个表达式

![008](008.png)

2)case后表达式的数据类型要与switch数据类型保持一致

![009](009.png)

![010](010.png)

注:

如果num2是一个具体的值，如20就可以，类似于20把值赋给num1，而num2赋值给num1不行

3)case后可以带多个表达式

![011](011.png)

4)case表达式如果后面是常量，则不能重复

![012](012.png)

![013](013.png)

注:可以用变量

![014](014.png)

7)switch可以不带表达式，类似于if-else

![015](015.png)

8)switch后可以直接声明一个变量，分号结束，不推荐

![016](016.png)

9)switch可以用fallthrough穿透

![017](017.png)

10)判断interface中变量实际指向的变量类型

![018](018.png)

四、for循环

1.for循环的三种写法

![019](019.png)

2.for循环遍历字符串

![020](020.png)

3.练习

![021](021.png)

![022](022.png)

五、while和do while

golang中没有while和do while，可以用for循环代替

![023](D:\Golang_Notes\Golang流程控制\023.png)

![024](D:\Golang_Notes\Golang流程控制\024.png)

六、经典练习

1.打印空心金字塔

![025](D:\Golang_Notes\Golang流程控制\025.png)

![026](D:\Golang_Notes\Golang流程控制\026.png)

2.统计3个班，每个班5名同学，求出各个班平均分，所有班级平均分，所有班级及格人数

![027](D:\Golang_Notes\Golang流程控制\027.png)

3.打印九九乘法表

![028](D:\Golang_Notes\Golang流程控制\028.png)

七、break

1.练习

![029](D:\Golang_Notes\Golang流程控制\029.png)

2.标签指明终止层数

![030](D:\Golang_Notes\Golang流程控制\030.png)

![031](D:\Golang_Notes\Golang流程控制\031.png)

注:不使用标签就是就近原则

八、continue

1.练习

![032](D:\Golang_Notes\Golang流程控制\032.png)

![033](D:\Golang_Notes\Golang流程控制\033.png)

2.break和continue综合练习

![034](D:\Golang_Notes\Golang流程控制\034.png)

九、goto

![035](D:\Golang_Notes\Golang流程控制\035.png)

![036](D:\Golang_Notes\Golang流程控制\036.png)

打印结果为:OK1 OK2 OK3

十、return(跳出所在方法或函数)

![037](D:\Golang_Notes\Golang流程控制\037.png)

