---
date: 2021-03-26 14:10:49
updated: 2021-03-26 14:10:49
title: 什么是 proc 文件夹
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 基本介绍

- proc 文件系统是一个伪文件系统，它只存在内存当中，而不占用外存空间。
- 它以文件系统的方式为访问系统内核数据的操作提供接口。
- Linux 内核提供了一种通过 proc 文件系统，在运行时访问内核内部数据结构、改变内核设置的机制
- 用户和应用程序可以通过 proc 得到系统的信息，并可以改变内核的某些参数
  - 由于系统的信息，如进程，是动态改变的，所以用户或应用程序读取 proc 文件时，proc 文件系统是动态从系统内核读出所需信息并提交的。
- 文件或子文件夹，并不是都是在你的系统中存在，这取决于你的内核配置和装载的模块

## 关键目录

- 在 proc 下还有三个很重要的目录
  - net
  - scsi
  - sys
- sys 目录是可写的，可以通过它来访问或修改内核的参数
- 例如，如果系统不支持 scsi，则 scsi 目录不存在。
- 还有的是一些以数字命名的目录，它们是进程目录。
  - 系统中当前运行的每一个进程都有对应的一个目录在 proc 下，以进程的 PID 号为目录名
  - 它们是读取进程信息的接口。而 self 目录则是读取进程本身的信息接口，是一个 link。

## /proc/[pid]/

### cmdline

- /proc/[pid]/cmdline 是一个只读文件，包含进程的完整命令行信息。如果该进程已经被交换出内存或者这个进程是 zombie 进程，则这个文件没有任何内容。该文件以空字符 null 而不是换行符作为结束标志

![error_loading](/gallery/2021-03-26-14-23-00.png)

### comm

- /proc/[pid]/comm 包含进程的命令名。举例如下：

![error_loading](/gallery/2021-03-26-14-24-26.png)

### cwd

- /proc/[pid]/cwd 是进程当前工作目录的符号链接。举例如下：

![error_loading](/gallery/2021-03-26-14-25-29.png)

### environ

- 显示进程的环境变量。举例如下：

![error_loading](/gallery/2021-03-26-14-26-17.png)

### exe

- 实际运行程序的符号链接。举例如下：

![error_loading](/gallery/2021-03-26-14-27-04.png)

### fd

- 包含进程打开文件的情况。举例如下
- 目录中的每一项都是一个符号链接，指向打开的文件，数字则代表文件描述符。

![error_loading](/gallery/2021-03-26-14-28-26.png)

### latency

- 显示哪些代码造成的延时比较大
- echo 1 > /proc/sys/kernel/latencytop

```bash
$ cat /proc/2406/latency

Latency Top version : v0.1
30667 10650491 4891 poll_schedule_timeout do_sys_poll SyS_poll system_call_fastpath 0x7f636573dc1d
8 105 44 futex_wait_queue_me futex_wait do_futex SyS_futex system_call_fastpath 0x7f6365a167bc
```

### maps

- 显示进程的内存区域映射信息。举例如下：
- 其中注意的一点是 [stack:] 是线程的堆栈信息，对应于 /proc/[pid]/task/[tid]/ 路径。

![error_loading](/gallery/2021-03-26-15-12-33.png)

### stack

- 示当前进程的内核调用栈信息，只有内核编译时打开了 CONFIG_STACKTRACE 编译选项，才会生成这个文件

### statm

- /proc/[pid]/statm 显示进程所占用内存大小的统计信息。包含七个值，度量单位是 page(page 大小可通过 getconf PAGESIZE 得到)。举例如下：
-

```bash
$ cat /proc/2406/statm
13720 2617 493 746 0 12007 0
a      b     c  d  e  f    g

a）进程占用的总的内存
b）进程当前时刻占用的物理内存
c）同其它进程共享的内存
d）进程的代码段
e）共享库(从2.6版本起，这个值为0)
f）进程的堆栈
g）dirty pages(从2.6版本起，这个值为0)
```

### status

- /proc/[pid]/status 包含进程的状态信息。其很多内容与 /proc/[pid]/stat 和 /proc/[pid]/statm 相同，但是却是以一种更清晰地方式展现出来

![error_loading](/gallery/2021-03-26-15-16-18.png)

- 关于信号（signal）的信息：SigQ 分为两部分（例如 0/31834），前面表示当前处在队列中的信号（0），后面则表示队列一共可以存储多少信号（31834）
- SigPnd 表示当前线程 pending 的信号
- ShdPnd 则表示整个进程 pending 的信号
- SigBlk、SigIgn 和 SigCgt 分别表示对信号的处理是阻塞，忽略，还是捕获
- 关于 Unix 信号的相关知识，可以参考 Unix: Dealing with signals

### syscall

- 显示当前进程正在执行的系统调用。举例如下：

```bash
$ cat /proc/2406/syscall
202 0xab3730 0x0 0x0 0x0 0x0 0x0 0x7ffff7f6ec68 0x455bb3
```

- 第一个值是系统调用号（202 代表 poll），后面跟着 6 个系统调用的参数值（位于寄存器中）
- 最后两个值依次是堆栈指针和指令计数器的值。
- 如果当前进程虽然阻塞，但阻塞函数并不是系统调用，则系统调用号的值为 -1
- 面只有堆栈指针和指令计数器的值。如果进程没有阻塞，则这个文件只有一个 running 的字符串。
- 内核编译时打开了 CONFIG_HAVE_ARCH_TRACEHOOK 编译选项，才会生成这个文件。

### wchan

- 显示当进程 sleep 时，kernel 当前运行的函数。举例如下

![error_loading](/gallery/2021-03-26-15-21-07.png)

## 子文件或子文件夹

```vim
/proc/buddyinfo 每个内存区中的每个order有多少块可用，和内存碎片问题有关

/proc/cmdline 启动时传递给kernel的参数信息

/proc/cpuinfo cpu的信息

/proc/crypto 内核使用的所有已安装的加密密码及细节

/proc/devices 已经加载的设备并分类


/proc/dma 已注册使用的ISA DMA频道列表

/proc/execdomains Linux内核当前支持的execution domains

/proc/fb 帧缓冲设备列表，包括数量和控制它的驱动

/proc/filesystems 内核当前支持的文件系统类型

/proc/interrupts x86架构中的每个IRQ中断数

/proc/iomem 每个物理设备当前在系统内存中的映射

/proc/ioports 一个设备的输入输出所使用的注册端口范围

/proc/kcore 代表系统的物理内存，存储为核心文件格式，里边显示的是字节数，等于RAM大小加上4kb

/proc/kmsg 记录内核生成的信息，可以通过/sbin/klogd或/bin/dmesg来处理

/proc/loadavg 根据过去一段时间内CPU和IO的状态得出的负载状态，与uptime命令有关

/proc/locks 内核锁住的文件列表

/proc/mdstat 多硬盘，RAID配置信息(md=multiple disks)

/proc/meminfo RAM使用的相关信息

/proc/misc 其他的主要设备(设备号为10)上注册的驱动

/proc/modules 所有加载到内核的模块列表

/proc/mounts 系统中使用的所有挂载

/proc/mtrr 系统使用的Memory Type Range Registers (MTRRs)

/proc/partitions 分区中的块分配信息

/proc/pci 系统中的PCI设备列表

/proc/slabinfo 系统中所有活动的 slab 缓存信息

/proc/stat 所有的CPU活动信息

/proc/sysrq-trigger 使用echo命令来写这个文件的时候，远程root用户可以执行大多数的系统请求关键命令，就好像在本地终端执行一样。要写入这个文件，需要把/proc/sys/kernel/sysrq不能设置为0。这个文件对root也是不可读的

/proc/uptime 系统已经运行了多久

/proc/swaps 交换空间的使用情况

/proc/version Linux内核版本和gcc版本

/proc/bus 系统总线(Bus)信息，例如pci/usb等

/proc/driver 驱动信息

/proc/fs 文件系统信息

/proc/ide ide设备信息

/proc/irq 中断请求设备信息

/proc/net 网卡设备信息

/proc/scsi scsi设备信息

/proc/tty tty设备信息

/proc/net/dev 显示网络适配器及统计信息

/proc/vmstat 虚拟内存统计信息

/proc/vmcore 内核panic时的内存映像

/proc/diskstats 取得磁盘信息

/proc/schedstat kernel调度器的统计信息

/proc/zoneinfo 显示内存空间的统计信息，对分析虚拟内存行为很有用

以下是/proc目录中进程N的信息

/proc/N pid为N的进程信息

/proc/N/cmdline 进程启动命令

/proc/N/cwd 链接到进程当前工作目录

/proc/N/environ 进程环境变量列表

/proc/N/exe 链接到进程的执行命令文件

/proc/N/fd 包含进程相关的所有的文件描述符

/proc/N/maps 与进程相关的内存映射信息

/proc/N/mem 指代进程持有的内存，不可读

/proc/N/root 链接到进程的根目录

/proc/N/stat 进程的状态

/proc/N/statm 进程使用的内存的状态

/proc/N/status 进程状态信息，比stat/statm更具可读性

/proc/self 链接到当前正在运行的进程
```

# 参考

- [1](https://www.linuxprobe.com/linux-proc-pid.html)
- [2](https://blog.csdn.net/zdwzzu2006/article/details/7747977)
