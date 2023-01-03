---
date: 2021-03-19 19:37:07
updated: 2021-03-19 19:37:07
title: Parted 命令使用手册
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 前置说明

- 在学习 parted 之前你至少应该知道什么是
  - 分区表
  - 分区
  - EFI
  - FLAG
  - 文件系统
- [原文链接](https://www.gnu.org/software/parted/manual/parted.html)

# parted 基本介绍

- GNU Parted 是用于创建和操作分区表的程序
- 这个文档假设读者有着 分区 以及 文件系统 的基础知识
- parted 设计目的是最小化的减少数据损失
- 比如提供诸如断电导致的数据丢失，parted 提供了很多数据安全监测工具
- 但是保险起见还是在使用 parted 之前对你的磁盘进行备份

# 基本知识介绍

## 对储存器的基本介绍

- 分区就是把一个存储设备划分为本地块的过程
- 以便于组织多文件系统以及分配操作系统
- 一个存储设备就好比是一个字节序列
- 它从 0 开始然后不断增长知道设备的最大值
- 字节通常一次读写一个扇区，而不是单独读写
- 每个扇区包含固定的字节数，字节数由设备决定。

```vim
+------------------------------------------------------------+
|            storage device with no partitions               |
+------------------------------------------------------------+
0 start                                                    end
```

## 对分区的基本介绍

- 为了存储多文件系统，一个存储设备可以被划分成多个分区
- 每个分区都可以看作是一个区域，其中包含一个真正的文件系统。
- 为了显示这些分区在设备上的位置，在开始处会写一个小表，如下图中的 PT 所示
- 这个表称为分区表或 disklab，它还存储每个分区的类型和一些标志。

```vim
+--+---------------+----------------+------------------------+
|PT|  Partition 1  |  Partition 2   |  Partition 3           |
+--+---------------+----------------+------------------------+
0 start                                                    end
```

## 模式介绍

- parted 有两种模式，不管那种模式都要以 # parted device 的格式为开头
  - 交互式
  - 命令行模式
- 其中 device 是要编辑的硬盘设备 （如果您很懒并且省略了 DEVICE 参数，Parted 将尝试猜测您想要哪个设备。）

## 一些说明

- 表示分区位置的数字可以是整数或小数。
- 后缀选择单位，可以是单位中描述的单位之一，CHS 和 compact 除外。
- Parted 将为您指定的位置计算合理范围（例如，当您以“g”指定位置时，范围为+/-500 MB）。
- 使用扇区单位“s”指定确切位置。在 parted-2.4 和更新版本中，IEC 二进制单位如“MiB”、“GiB”、“TiB”等也指定了确切的位置
- 如果您没有为命令提供参数，Parted 将提示您输入参数。例如：
  - (parted) mklabel
  - New disk label type? gpt
- Parted 总是在做一些潜在危险的事情之前警告您

### 命令行模式

- 在命令行模式，一般是一下的格式 # parted /dev/sda mklabel gpt mkpart P1 ext3 1MiB 8MiB
- 选项（比如 --help）只能在命令行模式下指定

### 交互式模式

- 在交互模式下，命令在提示下一次输入一个，并立即修改磁盘。例如：
  - (parted) mklabel gpt
  - (parted) mkpart P1 ext3 1MiB 8MiB
  - 允许使用明确的缩写。例如，您可以键入“p”而不是“print”，键入“u”而不是“units”
  - 命令可以用英语或您的母语（如果您的语言已被翻译）键入。这可能会产生歧义。命令不区分大小写。

# 使用 GNU Parted

## 命令行选项

- 当进入命令行模式下，parted 支持一下语法
  - parted [option] device [command [argument]]
- 不管是交互式还是命令行，命令都是不变的
- 只不过就是命令输入的形式不一样罢了
- GNU Parted 提供了一下命令
  1. • align-check
  2. • disk_set
  3. • help
  4. • mklabel
  5. • mkpart
  6. • name
  7. • print
  8. • quit
  9. • rescue
  10. • resizepart
  11. • rm
  12. • select
  13. • set
  14. • unit

# align-check

```bash
Command: align-check align-type n

Example:
(parted) align-check minimal 1
1 aligned

```

- 确定分区 n 的起始扇区是否符合磁盘所选的对齐标准
- 对齐类型必须是“minimal”、“optimal”或缩写。（最小/最佳）
  - I/O size (minimum/optimal): 512 bytes / 512 bytes
- 在 SHELL 模式下，如果分区不满足对齐要求，则以状态 1 退出，否则（包括在对齐数据不可用的旧内核上），继续处理任何剩余命令
- 如果不使用--script，则打印“N aligned”或“N not aligned”。

- 参考知识
  - [磁盘对齐与 4K 对齐](https://blog.51cto.com/9406836/2125670)
  - [磁盘性能二三事之—4k 对齐](https://developer.aliyun.com/article/169081)

# disk_set

```bash
Command: disk_set flag state

磁盘标志由“磁盘标志”行上的print命令显示。它们也作为机器模式下磁盘信息的最后一个字段输出。

(parted) disk_set pmbr_boot on
```

- 更改磁盘上的标志。标志可以是“开”或“关”。
- 根据您使用的磁盘标签，这些标志中的部分或全部可用：
  - ‘pmbr_boot’
  - （GPT）-此标志启用 GPT 的保护 MBR 分区上的引导标志。

![error_loading](/gallery/2021-03-19-22-19-22.png)

# help

```bash
Command: help [command]

Example:

(parted) help mklabel
Print help for the mklabel command.
```

- 用来查看指定命令的帮助

# mklabel

```bash
Command: mklabel label-type
Example:

(parted) mklabel msdos


```

- 这就是 分区表吧
- label-type
- bsd
- loop
- gpt
- mac
- msdos
- pc98
- sun
- 创建分区表 可以类比 Gparted 的

![error_loading](/gallery/2021-03-19-21-52-27.png)

# mkpart

- 创建新分区，而不在该分区上创建新文件系统。
- 这对于为 Parted 不支持的文件系统（或 LVM 等）创建分区很有用
- 你需要提供对应的文件系统类型，来创建一个分区
- 起点和终点是距磁盘起点的偏移量，即距磁盘起点的“距离”。
- part-type
  - primary
  - extended
  - logical
- fs-type
  文件系统类型？ [ext2]?  
  affs0 amufs2 fat16 linux-swap(v0)
  affs1 amufs3 fat32 linux-swap(v1)
  affs2 amufs4 freebsd-ufs nilfs2
  affs3 amufs5 hfs ntfs
  affs4 apfs1 hfs+ reiserfs
  affs5 apfs2 hfsx sun-ufs
  affs6 asfs hp-ufs swsusp
  affs7 btrfs jfs udf
  amufs ext2 linux-swap xfs
  amufs0 ext3 linux-swap(new) zfs
  amufs1 ext4 linux-swap(old)

- 创建这个分区的时候前提就是创建好分区表
- 也就是上面命令
- 注意使用“---”可防止将下面的“-1s”最后一个扇区指示符解释为无效的命令行选项。

```bash
Command: mkpart [part-type name fs-type] start end

$ parted -s /dev/sdX -- mklabel msdos \
    mkpart primary fat32 64s 4MiB \
    mkpart primary fat32 4MiB -1s
```

- 上面的代码创建了两 tole。第一个是不对齐的，很小，长度小于 4MiB。
- 第二个分区正好从 tole 4MiB 标记开始，一直延伸到设备的末端。
- 下一步通常是在第二 tole 个分区中创建文件系统：
  tole tole- $ mk tolefs.vfat /dev/sdX2
- [解决 parted 分区时 The resulting partition is not properly aligned for best performance 报警](https://blog.51cto.com/xiaosu/1590212)

# name

```bash
Command: name number name

Example:

(parted) name 2 'Secret Documents'
Set the name of partition 2 to ‘Secret Documents’.

```

- 设置分区号的名称（仅限 GPT、Mac、MIPS 和 PC98）。
- 名称可以用引号括起来。根据具体情况，shell 可能还需要用单引号括起
- 根据具体情况，shell 可能还需要用单引号括起来，这样 shell 就不会去掉双引号。
- 这个设置的其实就是分区名字
- 也就是 mkpart 命令要求的第一个参数一样的

![error_loading](/gallery/2021-03-19-23-21-59.png)

# print

```bash
Command: print [number]

Example:

(parted) print
Disk geometry for /dev/hda: 0.000-2445.679 megabytes
Disk label type: msdos
Minor    Start       End     Type      Filesystem  Flags
1          0.031    945.000  primary   fat32       boot, lba
2        945.000   2358.562  primary   ext2
3       2358.562   2445.187  primary   linux-swap
(parted) print 1
Minor: 1
Flags: boot, lba
File System: fat32
Size:            945.000Mb (0%)
Minimum size:     84.361Mb (0%)
Maximum size:   2445.679Mb (100%)
```

- 这里的number就是分区号，显示在第一列的那个数字
- 然后注意要么这个分区是本来就存在或者创建好的
- 要不然就是什么都不显示的

![error_loading](/gallery/2021-03-19-23-27-12.png)

# quit
```bash
Command: quit
Quits Parted.
```

- 只有在Parted退出之后，Linux内核才知道Parted对磁盘所做的更改。
- 但是，键入命令引起的更改可能会在键入命令后立即对磁盘进行更改。
- 但是，操作系统的缓存和磁盘的硬件缓存可能会延迟此操作。




# rescue （营救）

```bash
Command: rescue start end

Example:

(parted) print
Disk geometry for /dev/hdc: 0.000-8063.507 megabytes
Disk label type: msdos
Minor    Start       End     Type      Filesystem  Flags
1          0         8056 primary   ext3
(parted) rm
Partition number? 1
(parted) print
Disk geometry for /dev/hdc: 0.000-8063.507 megabytes
Disk label type: msdos
Minor    Start       End     Type      Filesystem  Flags

OUCH! We deleted our ext3 partition!!! Parted comes to the rescue...

(parted) rescue
Start? 0
End? 8056
Information: A ext3 primary partition was found at 0.031MB ->
8056.030MB.  Do you want to add it to the partition table?
Yes/No/Cancel? y
(parted) print
Disk geometry for /dev/hdc: 0.000-8063.507 megabytes
Disk label type: msdos
Minor    Start       End     Type      Filesystem  Flags
1          0.031   8056.032  primary   ext3

It’s back! :)

```

- 营救一个丢失的分区，这个分区以前大约位于起点和终点之间。
- 如果找到这样的分区，Parted将询问您是否要为其创建分区。
- 例如，如果您使用parted的rm命令意外地删除了分区，这将非常有用。

