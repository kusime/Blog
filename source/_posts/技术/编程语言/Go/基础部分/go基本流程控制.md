---
date: 2021-03-17 16:05:18
updated: 2021-03-17 16:05:18
title: go 基本流程控制
index_img: /gallery/2021-08-23-21-59-40.png
tags: 
  - Golang

categories:
  -  Golang

---

# 条件语句

## [基本 if， else，else if 的使用](/posts/code/go/流程控制/if.go)

- 条件语句不应该嵌套
- else 代码块必须紧跟着 if 代码块
- 多判断条件的可以使用 else if 代码来进行代替
- 支持初始化子语句
  - 只有一条
  - 在判断条件之前
  - 可以调用函数来对数值的初始化

# 选择语句

## [switch](/posts/code/go/流程控制/if.go)

- 主要就是为了多个条件判断会更加的简洁
- case 后面的不只限制于常数或者常量
  - 支持类型
- 相对于 C 语言来说，每个 case 后面默认添加了 break 语句
- 强制进行下一个 case 的话就使用 fallthrough 语句
- 可以直接比较对应的数值 比如 case 1 :就是默认有 case == 1 :
- 但是 case 也支持使用表达式判断一个范围 case a <= 100 :
- 一个 case 支持多个语句判断 case a==1,a==2:
  - 这个等价于 case a == 1:
  - fallthrough
  - case a == 2:
- fallthrough 的不会判断下一个条件是否成立
- 直接就运行下面 case 的语句了

## select

- select 涉及到通道
- select 随机选择一个 case 运行
- 涉及到并发编程
- 主要的作用就是防止通道阻塞
- make(chan int , 1024)
- case <- a
- 上面是遇到的不懂的语句

# 循环语句

## for

- for 一般有三个元素构成
- 初始化; 判断条件 ; 后置语句
- 其中判断条件必不可少
- 后置语句会在其内部的语句运行完的时候运行

```go
for i := 0; i < count; i++ {
    // ! code here
}
// ! 上面是有完整的

for i<count {
    // ! code here
}

```

## range 子语句

- 这个有点类似于 Python 中的 range(1,10)
- 作用类似于迭代器
- range 的对象可以是 字典 ，切片 ， 字符串
- 然后一般会有两个返回值，然后可以参考变量那个章节获取感兴趣的数值
- 没有要的返回值就直接在遍历结束后停止

---

| range 遍历的对象 | 第一个数值 | 第二个数值                   |
| ---------------- | ---------- | ---------------------------- |
| string           | index      | str[index]，返回为 rune 类型 |
| array/slice      | index      | str[index]                   |
| map              | key        | m[key]                       |
| channel          | element    | none                         |

---

# [defer](/posts/code/go/流程控制/defer.go)

- 这个是 go 特殊的控制语句
- 作用就是延时调用语句
- defer 只可以出现在函数内部
- defer 只有 defer 语句都全部执行，对应的函数才算结束
- 只有所有的 defer 执行完毕，才会执行 return
  - 回收资源
  - 清理工作

# break

- [基本使用](/posts/code/go/流程控制/break基本使用.go)
- [嵌套循环](/posts/code/go/流程控制/嵌套循环.go)

# continue

- continue 只可以用于 for 循环
- 用来跳转到指定代码块继续执行任务
- [基本用法](/posts/code/go/流程控制/continue基本用法.go)

# goto

- 对，，就是简单的跳到那个地方，没有什么智能的
- 然后用不好还可能出现死循环
