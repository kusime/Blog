---
date: 2021-03-25 15:24:06
updated: 2021-03-25 15:24:06
title: linux mount 基本介绍
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 基本介绍

- 名称
  mount - 挂载文件系统
- 此命令的标准格式是
  - mount -t type device dir
  - 原 dir 里面的 内容/属主/权限 将被屏蔽，直到此设备被卸载。
  - 如果只给出了 dir 或者只给出了 device ，那么将根据 /etc/fstab 的设置进行挂载
  - 为例避免目录和文件系统的混淆
    - 可以使用 --target(表示 dir)
    - --source(表示 device)

```bash

mount [-lhV]

 mount -a [-fFnrsvw] [-t vfstype] [-O optlist]

 mount [-fnrsvw] [-o option[,option]...]  device|dir

 mount [-fnrsvw] [-t vfstype] [-o options] device dir

```

## 关于 device 的说明

- 大多数的 device 是一个类似 /dev/sda1 这样的块设备文件名，但是对于 NFS 来说则可能是 knuth.cwi.nl:/dir 的样子
- 此外还可以通过 LABEL(-L) 或 UUID(-U) 来标识一个块设备(uuid 中必须使用小写字母)。
- 对于使用 GPT 分区格式的硬盘来说，还可以使用 PARTUUID 或 PARTLABEL 来标识。
- 推荐在 /etc/fstab 中使用 LABEL=label 来标记设备
- arch linux 的启动盘也是这么设定的
- LABEL 可读性更好，可移植性也更强
- 注意，mount 将 UUID 视为字符串(而不是 16 进制数)。UUID 中的字母必须小写。

## 自动挂载

- 将会挂载 /etc/fstab 中所有列出的所有设备(明确标记为"noauto"的除外)。
- 如果额外加上 -F 选项，那么将同时并行挂载多个文件系统
- 如果挂载一个 fstab 或 mtab 中的文件系统，那么只要指定设备或挂载点之一就可以了。

# 挂载规则

## /etc/mtab

- 维护一份当前已挂载文件系统的列表。
- 直接调用 mount 将会打印出此列表

## 一些说明

- 同时给出了 device(或 LABEL, UUID, PARTUUID, PARTLABEL) 和 dir ，那么 mount 将不会读取 /etc/fstab 中的设置。例如：
  - mount /dev/foo /dir
- 在 /etc/fstab 已有选项的基础上追加选项(追加到已有选项之后)
  - mount {device|dir }-o options
- /etc/mtab 替换成指向 /proc/mounts 的符号链接，这有助于提升挂载/卸载大量文件系统的速度。
  - 当然，这样做会导致无法使用"user"选项，而且也会丢失一些信息。

## 非 root 用户的挂载

- 通常只有 root 用户可以挂载文件系统，
- 在 fstab 中包含"user"选项，那么所有用户都可以挂载此文件系统。
  - /dev/cdrom /cd iso9660 ro,user,noauto,unhide
- 使用了"user"选项的文件系统，只有挂载该文件系统的用户才可以卸载它。
  - 若想允许任何其他用户卸载，那么可以使用"users"代替"user"来实现。
- "owner"选项与"user"类似，表示仅允许设备的属主挂载它。
- "group"选项与"owner"类似，表示仅允许设备的属组中的用户挂载它。

## 绑定挂载

- 可以将文件系统同时绑定到两个不同的位置：
  - mount --bind olddir newdir
  - fstab 中添加 bind 选项
  - 短格式 /olddir /newdir none bind
- 如果想要递归绑定整个目录树上所有的文件系统
  - mount --rbind olddir newdir
  - 短格式 mount -R olddir newdir
- 使用 --bind/--rbind 绑定挂载文件系统的时候，并不能改变其原有的挂载选项
  - 如果想要改变挂载选项，必须在绑定之后，再使用 remount 选项来修改：
    - mount --bind olddir newdir
      - 第一个命令将'bind'标记记录到 /etc/mtab 文件
    - mount -o remount,ro newdir
      - 第二个命令又会从 /etc/mtab 中读取到'bind'标记。
    - remount 的行为受 /etc/mtab 的控制。
- /etc/mtab 将不会被读取，这就意味着你必须在此情况下明确使用 bind 标记
  - mount --bind olddir newdir
  - mount -o remount,ro,bind olddir newdir
- 注意，remount,ro,bind 将会创建一个只读的挂载点，但是原始文件系统的超级块依然是可写的。也就是说，olddir 是可写的，但是 newdir 却是只读的。

## 移动

- 命令将一个目录移动到另一个地方，而保持文件的物理位置不变
  - mount --move olddir newdir
  - 短格式：mount -M olddir newdir
- 这将导致 olddir 中的内容完全转移到 newdir 中来访问，但是文件的真实物理位置保持不变。

- (1) olddir 必须是一个挂载点
- (2) olddir 不能位于带有"shared"属性的挂载点之下
- 可以使用"findmnt -o TARGET,PROPAGATION /dir"命令查看挂载点 /dir 的属性。

## 共享子树

- 可以为一个挂载点(可以包含子挂载点)设置传播类型标记
  - (shared, private, slave, unbindable)。
- 对于标记的说明

|            |                                                                                        |
| ---------- | -------------------------------------------------------------------------------------- |
| shared     | 表示允许创建镜像，一个镜像内的挂载和卸载操作会被自动传播到所有其他镜像中               |
| private    | 自动继承主挂载点中挂载和卸载操作，但是自身的挂载和卸载操作不会反向传播到主挂载点中。   |
| slave      | 表示既不继承主挂载点中挂载和卸载操作，自身的挂载和卸载操作也不会反向传播到主挂载点中。 |
| unbindable | 表示禁止对该挂载点进行任何绑定(--bind --rbind)操作。                                   |

- 基本说明

```bash
 支持的操作：
   mount --make-shared mountpoint
   mount --make-slave mountpoint
   mount --make-private mountpoint
   mount --make-unbindable mountpoint

 下面的命令表示递归的改变一个挂载点及其下的所有子挂载点的传播类型标记：
   mount --make-rshared mountpoint
   mount --make-rslave mountpoint
   mount --make-rprivate mountpoint
   mount --make-runbindable mountpoint

```

- (1) 每个 mount 命令只能修改一个传播类型标记，也就是说不可以一次指定多个传播类型标记。
- (2) mount 在进行 --make-\* 操作时不会读取 fstab(5) 文件，你必须在命令行上指定所有挂载选项。
- 从 util-linux-2.23 版本开始，可以在 fstab(5) 中的"挂载选项"字段设置传播类型标记：(private, slave, shared, unbindable, rprivate, rslave, rshared, runbindable)

# mount 挂载选项 (于 fstab 中)

## 文件系统无关的挂载选项

- 这部分选项的当前值可以通过 /proc/mounts 查看。而其中一部分的默认值由内核编译时的配置决定。
- 这里的选项与文件系统无关(适用于所有类型的文件系统)，并且都不可用于"rootflags="内核引导参数
- 下面的选项仅能在 /etc/fstab 文件中使用：

| option               | 解释                                                                                                           |
| -------------------- | -------------------------------------------------------------------------------------------------------------- |
| auto[default]/noauto | mount -a 是否自动挂载                                                                                          |
| group                | 允许非 root 用户挂载，如果该用户所属组之一匹配设备的属组                                                       |
| owner                | 该用户是此设备文件的宿主的话                                                                                   |
| user                 | 允许非 root 用户挂载此文件系统，此用户的名字将记入 mtab 中以便于随后再卸载                                     |
| async/sync           | 所有的 I/O 操作都异步/同步进行。这是默认值。                                                                   |
| dirsync              | 对目录的更新操作都同步进行。这将影响下列系统调用 ：“creat, link, unlink, symlink, mkdir, rmdir, mknod, rename" |
| atime/noatime        | (启动/禁止)更新文件与目录的 inode 访问时间                                                                     |
| diratime/nodiratime  | (启动/禁止)每一次访问目录都更新 inode 访问时间                                                                 |
| ro                   | 以只读模式挂载                                                                                                 |
| rw                   | 以读写模式挂载，这是默认值。                                                                                   |
| defaults             | 等价于使用如下默认选项：rw, suid, dev, exec, auto, nouser, async                                               |
| dev/nodev            | (启动/禁止) 使用其中的字符设备和块设备文件                                                                     |
| exec/noexec          | (启动/禁止) 直接执行其中的二进制文件                                                                           |
| suid/nouid           | (启动/禁止) SUID 和 SGID 位生效                                                                                |
| \_netdev             | 表明此文件系统位于网络上。用于防止网络未启用的时候就挂载                                                       |
| nofail               | 即使指定的设备不存在也不报错                                                                                   |
| remount              | [单独解释]                                                                                                     |
| 文件系统选项         | [文件系统选项] 各个文件系统有独有的配置，我单独挑几个出来                                                      |

## 文件系统有关的选项

- 这里好像就是 -t 选项后面的 也就是种类
- 下面的表格做一个说明

|  种类   |  说明   |
| --- | --- |
|    devpts |   devpts 是一个伪文件系统，一般挂载到 /dev/pts 目录，用于获取伪终端。  |
| proc    |  pro是c什么   |
|   tmpfs  |  tmpfs是什么  |
| ext2/3/4/vfat/ntfs    |   NaN  |
|  iso9660   |  ISO9660 是一种标准，描述了用于 CD-ROM 的文件系统结构。   |
| LOOP设备    |   LOOP设备是什么  |




## remount

- 重新挂载一个已经挂载了的文件系统而不修改其挂载点。通常用于更改挂载选项(比如从 ro 变为 rw)。
- remount 对于命令行选项和 fstab 中选项的处理方式和 mount 完全相同，仅在 device 和 dir 都明确指定的情况下，fstab 与 mtab 才会被忽略
  - mount -o remount,rw /dev/foo /dir
- 这个命令表示所有原来的挂载选项都将被替换，并且忽略 fstab 中预设的选项(除了 loop=)。
  - 这个命令表示将 fstab(或 mtab)中预设的选项与此处的命令行选项合并后得到新的挂载选项。
  - mount -o remount,rw /dir
