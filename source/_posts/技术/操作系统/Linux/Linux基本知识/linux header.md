---
date: 2021-03-18 16:53:26
updated: 2021-03-18 16:53:26
title: linux header 是什么
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 基本介绍

- [KernelHeaders](http://kernelnewbies.org/KernelHeaders)

![error_loading](/gallery/2021-03-18-16-56-58.png)

- 一个内核头包 应该包含 ’头文件‘ ，那些文件是包含在目标机器构建程序所需要的文件，比如 virtualbox 和 VMware 或者 nvidia driver 需要构建的一些模组
- 内核头文件包含 C 的头文件，用于指定 Linux 内核和用户空间库和程序，头文件指定 接口以及实例，以上是构建大多数标准程序以及重新构建 glibc 包.
- A kernel-devel package would contain the complete source code for the linux kernel, which is needed if you wish to recompile the kernel you are currently running.
- 'kernel-devel' package : 这个包提供内核头文件以及必要的 makefiles ，用来构建内核包

---

kernel-header：提供内核的信息，里面也就是内核的头文件以及 Kconfig 和 Makefile，可以看出，kernel-header 有统领内核的作用，同时，也是内核对外的一个接口，当需要向内核提供兼容的功能模块时，势必需要提供内核的信息。所以，在安装驱动时往往也需要 kernel-header。
由此可见，系统所正在运行的内核必须要与 kernel-header 版本一致。由于之前的内核为 title Fedora (2.6.25-14.fc9.i686)，所以没有成功。后来，通过 yum update kernel 更新 kernel 至与 kernel-header 相应的版本。重新后，新内核生效且是默认内核选项。这里网上有仁兄说了，yum 不是神仙，他只会把网络上最新的东西给你更新上，所以 kernel-header 和 kernel 通过 yum 安装，更新的时候务必小心。

---

区别：kernel-devel 包只包含用于内核开发所需的内核头文件以及 Makefile，而 kernel-souce 包含所有内核源代码。
如果仅仅是用你自己编写的模块开发的话，因为只需引用相应的内核头文件，所以只有 devel 包即可，如果你要修改现有的内核源代码并重新编译，那必须是 kernel-souce。
kernel-souce 在 RH 某些版本之后不再附带在发行版中了，必须自己通过 kernel-XXX.src.rpm 做出来。
kernel-devel 是用做内核的一般开发的，比如编写内核模块，原则上，可以不需要内核的原代码。kernel 则是专指内核本身的开发，因此需要内核的原代码。 
关于 kernel source 的有 kernel 和 kernel-devel 两个 rpm，其中 kernel rpm 包含源文件和头文件（就像 2.4 下的 kernel-source rpm），而 kernel-devel 则主要是头文件。

---

版权声明：本文为 CSDN 博主「terry01203」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
[原文链接：](https://blog.csdn.net/u012450329/article/details/54137564)

```vim
我的一些理解
1. 因为c语言的代码不是通用的(Python/golang)，意思就是跨平台的时候需要
专门针对目标机器去重新构建可执行的文件，然后这些头文件就指明这些函数需要哪些类型的数值
然后应用场景就比如vm虚拟机需要构建一些文件加载在内核中，是不是就需要去问问你这个内核
你这边的环境是怎么样的，然后我结合这些文件以及我的c语言源代码去构建

举个简单的例子就是虽然c语言代码写的都是
# include <stdio.h>
但是stdio这个具体的实现在不同机器的实现过程是不一样的，所以就有必要去申明具体的环境是整么样的
```

# 查看内核

```bash
uname -r
```

# 安装内核头

```bash
sudo apt-get install linux-headers-$(uname -r)
```
