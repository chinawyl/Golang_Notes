一、函数定义

![001](D:\Golang_Notes\Golang函数与包\001.png)

![002](D:\Golang_Notes\Golang函数与包\002.png)

二、包的使用

![003](D:\Golang_Notes\Golang函数与包\003.png)

![004](D:\Golang_Notes\Golang函数与包\004.png)

![005](D:\Golang_Notes\Golang函数与包\005.png)

注:

1.方法名要大写，小写为私有，挎包使用要大写，否则无法调用(变量也要大写)

2.文件名和包名最好一致

![006](D:\Golang_Notes\Golang函数与包\006.png)

![007](D:\Golang_Notes\Golang函数与包\007.png)

![008](D:\Golang_Notes\Golang函数与包\008.png)

3.引入包时不需要带src，因为环境变量$GOPATH已经配置了，编译器自动在src下开始

4.包名如果过长，可以取别名，但原来包名就不能用来访问变量和方法

![009](D:\Golang_Notes\Golang函数与包\009.png)

![010](D:\Golang_Notes\Golang函数与包\010.png)

5.在同一文件和包下不能有相同函数

6.要编译生成可执行文件，必须将一个包命名成main包，即package main(main包只能有一个)，编译后会把其他包打包成后缀为.a的库文件