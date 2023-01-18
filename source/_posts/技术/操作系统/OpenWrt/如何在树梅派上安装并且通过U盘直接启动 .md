---
date: 2021-02-21 00:55:05
updated: 2021-02-21 00:55:05
title: 树梅派上安装 ubuntu 通过 U 盘直接启动

index_img: /gallery/2021-08-23-20-04-07.png
tags:
  - Openwrt

categories:
  - Openwrt

---


## 如何在树梅派上安装 Ubuntu 20.10 并且通过 U 盘直接启动

### 前期准备

1.Ubuntu Desktop 20.10
2.64GB U disk
3.rpi-imager
4.EEPROM
5.rp firm.zip

### 1.更新树莓派的 EEPROM

#### 下载好最新的 EEPROM

```bash
cd ~/Desktop
mkdir rp4
cd rp4
wget https://github.com/raspberrypi/rpi-eeprom/releases/download/v2020.09.03-138a1/rpi-boot-eeprom-recovery-2020-09-03-vl805-000138a1.zip
mkdir EEPROM
cd  EEPROM
unzip ../rpi-boot-eeprom-recovery-2020-09-03-vl805-000138a1.zip
```

#### 复制到 TF 卡里面

```bash
cp * [ TF localtion]
```

#### 外部操作

- 拔出读卡器
- 树梅派关机
- 插入 SD 卡
- 树梅派通电
- 等待树梅派的绿色 LED 灯快速闪烁

---

### 2.使用 Rpi-Imater 来写入系统到 U 盘

#### 下载系统镜像

在此之前需要下载好 Uubntu Desktop 20.10 For ram 的镜像  
你可以通过下面的代码来进行下载

```bash
wget https://cdimage.ubuntu.com/releases/20.10/release/ubuntu-20.10-preinstalled-desktop-arm64+raspi.img.xz?_ga=2.131016527.1704084536.1612148758-346966349.1612148758
```

注意！这个下载来的文件是 IMG.XZ 文件，这个文件是不用解压的，直接可以使用 Imager 软件进行烧录

#### 安装，启动 Imager

```bash
sudo apt update
sudo apt install rpi-imager
sudo rpi-imager
```

如果上面的失败了

```bash
wget https://github.com/raspberrypi/rpi-imager/releases/download/v1.5/rpi-imager_1.5_amd64.deb
sudo chmod 755 rpi-imager_1.5_amd64.deb
sudo rpi-imager
```

#### 在 Imager 软件的操作

- 格式化 U 盘
- 把下载的 img.xz 文件烧录到 U 盘
- 烧录完成之后拔出 U 盘然后再插入
- 打开 system-boot 卷 【这个 U 盘的一个分区】

### 给刷入的 Ubuntu 安最新的 Firmware

#### 下载，更新，最新固件

```bash
cd ~/Desktop/rp4
mkdir firmware
cd firmware
wget https://github.com/raspberrypi/firmware/archive/1.20210111.master.zip
unzip firmware-1.20210111.master.zip
cd boot
cp * [system-boot localtion]
```

## 完成

- 在树莓派插入 U 盘
- 直接开机，等待一会就开机了
- 推荐使用 HDIM 连接一个显示器
