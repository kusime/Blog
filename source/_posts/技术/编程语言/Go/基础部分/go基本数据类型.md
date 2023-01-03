---
date: 2021-03-15 08:48:01
updated: 2021-03-15 08:48:01
title: go 基本数据类型
index_img: /gallery/2021-08-23-21-59-40.png
tags: 
  - Golang

categories:
  -  Golang

---

# 整型

## int 型

- 代表有符号整型

```go
int8
int16
int32
int64
```

## unsigned 型

- 代表无符号整型
- 后面的数值代表这对应数据占用的比特个数

```go
uint8
uint16
uint32
uint64
```

## 整型运算

- 整型支持常规的运算
- 取模用 % 号
- 比较符号可以使用

# 浮点型

## float32

- 精确到小数点的后 7 位
- 没有 float 类型
- 占用 4 个字节

## float64

- 精确到小数点的后 15 位
- 在比较浮点数的时候要记得测试！
- 尽量使用 float64，math 包很多用的就是这个

### [关于浮点声明](/posts/code/go/变量枚举常量介绍/浮点.go)

### [关于浮点精度](/posts/code/go/变量枚举常量介绍/浮点精度.go)

# 字符串

- 字符串声明之后就不可以改变了
- 类型
  - 解释型 双引号，包括转义字符
  - 非解释型 反引号， 不会解释转义字符，而且支持换行
- 每个字符串最小单位为字节
- 字符串可以被截取，语法和 Python 类似
- 字符串截取越界的话直接报错
- 字符串可以连接，拼接使用加号
- 可以理解字符串为字节数组

# 字符串操作

## [修改字符串](/posts/code/go/strings/修改字符串.go)

- 分为两种
  - 第一种为字节类型
  - 另一种为 utf-8 的多字节字符类型
  - 两种不同的修改操作有不同的方法

## [遍历字符串](/posts/code/go/strings/遍历字符串.go)

- 使用 range 方法来对多字节字符进行遍历
- 可以直接转化为 rune 类型来对其进行遍历

# strings

- 常用的字符串操作被 go 打包到 strings 中
- [子串查找](/posts/code/go/strings/strings包含判断.go)
- [index 返回](/posts/code/go/strings/strings搜引.go)
- [子串替换](/posts/code/go/strings/strings替换.go)
- [次数统计](/posts/code/go/strings/strings统计.go)
- [大小写转化](/posts/code/go/strings/strings大小写转化.go)
- [字符串分割 split](/posts/code/go/strings/strings分割.go)
- [插入](/posts/code/go/strings/strings插入.go)
- [修剪](/posts/code/go/strings/strings修剪.go)

# bool

- go 语言中只有 true，false 这种真假表达
- 使用 0，1 不可以在这个 go 语言中表达真假
- bool 不支持其他类型的转化，比如 bool -> int
- 可以在编译之前替换对应的值真假 x:= (1 == 2)
- golang 支持基本的 bool 运算
  - && and 运算
  - || or 运算
  - ! not 运算
- 在使用上面的运算的时候注意能化简就化简
- 推荐的 bool 变量的名称为 isFishing
