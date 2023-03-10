---
title: Openwrt 内网穿透
date: 2021-08-18 14:20:29
updated: 2021-08-18 14:20:29
index_img: /gallery/2021-08-23-20-04-07.png
tags:
  - Openwrt

categories:
  - Openwrt
---

# 基本原理

![图片描述](/gallery/2021-09-03-23-48-34.png)

# FRPS 配置

- 默认frps端口绑定,用来接收客户机(普通手机什么的)

![图片描述](/gallery/2021-09-03-23-50-37.png)

- token 的设定

![图片描述](/gallery/2021-09-03-23-51-50.png)

- 基本控制面板密匙设置

![图片描述](/gallery/2021-09-03-23-52-21.png)


- 然后再对于的云主机上开放
    - 7000 就是frps 允许端口
    - 以及要把内网服务映射出来的端口

# FPRC的设置

- 这里选择的是 Openwrt 的Frpc 的图形化配置
    - 开启基本frpc服务
    - 然后指定对应frps的端口,主机地址
    - 然后再上面设置的token

![图片描述](/gallery/2021-09-03-23-54-08.png)

## 添加基本 穿透 Server

- 点击添加,准备填入基本内容

![图片描述](/gallery/2021-09-03-23-55-26.png)

- 修改基本要穿透的协议
    - 如果是 ssh http ftp 这些东西
    - 然后就选择tcp就好了
    - 如果有些游戏是 udp传输的就是选择udp 穿透协议就好了

![图片描述](/gallery/2021-09-03-23-56-37.png)

- 按照上面提示,填写好对应的数据就好了
    - 对于Proxy-Protocol版本都可以
    - 然后内网主机就是这个Openwrt可以直接访问到的主机
    - 然后端口号就是对于的服务端口号
    - 名字需要唯一
![图片描述](/gallery/2021-09-03-23-57-10.png)

- 运行

![图片描述](/gallery/2021-09-03-23-59-26.png)
