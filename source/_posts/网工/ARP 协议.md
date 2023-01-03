---
title: ARP-DHCP-SNMP-DNS-HFC
index_img: /gallery/2021-09-02-08-25-15.png
date: 2021-09-06 20:16:35
updated: 2021-09-06 20:16:35
tags:
  - 网工
categories:
  - 网工
---

# ARP 协议

- 基本笔记
  - 数据链路层

![图片描述](/gallery/2021-09-06-20-22-14.png)

- 发送方询问 某个 IP 地址的 MAC 地址的时候
  - 源地址 MAC (以太网层级) : 询问者 MAC
  - 目的地址 : 广播
  - ARP 层级
    - Target MAC : 00x00x00x00...
    - Target IP : 要询问的 IP 地址
- 回答者
  - 源 MAC 地址为 : 回答方 的 MAC 地址
  - 目的 MAC 地址为 询问者 的 MAC 地址

# DHCP

- 基本报文的格式

![图片描述](/gallery/2021-09-06-20-29-05.png)

![图片描述](/gallery/2021-09-06-20-29-15.png)

![图片描述](/gallery/2021-09-06-20-29-24.png)

![图片描述](/gallery/2021-09-06-20-29-32.png)

- 基本工作原理

![图片描述](/gallery/2021-09-06-20-30-15.png)

- 运行端口
  - <font color="#FF0000">**客户 68 **</font>
  - <font color="#FF0000">**服务器 67 **</font>

# SNMP

- 基本介绍

> 简单网络管理协议 SNMP（Simple Network Management Protocol）用于网络设备的管理。网络设备种类多种多样，不同设备厂商提供的管理接口（如命令行接口）各不相同，这使得网络管理变得愈发复杂。为解决这一问题，SNMP 应运而生

- 运行端口号

![图片描述](/gallery/2021-09-06-20-46-26.png)

- 版本区别

![图片描述](/gallery/2021-09-06-20-33-57.png)

- 系统组成
  - NMS（Network Management System）、
  - SNMP Agent、
  - 被管对象 Management object
  - 管理信息库 MIB（Management Information Base）

## NMS

![图片描述](/gallery/2021-09-06-20-35-14.png)

## SNMP Agent

![图片描述](/gallery/2021-09-06-20-37-58.png)

## Managed Object

![图片描述](/gallery/2021-09-06-20-39-28.png)

## MIB

![图片描述](/gallery/2021-09-06-20-40-31.png)

# SNMP 查询

![图片描述](/gallery/2021-09-06-20-41-52.png)

# SNMP 报文

![图片描述](/gallery/2021-09-06-20-43-30.png)

![图片描述](/gallery/2021-09-06-20-43-57.png)

## SNMP Trap

![图片描述](/gallery/2021-09-06-20-45-17.png)

# DNS

- 基本介绍

![图片描述](/gallery/2021-09-06-20-49-36.png)

DNS 协议建立在 UDP 或 TCP 协议之上，默认使用 53 号端口

![图片描述](/gallery/2021-09-06-20-51-53.png)

- 递归

![图片描述](/gallery/2021-09-06-20-52-16.png)

- 迭代

DNS 返回下一个 DNS 服务器地址.然后让它再返回的地址上查询

- DNS 记录的类型

![图片描述](/gallery/2021-09-06-20-53-10.png)

# HFC

![图片描述](/gallery/2021-09-06-20-56-54.png)