---
title: Hexo添加一键部署
date: 2021-08-23 14:04:17
updated: 2021-08-23 14:04:17
index_img: /gallery/2021-08-23-18-01-53.png
tags:
  - HEXO
categories:
  - HEXO
---

[其他部署方式的参考](https://hexo.io/zh-cn/docs/one-command-deployment)
---

# Hexo添加一键部署(SFTP)

- 安装 `hexo-deployer-sftp`

```powershell
npm install hexo-deployer-sftp --save
```

- 在博客 `_config.conf` 添加下面描述的信息

```conf
deploy:
  type: sftp
  host: <host>
  user: <user>
  pass: <password>
  remotePath: [remote path]
  port: [port]
```

![参数对照](/gallery/2021-08-23-14-13-00.png)

## 关于部署注意事项

- 这里的部署含义就是把生成的public文件推送到指定的云服务器上
- 我这里就是采用的这种方式，Nginx作为远程的服务端
- 然后本地使用hexo生成并且推送生成的静态网页文件到blog中

---

- 如果有些修改没有成功的时候可以试一下
    - `hexo clean`
- 部署采用命令为(或者)
    - `hexo g -d`
    - `hexo d -g`

