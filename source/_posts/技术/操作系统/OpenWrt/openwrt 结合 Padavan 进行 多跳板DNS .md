---
date: 2021-09-01 10:25:08
updated: 2021-09-01 10:25:08
title: openwrt 结合 Padavan 配置多跳板 DNS(DoH)

index_img: /gallery/2021-08-23-20-04-07.png
tags:
  - Openwrt

categories:
  - Openwrt
---

# DNS over HTTPS 基本认识

> `DNS over HTTPS (DoH) is a protocol for performing remote Domain Name System (DNS) resolution via the HTTPS protocol. A goal of the method is to increase user privacy and security by preventing eavesdropping and manipulation of DNS data by man-in-the-middle attacks[1] by using the HTTPS protocol to encrypt the data between the DoH client and the DoH-based DNS resolver. By March 2018, Google and the Mozilla Foundation had started testing versions of DNS over HTTPS.[2][3] In February 2020, Firefox switched to DNS over HTTPS by default for users in the United States.[4]`

# 基本构架图

![基本构架图](/gallery/2021-09-01-10-27-55.png)

- 9.9.9.9 为学校操控的 DNS (转发什么的也好,反正这个就是被他们劫持了)
- eth1 - 10.3.x.x 为我网络出口地址,DNS 配置为学校发下来的 9.9.9.9

![学校分配劫持的DNS给我](/gallery/2021-09-01-10-40-49.png)

- 192.168.1.1 为第一台 SMartDns 服务器,运行端口为 53 标记为 DNS-1
- 66.66.66.1 为第二胎 SmartDns 服务器,运行端口为 53 标记为 DNS-2
- 为了保证他们正常运行的话,需要去转移 dnsmasq 的默认 dns 端口
- 目前这个体系不太需要原生 Dns..原生 dns 在我的理解唯一作用就是去解析 DoH 服务器的地址

# 基本 DNS 流量路径解释

## 如果被劫持

- DNS-1 查询 Doh Dns 地址,学校返回被劫持的地址 123.123.123.123
- 然后 https 检查证书,检查失败,解析失败
- DNS-2 询问 DNS-1 的解析地址,但是因为 DNS-1 解析失败,返回服务器无响应
- DNS-2 向客户端显示解析失败..
- **校园 DNS 劫持失败..**

## 如果正常返回解析路径

- 子网-2 设备机器向 DNS-2 请求解析,DNS-1 设备把请求转化为在子网-1 的 WAN 口
    - 如果当前还没有Doh解析的记录那么就先去询问Doh地址
    - 然后 Dns-1 向 Dns-2 返回非污染的 Doh地址
    - DNS-2 得到了 Doh 的服务器地址,之后就发送https请求出去,相当于正常访问页面
    - DNS-2 直接向Doh服务器 发送加密的dns解析
        - (这里就不经过 DNS-1的解析了,可以理解为DNS-1只是为了获得到非污染的Doh地址)
    - DNS-2 得到Doh服务器解析,返回解析地址给 子网-2 设备
    - 子网-2设备成功得到无污染解析的DNS
- WAN 口设置的 DNS 为 DNS-1,DNS-1收到请求,把请求转化为 Doh 请求
    - DNS-1没有做Doh服务器记录的缓存,用普通的UDP向学校服务器请求解析
    - 服务器返回正常的解析地址,返回地址通过DNS-1 的 https检查
    - 记录下来Doh的解析地址


## 为什么我要做2级转发

- 为了验证我想出来的这个体系能不能正常使用
- DNS-1 为一级DNS 可能发送UDP请求,然后这些可能纯在设备数量监测,DNS污染的可能性
- 这样的架构,只要设备在 子网-2 那么就一定可以保证,在这个子网获得到的地址一定是没被污染的地址

## 这个DNS解析网络的核心思想

- DNS-1 用来直接面对被学校污染的环境
- DNS-2 用于在 DNS-1 提供的保护(过滤的情况下)下提供一个安全的子网解析服务
    - 可以保证子网-2的DNS全部由Doh提供
    - 同时,如果说Doh需要外网的服务,在子网一的环境下也能提供较好的外围TCP连接的速度
    - 也就是说,在子网-2的设备的DNS访问是更加安全的,因为子网-2的DNS环境已经脱离了校园网了..

- 同时正常的语音通话一样的可以正常提供?..还是说UDP服务环境直接走中转代理?
    - **总之,二级子网的好处就是可以很大程度上脱离校园网被污染的环境.**

![通话正常](/gallery/2021-09-01-11-35-21.png)

- 这里指的污染的含义就是
    - 访问网站被劫持
    - 被校园网防火墙拦截
    - 可以绕过IPS的DNS拦截...
    - 可以绕过GFW的DNS污染..

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

# 效果展示

![raw.githubusercontent](/gallery/2021-09-01-11-37-35.png)

![绕过污染](/gallery/2021-09-01-11-40-15.png)

