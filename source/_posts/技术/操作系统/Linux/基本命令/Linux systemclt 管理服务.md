---
date: 2021-03-23 20:15:57
updated: 2021-03-23 20:15:57
title: Linux systemclt 管理服务
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 基本介绍

- Systemd 是一系列工具的集合，其作用也远远不仅是启动操作系统，它还接管了后台服务、结束、状态查询，以及日志归档、设备管理、电源管理、定时任务等许多职责
- 并支持通过特定事件（如插入特定 USB 设备）和特定端口数据触发的 On-demand（按需）任务。
- Systemd 的后台服务还有一个特殊的身份——它是系统中 PID 值为 1 的进程。
- Systemctl 接受服务（.service），挂载点（.mount），套接口（.socket）和设备（.device）作为单元。

# 一些扩展知识/一些优点

- Systemd 提供了 服务按需启动 的能力，使得特定的服务只有在真定被请求时才启动。
  - 这也是为什么 archlinux 会选择这个作为管理服务的控件原因
- 允许更多的进程并行启动
  - 在 SysV-init 时代，将每个服务项目编号依次执行启动脚本。
  - 而 Systemd 通过 Socket 缓存、DBus 缓存和建立临时挂载点等方法进一步解决了启动进程之间的依赖，做到了所有系统服务并发启动
  - 对于用户自定义的服务，Systemd 允许配置其启动依赖项目，从而确保服务按必要的顺序运行。
- 在 Systemd 之间的主流应用管理服务都是使用 进程树 来跟踪应用的继承关系的
- Systemd 则提供通过 CGroup 跟踪进程关系
- 通过 CGroup 不仅能够实现服务之间访问隔离，限制特定应用程序对系统资源的访问配额，还能更精确地管理服务的生命周期。
- 统一管理服务日志 ，Journald
  - 这个服务的设计初衷是克服现有 Syslog 服务的日志内容易伪造和日志格式不统一等缺点，Journald 用 二进制格式 保存所有的日志信息
  - Journald 还提供了一个 journalctl 命令来查看日志信息，这样就使得不同服务输出的日志具有相同的排版格式， 便于数据的二次处理。

# systemd 架构图

![error_loading](/gallery/2021-03-24-13-24-01.png)

# Unit 文件概述

- Systemd 可以管理所有系统资源，不同的资源统称为 Unit（单位）。
- Unit 文件统一服务的启/停、定时任务、设备自动挂载、网络配置、虚拟内存配置等
- Systemd 通过不同的文件后缀来区分这些配置文件。

## Unit 文件类型

1. .automount：用于控制自动挂载文件系统，相当于 SysV-init 的 autofs 服务
2. .device：对于 /dev 目录下的设备，主要用于定义设备之间的依赖关系
3. .mount：定义系统结构层次中的一个挂载点，可以替代过去的 /etc/fstab 配置文件
4. .path：用于监控指定目录或文件的变化，并触发其它 Unit 运行
5. .scope：这种 Unit 文件不是用户创建的，而是 Systemd 运行时产生的，描述一些系统服务的分组信息
6. .service：封装守护进程的启动、停止、重启和重载操作，是最常见的一种 Unit 文件
7. .slice：用于表示一个 CGroup 的树，通常用户不会自己创建这样的 Unit 文件
8. .snapshot：用于表示一个由 systemctl snapshot 命令创建的 Systemd Units 运行状态快照
9. .socket：监控来自于系统或网络的数据消息，用于实现基于数据自动触发服务启动
10. .swap：定义一个用户做虚拟内存的交换分区
11. .target：用于对 Unit 文件进行逻辑分组，引导其它 Unit 的执行。它替代了 SysV-init 运行级别的作用，并提供更灵活的基于特定设备事件的启动方式
12. .timer：用于配置在特定时间触发的任务，替代了 Crontab 的功能

# Systemd 目录介绍

- Unit 文件按照 Systemd 约定，应该被放置指定的三个系统目录之一中
- 这三个目录是有优先级的，如下所示，越靠上的优先级越高。
- 因此，在三个目录中有同名文件的时候，只有优先级最高的目录里的那个文件会被使用。
- Systemd 默认从目录 /etc/systemd/system/ 读取配置文件。但是，里面存放的大部分文件都是符号链接，指向目录 /usr/lib/systemd/system/，真正的配置文件存放在那个目录。

## 目录种类

```bash
/etc/systemd/system：系统或用户自定义的配置文件
/run/systemd/system：软件运行时生成的配置文件
/usr/lib/systemd/system：系统或第三方软件安装时添加的配置文件。
  ubuntu 16：被移到了 /lib/systemd/system
```

# SysV-init 运行级别与 Systemd Target 对应的 Unit 文件

- 这里就是 init {0|1|2|3|4|5|6} 和 systemd 的 target 单元的关系

![error_loading](/gallery/2021-03-24-13-39-21.png)

-

# 命令速览

```bash
ps -eaf | grep "[s]ystemd" #用来查看systemd进程的信息，可以发现其是PID为1的进程

systemd-analyze blame # 查看每个进程的启动时间
systemd-analyze critical-chain # 查看启动时的关键链
systemctl --failed  #  列出所有可用单元
systemctl list-* [serviceName.service]  #列出相关的单元，包括服务，挂载点，计时器。。等
systemctl list-units --type=target # 查看target 文件
systemctl {status|start|restart|stop|kill} serviceName.service #控制服务
systemctl {is-active|enable|disable} serviceName.service #控制启动时候的服务状态
systemctl {mask|unmask} #屏蔽/取消屏蔽 （其实这个叫标记）
systemctl show serviceName.service #查看服务配置细节
systemctl-cgtop # 按 CPU、内存、输入和输出列出控制组
systemclt set-default renlevel[3/5].target #设置多用户tty 、 图像界面 这里的效果和init 3、5是差不多的
$ systemctl -H root@rhel7.example.com status httpd.service # 显示远程主机的某个 Unit 的状态
```

# 查看 Unit 的状态

```bash
# 显示系统状态
$ systemctl status

# 显示单个 Unit 的状态
$ ystemctl status bluetooth.service

# 显示远程主机的某个 Unit 的状态
$ systemctl -H root@rhel7.example.com status httpd.service
```

# Unit 的管理

```bash
# 立即启动一个服务
$ sudo systemctl start apache.service

# 立即停止一个服务
$ sudo systemctl stop apache.service

# 重启一个服务
$ sudo systemctl restart apache.service

# 杀死一个服务的所有子进程
$ sudo systemctl kill apache.service

# 重新加载一个服务的配置文件
$ sudo systemctl reload apache.service

# 重载所有修改过的配置文件
$ sudo systemctl daemon-reload

# 显示某个 Unit 的所有底层参数
$ systemctl show httpd.service

# 显示某个 Unit 的指定属性的值
$ systemctl show -p CPUShares httpd.service

# 设置某个 Unit 的指定属性
$ sudo systemctl set-property httpd.service CPUShares=500

```

# 查看 Unit 的依赖关系

```bash
# 列出一个 Unit 的所有依赖，默认不会列出 target 类型
$ systemctl list-dependencies nginx.service

# 列出一个 Unit 的所有依赖，包括 target 类型
$ systemctl list-dependencies --all nginx.service
```

# 服务生命周期

## 服务的激活

```bash
systemctl {enable|start}
```

- systemctl enable：在 /etc/systemd/system/ 建立服务的符号链接，指向 /usr/lib/systemd/system/ 中
- systemctl start：依次启动定义在 Unit 文件中的 ExecStartPre、ExecStart 和 ExecStartPost 命令

## 服务的启动和停止

```bash
systemctl {start|stop|restart|kill|enable|disable|daemon-reload|reset-failed}
```

- 基本管理
  - systemctl start：依次启动定义在 Unit 文件中的 ExecStartPre、ExecStart 和 ExecStartPost 命令
  - systemctl stop：依次停止定义在 Unit 文件中的 ExecStopPre、ExecStop 和 ExecStopPost 命令
  - systemctl restart：重启服务
  - systemctl kill：立即杀死服务
- 服务的开机启动和取消
  - systemctl enable：除了激活服务以外，也可以置服务为开机启动
  - systemctl disable：取消服务的开机启动
- systemctl daemon-reload：Systemd 会将 Unit 文件的内容写到缓存中，因此当 Unit 文件被更新时，需要告诉 Systemd 重新读取所有的 Unit 文件
- systemctl reset-failed：移除标记为丢失的 Unit 文件。在删除 Unit 文件后，由于缓存的关系，即使通过 daemon-reload 更新了缓存，在 list-units 中依然会显示标记为 not-found 的 Unit。

## Target 管理

- Target 就是一个 Unit 组，包含许多相关的 Unit 。启动某个 Target 的时候，Systemd 就会启动里面所有的 Unit。
- 在传统的 SysV-init 启动模式里面，有 RunLevel 的概念，跟 Target 的作用很类似。不同的是，RunLevel 是互斥的，不可能多个 RunLevel 同时启动，但是多个 Target 可以同时启动。
- init x 这个命令是从属于 systemd 的，用来启动对应的 target

```bash
# 查看当前系统的所有 Target
$ systemctl list-unit-files --type=target

# 查看一个 Target 包含的所有 Unit
$ systemctl list-dependencies multi-user.target

# 查看启动时的默认 Target
$ systemctl get-default

# 设置启动时的默认 Target
$ sudo systemctl set-default multi-user.target

# 切换 Target 时，默认不关闭前一个 Target 启动的进程，systemctl isolate 命令改变这种行为，关闭前一个 Target 里面所有不属于后一个 Target 的进程
$ sudo systemctl isolate multi-user.target
```

## Target 与 SysV-init 进程的主要区别

- 默认的 RunLevel（在 /etc/inittab 文件设置）现在被默认的 Target 取代，位置是 /etc/systemd/system/default.target，通常符号链接到 graphical.target（图形界面）或者 multi-user.target（多用户命令行）
- 启动脚本的位置，以前是 /etc/init.d 目录，符号链接到不同的 RunLevel 目录 （比如 /etc/rc3.d、/etc/rc5.d 等），现在则存放在 /lib/systemd/system 和 /etc/systemd/system 目录。
- 配置文件的位置，以前 init 进程的配置文件是 /etc/inittab，各种服务的配置文件存放在 /etc/sysconfig 目录。现在的配置文件主要存放在 /lib/systemd 目录，在 /etc/systemd 目录里面的修改可以覆盖原始设置。


# 日志管理

- Systemd 通过其标准日志服务 Journald 提供的配套程序 journalctl 将其管理的所有后台进程打印到 std:out（即控制台）的输出重定向到了日志文件。
- Systemd 的日志文件是二进制格式的，必须使用 Journald 提供的 journalctl 来查看，默认不带任何参数时会输出系统和所有后台进程的混合日志。
- 默认日志最大限制为所在文件系统容量的 10%，可以修改 /etc/systemd/journald.conf 中的 SystemMaxUse 来指定该最大限制。

```bash
# 查看所有日志（默认情况下 ，只保存本次启动的日志）
$ sudo journalctl

# 查看内核日志（不显示应用日志）：--dmesg 或 -k
$ sudo journalctl -k

# 查看系统本次启动的日志（其中包括了内核日志和各类系统服务的控制台输出）：--system 或 -b
$ sudo journalctl -b
$ sudo journalctl -b -0

# 查看上一次启动的日志（需更改设置）
$ sudo journalctl -b -1

# 查看指定服务的日志：--unit 或 -u
$ sudo journalctl -u docker.servcie

# 查看指定服务的日志
$ sudo journalctl /usr/lib/systemd/systemd

# 实时滚动显示最新日志
$ sudo journalctl -f

# 查看指定时间的日志
$ sudo journalctl --since="2012-10-30 18:17:16"
$ sudo journalctl --since "20 min ago"
$ sudo journalctl --since yesterday
$ sudo journalctl --since "2015-01-10" --until "2015-01-11 03:00"
$ sudo journalctl --since 09:00 --until "1 hour ago"

# 显示尾部的最新 10 行日志：--lines 或 -n
$ sudo journalctl -n

# 显示尾部指定行数的日志
$ sudo journalctl -n 20

# 将最新的日志显示在前面
$ sudo journalctl -r -u docker.service

# 改变输出的格式：--output 或 -o
$ sudo journalctl -r -u docker.service -o json-pretty

# 查看指定进程的日志
$ sudo journalctl _PID=1

# 查看某个路径的脚本的日志
$ sudo journalctl /usr/bin/bash

# 查看指定用户的日志
$ sudo journalctl _UID=33 --since today

# 查看某个 Unit 的日志
$ sudo journalctl -u nginx.service
$ sudo journalctl -u nginx.service --since today

# 实时滚动显示某个 Unit 的最新日志
$ sudo journalctl -u nginx.service -f

# 合并显示多个 Unit 的日志
$ journalctl -u nginx.service -u php-fpm.service --since today

# 查看指定优先级（及其以上级别）的日志，共有 8 级
# 0: emerg
# 1: alert
# 2: crit
# 3: err
# 4: warning
# 5: notice
# 6: info
# 7: debug
$ sudo journalctl -p err -b

# 日志默认分页输出，--no-pager 改为正常的标准输出
$ sudo journalctl --no-pager

# 以 JSON 格式（单行）输出
$ sudo journalctl -b -u nginx.service -o json

# 以 JSON 格式（多行）输出，可读性更好
$ sudo journalctl -b -u nginx.serviceqq
 -o json-pretty

# 显示日志占据的硬盘空间
$ sudo journalctl --disk-usage

# 指定日志文件占据的最大空间
$ sudo journalctl --vacuum-size=1G

# 指定日志文件保存多久
$ sudo journalctl --vacuum-time=1years
```

# systemd 工具 基本介绍

- systemctl：用于检查和控制各种系统服务和资源的状态
- bootctl：用于查看和管理系统启动分区
- hostnamectl：用于查看和修改系统的主机名和主机信息
- journalctl：用于查看系统日志和各类应用服务日志
- localectl：用于查看和管理系统的地区信息
- loginctl：用于管理系统已登录用户和 Session 的信息 
- machinectl：用于操作 Systemd 容器
- timedatectl：用于查看和管理系统的时间和时区信息
- systemd-analyze 显示此次系统启动时运行每个服务所消耗的时间，可以用于分析系统启动过程中的性能瓶颈
- systemd-ask-password：辅助性工具，用星号屏蔽用户的任意输入，然后返回实际输入的内容
- systemd-cat：用于将其他命令的输出重定向到系统日志
- systemd-cgls：递归地显示指定 CGroup 的继承链
- systemd-cgtop：显示系统当前最耗资源的 CGroup 单元
- systemd-escape：辅助性工具，用于去除指定字符串中不能作为 Unit 文件名的字符
- systemd-hwdb：Systemd 的内部工具，用于更新硬件数据库
- systemd-delta：对比当前系统配置与默认系统配置的差异
- systemd-detect-virt：显示主机的虚拟化类型
- systemd-inhibit：用于强制延迟或禁止系统的关闭、睡眠和待机事件
- systemd-machine-id-setup：Systemd 的内部工具，用于给 Systemd 容器生成 ID
- systemd-notify：Systemd 的内部工具，用于通知服务的状态变化
- systemd-nspawn：用于创建 Systemd 容器
- systemd-path：Systemd 的内部工具，用于显示系统上下文中的各种路径配置
- systemd-run：用于将任意指定的命令包装成一个临时的后台服务运行
- systemd-stdio- bridge：Systemd 的内部 工具，用于将程序的标准输入输出重定向到系统总线
- systemd-tmpfiles：Systemd 的内部工具，用于创建和管理临时文件目录
- systemd-tty-ask-password-agent：用于响应后台服务进程发出的输入密码请求