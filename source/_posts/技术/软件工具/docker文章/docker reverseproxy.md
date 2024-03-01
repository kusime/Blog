---
title: Docker Reverse Proxy
date: 2024-03-01 18:13:28
tags:
  - Docker
categories:
  - Docker
---

# Docker Reverse Proxy


[GitHub - nginx-proxy/acme-companion: Automated ACME SSL certificate generation for nginx-proxy](https://github.com/nginx-proxy/acme-companion)
[GitHub - nginx-proxy/nginx-proxy: Automated nginx proxy for Docker containers using docker-gen](https://github.com/nginx-proxy/nginx-proxy)

## 说明

作用: 
1. 通过统一的端口，比如 443 提供对容器的https链接，然后http链接在内部网络，容器本身可以不需要提供https证书，简化容器构建
2. 子级域名区分服务，相比传统一个端口一个一个服务，使用了反向代理的话，就可以使用子域名去定向到对应的服务

例子
APP1 : CONTAINER1 RUNNING PORT 443  -> app1.example.com
APP2 :  CONTAINER2 RUNNING PORT 9000 -> app2.example.com


![NetWork](/gallery/2024-03-01-18-52-36.png)



## SetUp nginxproxy with acme 

下面启动了 nginx-proxy 以及 acme-companion ,然后声明了一个docker网络deploy_network。
挂载了本地卷用于存储共享证书信息
开放了80 和 443 用于 http-0 的验证以及反向代理服务的开放.



```yml
version: '3.2'
services:
  nginx-proxy:
    image: nginxproxy/nginx-proxy
    container_name: nginx-proxy
    ports:
      # for the https
      - "443:443"
      # for the certs
      - "80:80"
    volumes:
      # the shared volume
      - ./volumes/certs:/etc/nginx/certs
      - ./volumes/vhost:/etc/nginx/vhost.d
      - ./volumes/html:/usr/share/nginx/html
      # docker deamon supervisor
      - /var/run/docker.sock:/tmp/docker.sock:ro
    networks:
      - deploy_network

  nginx-proxy-acme:
    image: nginxproxy/acme-companion
    container_name: nginx-proxy-acme
    volumes:
      # the shared volume
      - ./volumes/certs:/etc/nginx/certs
      - ./volumes/vhost:/etc/nginx/vhost.d
      - ./volumes/html:/usr/share/nginx/html
      # docker deamon supervisor
      - /var/run/docker.sock:/var/run/docker.sock:ro
      # acme script
      - ./volumes/acme:/etc/acme.sh

    environment:
      - DEFAULT_EMAIL=youremail@address.com
      - NGINX_PROXY_CONTAINER=nginx-proxy

    depends_on:
      - nginx-proxy

    networks:
      - deploy_network

networks:
  deploy_network:
    driver: bridge
    name: deploy_network
```


## Apply Reverse Proxy to the Docker container

By setting the environment variable

```yml
services:
  your_service_name: 
    # build:
    #   context: ./src/backend/
    #   dockerfile: Dockerfile
    image: your_image:your_tag
    container_name: your_container_name
    # Define the Env
    environment:
        # the services that run in the container 
        VIRTUAL_PROTO: http
        VIRTUAL_PORT: 80
        # the sub domain of this service want to use for
        VIRTUAL_HOST: your.sub.domain.com
        LETSENCRYPT_HOST: your.sub.domain.com

    # connect to the deploy_network
    networks:
      - deploy_network

# define the networks that will be connect to the connect where the nginxproxy 
networks:
  deploy_network:
    driver: bridge
    name: deploy_network
```

# NOTE

定义在环境变量里面的

```yml
VIRTUAL_PROTO: http
VIRTUAL_PORT: 80
```

这两个参数，第一个就是协议，是http 还是https。 port的话就是在容器上监听的是哪一个端口。

通过使用这样的反向代理，我们只需要暴露443端口就可以实现多个服务的部署。而不是一个服务一个端口。