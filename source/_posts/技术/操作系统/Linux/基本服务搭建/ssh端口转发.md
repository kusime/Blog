---
date:  2021-03-13 12:07:21
updated:  2021-03-13 12:07:21
title:  ssh 端口转发
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# ssh 本地转发

```bash
ssh -L local_port:remote_addr1:remote_port -N user@remote_addr2
```

## 本地参数说名

| 参数         | 解释                                     |
| ------------ | ---------------------------------------- |
| remote_port  | 对应 addr1 主机的服务端口地址            |
| -L           | 指明转发类型                             |
| user         | 登入 remote_addr2 的用户名               |
| remote_addr2 | 要登入的主机地址                         |
| remote_addr1 | 这个可以为 remote_addr2 能链接到到的主机 |
| localport    | 远程服务转发到本地的哪个端口             |
| -N           | 链接后不打开虚拟终端                     |

## 例子 （把远程 sshd 服务转移到本地）

```bash
ssh -L 7878:ming-cloud.com:22 root@ming-cloud.com
#这个在本地做监听7878的程序是ssh
```

# ssh 远程转发

```bash
ssh -R remote_prot:local_addr:local_prot -N user@remote_addr
```

## 远程参数说名

| 参数        | 解释                                        |
| ----------- | ------------------------------------------- |
| -R          | 指明转发类型                                |
| -N          | 链接后不打开虚拟终端                        |
| remote_prot | 本地的服务转发到远程的端口                  |
| local_addr  | 这个可以为 运行 ssh 的主机 能链接到到的主机 |
| local_prot  | 对应 local_addrn 提供服务的端口             |
| user        | 登入 remote_addr 的用户名                   |
| remote_addr | 要登入的主机地址                            |

## 例子 （把本地的 pulseaudio 服务转移到远程）

```bash
ssh -R 7979:127.0.0.1:4713 -N root@ming-cloud.com
#这个在远程做监听7979的程序是sshd
```

### 一个提醒

- 在执行了上面的代码后
- 会把本地声音服务转移到远程主机的 7979 端口
- 但是 7979 默认 sshd 只会监听 127.0.0.1 这个地址

```bash
sshd  366948  root    9u  IPv6 2411190   TCP [::1]:7979 (LISTEN)
sshd  366948  root   10u  IPv4 2411191   TCP 127.0.0.1:7979 (LISTEN)
```

- 这就会导致只有本机能够享受到这个转发端口
- 如果我们想要另一个能访问到远程地址的主机
- 享受到这个端口，那么就要修改 sshd 的配置

```bash
sudo vim /etc/ssh/sshd_config
sudo service sshd restart
```

- 在上述文件添加或者修改

```vim
AllowTcpForwarding yes
GatewayPorts yes
ListenAddress 0.0.0.0
```

- 可以看到结果

```bash
sshd  368025  root    9u  IPv4 2418891  TCP *:7979 (LISTEN)
sshd  368025  root   10u  IPv6 2418892  TCP *:7979 (LISTEN)
```
