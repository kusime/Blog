---
date: 2021-03-24 21:37:03
updated: 2021-03-24 21:37:03
title: linux timer 文件简述
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 基本介绍

- 以 ".timer" 为后缀的单元文件， 封装了一个由 systemd 管理的定时器， 以支持基于定时器的启动
- 文件由组成
  - 基本段
  - Timer 段
- 每个定时器单元都必须有一个与其匹配的单元， 用于在特定的时间启动
- 匹配的单元可以通过 Unit= 选项(见下文)明确指定。 若未指定，则默认是与该单元名称相同的 .service 单元
  - 例如 foo.timer 默认匹配 foo.service 单元。

# 注意事项

- 如果在启动时间点到来的时候，匹配的单元已经被启动， 那么将不执行任何动作，也不会启动任何新的服务实例。
- 因此，那些设置了 RemainAfterExit=yes(当该服务的所有进程全部退出之后，依然将此服务视为处于活动状态) 的服务单元一般不适合使用基于定时器的启动
- 因为这样的单元仅会被启动一次，然后就永远处于活动(active)状态了

# Timer 常用段的解释

| 参数名字           | 解释                                                                                                                                                                                                                                |
| ------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Unit               | 该定时器单元的匹配单元， 也就是要被该定时器启动的单元。参数是一个不以 ".timer" 结尾的单元名                                                                                                                                         |
| OnCalendar         | 定义基于挂钟时间(wallclock)的日历定时器，值是一个日历事件表达式， 这是与传统 cron 任务类似的定时器                                                                                                                                  |
| RandomizedDelaySec | 将此单元的定时器随机延迟一小段时间， 这一小段时间的长度 介于零到该指令设置的时间长度之间， 以均匀概率分布。 默认值是零，表示不延迟                                                                                                  |
| Persistent         | 若设为"yes"，则表示将匹配单元的上次触发时间永久保存在磁盘上。 这样，当定时器单元再次被启动时， 如果匹配单元本应该在定时器单元停止期间至少被启动一次， 那么将立即启动匹配单元。 这样就不会因为关机而错过必须执行的任务。 默认值为 no |

# 实例查看

```vim
[Unit]
Description=Daily apt download activities

[Timer]
OnCalendar=*-*-* 6,18:00
RandomizedDelaySec=12h
Persistent=true

[Install]
WantedBy=timers.target

```
