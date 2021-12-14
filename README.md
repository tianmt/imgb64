# image 和 base64 转换工具（cmdline）

是不是总是在找在线转换工具？

是不是在线转换工具有些时候也不太友好？

自己编一个离线版本吧，加到环境变量里，做到随手转换；

自己还可以随便改代码，得到自己想要的结果；



## 一、用法

> 此处展示当前目录下使用，把编译得到的二进制文件放到 `user/local/bin` 目录下就可以全局使用了

```shell
tianmt@~$ go build imgb64.go
tianmt@~$ ./imgb64 -h
PRINT HELP INFO.
Usage: imgb64 [-bih] [-ip image_path] [-bs base64_string] [-bp base64_path] [-tp target_path]
Options:
  -b	base64 string to image.
  -bp string
    	base64 path.
  -bs string
    	base64 string.
  -h	help info.
  -i	image to base64 string.
  -ip string
    	image path.
  -tp string
    	target path.
```



## 二、使用示例

### 1、图片转 base64，打印数据，并保存到文件

```shell
tianmt@~$ imgb64 -i -ip ./test.jpg -tp test.b64
```



### 2、base64 转图片

```shell
tianmt@~$ imgb64 -b -bp ./test.b64 -tp test.jpg
```



### 3、base64 string 转图片

```shell
tianmt@~$ imgb64 -b -bs 你的base64的字符串 -tp test.jpg
```