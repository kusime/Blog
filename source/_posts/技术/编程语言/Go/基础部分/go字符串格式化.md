---
date: 2021-03-16 16:23:26
updated: 2021-03-16 16:23:26
title: go 字符串格式化
index_img: /gallery/2021-08-23-21-59-40.png
tags: 
  - Golang

categories:
  -  Golang

---

- [基本的格式化指令](#基本的格式化指令)
- [关于一些字符串的小提示](#关于一些字符串的小提示)

# 基本的格式化指令

| 格式化指令 | 含 义                                                    |
| ---------- | -------------------------------------------------------- |
| %%         | %字面常量                                                |
| %b         | 二进制整数 int->bin                                      |
| %c         | unicode 字符                                             |
| %d         | 十进制                                                   |
| %o         | 八进制                                                   |
| %x         | 小写十六进制                                             |
| %X         | 大写的十六进制                                           |
| %U         | Unicode 的字符码                                         |
| %s         | 输出原生的 utf-8 表示的字符                              |
| %t         | true / false 的方式输出 bool 数值                        |
| %v         | 默认格式输出，如果方法存在使用 Strings()方法输出自定义值 |
| %T         | 输出值的类型                                             |

# 关于一些字符串的小提示

- go 中可以用 ' （单引号）来创建字符
- 字符串支持切片操作
- 切片是通过字节进行 index 的
- 所以保险的方法就是使用 range 操作符
