---
date: 2021-02-02  15:24:06
updated: 2021-02-02  15:24:06
title:  Linux 路由表
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---


## Linux kernel route table 配置

### 官方的帮助文档

```vim
Usage: route [-nNvee] [-FC] [<AF>]           List kernel routing tables
       route [-v] [-FC] {add|del|flush} ...  Modify routing table for AF.

       route {-h|--help} [<AF>]              Detailed usage syntax for specified AF.
       route {-V|--version}                  Display version/author and exit.

        -v, --verbose            be verbose
        -n, --numeric            don't resolve names
        -e, --extend             display other/more information
        -F, --fib                display Forwarding Information Base (default)
        -C, --cache              display routing cache instead of FIB

  <AF>=Use -4, -6, '-A <af>' or '--<af>'; default: inet
  List of possible address families (which support routing):
    inet (DARPA Internet) inet6 (IPv6) ax25 (AMPR AX.25)
    netrom (AMPR NET/ROM) ipx (Novell IPX) ddp (Appletalk DDP)
    x25 (CCITT X.25)

```

## route

查看上面的东西就可以知道，主要用的就是 route 命令的 add | del 分支
也就是添加和删除

### route add

#### route add 帮助

```vim
inet_route [-vF] add {-host|-net} Target[/prefix] [gw Gw] [metric M]
                       [netmask N] [mss Mss] [window W] [irtt I]
                       [mod] [dyn] [reinstate] [[dev] If]
inet_route [-vF] add {-host|-net} Target[/prefix] [metric M] reject
```

#### route add 说明

按照下面的写好参数，匹配到的路由规则就会被添加

### route del

#### route del 帮助

```vim
Usage: inet_route [-vF] del {-host|-net} Target[/prefix] [gw Gw] [metric M] [[dev] If]
```

#### route del 说明

按照下面的写好参数，匹配到的路由规则就会被删除

### 各个参数的解释

#### {-host|-net}

- -host 参数指定可以个主机，指定主机的时候不需要指定掩 🏇
- -net 参数指定一个网段的所有主机

#### Target[/prefix]

- Target 部分为指定的主机地址或者网段号（exp. 10.1.0.0 ）
- [/prefix] 部分指定网段的掩 🏇 ([32/24/16/8] 为常用的掩 🏇 )

#### gw [GW]

- gw 这个是固定的语法，只要指定网关地址就要写这个
- [Gw] 这个部分指定网关的 IP 地址

#### 小提示

网关的地址必须和指定的 网段的地址/主机地址 是在和网卡地址同一个网段上的，  
黄菊花说就是网卡必须要和网关逻辑上能直接通信

```bash
sudo route del -net 10.1.0.0/16 gw 192.168.0.1 wlan0 # wlan0_ip: 192.168.0.108
```

上面的语句会被允许，因为网卡理论上能和指定的网关直接进行通信

```bash
$sudo route del -net 10.1.0.0/16 gw 192.168.1.1 wlan0 # wlan0_ip: 192.168.0.108
SIOCADDRT: Network is unreachable
```

上面的就会出错了，因为路由匹配到的下一跳网卡不能直接和对应的网关通信

我不确定掩 🏇 会不会影响这个最后的结果  
但是我知道的就是 1.1 不可以直接和 0.1 直接通信的

#### metric M

这个其实就是设定一个优先级，M 为一个正整数

- 不同的网卡贡献同一个优先级排序
- 推荐把容易匹配到的放在路由表优先级最高的位置
- 数字越大，优先级越低

#### [dev]

- 填入对于的网卡
- 这个可以理解为，符合对应条件的网卡走的是那个出口
- 路由表的作用在我理解来看主要就是匹配数据出口
- LInux 贡献数据包流入接口

- 因为输入入口是一个整体，所以就算是 wlan0 口的数据流量，也可以匹配对应的 eth0 的路由规则
- 在我看来，route 表的主要目的就是为了指定数据是从那里出去，还有就是路由表不会修改数据报文
- 路由表在 iptables 的 kernel 空间占有两个匹配位置，分别是 PREOUTING 之后 和 OUTPUT 之前
- 可以理解为路由表指定数据怎么走？走哪里？防火墙指定数据流要不要？改不改？

### 比较特殊的情况

#### default

```bash
sudo route add/del default gw [gateway_addr] [dev]
```

这个 **default** 就相当于 -net 0.0.0.0/0  
他的意思就是任意地址都匹配到这个网关

如果使用 _default_ 的时候没有指定对应的网卡，那么 route 命令就会自动匹配到能和指定的网关直接通信的网卡,而且不指定优先级的时候，默认最高优先级

```vim
?》sudo route del -net 0.0.0.0 gw 10.1.0.1 #

？》route
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         10.1.0.1        0.0.0.0         UG    0      0        0 enp7s0
0.0.0.0         192.168.0.1     0.0.0.0         UG    20600  0        0 wlp0s20f3
10.1.0.0        0.0.0.0         255.255.0.0     U     100    0        0 enp7s0
169.254.0.0     0.0.0.0         255.255.0.0     U     1000   0        0 wlp0s20f3
192.168.0.0     0.0.0.0         255.255.255.0   U     600    0        0 wlp0s20f
```

#### route -n

```vim
?>route
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
default         bogon           0.0.0.0         UG    600    0        0 wlp0s20f3
10.1.0.0        0.0.0.0         255.255.0.0     U     100    0        0 enp7s0
link-local      0.0.0.0         255.255.0.0     U     1000   0        0 wlp0s20f3
192.168.0.0     0.0.0.0         255.255.255.0   U     600    0        0 wlp0s20f3

```

可以看见上面有什么 default blogon 的单词但不知道具体是什么地址？

```vim
?>route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.0.1     0.0.0.0         UG    600    0        0 wlp0s20f3
10.1.0.0        0.0.0.0         255.255.0.0     U     100    0        0 enp7s0
169.254.0.0     0.0.0.0         255.255.0.0     U     1000   0        0 wlp0s20f3
192.168.0.0     0.0.0.0         255.255.255.0   U     600    0        0 wlp0s20f3
```

可以看见对应单词被替换曾具体的 IP 地址了，但是比较有意识的就是 route 命令会对 default 进行一个解释但是不会对 blogon 进行解释  
（我以为他会自动解释到这个网络地址的网关地址 这个理论上来说是可以从 dhcp 的信息的来的） ：D

#### SIOCADDRT: Invalid argument

```bash
$sudo route add -net 10.1.0.0  gw 0.0.0.0 netmask 255.255.0.0 enp7s0
SIOCADDRT: Invalid argument
```

出现这个问题就是但我想添加一个路由到这个网段，然后表示这这个网段上的机器就是直接连接的时候出现的一个错误，那么解决方法就是

```bash
sudo route add -net 10.1.0.0  netmask 255.255.0.0 enp7s0
sudo route add -net 10.1.0.0/16 enp7s0
#对于/16和255.255.0.0这个是等孝的
```

也就是说添加直接可连接的路由规则的时候是不需要指定网关为 0.0.0.0 的 当然，后面的设备也是可以不用指定的，因为这个会匹配对应的网卡

对于同样的问题还存在在于以下

```bash
1： sudo route del -net 10.1.0.0 gw 0.0.0.0 #不指定网络☯🏇
SIOCDELRT: Invalid argument
2：  sudo route del -net 10.1.0.0 #不指定网络☯🏇
SIOCDELRT: Invalid argument
```

其实这个也不需要指定对应的网关来进行一个删除，只需要匹配到网络养马，和对应的网段号就可了所以可以删除对应的直连路由的命令为

```bash
1：sudo route add -net 10.1.0.0/16 enp7s0
2： sudo route del -net 10.1.0.0 netmask 255.255.0.0
3：sudo route del -net 10.1.0.0/16 gw 0.0.0.0
4：sudo route add -net 10.1.0.0/16
```
