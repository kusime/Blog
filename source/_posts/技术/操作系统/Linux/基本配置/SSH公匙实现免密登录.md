---

date: 2021-01-21  10:30:18
updated: 2021-01-21  10:30:18
title: ssh 免密登录
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# ssh 免密登录


![图片描述](/gallery/2021-11-22-23-24-26.png)
### 操作

- 在客户机

```bash
cd ~/.ssh
ssh-keygen -t rsa -C "your email"  # 然后三个回车
cat id_rsa.pub # 然后复制这个公匙
```

- 在服务机

```bash
cd ~/.ssh
vim authorized_keys #复制客户机的公匙
```

- 效果
- 实现客户机向服务机子链接不需要输入密码
- 但是服务机子链接客户机要客户机子的密码
