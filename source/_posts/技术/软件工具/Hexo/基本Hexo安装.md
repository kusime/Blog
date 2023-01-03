---
title: 基本Hexo安装
date: 2021-08-23 13:42:59
updated: 2021-08-23 13:42:59
index_img: /gallery/2021-08-23-18-05-09.png
tags:
  - HEXO
categories:
  - HEXO
---


本文章的描述环境为 Windows 10
---

# 基本Hexo安装 

- Hexo 以及配套的各种插件安装的环境为`npm` ,[git](http://git-scm.com/)
    - npm会在安装nodejs的自动安装
- 所有对应的只要安装 Node.js 就好了！
    - 然后目前比较稳定并且不会报错的版本为 [v12.22.5](https://nodejs.org/dist/latest-v12.x/node-v12.22.5-x64.msi)

```powershell
npm config set registry https://registry.npm.taobao.org/
npm install -g hexo-cli

hexo init <folder>
cd <folder>
npm install
```

- 一些注意事项

![Nodejs安装](/gallery/2021-08-23-13-52-22.png)

# 安装 Fluid 主题

- [这个主题](https://github.com/fluid-dev/hexo-theme-fluid)也是本博客选择的主题

- 下载对应的主题发行版本，解压到 `hexo init` 之后的文件夹 `themes/`中

![文件夹](/gallery/2021-08-23-14-01-00.png)

- 修改hexo默认配置文件

![_config.conf](/gallery/2021-08-23-14-02-21.png)

---

![修改themes为fluid](/gallery/2021-08-23-14-03-06.png)

## [其他Fluid的主题配置参考](https://hexo.fluid-dev.com/docs/)


![5.0hexo的主题配置注意事项](/gallery/2021-08-23-14-05-51.png)
