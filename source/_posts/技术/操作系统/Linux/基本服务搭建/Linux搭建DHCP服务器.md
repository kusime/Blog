---
date: 2021-02-02  19:10:25
updated: 2021-02-02  19:10:25
title: 搭建 DHCP 服务器
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

## 如何在 Linux 上搭建 DHCP 服务器？

不同人的机器安装 dhcpd 的方式不一样
我的环境是 Ubuntu 20.04LTS

### 安装 DHCPD 服务器

```bash
sudo apt install dhcpd
```

### 配置 DHCP 服务器

```bash
sudo vim /etc/dhcp/dhcpd.conf
```

#### 基本配置

```vim
# dhcpd.conf
#
# Sample configuration file for ISC dhcpd
#
# Attention: If /etc/ltsp/dhcpd.conf exists, that will be used as
# configuration file instead of this file.
#

# option definitions common to all supported networks...option domain-name "example.org";option domain-name-servers ns1.example.org, ns2.example.org;

default-lease-time 600;
max-lease-time 7200;

# The ddns-updates-style parameter controls whether or not the server will
# attempt to do a DNS update when a lease is confirmed. We default to the
# behavior of the version 2 packages ('none', since DHCP v2 didn't
# have support for DDNS.)
ddns-update-style none;

# If this DHCP server is the official DHCP server for the local
# network, the authoritative directive should be uncommented.
#authoritative;

# Use this to send dhcp log messages to a different log file (you also
# have to hack syslog.conf to complete the redirection).
#log-facility local7;

# No service will be given on this subnet, but declaring it helps the
# DHCP server to understand the network topology.

# This is a very basic subnet declaration.

subnet 10.1.0.0 netmask 255.255.0.0 {
  range 10.1.0.2 10.1.255.255;
  option domain-name-servers 10.1.0.1;
}

# This declaration allows BOOTP clients to get dynamic addresses,
# which we don't really recommend.

#subnet 10.254.239.32 netmask 255.255.255.224 {
#  range dynamic-bootp 10.254.239.40 10.254.239.60;
#  option broadcast-address 10.254.239.31;
#  option routers rtr-239-32-1.example.org;
#}

# host statements.   If no address is specified, the address will be
# allocated dynamically (if possible), but the host-specific information
# will still come from the host declaration.

#host passacaglia {
#  hardware ethernet 0:0:c0:5d:bd:95;
#  filename "vmunix.passacaglia";
# Fixed IP addresses can also be specified for hosts.   These addresses
# should not also be listed as being available for dynamic assignment.
# Hosts for which fixed IP addresses have been specified can boot using
# BOOTP or DHCP.   Hosts for which no fixed address is specified can only
# be booted with DHCP, unless there is an address range on the subnet
# to which a BOOTP client is connected which has the dynamic-bootp flag
# set.
#host fantasia {
#  hardware ethernet 08:00:07:26:c0:a5;
#  fixed-address fantasia.kusime.icu;
#}

# You can declare a class of clients and then do address allocation
# based on that.   The example below shows a case where all clients
# in a certain class get addresses on the 10.17.224/24 subnet, and all
# other clients get addresses on the 10.0.29/24 subnet.

#class "foo" {
#  match if substring (option vendor-class-identifier, 0, 4):  "SUNW";
#}

#shared-network 224-29 {
#  subnet 10.17.224.0 netmask 255.255.255.0 {
#    option routers rtr-224.example.org;
#  }
#  subnet 10.0.29.0 netmask 255.255.255.0 {
#    option routers rtr-29.example.org;
#  }
#  pool {
#    allow members of "foo";
#    range 10.17.224.10 10.17.224.250;
#  }
#  pool {
#    deny members of "foo";
#    range 10.0.29.10 10.0.29.230;
#  }
#
```

```vim
subnet 10.1.0.0 netmask 255.255.0.0 {
  range 10.1.0.2 10.1.255.255;
  option domain-name-servers 10.1.0.1;
}
```

#### 对于以上配置的说明

这个配置让 DHCP 服务器工作在 10.1.0.0/16 这个网段上
分派地址的范围是 10.1.0.2 10.1.255.255;
域名解析服务器是 10.1.0.1；

### 对本地网络的设置

- 设置本机地址为 10.1.0.1
- netmask 为 255.255.0.0
- gateway 为 10.1.0.1

### 启动 dhcp

```bash
sudo dhcpd
```

#### 如果启动成功

```vim
Internet Systems Consortium DHCP Server 4.4.1
Copyright 2004-2018 Internet Systems Consortium.
All rights reserved.
For info, please visit https://www.isc.org/software/dhcp/
Config file: /etc/dhcp/dhcpd.conf
Database file: /var/lib/dhcp/dhcpd.leases
PID file: /var/run/dhcpd.pid
Wrote 1 leases to leases file.


Listening on LPF/enp7s0/98:fa:9b:49:90:31/10.1.0.0/16
Sending on   LPF/enp7s0/98:fa:9b:49:90:31/10.1.0.0/16
Sending on   Socket/fallback/fallback-net
```

#### 如果重复启动

```vim
Internet Systems Consortium DHCP Server 4.4.1
Copyright 2004-2018 Internet Systems Consortium.
All rights reserved.
For info, please visit https://www.isc.org/software/dhcp/
Config file: /etc/dhcp/dhcpd.conf
Database file: /var/lib/dhcp/dhcpd.leases
PID file: /var/run/dhcpd.pid
There's already a DHCP server running.

If you think you have received this message due to a bug rather
than a configuration issue please read the section on submitting
bugs on either our web page at www.isc.org or in the README file
before submitting a bug.  These pages explain the proper
process and the information we find helpful for debugging.

exiting.
```

---

```bash
$ cat /var/run/dhcpd.pid
1231
$ sudo kill -9 1231
```

#### 如果 dhcpd.leases 冲突

```vim
Internet Systems Consortium DHCP Server 4.4.1
Copyright 2004-2018 Internet Systems Consortium.
All rights reserved.
For info, please visit https://www.isc.org/software/dhcp/
unable to create icmp socket: Operation not permitted
Config file: /etc/dhcp/dhcpd.conf
Database file: /var/lib/dhcp/dhcpd.leases
PID file: /var/run/dhcpd.pid
Can't open /var/lib/dhcp/dhcpd.leases for append.

If you think you have received this message due to a bug rather
than a configuration issue please read the section on submitting
bugs on either our web page at www.isc.org or in the README file
before submitting a bug.  These pages explain the proper
process and the information we find helpful for debugging.

exiting.
```

---

```bash
sudo rm /var/lib/dhcp/*
sudo touch /var/lib/dhcp/dhcpd.leases
sudo dhcpd
```
