---
title: docker 部署wordpress
index_img: /gallery/2021-08-23-21-34-28.png
date: 2023-01-04 21:51:12
tags:
  - Docker
categories:
  - Docker
---

# docker 部署 wordpress

只要在这个 docker compose 下面去配置账户和用户名密码就可以了，这里绑定的是路径 Volume。可以改成命名卷轴，可以保证持久化存储，这里保护了数据库，以及 wp 本体。

{% note danger %}
注意，这里不做平台迁移的保证，就是数据库文件，不能直接像文件一样转移的，所以我个人不推荐这个部署方式，但是确实能够做到容器化部署 wp。除非能够有更好的数据库备份方案，否则不推荐
{% endnote %}

```yml
mcan you repeat the gameversion: "3.1"
# https://stackoverflow.com/questions/50985298/volume-path-or-mount-in-windows-container
services:
  wordpress:
    image: wordpress
    restart: always
    ports:
      - 8080:80
    environment:
      WORDPRESS_DB_HOST: db
      WORDPRESS_DB_USER: kusime
      WORDPRESS_DB_PASSWORD: passwd
      WORDPRESS_DB_NAME: KusimeBlog
    volumes:
      - C:\Users\kusime\Desktop\Blog\wordpress:/var/www/html

  db:
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_DATABASE: KusimeBlog
      MYSQL_USER: kusime
      MYSQL_PASSWORD: passwd
      MYSQL_RANDOM_ROOT_PASSWORD: "1"
    volumes:
      - C:\Users\kusime\Desktop\Blog\database:/var/lib/mysql

volumes:
  wordpress:
  db:
```

# 运行，以及效果

```cmd
docker-compose-v1.exe -f Blog.yaml up
docker-compose-v1.exe -f Blog.yaml down
docker-compose-v1.exe -f Blog.yaml start
```
