---

date: 2021-02-13  10:30:18
updated: 2021-02-13  10:30:18
title: 关于链接 gitraw 出现 refuse 的问题解决
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

#　解決方法

```vim
https://site.ip138.com/raw.Githubusercontent.coｍ
#get the ip from the site
```

```bash
sudo echo " {ip you get}    raw.githubusercontent.com" >>/etc/resolv.conf
```
