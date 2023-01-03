---

date: 2021-03-11  10:30:18
updated: 2021-03-11  10:30:18
title: 投屏到 Linux 上
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---
# 安装必要的软件

<!-- ######## BASH Start########  -->

```bash
sudo snap install scrcpy
sudo apt install adb
#hello world
#incle
```

<!-- ######## BASH End  ########  -->

## 基本命令的说明

### scrcpy [option]

- --always-on-top #保证窗口在最前
- --bit-rate value #比特率
- --display id # 多显示器的显示器 id
- --fullscreen #投屏全屏
- --max-fps value # 最大的帧率
- --serial serial # 多设备投屏下指定程序的投屏设备
