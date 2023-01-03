---

date: 2021-02-16  10:30:18
updated: 2021-02-16  10:30:18
title: 配置一个 ubuntu
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---
## 基本软件安装

### 清华源

<!-- ######## VIM Start########  -->

```vim
# 默认注释了源码镜像以提高 apt update 速度，如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted universe multiverse
```

<!-- ######## VIM End  ########  -->

### 基本软件安装

<!-- ######## BASH Start########  -->

```bash
sudo apt install -y wine
sudo apt install -y conky
sudo apt install -y gedit
sudo apt install -y net-tools python
sudo apt install -y python-is-python3
sudo apt install terminator vim
sudo apt install  fcitx-config-gtk
sudo apt install pip
sudo apt install aptitude
sudo apt install tget
```

<!-- ######## BASH End  ########  -->

### 用户名改变了

<!-- ######## BASH Start########  -->

```bash
sed -i s/raw_user_name/new-user-name/gc file-name
```

<!-- ######## BASH End  ########  -->

### 安装输入法

<!-- ######## BASH Start########  -->

```bash
firefox "https://ime.sogoucdn.com/dl/index/1612260778/sogoupinyin_2.4.0.3469_amd64.deb?st=pb_RWPuATD8i2plgFGCtDg&e=1613556507&fn=sogoupinyin_2.4.0.3469_amd64.deb"
sudo dpkg -i sogoupinyin_2.4.0.3469_amd64.deb
sudo apt install -f
sudo dpkg -i sogoupinyin_2.4.0.3469_amd64.deb
im-config #选择fix那个
#然后到设置里面选择输入法框架
reboot
fcitx-config-gtk
#然后添加搜狗输入法的拍到第一
```

<!-- ######## BASH End  ########  -->
