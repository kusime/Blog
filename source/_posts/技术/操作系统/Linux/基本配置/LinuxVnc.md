---
title: Linux Vnc and ScreenShot
index_img: /gallery/2021-11-23-00-03-05.png
date: 2021-11-22 23:18:20
updated: 2021-11-22 23:18:20
tags:
  - Linux
categories:
  - Linux
---

# 基本組件安裝

```bash
sudo apt install xclip x11vnc 
# x11vnc 应该是自带的，而且是直接分享本机的大屏幕
# which is better than vncserver
# vncserver configuratino is quit complex and 
# even do not work so just just run x11vnc

# Runing vncserver
x11vnc

```

## 关于如何截图

本来说是用那个`flameshot` 的然后我在准备给这个东西设置快捷键的时候
发现你妈系统自带了可还行。而且还是支持区域选择的东西

![图片描述](/gallery/2021-11-22-23-47-29.png)

By the way . You Can Use `Shutter` for ScreenShot
