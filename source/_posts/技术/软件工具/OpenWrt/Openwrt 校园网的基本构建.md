---
date: 2021-08-28 14:20:29
updated: 2021-08-28 14:20:29
title: Openwrt 校园网的基本构建
index_img: /gallery/2021-08-23-20-04-07.png
tags:
  - Openwrt

categories:
  - Openwrt
---

# 基本情景说明

- 学校的校园网认证制度的大幅度升级
  - 弱口令限制
  - 共享设备监测
- 这意味者
  - 需要一个正常账号
  - 再进行转发的时候需要更加的细致
  - 尽量要简化出口 Ip 变化的成本减少
  - 断电重启的时候,需要自动化配置上网

## 我所准备的硬件

- 树莓派 4B
- 五口全千兆交换机
- 千兆 Usb 有线网卡
- 机械硬盘
- ~~Tp link ax3200~~
- AC2100 Redmi
  - 后面我还是选择了这个路由,理由很简单
  - 有了老毛子之后能支持 160Mhz 的频宽
  - 硬件级别加速 NAT
  - 能学习新的路由固件的使用
  - 相比 TP 可操作性没有这么高
  - 我结合老毛子搭建了双跳板 DNS

## 我的基本构架图

![基本构架图](/gallery/2021-08-28-23-11-44.png)

{% note warning %}
![9 月 1 日更新](/gallery/2021-09-01-09-48-10.png)
{% endnote %}

- 这里只是对网关和子网结构做了一些小小的调整

# 目前使用到的网络技术

- 基本子网配置
- SNAT (MASQUERADE) 解决网络设备限制 ... SmartDns
- DoH (Dns Over Https) 可能环节系统监测共享网络的行为
- Mac 伪装 用于新 IP 的获取以及对系统监测的回避
- WebDav 用于提供局域网盘
- BBR Tcp 网络阻塞算法,(因为 Wifi 可以被视作为有限客户端,所以不会减少 wifi 的吞吐)
- Qos 流量整形
- 路由基本知识

# 需要编写的脚本

- 那么就是简单的 Post 表单发送
  - 我尝试了 `MentoHUST` 进行捷锐校园网的登陆,但是失败了 我查了相关的资料
    - 这里其他学校可能开放了 802.1x 的一个认证方式
    - 但是我们学校还没有开这个,估计也是知道有这个软件然后做了限制
    - 所以说我自己就编写了个脚本
  - 所以这里还是使用 Python 进行一个登录的请求
    - 28 后面几天我又开发了新的功能
    - 目前脚本可以实现自动登录
    - 30 分钟自动尝试更换 MAC (有失败的可能,但是比较少,但是也差不多了)
    - 自主下线,自主取消注销更换 MAC 之前的注册
  - 然后把这个脚本作为 `crontab` 的定时任务
    - 这里结合脚本的状态监测功能可以实现通常情况下的不报错
    - 单次调用能正常退出的效果
  - ~~或者还可以多编写一些代码进行常驻,进行网口状态的监测等等~~
    - 可以写,但是目前没得时间了...

## ps

- 下面描述的我都实现了

  - ~~我现在的路由器还没有到,所以可能对于路由器子网的配置需要解决以下问题~~
  - ~~首先是能否实现 2 个局域网子网进行通信~~
  - ~~能否保证正常的 Nat 行为~~

- 这里我查了相关的资料,说是因为断电会导致这个原因,应该是自我保护还是什么的
  - 好像我现在这个 Sd 卡有点问题,有时候就直接 `readonly filesystem了`然后就要去 Linux 进行 Check 很麻烦
  - 不过现在概率还是有点小,要是实在不行我就去买一个好一点的卡..
  - 然后重新刷入新的 OpenWrt 然后把现在的备份还原就好了..
  - 要是实在不行就按照现在博客再去重新配置一便...操
  - 但是再怎么说树莓派的失败成本要比 AC2100 的失败成本小的多...

![备份还原](/gallery/2021-09-01-10-00-29.png)

# 我知道但是我还不会的用途

- Docker 一个虚拟化应用平台,我树莓也有 4Gb 的运存,然后 16Gb 的 SD 卡,或者我自己再去买个 U 盘挂上去做 overlay

  - 这样就意味着我可以提供丰富的局域网服务内容
  - 我会拥有基本的 Docker Os (Ubuntu , ArchLinux) 等等
  - 我会有一个安全的内外环境进行一些网络实验,服务器配置(Nginx 学习)
  - 同时还不会直接对我的电脑,云服务器 产生潜在的可能性质

- 结合智能家具,智能插座等等,进行远程的宿舍控制(虽然有点黑人小哥的行为 hhhh)

  - 学习一般都去图书馆啊(#`O′)

- 学习网工的时候,有着一个能够直接实际操作的实验环境,知识点和实际操作的结合
  - 总之 openwrt 我也不能说是完全精通了
  - 只能说,我按照基本的知道和我之前的基础打造了现有的环境
  - 但是随着我知识的积累我会对每个配置更加了解,便于直接提高的我的技术水平

# OpenWrt 基本配置(不包含DNS,DNS单独一个配置)

- SNAT 配置

```bash
iptables -t nat -A POSTROUTING -s 192.168.1.0/24 -o eth1 -j MASQUERADE
```

- eth1 接口配置

![eth1接口配置](/gallery/2021-09-01-10-05-47.png)

- eth0 接口配置
  - 下面部分记得开 DHCP 就好

![eth0 接口配置](/gallery/2021-09-01-10-06-52.png)

- 防火墙域的设置

![防火墙域的设置](/gallery/2021-09-01-10-07-51.png)

- 静态路由

![静态路由](/gallery/2021-09-01-10-09-38.png)

- 静态DHCP

![静态DHCP](/gallery/2021-09-01-10-10-39.png)


- 静态的硬件配置


```bash
# /etc/config/network
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
	option ipaddr '192.168.1.1'
	option netmask '255.255.255.0'
	option ip6assign '60'
	option auto '1'

config interface 'Lan_Out'
	option proto 'dhcp'
	option ifname 'eth1'
	option force_link '1'
	option hostname 'Xiaomi_3'
	option macaddr '32:60:1A:1B:91:58'

config route 'default'
	option interface 'Lan_Out'
	option target '0.0.0.0/0'
	option table '0'
	option gateway '10.3.255.254'
```

- 磁盘挂载
  - 可以参考 OpenWrt分类下面的磁盘挂载文章

![磁盘挂载](/gallery/2021-09-01-10-17-49.png)

- WebDav
  - 我反正现在是用不来Samba...

![WebDAV](/gallery/2021-09-01-10-18-31.png)

- 网络加速

![网络加速](/gallery/2021-09-01-10-19-53.png)

- 流量整形 SQM Qos

![SQM Qos](/gallery/2021-09-01-10-20-27.png)

