---
date: 2021-03-24 21:38:18
updated: 2021-03-24 21:38:18
title: linux unit 文件通用段解释
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 基本介绍

- 这里介绍 Unit 段 和 Install 段
- Unit 和 Install 段|所有 Unit 文件通用，用于配置服务（或其它系统资源）的描述、依赖和随系统启动的方式

# 什么是服务单元（Unit） ？

- Unit 是 Systemd 管理系统资源的基本单元，可以认为每个系统资源就是一个 Unit，并使用一个 Unit 文件定义
- 在 Unit 文件中需要包含相应服务的描述、属性以及需要运行的命令
- 一个单元配置文件可以描述如下内容之一|系统服务(.service)、挂载点(.mount)、sockets(.sockets) 、系统设备(.device)、交换分区(.swap)、文件路径(.path)、启动目标(.target)、由 systemd 管理的计时器(.timer)

# Unit 段介绍

| 配置名字      | 解释                                                                                                                             |
| ------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| Description   | 描述这个 Unit 文件的信息                                                                                                         |
| Documentation | 指定服务的文档，可以是一个或多个文档的 URL 路径                                                                                  |
| Requires      | 依赖的其它 Unit 列表，列在其中的 Unit 模板会在这个服务启动时的同时被启动。并且，如果其中任意一个服务启动失败，这个服务也会被终止 |
| Wants         | 与 Requires 相似，但只是在被配置的这个 Unit 启动时，触发启动列出的每个 Unit 模块，而不去考虑这些模板启动是否成功 ，是一种期望    |
| After         | 与 Requires 相似，但是在后面列出的所有模块全部启动完成以后，才会启动当前的服务                                                   |
| Before        | 与 After 相反，在启动指定的任务一个模块之间，都会首先确证当前服务已经运行                                                        |
| Binds To      | 与 Requires 相似，失败时失败，成功时成功，但是在这些模板中有任意一个出现意外结束或重启时，这个服务也会跟着终止或重启             |
| Part Of       | 一个 Bind To 作用的子集，仅在列出的任务模块失败或重启时，终止或重启当前服务，而不会随列出模板的启动而启动                        |
| OnFailure     | 当这个模板启动失败时，就会自动启动列出的每个模块                                                                                 |
| Conflicts     | 与这个模块有冲突的模块，如果列出的模块中有已经在运行的，这个服务就不能启动，反之亦然                                             |

# Install 段介绍

- 只在 systemctl enable/disable 操作时有效
- 一般只配置一个 WantedBy 指令
- 如果不期望服务开机自启动，则 Install 段落通常省略

| 配置名字        | 解释                                                                                                                                                                                                                       |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| WantedBy        | 它的值是一个或多个 Target，当前 Unit 激活时（enable）符号链接会放入 /etc/systemd/system 目录下，以 TargetName + .wants 后缀构成的子目录中                                                                                  |
| RequiredBy      | 和 Unit 段的 Wants 作用相似，只有后面列出的不是服务所依赖的模块，而是依赖当前服务的模块。它的值是一个或多个 Target                                                                                                         |
| Also            | 当前 Unit enable/disable 时，同时 enable/disable 的其他 Unit                                                                                                                                                               |
| Alias           | 当前 Unit 可用于启动的别名                                                                                                                                                                                                 |
| DefaultInstance | 当是一个模板服务配置文件时(即文件名为Service_Name@.service)，该指令指定该模板的默认实例。例如trojan@.service中配置了 DefaultInstall=server 时，systemctl enable trojan@.service时将创建名为trojan@server.service的软链接。 |
