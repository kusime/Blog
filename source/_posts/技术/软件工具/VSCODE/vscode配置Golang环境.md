---
date:  2021-03-14 10:37:27
updated:  2021-03-14 10:37:27
title:  vscode 配置 Golang 环境
index_img: /gallery/2021-08-23-18-47-02.png
tags:
  - Vscode

categories:
  - Vscode
---

# 安装 golang in linux

```bash
firefox https://golang.google.cn/dl/
# 打开上面的网站,然后选择下载最新的linux tar包
firefox https://golang.google.cn/doc/install#install
#按照上面的教程就可以直接安装到最新版本的
# 经过我的测试，官方提供的方法是可以世界升级版本的
# 但是要注意下载的最新包和运行命令的目录要一样
go version
```

# [vscode 安装和基本配置](/2021/03/14/Vscode/vscode文章/vscode简明配置/)

# 安装必要插件

1. 打开 vscode
2. 在侧边栏找到 ![error_loading](/gallery/2021-03-14-10-17-28.png)
3. 然后在搜索框查找 go
4. 找到排名第一的，直接点击安装就好
5. 然后在你准备运行 go 的时候右下角会提示要安装其他的东西
6. 直接点击 INSTALL ALL 就好了
7. 可能会因为墙的原因安装失败
8. 但是 bing 可以帮你解决哈哈哈
9. 然后运行 go 是用 F5，就是调试

## 如果还是运行不了

```bash
firefox https://github.com/go-delve/delve/tree/master/Documentation/installation
# 可以手动安装一下这个
# 还是不行就试一下
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
go install github.com/go-delve/delve/cmd/dlv@latest
go install github.com/Go-zh/tools/cmd/gopls@latest
```
