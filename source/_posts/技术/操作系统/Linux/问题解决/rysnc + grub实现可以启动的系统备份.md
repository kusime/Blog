---
title: rysnc + grub实现可以启动的系统备份
date: 2022-04-02 21:36:55
tags:
  - Linux
categories:
  - Linux
---

# rysnc + grub实现可以启动的系统备份

这里是实现效果

![root分区为U盘](/gallery/2022-04-02-21-38-32.png)

![root分区为U盘](/gallery/2022-04-02-21-39-42.png)

## 使用到的基础知识

1. grub 安装，配置
2. UEFI 启动的基本原理
3. X11 桌面系统的配置
4. rsync 的使用
5. Linux 各文件夹的作用
6. Linux 挂载点

# 1.rsync 的基本配置

基本用到的参数就是 <font color="#FF0000">**-aP**</font>,`a` 就是归档模式，保留文件的链接，权限，属性等等. `-P` 就是显示文件传输的进度。

---

然后为了保证我们做的是 <font color="#FF0000">** 镜像备份 **</font> ,换句话说就是 我需要备份的系统有些东西不要了，那么我也需要备份里面那些不要的东西被删除。所以我们需要添加 <font color="#FF0000">**\-\-delete**</font>

所以接下来的操作就是指定 `--exclude=` 参数，来指定我们不需要在备份了的系统里面备份的文件，意思就是，Linux的 <font color="#FF0000">**启动，挂载，显示配置，进程管理**</font> 等等都是通过文件配置的，但是很明显，我们的U盘环境不同，所以我们需要保留那些有差异的地方，以保证我们的系统能够正常的被运行，同时不出现问题。

## 指定需要排除的文件夹

```vim

--- 使用的参数解释

-a # archive mode
-P # show Progeress
--delete # make sure dest match source mirror

--- 通常来说可用忽略的东西

--exclude=/media/*  # except media devices
--exclude=/sys/*  # except sys runtime
--exclude=/proc/* # except proc progress
--exclude=/mnt/*  # except mnt devices
--exclude=/tmp/*  # except temp file
--exclude=/dev/*  # except devices block runtime
--exclude=/run/*  # except runtime lock ..
--exclude=/home/<username>/.config/google-chrome/Default/GPUCache/* # except UnusedCache
--exclude=/var/log/* # excep log file to sync

--- 在首次制作启动盘的时候需要注意的一些东西
--exclude=/etc/fstab  # except mountpoint effect backup boot
--exclude=/boot/grub/* # except grub menu effect backup boot
--exclude=/boot/efi/* # except boot efi entry point but no mounting would be not effected
--exclude=/etc/X11/* #X11 will effect portable gui Desktop boot 
```

### 注意事项

关于什么时候排除是有讲究的，这么理解，如果我第一次就忽略的时候，那么就会导致文件确实，可能是文件夹不存在，等等，所以下面单独列出来的东西是需要注意的，如果在之后的成功启动后那么就可用放心的给起添加上去了。

```vim
--exclude=/etc/fstab  # except mountpoint effect backup boot
--exclude=/boot/grub/* # except grub menu effect backup boot
--exclude=/boot/efi/* # except boot efi entry point but no mounting would be not effected
--exclude=/etc/X11/* #X11 will effect portable gui Desktop boot 
```

## 最终命令

<font color="#FF0000">**1.第一次备份**</font>

```vim
sudo rsync -Pa --delete / </target/devices/partition2> --exclude=/media/* --exclude=/sys/* --exclude=/proc/* --exclude=/mnt/* --exclude=/tmp/* --exclude=/dev/*  --exclude=/run/*  --exclude=/home/kusime/.config/google-chrome/Default/GPUCache/* 

# 会复制和主机一样的boot选项，挂载选项，grub启动配置 --- 可能导致系统启动失败
# 会复制和主机一样的 X11 --- 可能导致GUI桌面启动不正常
# 注意 EFI 分区的挂载时机，以及后续是否自动挂载
# 注意 ROOT 分区挂载位置。
```

<font color="#FF0000">**2.后续配置好X11和BOOT后备份**</font>

```vim
sudo rsync -Pa --delete / </target/devices/partition2> --exclude=/media/* --exclude=/sys/* --exclude=/proc/* --exclude=/mnt/* --exclude=/tmp/* --exclude=/dev/*  --exclude=/run/* --exclude=/etc/fstab --exclude=/boot/grub/* --exclude=/boot/efi/* --exclude=/home/kusime/.config/google-chrome/Default/GPUCache/* --exclude=/var/log/* --exclude=/etc/X11/*

# 会排除 X11 , BOOT , EFI ,GRUB , 保证在配置好能够启动的配置后不会再次被原来的主机配置给失效
```



# 配置UEFI启动

## 分区工作

1. 首先做的就是配置好分区

![图片描述](/gallery/2022-04-02-22-11-51.png)

2. 打好1分区的FLAG

![图片描述](/gallery/2022-04-02-22-12-52.png)

3. 记录好分区 UUID

![EFI](/gallery/2022-04-02-22-13-51.png)

![ROOT](/gallery/2022-04-02-22-14-38.png)

## 挂载好EFI分区然后开始安装grub

```vim
sudo apt install grub-efi-amd64 # fix grub x86_64 dependencies

sudo grub-install --target=x86_64-efi --bootloader-id=Kusime --root-directory=/<devices>/<partition2>/<mountpoint>/boot --efi-directory=/<devices>/<partition1>/<mountpoint>

# --bootloader-id= #define bootloader entry point 
# --efi-directory= #defince EFI partition ! Note , the format of EFI partition is /<p1>/EFI/<bl-id>/xxx.efi
# --root-directory= # this prama defince would be <directory>/grub/<some-grub-file-or-floder>
```

上述操作完成之后，应该就可用看见 `/<devices>/<p1>/EFI/` 下面有启动文件，如果少了 `BOOT` 那么就直接从现有电脑里面复制过来
然后以及 `/<devices>/p2/<root>/grub` 会创建

## <可选> 修复每次开机默认进入 GRUB

之前查了说问题出在 `/devices/p1/EFI/<BL-ID>/grubx64.efi` 这个文件上面，黑屏的是 1.7M 的，但是我们需要的是来自于 `/boot/grub/x86_64-efi/grub.efi` 的 0.4M 的文件，要做的就是直接替换就可以。

然后说这一步不是必须的是因为EFI的入口点依赖与机器，也就是说我们创建的 <BL-id>这个入口换机器会不认，就每次进入到那个GRUB命令行，但是没关系，一样的可用启动

## 解决开机进入黑屏

出现黑屏就是说GRUB没能成功找到下一阶段的配置文件在哪里，所以我们就手动指定就好了，所以我们提前把配置文件写好到EFI里面，之后每次进入黑屏加载那个文件就可以了。

```vim
search.fs_uuid <root-uuid> root hd0,gpt2 
set prefix=($root)'/boot/grub'
configfile $prefix/grub.cfg
```

<font color="#FF0000">**root-id 那里就填写ROOT分区的UUID就好**</font>

然后每次到黑屏的时候就输入

```vim
configfile (hd0,gpt1)/efi/EFI/<BL-ID>/grub.cfg
```

然后就可以见到熟悉的GRUB界面了

## 修改/boot/grub/grub.cfg 配置

这里就是承接上面阶段的 grub的下一阶段的配置，从这里开始就要真正加载kernel了。但是我们要做的操作也十分的简单，就是把原来系统的ROOT的UUID 替换成我们U盘的UUID

```vim
sudo sed s/<raw-root-uuid>/<new-root-uuid>/ /boot/grub/grub.cfg
```

## 修改 /etc/fstab 配置

需要把一些非必要的，依赖主机的挂载点给删除掉。最好是只保留 BOOT ROOT 这两挂载点。

```vim
sudo sed s/<raw-BOOT-uuid>/<new-BOOT-uuid>/ /etc/fstab
sudo sed s/<raw-root-uuid>/<new-root-uuid>/ /etc/fstab
sudo vim /etc/fstab
```

# 后续注意事项

## GUI图像桌面启动不了

这里主要就是收到 NVIDIA 驱动的影响，所以我们要做的就是撤销那些依赖与 NVIDIA 配置的东西，所以我们需要做的就是去重命名（备份）。
涉及到的主要文件有
```vim
/etc/X11/xorg.conf
/etc/X11/xorg.conf.d/10-nvidia-drm-outputclass.conf
/etc/X11/nvidia-xorg.conf
```

所以我们就只要做好重命名

```vim
sudo mv /etc/X11/xorg.conf /etc/X11/xorg.conf~
sudo mv /etc/X11/xorg.conf.d/10-nvidia-drm-outputclass.conf /etc/X11/xorg.conf.d/10-nvidia-drm-outputclass.conf~
sudo mv /etc/X11/nvidia-xorg.conf /etc/X11/nvidia-xorg.conf~
```

## 从备份还原

呃呃，还别说，这个我还真的做过了，因为之前手贱直接 `sudo apt autoremove` 然后直接系统包全部直接不见了。。然后就从U盘里面还原。。

首先先说结论
1. <font color="#FF0000">**这个可行，能够做好一次不差的成功还原**</font>
2. <font color="#FF0000">**不能够在需要还原的系统直接挂载U盘还原**</font>
3. <font color="#FF0000">**需要从U盘成功启动进入系统，然后运行还原命令**</font>

## 备份还原

<font color="#FF0000">**0. 开机进入我们的备份系统**</font>

1. 挂载需要还原的分区，记住还原分区挂载点

![图片描述](/gallery/2022-04-02-22-51-57.png)

2. 运行不会影响原来系统的 X11 fstab grub boot 的还原命令

```vim
sudo su # enter root promote
sudo rsync -Pa --delete  / /media/kusime/3c63c5bf-0299-40da-9a8c-5a66eb23ac12  --exclude=/media/* --exclude=/sys/* --exclude=/proc/* --exclude=/mnt/* --exclude=/tmp/* --exclude=/dev/*  --exclude=/run/* --exclude=/etc/fstab --exclude=/boot/grub/* --exclude=/boot/efi/* --exclude=/home/kusime/.config/google-chrome/Default/GPUCache/* --exclude=/var/log/* --exclude=/etc/X11/*

# / is source means from this Udisk -> partition
# /media/kusime/3c63c5bf-0299-40da-9a8c-5a66eb23ac12 is your destination Root partition
```