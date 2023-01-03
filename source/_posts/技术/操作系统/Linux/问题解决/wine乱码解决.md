---
date: 2021-02-15T00:29:23
updated: 2021-02-15T00:29:23
title: wine 乱码解决
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

## 解决方案

### new.reg

```reg
REGEDIT4
[HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\FontSubstitutes]
"Arial"="simsun"
"Arial CE,238"="simsun"
"Arial CYR,204"="simsun"
"Arial Greek,161"="simsun"
"Arial TUR,162"="simsun"
"Courier New"="simsun"
"Courier New CE,238"="simsun"
"Courier New CYR,204"="simsun"
"Courier New Greek,161"="simsun"
"Courier New TUR,162"="simsun"
"FixedSys"="simsun"
"Helv"="simsun"
"Helvetica"="simsun"
"MS Sans Serif"="simsun"
"MS Shell Dlg"="simsun"
"MS Shell Dlg 2"="simsun"
"System"="simsun"
"Tahoma"="simsun"
"Times"="simsun"
"Times New Roman CE,238"="simsun"
"Times New Roman CYR,204"="simsun"
"Times New Roman Greek,161"="simsun"
"Times New Roman TUR,162"="simsun"
"Tms Rmn"="simsun"
```

```bash
sudo cp /media/WINDOWS/Windows/Fonts/* ~/.wine/windows/Fronts #复制字体 （宋体就好了）
sudo apt-get install fonts-droid-fallback ttf-wqy-zenhei ttf-wqy-microhei fonts-arphic-ukai fonts-arphic-uming #安装字体
regedit new.reg # 复制上面的东西新建文件
```

## 具体的软件要求

### Key Rewrite

- 汉化字体

### Death Cell

- 不需要函数 dll 的添加
- 不需要汉化字体

### QQ

- 需要汉化字体
- winecfg 添加（msvcp60 riched20 riched32）

## 参考

- [ ubuntu 超级详细的 wine5.0 攻略）](https://blog.csdn.net/cxrshiz7890/article/details/106037534)
- [wine 乱码](https://www.cnblogs.com/xiangzi888/archive/2011/09/21/2184410.html)

## 更新到 wine 6.0

```bash
sudo add-apt-repository 'deb https://dl.winehq.org/wine-builds/ubuntu/ focal main'
sudo apt install --install-recommends winehq-stable
```

## 一些说明

- wine 版本不影响其虚拟机中的内容，就是~/.wine 中的内容
- 因为实在没有办法安装 PR AE 然后自带的剪辑软件菜的死
- 所以我打算之后去学习 Blender
