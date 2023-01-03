---

date: 2021-01-11
updated: 2021-01-11
title: 添加 ssh 公匙到 GitHub 中
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

## 基本操作

### 前期基本配置

```bash
sudo apt install git
git config --global user.email "2353863800@qq.com"
git config --global user.name "kusime"
cd [repo-location]
git init
git add .
git commit -m "firtst commit"
git remote add origin [remote url location]
git push -u origin master
```

### ssh-keygen

- 语法

```bash
ssh-keygen -t rsa -C "email@exanple.com"
cd ~/.ssh
cat rsa.pub
```

### 操作

- 操作
  - 打开https://github.com/settings/ssh/new
  - 添加标题和复制的内容就 ok 了

### 碎碎念

- 这里其实就是非对称加密的使用例子
  - 使用 ssh-keygen 进行生产的时候，会产生一个公匙和私匙
  - 传输到 GitHub 的就是我们的公匙
  - 然后 GitHub 使用我们的公匙来加密传输的信息
  - 到达我们机子上我们拿私匙进行解密
  - 这样通信就建立了
- 支持的加密方式

```bash
ssh-keygen -t [mehod]
dsa      ecdsa    ed25519  rsa
```
