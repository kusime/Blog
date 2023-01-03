---
date: 2021-02-23  19:55:56
updated: 2021-02-23  19:55:56
title: 安卓运行 ubuntu
index_img: /gallery/2021-08-23-19-56-51.png
tags:
  - Android

categories:
  - Android

---

- [说明](#说明)
- [文件准备](#文件准备)
  - [必要的文件安装](#必要的文件安装)
- [电脑制作系统镜像](#电脑制作系统镜像)
- [运行制作的系统镜像（包含网络映射）](#运行制作的系统镜像包含网络映射)

## 说明

- 因为手机制作安装镜像极其的慢
- 而且 termux 使用不了 KMV 技术
- 所以我的思路就是在电脑上制作好了启动镜像到手机上运行

## 文件准备

- 电脑（Linux）
- ubuntu-server 镜像
- termux
- qemu

### 必要的文件安装

- 电脑
<!-- ######## BASH Start########  -->

```bash
sudo apt install qemu qemu-kvm
```

<!-- ######## BASH End  ########  -->

- termux （手机）
<!-- ######## BASH Start########  -->

```bash
sed -i 's@^\(deb.*stable main\)$@#\1\ndeb https://mirrors.tuna.tsinghua.edu.cn/termux/termux-packages-24 stable main@' $PREFIX/etc/apt/sources.list
sed -i 's@^\(deb.*games stable\)$@#\1\ndeb https://mirrors.tuna.tsinghua.edu.cn/termux/game-packages-24 games stable@' $PREFIX/etc/apt/sources.list.d/game.list
sed -i 's@^\(deb.*science stable\)$@#\1\ndeb https://mirrors.tuna.tsinghua.edu.cn/termux/science-packages-24 science stable@' $PREFIX/etc/apt/sources.list.d/science.list
apt update && apt upgrade

pkg install x11-repo unstable-repo
pkg install qemu-utils qemu-system-x86_64
```

<!-- ######## BASH End  ########  -->

## 电脑制作系统镜像

<!-- ######## BASH Start########  -->

```bash
cp where/ubuntu-20.04.2-live-server-amd64.iso .
qemu-img create -f qcow2 disk.img 10g
qemu-system-x86_64 -m 1024 -boot d -enable-kvm -smp 8 -net nic -net user -hda disk.img -cdrom ubuntu-20.04.2-live-server-amd64.iso
```

<!-- ######## BASH End  ########  -->

- 复制 disk.img 到手机

## 运行制作的系统镜像（包含网络映射）

<!-- ######## BASH Start########  -->

- 手机运行

```bash
qemu-system-x86_64 -hda disk.img -boot d -m 1024 -smp 8 -netdev user,id=nde1,hostfwd=tcp::2222-:22 -device e1000,netdev=nde1,id=d-net1 -nographic --vnc :1
```

<!-- ######## BASH End  ########  -->

- 电脑运行
<!-- ######## BASH Start########  -->

```bash
qemu-system-x86_64 -hda disk.img -boot d -enable-kvm -m 1024 -smp 8 -netdev user,id=nde1,hostfwd=tcp::2222-:22 -device e1000,netdev=nde1,id=d-net1
```

<!-- ######## BASH End  ########  -->
