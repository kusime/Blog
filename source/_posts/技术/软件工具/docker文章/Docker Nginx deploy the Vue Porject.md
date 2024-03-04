---
title: Docker Nginx deploy the Vue Porject
date: 2024-03-01 20:43:26
updated: 2024-03-01 21:00:12
tags:
  - Docker
categories:
  - Docker
---

# Ref

[Vue.jsプロジェクトにおけるnginxの設定とDockerによるコンテナ化の例 - Qiita](https://qiita.com/kwashi/items/da5420dca2d1b3c53ee7)

[Deployment | Vue CLI](https://cli.vuejs.org/guide/deployment.html#docker-nginx)



# Docker Nginx deploy the Vue Porject

使用Hash history 的模式就不需要对nginx服务器进行特殊的配置，但是如果用的是普通URL的方式来部署的话就需要配置一下


## Nginx with 80 with Normal URL

`default.conf`

```conf
server {
    listen 80 ;
    listen [::]:80 ;
    server_name hsc.kusime.icu;
    index index.html index.htm;
    location / {
        root /usr/share/nginx/html;
        try_files $uri $uri/ /index.html;
    }
}

```


## Nginx with 80 redirect and 443 with Normal URL

`default.conf`

```conf

server {

    listen 80 ;
    listen [::]:80 ;
    server_name your.domain.com;

    if ($host = your.domain.com) {
        return 301 https://your.domain.com$request_uri;
    }
}

server {
    server_name your.domain.com;
    listen 443 ssl;
    listen [::]:443 ssl ;
    # Index Page Set ------
    # https set
    ssl_certificate /etc/nginx/ssl/server.crt;
    ssl_certificate_key /etc/nginx/ssl/server.key;

    #Root Index Set -------
    root /usr/share/nginx/html; # location
    # error_page set
    error_page 404 405 406 /404.html;


    index index.html index.htm;


    location / {
        root /usr/share/nginx/html;
        try_files $uri $uri/ /index.html;
    }

}
```



# Dockerfile build

```Dockerfile
# 第一阶段：构建前端资源的 Node 镜像
FROM node AS builder
# 设置工作目录
WORKDIR /app
# 复制当前目录中的所有文件到镜像中的工作目录
COPY package*.json ./
# 安装 Node 依赖
RUN npm install
# 复制源代码
COPY . .
# 构建前端资源
RUN npm run build

# 第二阶段：使用 nginx 服务前端资源
FROM nginx:alpine
# 设置工作目录为 nginx 的静态资源目录
WORKDIR /usr/share/nginx/html
# 移除默认的 nginx 静态资源
RUN rm -rf ./*
# 从构建阶段的 builder 中复制静态资源到 nginx
COPY --from=builder /app/dist .

# HTTPS support
COPY ./https/certificate.crt /etc/nginx/ssl/server.crt
COPY ./https/private.key /etc/nginx/ssl/server.key
# HTTPS Ngnix Conf and VUE project deployment
COPY nginx_config/nginx.conf /etc/nginx/nginx.conf
COPY nginx_config/default.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 4000
# 定义容器启动命令
CMD ["nginx", "-g", "daemon off;"]

```

## HTTPS support

这个部分其实用卷共享也可以，只需要名字一样就可以了 `server.crt` and `server.key`

{% note success %}
这个部分其实可以使用ACME 去进行管理更新
{% endnote %}


# With certbot support

[docker環境で let&#39;s encrypt のssl証明書を取得する - Qiita](https://qiita.com/mttt/items/aa2ba3a0677a803d0436)

## get certificate from letsencrypt


![目录结构](/gallery/2024-03-01-21-23-07.png)

`new certificate`

```bash
docker run --rm -it \
  -v "${PWD}/docker-data/certbot/certs/:/etc/letsencrypt/" \
  -v "${PWD}/docker-data/certbot/logs/:/var/log/letsencrypt/" \
  -p 80:80 \
  certbot/certbot certonly --standalone -d domail.name
```

`renew certificate`

```bash
docker run --rm -it \
  -v "${PWD}/docker-data/certbot/certs/:/etc/letsencrypt/" \
  -v "${PWD}/docker-data/certbot/logs/:/var/log/letsencrypt/" \
  -p 80:80 \
  -p 443:443 \
  certbot/certbot renew
```

## Dockerfile

关闭复制证书

```Dockerfile
# 第一阶段：构建前端资源的 Node 镜像
FROM node AS builder
# 设置工作目录
WORKDIR /app
# 复制当前目录中的所有文件到镜像中的工作目录
COPY package*.json ./
# 安装 Node 依赖
RUN npm install
# 复制源代码
COPY . .
# 构建前端资源
RUN npm run build

# 第二阶段：使用 nginx 服务前端资源
FROM nginx:alpine
# 设置工作目录为 nginx 的静态资源目录
WORKDIR /usr/share/nginx/html
# 移除默认的 nginx 静态资源
RUN rm -rf ./*
# 从构建阶段的 builder 中复制静态资源到 nginx
COPY --from=builder /app/dist .

# HTTPS support
# COPY ./https/certificate.crt /etc/nginx/ssl/server.crt
# COPY ./https/private.key /etc/nginx/ssl/server.key
# HTTPS Ngnix Conf and VUE project deployment
COPY nginx_config/nginx.conf /etc/nginx/nginx.conf
COPY nginx_config/default.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 4000
# 定义容器启动命令
CMD ["nginx", "-g", "daemon off;"]

```

## compose.yml

```yml
services:
  service_name:
    image: imagename
    container_name: container_name
    ports:
      - "80:80"
      - "443:443"

    volumes:
      # TLS configuration
      - ./docker-data/certbot/certs/:/etc/letsencrypt
```


## default.conf

```conf

server {

    listen 80 ;
    listen [::]:80 ;
    server_name your.domain.com;

    if ($host = your.domain.com) {
        return 301 https://your.domain.com$request_uri;
    }
}

server {
    server_name your.domain.com;
    listen 443 ssl http2;
    listen [::]:443 ssl http2;
    # Index Page Set ------
    # https set
    ssl_certificate /etc/letsencrypt/live/your.domain.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/your.domain.com/privkey.pem;

    #Root Index Set -------
    root /usr/share/nginx/html; # location
    # error_page set
    error_page 404 405 406 /404.html;


    index index.html index.htm;


    location / {
        root /usr/share/nginx/html;
        try_files $uri $uri/ /index.html;
    }

}

```



# nginx.conf

```conf
user root;#设置 nginx 服务的系统使用的用户
worker_processes auto; #工作进程数，一般与 CPU 相等
pid /run/nginx.pid; #nginx 服务启动时候的 pid
include /etc/nginx/modules-enabled/*.conf;

events {
    worker_connections 1025; #每个进程运行最大连接数，进程是上面的 worker_processes 指定的 一般是cpu1的核心个数
    multi_accept on; # 精准投喂请求的连接到对应的工作i线程
}

http {


    # server 块 即是一个虚拟主机，需配置域名和端口，也只处理对应主机域名的 http 请求，
    # 内可包含多个 location 块。
    ##
    # Basic Settings
    ##
    sendfile on;
    tcp_nopush on;
    tcp_nodelay on;
    keepalive_timeout 65;
    types_hash_max_size 2048;
    server_tokens off; #隐藏服务器版本
    client_body_buffer_size 16k; # 客户体 缓冲区


    include /etc/nginx/mime.types;
    default_type application/octet-stream; # 默认传输类型

    ##
    # SSL Settings
    ##

    ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3; # Dropping SSLv3, ref: POODLE
    ssl_prefer_server_ciphers on;

    ##
    # Logging Settings
    ##

    access_log /var/log/nginx/access.log;
    error_log /var/log/nginx/error.log;

    ##
    # Gzip Settings
    ##

    gzip on;

    gzip_vary on;
    gzip_proxied any;
    gzip_comp_level 6;
    gzip_buffers 16 8k;
    gzip_http_version 1.1;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;
    gzip_min_length 20;
    ##
    # Virtual Host Configs
    ##

    # https://qiita.com/kwashi/items/da5420dca2d1b3c53ee7
    include /etc/nginx/conf.d/*.conf;
    include /etc/nginx/sites-enabled/*;
}


```

