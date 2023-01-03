---
date: 2021-03-22 08:05:10
updated: 2021-03-22 08:05:10
title: go 函数基本介绍
index_img: /gallery/2021-08-23-21-59-40.png
tags: 
  - Golang

categories:
  -  Golang

---

# 函数基本声明

- 不支持 嵌套，重载，默认参数
- 特殊功能
  - 多返回值
  - 延迟语句
  - 可变参数

```go
func funcName(input1 type1 , input1 type2 )(output1 type1 ,output2 type2){
    return value1 ,value2
}
```

- 特殊的声明方式
  - 只有一个返回值

```go
func funcNAME(int,int) int {
    // ! 上面是没有声明函数输入的参数名字
    // ! 所以结合之前的，传递参数的时候可能会出现点问题
}
```

# [函数类型](/posts/code/go/函数/go函数类型.go)

- 这里的函数类型是指一种声明方式
- 这里就可以支持函数作为参数传入,不过也只是和类型那一个参数对等的地位
- 这里有点类似与 C 语言中的 type def
- 函数值是可以相互比较的
  - 引用的都是相同函数
  - 或者函数返回至是 nil

```go
type functionType func (int,int) int

add := functionType

这样就是把变量指向addNum函数签名

```

## 关于公有私有

- 如果函数名字开头为大写就是公开
- 如果是小写就是私有
- 这个规则对于变量，等实体厚实有用的

## [可变参数](/posts/code/go/函数/函数可变参数.go)

- 为了对可变参数的支持，应该这样声明
- ...int 这个是可变参数，需要指定类型
- 这里指的可变是类似于切片，数组，等可以变长的数据类型

```go
func MyFunction(arg ...int) int {
    // ! code

}
```
