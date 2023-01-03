---
date:  2021-03-17 21:02:26
updated:  2021-03-17 21:02:26
title:  Pulse Audio 基本设置
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# pulseaudio 运行模式介绍

![error_loading](/gallery/2021-03-17-21-18-46.png)

- 真正播出声音的是 server 端
- 默认 server 不支持 tcp，但是可以手动打开
- 通过 tcp 可以实现音频流的无线传输
- 好像 pulse audio 默认就是 tcp 传输的
- 比如主流的播放器 ，解析音频文件为音频流，然后 pulse client 对音频流进行 tcp 封包
  - 比较准确的说法就是，vlc 解析文件到 pulse client
  - 然后 client 把数据发到 server
  - 但是 client 一般我只关系其服务器在哪里
  - 然后 vlc 会自动的对 client 进行配置链接
  - 也就是说，一般不关心 client
- 然后传输到 server 端，然后解析播放到外界

# pulseaudio 基本配置结构

```bash
[oh-my-zsh] Random theme 'blinks' loaded

kusime@mylinux ~
  % cd /etc/pulse                                                         !7820

kusime@mylinux /etc/pulse
  % ls                                                                    !7821
client.conf  client.conf.d  daemon.conf  default.pa  system.pa

```

- client.conf 对客户端配置
- default.pa 对服务端进行配置

# 打开 pulse audio 的 tcp 监听

- 这个主要应用场景为，在另一个 linux 发送音频到本机
- 然后本机播放另一个 linux 上的数据
- auth-ip-acl=host;netrange/netmask;
  - 这个用来决定那个主机 ip 或者指定网段的主机允许发送数据到本机播放

```bash
sudo vim /etc/pulse/default.pa

--vin 编辑一下文件-----

### Network access (may be configured with paprefs, so leave this commented
### here if you plan to use paprefs)
load-module module-esound-protocol-tcp auth-anonymous=1
load-module module-native-protocol-tcp auth-ip-acl=127.0.0.1;192.168.0.0/16;20.20.1.0/24
#load-module module-zeroconf-publish

```

# 指定 PULSE_SERVER 的地址

- 结合上面的架构来说就是用于指定这些音频数据要发到哪里去
- 最后其实就会在指定的主机上播放数音乐了

```bash
export PULSE_SERVER=tcp:ip_addr:port
```

- 然后在这个 shell session 的话，有需要音频服务的就会去链接这个地址

- **注意的是对应的服务端要结合上个小结配置好**

## 配置默认服务地址

- 如果不进行配置的话，就会默认在本地服务
- 但是可以直接设置音频服务的地址
- 就不用每次都 export 那个命令了
- 操作如下

```bash
sudo vim /etc/pulse/client.conf

-------
; default-sink =
; default-source =
default-server =tcp:ip_addr:port
; default-dbus-server =

```

# 一些说明

- 结合之前的 ssh 端口转发以及 frp
- 就可以实现神奇的事情
- 但是不管整么样都要记住
  1. 真的播出音频的就是 server 端
  2. server 会监听 4713 这个端口
