---
date: 2021-03-21 17:27:44
updated: 2021-03-21 17:27:44
title: 安装基本 arch linux
index_img: /gallery/2021-08-23-19-28-51.png
tags:
  - Linux

categories:
  - Linux
---

# 注意事项

- 这里的教程可以保证你能正常的安装好 arch linux 并且启动
- 但是因为为了文章的简洁起见
- 这里把 chroot 中的一些步骤省略掉了
- 目的有下
  - 要是没安装成功不就白输入那又臭又长的命令了吗？
  - 就算那个时候没有做其实安装好了之后在做也不急

# 前期环境介绍

- 检查是否对应的磁盘都正常接入主机了
  - 我这里是把 U 盘接入虚拟机然后把系统安装到 U 盘中的
  - 虚拟机的配置可以参考另外一个文章
- 正常的话应该是下面这个样子的

![error_loading](/gallery/2021-03-21-17-34-53.png)

# 基本分区过程

- 这里涉及到 parted 基本操作
- 我的分区方案是
  - EFI 5%
  - ROOT 90%
  - SWAP 5%
- 步骤可以参考下面的图片
  - 这里默认已经创建了分区表
  - 没有分区表可以输入 mktable gpt

![error_loading](/gallery/2021-03-21-17-42-55.png)

# 创建文件系统

- 刚刚那里只是创建了基本的分区一个文件系统标志
- 但是还没有完成
- 下面创建文件系统

![error_loading](/gallery/2021-03-21-17-45-59.png)

# 挂载文件系统

- 挂载点可以随便
- 但是要确保被挂载的文件夹要存在
- 这里方便起见就直接挂载到/mnt 这个目录下了
- 注意这里暂时只挂载 ROOT 分区

# 调整系统镜像源

- 把清华源直接放到最前面

```vim
vim /etc/pacman.d/mirrorlist
```

![error_loading](/gallery/2021-03-21-17-49-44.png)

# 安装最基本系统

- 确保之前的分区都挂载好了，然后根据挂载点修改对应的命令

```vim
pacstrap /mnt base linux linux-firmware net-tools dhcpcd grub sudo efibootmgr vim
```

- 成功的话会是下面这个图

![error_loading](/gallery/2021-03-21-17-54-44.png)

# 挂载 EFI 分区，然后重复一遍安装最基本系统

- 这里是我自己的经验
- 怎么安装随便你的
- 关键是重新安装一遍也不用重新下载那些东西了
- 就是重新生成了一些东西罢了

![error_loading](/gallery/2021-03-21-17-56-15.png)

# 生成 fstab 文件

- 这里就是为了之后启动设备的时候能自动挂载到这些盘
- 这里使用自动脚本
- 但是确保两个分区都挂载好了

```bash
genfstab -U /mnt >>/mnt/etc/fstab
```

![error_loading](/gallery/2021-03-21-18-00-26.png)

# 安装 GRUB 启动器

- 确保之前的操作都完成了
- 然后分区都正确的挂载好了
- 然后切换操作环境

```bash
arch-chroot /mnt
```

![error_loading](/gallery/2021-03-21-18-01-52.png)

## 检查对应的分区是否都正确的挂载了

- 这一步是为了保险
- 如果你自信当我没说

![error_loading](/gallery/2021-03-21-18-04-57.png)

## 初始化启动镜像

- 这一步是为了保险
- 生成也很快的所以做一下没有坏处

```bash
mkinitcpio -P
```

## 运行安装脚本

- 这一步就是关键了
- 那个 bootloader-id 那个地方是可以随便输入的

```bash
grub-install --target=x86_64-efi --root-directory=/boot --bootloader-id=mingboot
```

- 没啥问题的话就是下面的图片

![error_loading](/gallery/2021-03-21-18-08-21.png)

## 调整 grub 文件夹的位置

- 如果你不想开机就见到 grub 命令行的话就照着做吧
- 如果你直接在 boot 文件夹就可以直接看到 grub 文件夹就可以跳过这一步

```bash
cp  -r /boot/boot/grub /boot
```

- 实现结果图片

![error_loading](/gallery/2021-03-21-18-11-28.png)

## 生成 grub.cfg 文件

- 为了保险起见就直接在两个文件夹都生成一份就好了
- 反正也不耗时间

```bash
grub-mkconfig -o /boot/boot/grub/grub.cfg
grub-mkconfig -o /boot/grub/grub.cfg
```

- 最后的效果图

# 结束步骤

- 退出 chroot 环境

```bash
exit
```

- 卸载挂载的磁盘

```bash
umount /dev/sda*
```

![error_loading](/gallery/2021-03-21-18-15-56.png)

- 关机重启

# 查看最后的效果

![error_loading](/gallery/2021-03-21-18-17-33.png)

# 结束语

- 这里是在虚拟机里面实现了把 archlinux 安装在 U 盘中的过程
- 这里也不用写入安装镜像到 U 盘，直接使用 VMware 虚拟的那一套硬件就好了
- 成功启动后配置可以 查看这个文章

# 如果在 boot 选项哪里没有成功启动的话就自己在 bois 新建一个启动项
