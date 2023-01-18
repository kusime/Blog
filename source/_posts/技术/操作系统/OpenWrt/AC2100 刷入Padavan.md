---
date: 2021-09-01 11:41:59
updated: 2021-09-01 11:41:59
title: AC2100 刷入Padavan和一些基本配置
index_img: /gallery/2021-08-23-20-04-07.png
tags:
  - Openwrt

categories:
  - Openwrt
---


# 刷入第三方固件

## 刷入 breed

- 我28还是29买的新路由,逼站上面的哪个教程不得行了
    - 就是哪个反弹SHELL哪个,用PPPOE拨号破解
    - 首先是他们抓包脚本报错了
    - 然后就是要进行的操作也太多了...

{% note warning %}
[刷 Breed](https://www.chenxie.net/archives/2432.html#Padavan)
{% endnote %}

![教程-1](/gallery/2021-09-01-11-46-01.png)

说明：

-  上面他的教程写错一个地方,就是STOK哪里不要书名号,直接在STOK位置去掉书名号然后填入STOK就好了
-  然后路由器确实会自动重启,然后进入Breed
-  !!!这个方法需要路由器可以正常联网!!!

此方法来自：

– openwrt 官网 https://openwrt.org/toh/xiaomi/xiaomi_redmi_router_ac2100

– 恩山论坛 https://www.right.com.cn/forum/thread-4066963-1-1.html

```bash
http://192.168.31.1/cgi-bin/luci/;stok=<STOK>/api/misystem/set_config_iotdev?bssid=Xiaomi&user_id=longdike&ssid=%0Acd%20%2Ftmp%0Acurl%20-o%20B%20-O%20https%3A%2F%2Fbreed.hackpascal.net%2Fbreed-mt7621-xiaomi-r3g.bin%20-k%0A%5B%20-z%20%22%24(sha256sum%20B%20%7C%20grep%20242d42eb5f5aaa67ddc9c1baf1acdf58d289e3f792adfdd77b589b9dc71eff85)%22%20%5D%20%7C%7C%20mtd%20-r%20write%20B%20Bootloader%0A
```

## 刷入Padavan

![刷入Padavan](/gallery/2021-09-01-11-53-04.png)

- 这一步没有什么坑.然后接下来的事情就是去玩老毛子了~~

# 刷入后的基本配置

{% note warning %}
这里我就是记录一下我自己这套网络环境下的配置,直接照搬你可能上不了网的..
{% endnote %}


## 无线配置

- 我因为设备都在寝室,所以我直接关掉了

![2.4G](/gallery/2021-09-01-11-58-18.png)

- 5G Wlan 基本配置

![5G-160Mhz-c36](/gallery/2021-09-01-11-59-04.png)

- 5G Wlan 高级配置

![4T-4R-MU](/gallery/2021-09-01-12-00-45.png)

- 实际效果
    -  信号理想情况下,1.7 Gbps 的连接速度 (但是需要网卡的支持)
    -  WebDav读取 100Mb/s的速度

## Wan口配置

![Wan](/gallery/2021-09-01-12-02-43.png)

## Lan口配置

![LAN](/gallery/2021-09-01-12-03-11.png)

![Lan-Dhcp](/gallery/2021-09-01-12-03-40.png)

## 防火墙配置

- 这里有好多参数我看不懂,等之后学了再单独做记录..

![Netfilter](/gallery/2021-09-01-12-04-45.png)

## 系统配置

- 工作模式

![工作模式](/gallery/2021-09-01-12-05-57.png)


## DNS-1

- 基本配置

![基本配置](/gallery/2021-09-01-12-08-27.png)

- 上游服务器配置

![上游服务器配置](/gallery/2021-09-01-12-09-25.png)

## DNS-2 

- 基本配置

![SmartDns](/gallery/2021-09-01-12-11-45.png)

![Dnsmasq](/gallery/2021-09-01-12-12-26.png)

- 上游服务器配置

![上游服务器配置](/gallery/2021-09-01-12-13-03.png)
