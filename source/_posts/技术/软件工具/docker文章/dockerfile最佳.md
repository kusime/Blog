---

date: 2021-02-11 00:30:32
updated: 2021-02-11 00:30:32
title: dockerfile 最佳实践
index_img: /gallery/2021-08-23-21-34-28.png
tags: 
  - Docker

categories:
  -  Docker

---

- [docker file 最佳实践](#docker-file-最佳实践)
  - [注意事项](#注意事项)
- [普遍性的建议](#普遍性的建议)
  - [做测试](#做测试)
  - [理解上下文](#理解上下文)
  - [使用.dockerignore](#使用dockerignore)
  - [使用 docker 多阶段构建](#使用-docker-多阶段构建)
- [不要安装不必要的包](#不要安装不必要的包)
  - [分离应用](#分离应用)
  - [最小化层数](#最小化层数)
- [排列好行](#排列好行)
- [利用构建缓存](#利用构建缓存)
  - [缓存构建规则](#缓存构建规则)

## docker file 最佳实践

### 注意事项

- 介绍
  - 这篇文章包含了使用 dockerfile 构建镜像的最佳实践方式
  - 在此之前，请先学习[docker file](http://www.ming-cloud.com/zh/posts/docker-2-/)和[docker build](http://localhost:1313/zh/posts/docker-3-/)的基本内容
- 碎碎念
  - 每一句 dockerfile 里面的指令会产生一个只读镜像层
  - 每层是先前层变化的增量
- 涉及到的
  - docker [构建镜像的过程](https://docs.docker.com/storage/storagedriver/)

## 普遍性的建议

### 做测试

- 使用 dockerfile 创建镜像不要直接发布了
  - 推荐在本机上测试，重构后在最后发布

### 理解上下文

- docker build 运行所在的目录就是 “上下文路径”
  - 默认情况下 dockerfile 会在上下文根目录
  - 可以使用-f 来指定不同位置（不是上下文）
- 无论 Dockerfile 实际位于何处，当前目录中文件和目录的所有递归内容都将作为构建上下文发送到 Docker 守护
- 建议
  - 创建一个专门存放 dockerfile 的文件夹
  - 创建一个专门存放构建材料的
  - 在构建的时候，使用 --file 指定 dockerfile
  - 最后一个参数指定这个版本要用到的上下文
- 例子

```bash
$[~/Desktop/myproject]tree
.
├── v1
│   └── hello
├── v2
│   └── hello2
└── dockerfiles
    └── Dockerfile
$[~/Desktop/myproject]docker build --file dockerfile/Dockerfile -t hello:v1 v1

$[~/Desktop/myproject]docker build --file=./dockerfiles/v1 --tag=busybox:test context/v1
Sending build context to Docker daemon  2.607kB

```

- 解释
  - 这样子的话只会传输指定的文件夹
  - 如果依赖 . 来表示上下文，可能会出现不必要的文件传入
  - 这样子文件目录也美观
  - dockerfile 的名字也不唯一，只要--file 指定好就 ok

### 使用.dockerignore

- 作用
  - 在构建目录上下文排除不必要的文件
- 相关文档
  - [官方的](https://docs.docker.com/engine/reference/builder/#dockerignore-file)
  - [我整理的](http://www.ming-cloud.com/zh/posts/docker-2-/#dockerignore-%E6%96%87%E4%BB%B6)

### 使用 docker 多阶段构建

- [docker 多阶段构建](https://www.ming-cloud.com/zh/posts/docker-5/)

## 不要安装不必要的包

- 为了减小复杂度，依赖，大小，构建时间
  - 比如说你不需要在数据库镜像安装文本编辑器

### 分离应用

- 每一个容器应该有其注重点
  - 比如说页面可以分解为
    - 数据库
    - 服务端
- 好处
  - 可以方便的缩放
  - 可以方便的复用
- 解决
  - 可以选择性的分离
  - 然后通过 docker 网络保证联系

### 最小化层数

- 只有 RUN COPY ADD 会添加图层
  - 其他会创建中间层
  - 中间层不会增加构建大小
- 建议
  - 尽量使用[多阶段构建](https://www.ming-cloud.com/zh/posts/docker-5/)
  - 并且只添加必要文件

## 排列好行

- 排版好会帮助阅读和维护
  - 例子

```vim
  RUN apt-get update && apt-get install -y \
  bzr \
  cvs \
  git \
  mercurial \
  subversion \
  && rm -rf /var/lib/apt/lists/*
```

## 利用构建缓存

- 构建映像时，Docker 会逐步执行 Dockerfile 中的指令，并按指定的顺序执行每个指令。
- 在检查每条指令时，Docker 会在其缓存中查找可重用的现有映像，而不是创建新的（重复的）映像。

- 如果不想用缓存构建
  - you can use the --no-cache=true option on the docker build
- 如果想用缓存构建
  - 那么你就要知道什么时候能，什么时候不能

### 缓存构建规则

- 如果基镜像已经在缓存中
  - 下一条指令会比较所有基镜像的子镜像
  - 查看是否其中有使用完全相同的指令构建的
  - 如果有，使用缓存，没有，重新构建
- 在大多数情况下，只需将 Dockerfile 中的指令与子映像之一进行比较就足够了
  - 。但是，某些说明需要更多的检查和解释。
- 对于 ADD 和 COPY 指令，将检查图像中文件的内容
  - 并为每个文件计算一个校验和。
  - 在这些校验和中不考虑文件的最后修改时间和最后访问时间
  - 在缓存查找期间，将校验和与现有映像中的校验和进行比较。
  - 如果文件中的任何内容（例如内容和元数据）发生了更改，则缓存将无效。
- 除了 ADD 和 COPY 命令以外
  - 缓存检查不会查看容器中的文件来确定缓存是否匹配。
  - 例如，在处理 RUN apt-get -y update 命令时
    - 不会检查容器中更新的文件，以确定是否存在缓存命中
    - 在这种情况下，仅使用命令字符串本身来查找匹配项。
- 缓存无效后，所有后续 Dockerfile 命令都会生成新映像，并且不使用缓存。
