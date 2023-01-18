---
title: Openwrt 更多的软件源以及问题修复
index_img: /gallery/2021-08-23-20-04-07.png
date: 2023-01-18 14:42:49
tags:
  - Openwrt
categories:
  - Openwrt
---

#  Openwrt 启动问题修复

首先是 UEFI 启动,之后有时间去搞一个网盘去...这里需要配置TF卡里面的 cmdline , 把那个设备改名字, 如果是TF卡的二分区 那么就是mmcblkp2 这样样子,但是我这边是插入的U盘所以直接找 /dev/sda1 , 或者 /dev/sda2 ,这个看系统安装在哪个分区,因为是之前搞的了,我不记得了,而且还有一堆关键文件这样子hhhh

然后还有一个关键点就是 /etc/config/network 这里. 在这个地方可以修改网络.

```vim
config interface 'loopback'
        option ifname 'lo'
        option proto 'static'
        option ipaddr '127.0.0.1'
        option netmask '255.0.0.0'

config globals 'globals'
        option ula_prefix 'fd08:54e2:ee22::/48'

config interface 'lan'
        option ifname 'eth0'
        option proto 'static'
        option ipaddr '192.168.2.31'
        option netmask '255.255.255.0'
        option ip6assign '60'
        option gateway '192.168.2.2'

config interface 'Lan_Out'
        option proto 'dhcp'
        option ifname 'eth1'
        option force_link '1'
        option hostname 'Iqoo_63'
        option macaddr '63:3D:01:6C:72:2A'

config route 'default'
        option interface 'Lan_Out'
        option target '0.0.0.0/0'
        option table '0'
        option gateway '10.3.255.254'
```


## 卡在luci文件下载

ssh 链接上去运行
```bash
firstboot # 不会修改你的ROOT密码
reboot
```
然后就可以进入登录界面了.


## AdGuard configuration

然后记得修改

```vim
 cat /usr/bin/AdGuardHome.yaml
bind_host: 0.0.0.0
bind_port: 4044
beta_bind_port: 0
users:
- name: Kusime
```

保证DNS服务的正常运行


## 添加更多的软件源

这样子就可以添加很多包了!

![修改luci成为package](/gallery/2023-01-18-14-50-44.png)


```bash
opkg update
opkg install git 
```


## iptablet configuration

```vim
iptables -t nat -A POSTROUTING -s 192.168.2.2/30 -o eth1 -j MASQUERADE
iptables -t nat -A POSTROUTING -s 192.168.0.0/22 -o eth1 -j MASQUERADE
```