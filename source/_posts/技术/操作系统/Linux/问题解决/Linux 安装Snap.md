---
title: Linux 解决snap安装错误
index_img: /gallery/2021-08-23-19-18-20.png
date: 2021-12-02 19:00:39
updated: 2021-12-02 19:00:39
tags:
  - Linux
categories:
  - Linux
---

# Linux 解决snap安装错误



## Error Detail
{% note danger %}
```bash
$ sudo apt update
sudo apt install snapd
sudo snap install core
Hit:1 http://dl.google.com/linux/chrome/deb stable InRelease
Hit:2 http://packages.microsoft.com/repos/code stable InRelease                                         
Hit:3 https://dl.winehq.org/wine-builds/debian stable InRelease                                         
Hit:4 http://mirrors.ustc.edu.cn/kali kali-rolling InRelease                                            
Hit:5 http://deb.debian.org/debian sid InRelease                                                        
Ign:6 https://apt.releases.hashicorp.com kali-rolling InRelease                    
Err:7 https://apt.releases.hashicorp.com kali-rolling Release
  404  Not Found [IP: 151.101.110.49 443]
Reading package lists... Done
E: The repository 'https://apt.releases.hashicorp.com kali-rolling Release' does not have a Release file.
N: Updating from such a repository can't be done securely, and is therefore disabled by default.
N: See apt-secure(8) manpage for repository creation and user configuration details.
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
snapd is already the newest version (2.51.7-2+b1).
0 upgraded, 0 newly installed, 0 to remove and 244 not upgraded.
error: cannot communicate with server: Post "http://localhost/v2/snaps/core": dial unix /run/snapd.socket: connect: no such file or directory
```
{% endnote %}

## 解决方法

```bash
service snapd start # 启动Snapd
sudo systemctl start snapd.service # 启动
sudo systemctl enable snapd.service # auto launch
```

