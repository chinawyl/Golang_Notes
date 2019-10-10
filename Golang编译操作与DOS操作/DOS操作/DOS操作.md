一、目录操作

1.查看当前目录内容:dir

2.切换目录:cd /d f:(切换到f盘)

注:切换到下层目录用\，cd d:\Test\test1

3.顶级目录:cd \，上层目录:cd ..

4.创建目录:md 目录名 (创建多个目录空格隔开)

5.删除空目录:rd 目录名

6.删除目录及子目录:rd /q/s 目录名 (不带询问)

7.删除目录及子目录:rd /s 目录名 (带询问)

二、文件操作

1.新建文件:echo 追加内容 > 路径\新建文件名

例:echo abc > d:\Test\test1\abc.txt

2.复制文件 

copy 被拷贝文件 拷贝路径(可在最后指定名字)

例:copy abc.txt d:\Tests\a.txt

3.移动文件

move 被移动文件 移动路径(可在最后指定名字)

例:move abc.txt f:

4.删除文件

del 文件名 删除路径

例:del abc.txt f:

5.删除所有txt文件

del *.txt

三、其他指令

1.清屏:cls

2.退出:exit