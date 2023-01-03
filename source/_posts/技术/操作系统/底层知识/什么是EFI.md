---
date: 2021-03-19 16:57:17
updated: 2021-03-19 16:57:17
title: 什么是 EFI
index_img: /gallery/2021-08-23-20-13-43.png

tags:
  - System

categories:
  - System

---

- [EFI 简单介绍](#efi-简单介绍)
- [UEFI 简单介绍](#uefi-简单介绍)
- [扩展知识](#扩展知识)
- [与操作系统的关系](#与操作系统的关系)
- [什么是 efi 文件](#什么是-efi-文件)
- [为什么 UEFI 文件系统格式必须为 FAT](#为什么-uefi-文件系统格式必须为-fat)

# EFI 简单介绍

- EFI 是可扩展固件接口（Extensible Firmware Interface）的缩写
- 是英特尔公司推出的一种在未来的类 PC 的电脑系统中替代 BIOS 的升级方案。
- EFI 在概念上非常类似于一个低阶的操作系统，并且具有操控所有硬件资源的能力
- 它只是硬件和预启动软件间的接口规范
- 当 EFI 所有组件加载完毕时，系统可以开启一个类似于操作系统 Shell 的命令解释环境，在这里，用户可以调入执行任何 EFI 应用程序
- 对于 EFI 应用程序的功能并没有任何限制，任何人都可以编写这类软件

# UEFI 简单介绍

- 这个没那么多事,efi 的叫法是英特尔当初自己所属的叫法。UEFI 是它捐出去后的叫法。
- EFI 也不再属于英特尔，而是属于了 Unified EFI Form 的国际组织
- EFI 在 2.0 后也遂改称为 UEFI，UEFI，其中的 EFI 和原来是一个意思
- 与前身 EFI 相比，UEFI 主要有以下改进
  - 到了 UEFI，则是拥有了完整的图形驱动，无论是 PS/2 还是 USB 键盘和鼠标，UEFI 一律是支持的
  - 其次，UEFI 具有一个独特的功能，安全启动，而 EFI 是没有安全启动的
- UEFI 用了 GPT 分区表
- 硬盘上的 ESP 分区是用来储存 bootmgfw.efi 等 EFI 文件的

  ![error_loading](/gallery/2021-03-19-17-09-45.png)

# 扩展知识

- UEFI 是用模块化，C 语言风格的参数堆栈传递方式，动态链接的形式构建的系统
- 组成
  - Pre-EFI 初始化模块（PEI）
  - UEFI 驱动程序执行环境（DXE）
  - UEFI 驱动程序（UEFI driver）
  - 兼容性支持模块（CSM）
  - UEFI 高层应用（UEFI Application）
  - GUID 磁盘分区表
  - 系统管理模式（SMM）

# 与操作系统的关系

- UEFI 在概念上类似于一个低阶的操作系统，并且具有操控所有硬件资源的能力
- UEFI 环境下不提供中断的机制，也就是说每个 UEFI 驱动程序必须用轮询的方式来检查硬件状态，并且需要以解释的方式运行
- UEFI 应用程序（UEFI Application）和 UEFI 驱动程序（UEFI driver）是 PE 格式的.efi 文件，可用 C 语言编写。
- UEFI 固件区分架构，在 UEFI 引导模式下，通常只能运行特定架构的 UEFI 操作系统和特定架构的 EFI 应用程序
- 比如，采用 64 位 UEFI 固件的 PC，在 UEFI 引导模式下只能运行 64 位操作系统引导程序；
- 既可以运行 16 位的操作系统（如 DOS），也可以运行 32 位操作系统和 64 位操作系统。

# [什么是 efi 文件](https://blog.csdn.net/Pedroa/article/details/53842115)

- 无论是从功能划分还是各个文件划分，我们都可以看到这种模块化思想
- 不同的人独立开发自己的模块，通过统一的接口添加整合到一个 firmware file 里
- efi 文件就是一个小的 module ，是一种可执行文件类型
- efi 文件作为可以被动态载入然后被执行
- EFI 里的是启动文件。用于 GPT 分区

![error_loading](/gallery/2021-03-19-17-24-28.png)

# [为什么 UEFI 文件系统格式必须为 FAT](https://zhuanlan.zhihu.com/p/25992179)

- UEFI 规定 UEFI 系统分区 EFI system partition (ESP)采用 FAT32 格式，同时支持 FAT12/FAT16 作为移动介质的文件系统
- UEFI 的 FAT 分区和普通的 FAT 分区有个比较小的区别，即 OSType 不同。ESP 在 GPT（不再本文范围内）分区表中有个特殊的 GUID:
- ESP 对 UEFI 启动很重要，UEFI 的操作系统引导程序是以后缀名为.efi 的文件存放在 ESP 分区中的
- 启动操作系统本质上就是运行 ESP 分区内的 app：bootloader 而已
- UEFI 分区内仅仅存储各种 bootloader 和 BIOS 错误恢复程序，操作系统和数据本身是不放在这里的，完全没有问题。


# 什么是ESP

- ESP (EFI system partition)就是 EFI 分区的意思
- EFI支持的是GPT分区表（https://en.wikipedia.org/wiki/GUID_Partition_Table）如果是GPT分区表，那么ESP几乎是必须的，否则EFI会不认。
- ESP 分区是 EFI 启动需要的。EFI 启动对应的分区表类型我记得叫做 GPT 分区表。最多支持 256 个分区。
- 非 EFI 引导就不需要 ESP 分区了。
- EFI仅仅是推荐划分出专门的ESP分区。
- 主板会通过枚举的方式访问每一个能够访问到的分区中的EFI/BOOT文件夹。
- 由于大部分固件只能读取FAT32的内容，所以也就造成了“只能读取ESP”的错觉。

# 什么是安全启动

- 就是验证固件，硬件的签名
- 如果失败就不启动
- 一般刷机都要关掉这个的