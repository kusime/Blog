---
date: 2021-03-14 15:07:12
updated: 2021-03-14 15:07:12
title: go 注意事项-1
index_img: /gallery/2021-08-23-21-59-40.png
tags: 
  - Golang

categories:
  -  Golang

---

## 推荐使用驼峰命名

```go
const var_dona_do_this = 11
const VarDoThis = 11
```

## 不要以“\n”结尾

```go
fmt.Println(donaDoThis,"\n")
fmt.Println(DoThis)
```

## 变量，包声明了就一定要使用

```go
var v1 int = 44
var v2 string = "ming"
fmt.Println(v1)
 v2 declared but not used
```

## 重复声明 main

- 在一个文件夹中，多个 go 文件
- 只可以有一个 main 函数
- 但是做单文件运行 vscode
  - go run 的时候可以忽略这个问题
- 多文件编译的时候
  - go build 不可以忽略这个
- [参考-1](https://www.jianshu.com/p/bc2bcfaf2a0f)

## init 函数的理解

- init 函数会在 main 函数运行之前运行
- 一般 init 用作初始化全局变量
- 所以其常量，变量也好，声明都不应该在 init 函数里面
- 而是在外面声明了某些量，然后使用 init 对其初始化
- [参考](/posts/code/go/函数/init解释.go)

## 比较数据的时候

- 确保两个比较的对象要是同一个数据类型
- 两个不同的数据类型需要转换
- 但是可以使用无数据类型 （比如说单独一个数据 2）
- 比较的对象不可以超过对象的最大范围

```go
package main
import "fmt"
var i8 int8
var i32 int32

func main() {
  i8 = 2
  i32 = 2
  if i8 == 2 || i32 == 2 {
    // ! 上面两个和常数比较是没问题的
    fmt.Println("yes")
  }
  // if i8 < 4444 { error }

}
```

## 字符串和字符的区别

```go
'a' 这个为字符，这个占用一个字节
'啊' 这个为utf-8的字符类型，占用三个字节
"a" 这个被成为无类型字符串，是两个字节，后面有一个 '\0'标志字符串结尾
对于修改单字符的时候要注意赋值的是字符，而不是字符串

31526
符
就比如说这个，上面的数字就是对应文字的字节码
当有这个字节码输出的时候，直接可以使用fmt.Printf 的 “%c”
来对其进行格式化，格式化的过程可以理解为C语言的printf的过程
```
