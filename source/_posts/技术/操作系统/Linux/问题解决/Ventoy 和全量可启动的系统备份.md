---
title: 制作一个基于 Ventoy 而且可以启动的全量备份系统
index_img: /gallery/2021-08-23-19-18-20.png
date: 2023-01-05 10:08:11
tags:
  - Linux
categories:
  - Linux
---

# 制作一个基于 Ventoy 而且可以启动的全量备份系统

{% note success %}
[准备, 安装 Ventoy](https://www.ventoy.net/cn/)
{% endnote %}

```bash
sudo ./Ventoy_WebUI.sh
```

## Ventoy 分区的基础

{% note danger %}
在 webUI 设置的时候要注意保留足够的大小,Ventoy 的分区格式是

1. ISO 分区,后面我们会用到这个分区来存储 grub 文件
2. FAT32 分区,这个系统只分了 32MB,而且基本用满了,但是不影响,我看了本身 Ventoy 就是用的 Grub2 来启动的.所以我们要做的就是设置好分区的 UUID,以及 fstab 等.
3. BTRFS 分区, 这个是第三个分区,我们用它来存储我们的系统备份.
   {% endnote %}

# 拷贝文件

{% note danger %}
注意: 只能使用 rsync 进行传输,如果直接用 CP 的话,会破坏链接等,导致启动失败...
使用命令:

```bash
sudo rsync -a -vv --progress -r  [SRC] [DEST]
```

{% endnote %}

## 启动到 ubuntu 或者其他 Linux

这里随便复制一个有图像界面的安装 ISO 到第一个分区,然后启动过去

## 使用 RSYNC 进行文件复制

1. 创建 BTRFS sub volume

```bash
cd /your/vendor/partition/3
sudo btrfs subvolume create @
sudo btrfs subvolume create @home
# Manjaro
sudo btrfs subvolume create @cache
sudo btrfs subvolume create @log
```

2. 开始复制文件

```bash
sudo rsync -a -vv --progress -r  /your/system/partition/@/* ./@/
sudo rsync -a -vv --progress -r  /your/system/partition/@home/* ./@home/
# Manjaro
sudo rsync -a -vv --progress -r  /your/system/partition/@cache/* ./@cache/
sudo rsync -a -vv --progress -r  /your/system/partition/@log/* ./@log/
```

![实例](/gallery/2023-01-05-10-23-21.png)

# 修改替换 UUID

{% note danger %}
这一步主要是修改 Linux 启动的一些找盘的参数,这里可以记住原来的 UUID,然后直接`sudo blkid` 查看新的 UUID,这里注意,我用的 Manjaro 系统他系统自动挂载了一些 subVolume,所以我们也要创建好,或者要是不想文件丢失的话一样的向上面传输过来.

{% endnote %}

## /etc/fstab

{% note danger %}
如果有注册了 Swap 分区,那么请注释掉
{% endnote %}

```vim
# /etc/fstab: static file system information.
#
# Use 'blkid' to print the universally unique identifier for a device; this may
# be used with UUID= as a more robust way to name devices that works even if
# disks are added and removed. See fstab(5).
#
# <file system>             <mount point>  <type>  <options>  <dump>  <pass>
#UUID=EC68-82E7                            /boot/efi      vfat    umask=0077 0 2
UUID=e011abbf-bf3d-486d-9c01-713365092dd8 /              btrfs   subvol=/@,defaults,discard=async,ssd 0 0
UUID=e011abbf-bf3d-486d-9c01-713365092dd8 /home          btrfs   subvol=/@home,defaults,discard=async,ssd 0 0
UUID=e011abbf-bf3d-486d-9c01-713365092dd8 /var/cache     btrfs   subvol=/@cache,defaults,discard=async,ssd 0 0
UUID=e011abbf-bf3d-486d-9c01-713365092dd8 /var/log       btrfs   subvol=/@log,defaults,discard=async,ssd 0 0
#UUID=d7b70ec6-23c2-4fa3-8c33-41fadc89f29c swap           swap    defaults,noatime 0 0
tmpfs                                     /tmp           tmpfs   defaults,noatime,mode=1777 0 0
```

## /etc/default/grub

{% note danger %}
这里如果有 `resume=UUID=[UUID]` 的话请删掉
{% endnote %}

```vim
GRUB_DEFAULT=saved
GRUB_TIMEOUT=5
GRUB_TIMEOUT_STYLE=hidden
GRUB_DISTRIBUTOR="Manjaro"
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash apparmor=1 security=apparmor udev.log_priority=3 i915.enable_psr=0 ibt=off"
GRUB_CMDLINE_LINUX=""
```

## /boot/grub/grub.cfg

{% note danger %}
这里注意上述两点,有注册 SWAP 删除 SWAP 注册,然后替换所有 UUID
{% endnote %}

# 和 Ventoy 一起运行

上述操作就完成了所有的前期准备,但是接下来要处理的就是如何让 Ventoy 来帮助启动我们的系统了.这里利用了 Ventoy 提供的一个叫做 ExMenu 的功能.

把 `/boot/grub/grub.cfg` 复制到 第一个分区 的 ventoy 文件夹下(这个要自己创建的). 然后再把名字命名为 ventoy_grub.cfg

![例子](/gallery/2023-01-05-10-28-57.png)

然后启动的时候,直接 `F6` 然后就可以启动了

# 注意事项

注意,启动之后, `update-grub` 只会更新 `/boot/grub/grub.cfg` 这个文件,但是我 Boot 启动的时候是用的 `/ISO_Partition/ventoy/ventoy_grub.cfg` 这个文件所以更新 Grub 也要记得重新复制一下文件...
