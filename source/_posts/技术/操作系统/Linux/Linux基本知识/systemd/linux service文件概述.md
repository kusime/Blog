---
date: 2021-03-24 21:37:03
updated: 2021-03-24 21:37:03
title: linux service 文件简述
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 基本介绍

- 服务（Service）类型的 Unit 文件（后缀为 .service）特有的，用于定义服务的具体管理和操作方法
- 从属于 systemd 的服务控制 **|封装守护进程的启动、停止、重启和重载操作，是最常见的一种 Unit 文件**
- 这个文件 从属于 systemd 控制单元
- 文件由组成
  - 基本段
  - Service 段

# Service 段

- 用来 Service 的配置，只有 Service 类型的 Unit 才有这个区块。它的主要字段分为服务生命周期和服务上下文配置两个方面

## Type 部分

| Type=   | 解释                                                   |
| ------- | ------------------------------------------------------ |
| simple  | 默认值，执行 ExecStart 指定的命令，启动主进程          |
| forking | 以 fork 方式从父进程创建子进程，创建后父进程会立即退出 |
| oneshot | 一次性进程，Systemd 会等当前服务退出，再继续往下执行   |
| dbus    | 当前服务通过 D-Bus 启动                                |
| notify  | 当前服务启动完毕，会通知 Systemd，再继续往下执行       |
| idle    | 若有其他任务执行完毕，当前服务才会运行                 |

## 声明周期部分

| 参数名称        | 解释                                                                                                                                                                                                                                                 |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| RemainAfterExit | 值为 true 或 false（默认）。当配置为 true 时，Systemd 只会负责启动服务进程，之后即便服务进程退出了，Systemd 也仍然会认为这个服务还在运行中。这个配置主要是提供给一些并非常驻内存，而是启动注册后立即退出，然后等待消息按需启动的特殊类型服务使用的。 |
| ExecStart       | 启动当前服务的命令                                                                                                                                                                                                                                   |
| ExecStartPre    | 启动当前服务之前执行的命令                                                                                                                                                                                                                           |
| ExecStartPos    | 启动当前服务之后执行的命令                                                                                                                                                                                                                           |
| ExecReload      | 重启当前服务时执行的命令                                                                                                                                                                                                                             |
| ExecStop        | 停止当前服务时执行的命令                                                                                                                                                                                                                             |
| ExecStopPost    | 停止当其服务之后执行的命令                                                                                                                                                                                                                           |
| RestartSec      | 自动重启当前服务间隔的秒数                                                                                                                                                                                                                           |
| Restart         | 定义何种情况 Systemd 会自动重启当前服务 always（总是重启）、on-successon-failure、on-abnormal、on-abort、on-watchdog                                                                                                                                 |
| TimeoutStartSec | 启动服务时等待的秒数，这一配置对于使用 Docker 容器而言显得尤为重要，因其第一次运行时可能需要下载镜像，严重延时会容易被 Systemd 误判为启动失败杀死。通常，对于这种服务，将此值指定为 0，从而关闭超时检测                                              |
| TimeoutStopSec  | 停止服务时的等待秒数，如果超过这个时间仍然没有停止，Systemd 会使用 SIGKILL 信号强行杀死服务的进程                                                                                                                                                    |

## 上下文配置

| 参数名称         | 解释                                                                                                                                                                                                        |
| ---------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Environment      | 为服务指定环境变量                                                                                                                                                                                          |
| EnvironmentFile  | 指定加载一个包含服务所需的环境变量的列表的文件，文件中的每一行都是一个环境变量的定义                                                                                                                        |
| Nice             | 服务的进程优先级，值越小优先级越高，默认为 0。其中 -20 为最高优先级，19 为最低优先级                                                                                                                        |
| WorkingDirectory | 指定服务的工作目录                                                                                                                                                                                          |
| RootDirectory    | 指定服务进程的根目录（/ 目录）。如果配置了这个参数，服务将无法访问指定目录以外的任何文件                                                                                                                    |
| User             | 指定运行服务的用户                                                                                                                                                                                          |
| Group            | 指定运行服务的用户组                                                                                                                                                                                        |
| MountFlags       | 服务的 Mount Namespace 配置，会影响进程上下文中挂载点的信息，即服务是否会继承主机上已有挂载点，以及如果服务运行执行了挂载或卸载设备的操作，是否会真实地在主机上产生效果。可选值为 shared、slaved 或 private |

### MountFlags

| 可选值为 | 解释                                                                                                                    |
| -------- | ----------------------------------------------------------------------------------------------------------------------- |
| shared   | 服务与主机共用一个 Mount Namespace，继承主机挂载点，且服务挂载或卸载设备会真实地反映到主机上                            |
| slave    | 服务使用独立的 Mount Namespace，它会继承主机挂载点，但服务对挂载点的操作只有在自己的 Namespace 内生效，不会反映到主机上 |
| private  | 服务使用独立的 Mount Namespace，它在启动时没有任何任何挂载点，服务对挂载点的操作也不会反映到主机                        |
| 上       |

## 资源分配

| 参数名称    | 解释         |
| ----------- | ------------ |
| LimitCPU    | CPU          |
| LimitSTACK  | 程序堆栈     |
| LimitNOFILE | 文件句柄数量 |
| LimitNPROC  | 子进程数量   |

# 注意

    - 如果在 ExecStart、ExecStop 等属性中使用了 Linux 命令,则必须要写出完整的绝对路径
    - 对于 ExecStartPre 和 ExecStartPost 辅助命令，若前面有个 “-” 符号，表示忽略这些命令的出错
    - 因为有些 “辅助” 命令本来就不一定成功，比如尝试清空一个文件，但文件可能不存在

# Unit 文件占位符

- 在 Unit 文件中，有时会需要使用到一些与运行环境有关的信息，例如节点 ID、运行服务的用户等。这些信息可以使用占位符来表示，然后在实际运行被动态地替换实际的值。
- 这个优点类似于 shell 中的 $

| 参数名称 | 解释                                                                  |
| -------- | --------------------------------------------------------------------- |
| %n       | 完整的 Unit 文件名字，包括 .service 后缀名                            |
| %p       | Unit 模板文件名中 @ 符号之前的部分，不包括 @ 符号                     |
| %i       | Unit 模板文件名中 @ 符号之后的部分，不包括 @ 符号和 .service 后缀名   |
| %t       | 存放系统运行文件的目录，通常是 “run”                                  |
| %u       | 运行服务的用户，如果 Unit 文件中没有指定，则默认为 root               |
| %U       | 运行服务的用户 ID                                                     |
| %h       | 运行服务的用户 Home 目录，即 %{HOME} 环境变量的值                     |
| %s       | 运行服务的用户默认 Shell 类型，即 %{SHELL} 环境变量的值               |
| %m       | 实际运行节点的 Machine ID，对于运行位置每个的服务比较有用             |
| %b       | Boot ID，这是一个随机数，每个节点各不相同，并且每次节点重启时都会改变 |
| %H       | 实际运行节点的主机名                                                  |
| %v       | 内核版本，即 “uname -r” 命令输出的内容                                |
| %%       | 在 Unit 模板文件中表示一个普通的百分号                                |

# Unit 模板

- 在现实中，往往有一些应用需要被复制多份运行。例如，用于同一个负载均衡器分流的多个服务实例，或者为每个 SSH 连接建立一个独立的 sshd 服务进程。
- Unit 模板文件的写法与普通的服务 Unit 文件基本相同，不过 Unit 模板的文件名是以 @ 符号结尾的。通过模板启动服务实例时，需要在其文件名的 @ 字符后面附加一个参数字符串。

## 启动 Unit 模板的服务实例

- 在服务启动时需要在 @ 后面放置一个用于区分服务实例的附加字符参数，通常这个参数用于监控的端口号或控制台 TTY 编译号。
- 这个就好比启动同一个 apache 服务 ，但是每个服务运行在不同端口
  - 但是就算是不同端口，启动过程基本上是相似的
  - 所以为了简单起见就直接套用模板
  - 然后用@后面的参数来决定启动的服务实例的运行参数
  - systemctl start apache@8080.service
- Systemd 在运行服务时，总是会先尝试找到一个完整匹配的 Unit 文件，如果没有找到，才会尝试选择匹配模板
  - System 首先会在约定的目录下寻找名为 apache@8080.service 的文件
  - 如果没有找到，而文件名中包含 @ 字符，它就会尝试去掉后缀参数匹配模板文件
  - 对于 apache@8080.service，systemd 会找到 apache@.service 模板文件，并通过这个模板文件将服务实例化。

# 实例查看

```bash
[Unit]
Description=My Advanced Service Template
After=etcd.service docker.service
[Service]linux timer 文件简述.md
TimeoutStartSec=0
ExecStartPre=-/usr/bin/docker kill apache%i
ExecStartPre=-/usr/bin/docker rm apache%i
ExecStartPre=/usr/bin/docker pull coreos/apache
ExecStart=/usr/bin/docker run --name apache%i -p %i:80 coreos/apache /usr/sbin/apache2ctl -D FOREGROUND
ExecStartPost=/usr/bin/etcdctl set /domains/kusime.icu/%H:%i running
ExecStop=/usr/bin/docker stop apache1
ExecStopPost=/usr/bin/docker rm apache1
ExecStopPost=/usr/bin/etcdctl rm /domains/kusime.icu/%H:%i
[Install]
WantedBy=multi-user.target
```
