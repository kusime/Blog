---
date: 2021-03-24 21:05:06
updated: 2021-03-24 21:05:06
title: linux socket 文件简述
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 基本介绍

- 以 ".socket" 为后缀的单元文件,封装了一个用于进程间通信的套接字(socket)或管道(FIFO),以支持基于套接字的启动。
- 这个文件 从属于 systemd 控制单元
- 文件由组成
  - 基本段
  - Socket 段
  - 每个套接字单元 都必须有一个与其匹配的 服务单元以处理该套接字上的接入流量。 匹配的 .service 单元名称默认与对应的 .socket 单元相同,但是也可以通过 Service= 选项(见下文)明确指定
- foo.socket 单元文件内并不隐含 WantedBy=foo.service 或 RequiredBy=linux service 文件概述 为了避免这个问题,可以向 foo.service 服务单元文件中 明确添加一个 Requires=foo.socket 依赖。

# service 参数介绍

| 配置名字                | 解释                                                                                                                                                        |
| ----------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ListenStream=           | 字节流(SOCK_STREAM) [套接字注意事项和格式](#套接字注意事项和格式)                                                                                           |
| ListenDatagram=         | 数据报(SOCK_DGRAM) [套接字注意事项和格式](#套接字注意事项和格式)                                                                                            |
| ListenSequentialPacket= | 顺序包(SOCK_SEQPACKET) [套接字注意事项和格式](#套接字注意事项和格式)                                                                                        |
| ListenFIFO=             | 设置一个要监听的文件系统管道(FIFO),必须使用绝对路径。 该选项的行为与 ListenDatagram=选项 非常相似。                                                         |
| ListenSpecial=          | 设置一个要监听的特殊文件,必须使用绝对路径。 该选项的行为与 ListenFIFO= 选项非常相似。 此选项可用于监听 /dev 中的字符设备文件以及 /proc 与 /sys 中的特殊文件 |
| ListenNetlink=          | 设置一个要监听的 Netlink 套接字。 必须设为一个合理的 AF_NETLINK 名称(例如 audit 或 kobject-uevent)                                                          |
| ListenMessageQueue=     | 设置一个 要监听的 POSIX 消息队列。 必须设为一个合理的 POSIX 消息队列名称(以"/"开头)。 该选项的行为与 ListenFIFO= 选项非常相似                               |
| BindIPv6Only=           | 可设为 default, both, ipv6-only 之一默认值 default 则表示使用 /proc/sys/net/ipv6/bindv6only 中的设置 (默认值等价于 both)。                                  |
| BindToDevice=           | 将套接字绑定到特定的网络接口。 如果设置了此选项，那么仅接收指定的网络接口上的流量,并且会自动创建到该网络接口的 device 单元                                  |
| SocketUser=             | 接受一个 UNIX 用户                                                                                                                                          |
| SocketGroup=            | 接受一个 UNIX 组 名称。 设置 AF_UNIX 套接字文件与 FIFO 管道文件的属主与属组默认未设置这两个选项,表示文件的属主与属组都是 root                               |
| SocketMode=             | 设置创建文件节点时的 默认访问模式,仅用于文件系统上的套接字文件与管道文件。 默认值是 0666                                                                    |
| DirectoryMode=          | 当监听文件系统上的套接字或管道时,将会自动创建所需的父目录。 此选项设置了创建父目录时的 默认访问模式。 默认值是 0755                                         |
| Accept=                 | [参数解释](#accept)                                                                                                                                         |

## Accept 参数解释

- 若设为 yes ,则会为每个进入的连接派生一个服务实例,并且仅将该连接传递给服务实例。
- 若设为 no(默认值),则会将套接字自身传递给匹配的服务单元。并且仅为所有进入的连接派生单独一个服务实例，出于性能考虑,守护进程应该改造为仅使用 Accept=no 的设置

# 套接字注意事项和格式

- 当字节流 SOCK_STREAM (也就是 ListenStream=) 应用于 IP 套接字时|其含义是 TCP 套接字
- 当数据报 SOCK_DGRAM (也就是 ListenDatagram=) 应用于 IP 套接字时|其含义是 UDP 套接字。
- 当其中任意一个套接字上有流量进入时,对应的服务都将被启动,并且所有这些选项列出的套接字都将被传递给对应的服务
- 如果多次使用这些选项,那么当其中任意一个套接字上有流量进入时,对应的服务都将被启动,
- 并且所有这些选项列出的套接字都将被传递给对应的服务,无论这些套接字上是否有流量进入。
- 如果为某个选项指定了一个空字符串,则表示 清空该选项之前设置的所有监听地址。

| 配置名字                                       | 解释                                                                                                                                       |
| ---------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| 如果地址以 "/" 开头                            | 那么将被视为文件系统上的一个 UNIX 套接字(AF_UNIX)。                                                                                        |
| 如果地址以 "@" 开头                            | 那么将被视为抽象名字空间内的一个 UNIX 套接字(AF_UNIX)。                                                                                    |
| 如果地址是一个纯数字                           | 那么将被视为在 IPv6 地址上监听的一个端口号。                                                                                               |
| 如果地址是一个符合"v.w.x.y:z"格式的字符串      | 那么将被视为 在指定的 IPv4 地址"v.w.x.y"的"z"端口上监听。                                                                                  |
| 如果地址是一个符合"[x]:y"格式的字符串          | 那么将被视为在指定的 IPv6 地址"x"的"y"端口上监听。 绑定的端口也可能同时接受 IPv6 与 IPv4 连接。这取决于 BindIPv6Only= 选项(见下文)的设置。 |
| 如果地址是一个符合 "vsock[^1]:x:y"格式的字符串 | 那么将被视为一个 AF_VSOCK(Linux VM socket) 套接字并在指定的 CID(地址) "x" 端口 "y" 上监听                                                  |

[^1]: virtio-vsock 是一个半虚拟化设备|它可以让宿主机内的应用程序与客户机内的应用程序直接使用套接字 API 进行通信。

# 实例查看

下面是 "/usr/lib/systemd/system/ssh.socket" 文件

```vim
[Unit]
Description=OpenBSD Secure Shell server socket # 描述
Before=ssh.service # 在ssh.service 启动知乎
Conflicts=ssh.service #
ConditionPathExists=!/etc/ssh/sshd_not_to_be_run #

[Socket]
ListenStream=22 # 那么将被视为在 IPv6 地址上监听的一个端口号。，然后没有BindIPv6Only 默认就是都监听
Accept=yes # 也就是每个ssh链接都会创建一个单独的ssh.service 实例对它进行服务
# 并且只把这个链接传递给服务实例

[Install]
WantedBy=sockets.target
```
