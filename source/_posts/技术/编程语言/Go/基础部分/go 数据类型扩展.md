---
date: 2021-03-16 17:26:55
updated: 2021-03-16 17:26:55
title: go 数据类型扩展
index_img: /gallery/2021-08-23-21-59-40.png
tags: 
  - Golang

categories:
  -  Golang

---

# 强制转化

1. 可以强制转化
2. 要主意转化过程中的精度导致数据丢失的问题
3. 比如 1.1(float)->1(int)

```go
// ! 基本语法
NewTypeName(RawData)
```

# 类型别名

- byte 就是 int8
- rune 就是 int32
- go 允许自定别名
- 这个可以理解为 c 语言中的 typedef
- 但是注意和原生类型操作的时候啊就要主要类型转化

```go
import (
  "fmt"
)
type(
  字符串 string
)
func main() {
  var str 字符串
  // ! 这个时候 字符串 就作为类型
  var rawstr string
  rawstr:="原始"
  // ! rawstr+str 这个操作是不允许的！！
  string(str)
  // ! 其实，不管什么数据就是某个固定长度的字节数组，然后所谓的强制转化就是
  // ! 强制套用某套解释这个字节数组
}
```

# 指针

- 在 golang 中的指针和 C 语言中的指针有很大的相似之处
- golang 不支持指针运算
  - 这个理解为在 c 语言可以对指针进行加减
  - 然后对应的指针就跳到对应数据类型大小的下个地址上
- golang 不支持 ->运算符
  - 这个符号就是在结构体去成员的时候
- nil 指针就是空指针，类似于 C 语言中的 NUll 指针
  - 就是指明这个指针不指向储存器中的位置
  - 可以使用 (prt == nil)来判断是否为空指针
- golang 的指针也可以拿来指针套娃
- golang 也有数组这个概念，然后对应的数组指针操作和 C 语言就很像了

```go
a:="Ming-Cloud"
var ptr * string =&a // ! 获取到对应的指针地址
getValue=*ptr // ! 得到指针地址的值
```

## [golang 指针的一些参考](/posts/code/go/变量枚举常量介绍/指针.go)
