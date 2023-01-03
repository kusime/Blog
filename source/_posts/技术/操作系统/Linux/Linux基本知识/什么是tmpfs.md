---
date: 2021-03-27 09:38:11
updated: 2021-03-27 09:38:11
title: 什么是 tmpfs 文件系统
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 简单理解

- tmpfs 就是建立在内存之上的文件系统
- 也就是说，可以想正常的文件系统（ext4）一样
- 直接对内存进行读写的操作
- 然后如果内存充足的话一般直接写到 tmpfs 文件系统的都是
- 直接写入内存的，当内存不够的时候就还是会写入到 swap 分区里面

# 基本介绍

- tmpfs,临时文件系统，是一种基于内存的文件系统
- tmpfs 可以使用 RAM，但它也可以使用 swap 分区来存储
- 而 tmpfs 是一个文件系统，并不是块设备，只是安装它，就可以使用了
- tmpfs 是最好的基于 RAM 的文件系统。

# 虚拟内存

- Linux 内核的虚拟内存资源同时来源于您的 RAM 和交换分区
- tmpfs 并不是存在于一个底层块设备上面。因为 tmpfs 是直接建立在 VM 之上的，您用一个简单的 mount 命令就可以创建 tmpfs 文件系统

```bash
mount tmpfs /mnt/tmpfs -t tmpfs
```

# 优势

- 动态文件系统大小
- tmpfs 最初会只有很小的空间，但随着文件的复制和创建，tmpfs 文件系统驱动程序会分配更多的 VM
- 并按照需求动态地增加文件系统的空间
- tmpfs 的另一个主要的好处是它闪电般的速度。因为典型的 tmpfs 文件系统会完全驻留在 RAM 中
- 当更多空闲的 VM 资源可以使用时，这部分 tmpfs 文件系统会被移动到 RAM 中去

# 使用

- 默认情况下，tmpfs 会 mount 到/dev/shm 目录。使用 tmpfs，就是说你可以使用这个目录，这个目录就是 tmpfs，如你写临时文件到此目录，这些文件实际上是在 VM 中。
- 要使用 tmpfs，您要在内核配置时，启用“Virtual memory file system support”。
- 为防止 tmpfs 使用了全部 VM，有时候要限制其大小。要创建一个最大为 32 MB 的 tmpfs 文件系统，键入：

```bash
mount tmpfs /dev/shm -t tmpfs -o size=32m
```
