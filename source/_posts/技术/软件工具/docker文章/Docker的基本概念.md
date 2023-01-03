---
date: 2021-02-14 00:30:32
updated: 2021-02-14 00:30:32
title: docker 的基本概念
index_img: /gallery/2021-08-23-21-34-28.png
tags: 
   - Docker

categories:
   -  Docker

---

## Docker 基本概念

### Docker 介绍

Docker 是一个轻量化的,快速的一个虚拟化框架,可以实现快速交付,响应部署,未来的作用主要在云端,docker 解决了传统虚拟机框架的慢速,低效,docker 可以实现对系统资源的充分利用而且其创建,销毁容器的开销十分的小,所以可以实现快速的弹性部署,这个也是未来集群应用的一个好处,docker 的镜像比较的小,可以理解为只包含运行程序的最小环境,这也为集群的部署创建了比较好的环境,docker 的镜像还可以推送到远程仓库,可以和多人进行分享,实现多人合作的功能,还有 docker 解决了因为产品环境的问提,可以实现快速部署,一个镜像本身就包含了全部的环境只要安装了 docker 就可以直接把镜像加载在容器中运行

### Docker 基本概念

#### Docker Deamon

Docker Deamon 是 Docker 的守护进程,主要处理 DockerClient 的请求,管理镜像,网络,物理卷,是 docker 运行的核心 dockerdeamon 可以是运行在远端也可以是运行在本地,dockerdeamon 可以和其他的 dockerdeamon 进行通信

#### Docker Client

Docker Client 就是主要和我们打交道的东西,也就是 Dcoker Cli 我们常用的 docker 命令就是输入到 docker client 中,然后 client 把对应的命令传送给 docker deamon

#### Docker regitries

Docker Regitries 是用来存储镜像的,docker 官方的有 docker hub,这个是最大的共有库,主要存放的是官方的镜像,个人可以拥有自己的私有或者共有的仓库,使用 docker pull 命令的时候会从指定的仓库拉镜像,docker pull 的时候就是推送到指定的仓库(拉取的仓库用的链接好像是阿里云的镜像加速器,这个地址好像是和 push 的仓库地址是不一样的)

#### Docker Object

在 docker 的世界中,常见的对象是 IMAGES,CONTAINER.NETWORKS.VOLUMES.PLUGINS....

#### IMAGES

IMAGES 是用于构建容器的模板,是只读的一个对象,IMAGES 是分层的,每一层对于一个改变(涉及到 dockerfile),是增量增加的方式,所以出来的镜像是十分轻量的,(对于容器可理解为是镜像的最顶部的高一个基本的新建的一个层)
那个层是可读写的),一个镜像通常是由另一个镜像变化来的,也就是说从官方的下载来的镜像就是一白纸,然后我们就用 dockerfile 这个笔来写,然后写出来的东西就是基于那个原来的白纸是不？所以就说这个画是基于那个白纸画出来的,然后对于分层就可以理解为,你画的每一笔就是一个相对于白纸的增量是不？然后每一层镜像就记录一个变化,最后在你画了很多笔之后,最后形成的,可以理解为*最上层*,那么这个最新的画,容器就拿这个来展示了：)

#### CONTAINERS

CONTAINERS 这个就是实在 docker 运行的对象，是跑起来的东西,可以理解为画框,CONTAINERS 中的改变不会影响所使用到的 IMAGES,可以理解为 CONTAINERS 是在所使用的镜像层复制了一层,然后让这一层来可读写,运行,CONTAINER 相对于传统虚拟机来说这个就是轻量的虚拟机,轻量到只运行指定程序的必须依赖...容器和容器之间,容器和宿主机之间,是相互独立的,我们可以对容器的网络,存储,子系统,进行定制化的独立,甚至可以分配其对于可用的系统资源的多少..._没有保存的容器配置,是不会被保存的,这个和镜像的制度性相一致_
