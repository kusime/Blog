---

date: 2021-02-23 16:09:50
updated: 2021-02-23 16:09:50
title: 树莓派刷入 Openwrt

index_img: /gallery/2021-08-23-20-04-07.png
tags:
  - Openwrt

categories:
  - Openwrt

---


## 准备必要的文件

- 资源地址
  - [openwrt 项目地址]("https://hub.fastgit.org//SuLingGG/OpenWrt-Rpi")
  - [我的树莓固件下载的地址]("https://github.com/SuLingGG/OpenWrt-Rpi/actions/workflows/build-rpi4-lean-openwrt.yml?query=is%3Asuccess")
- 按照对应的操作下载到对应的文件
  - 下载页面基本上是编译好的，所以直接选择最新的那个
  - 越靠前的就新
  - 然后提供下载的网址也不是 github
- [我的下载](https://wetransfer.com/downloads/c511123c8d3a7bbdb2c3a7c7cb5786a320210310231453/cf63c9/grid)
  - 这里看好选择一个基本的镜像就好了

<!-- ######## BASH Start########  -->

```bash
sudo apt install rpi-imager
```

<!-- ######## BASH End  ########  -->

## 刷入过程

- 打开刚刚安装的刷入程序
- 插入 sd 卡然后格式化，刷入

## 修改登入密码

- 没有显示器也没关系
- 直接插入键盘到树莓 π
- 输入 passwd
- 然后输入两次一样的密码
- 然后就可以登入到树莓派了
