---
date: 2021-03-25 14:30:41
updated: 2021-03-25 14:30:41
title: linux fstab 基本介绍
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---



# 基本介绍

- 文件可用于定义磁盘分区，各种其他块设备或远程文件系统应如何装入文件系统。
- 每个文件系统在一个单独的行中描述
- /etc/fstab 文件用于指定在开机时自动挂载的分区、文件系统、远程文件系统或块设备，以及它们的挂载方式。
- 执行 mount -a 操作也可以重新挂载/etc/fstab 中的所有挂载项。

## 与 systemd 的关系

- 这些定义将在引导时动态地转换为系统挂载单元(mount unit)，并在系统管理器的配置重新加载时转换
  - 使用 systemd 系统时，systemd 接管了挂载/etc/fstab 的任务
  - 在系统启动的时候，systemd 会读取/etc/fstab 文件并通过 systemd-fstab-generator 工具将该文件转换为 systemd unit 来执行，从而完成挂载任务

## 被 systemd 托管的好处

- 比如，systemd.mount 可以让那些要求在网络可用时的文件系统在网络已经可用的情况下才去挂载，还可以定义等待网络可用的超时时间，从而避免在开机过程中长时间卡住。
- systemd 可以让某个挂载项自动开始挂载和自动卸载，而不是在开机时挂载后永久挂载在后台。
- 这里可以查找挂载服务单元

![error_loading](/gallery/2021-03-25-14-39-52.png)

![error_loading](/gallery/2021-03-25-14-41-18.png)

# 文件构成

```bash
# <file system> <mount point>   <type>  <options>       <dump>  <pass>
# / was on /dev/nvme0n1p2 during installation
UUID=9df13d69-2935-4d85-be8c-895ef97b3a06 /               ext4    errors=remount-ro 0       1
# /boot/efi was on /dev/nvme0n1p1 during installation
UUID=5564-F813  /boot/efi       vfat    umask=0077      0       1
/swapfile                                 none            swap    sw              0       0

```

| 列             | 参数说明                                                                                                                    |
| -------------- | --------------------------------------------------------------------------------------------------------------------------- |
| file system 列 | 用于标识哪个设备需要被挂载。[可选的格式](#可选的格式)                                                                       |
| mount point    | 指明挂载点                                                                                                                  |
| type           | 指明文件系统的类型 如 ext4                                                                                                  |
| options        | 多个参数使用,隔开 |
| dump           | 是否做 dump 0 关 1 开 |
| pass           | fsck 读取 pass 的数值来决定需要检查的文件系统的检查顺序[^what] |

- 允许的数字是0, 1, 和2。 
  - 根目录应当获得最高的优先权 1, 
  - 其它所有需要被检查的设备设置为 2
  - 0 表示设备不会被 fsck 所检查。


# 可选的格式

- |    参数说明    |              举例               |
  | :------------: | :-----------------------------: |
  | 文件系统 LABEL |            LABEL=EFI            |
  | 文件系统 UUID  | UUID=0a3407de-xxxx-848e92a327a3 |
  | GPT 分区 LABEL |         PARTLABEL=HOME          |
  |    GPT UUID    |    PARTUUID=98a81274-xxxx-0     |

# 参考资料

- [1](https://www.junmajinlong.com/linux/systemd/systemd_fstab/)
- [2](http://www.jinbuguo.com/systemd/systemd.mount.html#)
