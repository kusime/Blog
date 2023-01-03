---
date:  2021-02-21 19:10:25
updated:  2021-02-21 19:10:25
title:  命令行配置 wifi
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---



## 相关的命令

### 查看现有网络设备

<!-- ######## CODE Start########  -->

```bash
sudo ifconfig  [-a]
sudo nmcli dev
```

<!-- ######## CODE End  ########  -->

### 控制网卡的开关

<!-- ######## CODE Start########  -->

```bash
sudo ifconfig [interface] [up|down]
```

<!-- ######## CODE End  ########  -->

### 使用命令搜索 wifi

<!-- ######## CODE Start########  -->

```bash
sudo iwlist |grep "ESSID"
sudo nmcli dev wifi
```

<!-- ######## CODE End  ########  -->

### 查看网卡的详细信息

<!-- ######## CODE Start########  -->

```bash
sudo nmcli device show
```

<!-- ######## CODE End  ########  -->

### 使用命令连接 WIFI

<!-- ######## CODE Start########  -->

```bash
sudo nmcli dev wifi [SSID] [PASSWD] #有密码的
sudo nmcli dev wifi [SSID] #没有密码的
```

<!-- ######## CODE End  ########  -->
