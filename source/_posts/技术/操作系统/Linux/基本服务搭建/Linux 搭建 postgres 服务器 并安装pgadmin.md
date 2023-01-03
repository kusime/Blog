---
title: Pgadmin 和postgres 的安装配置
index_img: /gallery/2021-08-23-19-18-20.png
date: 2021-12-03 18:15:55
updated: 2021-12-03 18:15:55
tags:
  - Linux
categories:
  - Linux
---

# Pgadmin 和 postgres 的安装配置 <!--- ANCHOR - Pgadmin 和 postgres 的安装配置 -->

[Document](https://www.postgresql.org/docs/14/tutorial-createdb.html)

## 基本安装 <!--- ANCHOR - 基本安装 -->

下面的是官方的命令，所以实际上 kali （我用的是不能安装的因为官方源头里面是没有kali-rolling代号的）
```bash
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get -y install postgresql
```

## install In Kali-Rolling <!--- ANCHOR - install In Kali-Rolling -->

所以说原理就是把代号改为他们官方源头里面有的代号，其实也就是挂羊头卖狗肉的说法，
但是因为都是一个爸爸生出来的Linux，安装也没有问题，然后我去查了写资料能运行的就是
Debian 代号为 bullseye 是可用正常安装的

```vim
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt bullseye-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get -y install postgresql
```

##  insall pgAdmin4 in Kali-Rolling <!--- ANCHOR -  insall pgAdmin4 in Kali-Rolling -->

同理修改

```vim
sudo sh -c 'echo "deb https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/bullseye pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list && apt update '

sudo apt update                                                                                                                                          
sudo apt install pgadmin4

```
## Register postgresql service

```vim
sudo service postgresql start

sudo systemctl enable postgresql
```

### 坑1 手册的坑 <!--- ANCHOR - 坑1 手册的坑 -->

1. 手册阅读不仔细。。我大概就只是简单的打了命令上去发现不行。。

![图片描述](/gallery/2021-12-03-18-46-32.png)

2. 其实已经说了需要你指定自己用户名，默认是使用自己电脑用户名来登录的

![图片描述](/gallery/2021-12-03-19-07-17.png)

3. 但是这个还没有完。然后也还是要指定端口的。。
4. 所以组合成的命令就是要提供具体的端口号和用户名然后在提示下输入命令创建一个数据库

{% note danger %}
![图片描述](/gallery/2021-12-03-18-36-42.png)

其实手册上面的东西也没问题。。。

```bash
createdb -h localhost -p 5432 -U postgres kusime  # kusime is the database name you want to creat
```
![图片描述](/gallery/2021-12-03-18-43-18.png)

---

![图片描述](/gallery/2021-12-03-18-50-06.png)

{% endnote %}

# Pgadmin Connect Server <!--- ANCHOR - Pgadmin Connect Server -->

![图片描述](/gallery/2021-12-03-19-08-52.png)

![图片描述](/gallery/2021-12-03-19-09-12.png)

![图片描述](/gallery/2021-12-03-19-09-33.png)

## Change Default Passowrd <!--- ANCHOR - Change Default Passow0rd -->

![图片描述](/gallery/2021-12-03-19-10-17.png)

# Basic Useage <!--- ANCHOR - Basic Useage -->

## Use command line Connect To Database <!--- ANCHOR - Use command line Connect To Database -->

```bash
psql -h localhost -U postgres
```

![图片描述](/gallery/2021-12-03-19-24-07.png) 

![图片描述](/gallery/2021-12-03-19-26-13.png)

