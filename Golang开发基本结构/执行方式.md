一、执行方式

首先cd到被执行文件main文件夹，然后如下图执行

方法一.先通过go build编译，再执行执行编译文件

方法二.直接go run执行，经过底层编译

![001](D:\Golang_Notes\Golang开发基本结构\001.png)

两种方法区别:

![002](D:\Golang_Notes\Golang开发基本结构\002.png)

二、指定编译文件名(shift+alt+下箭头是快速复制)

go build -o 指定编译文件名 被编译文件

![003](D:\Golang_Notes\Golang开发基本结构\003.png)