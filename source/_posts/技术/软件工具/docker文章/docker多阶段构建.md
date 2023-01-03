---

date: 2021-02-13 00:30:32
updated: 2021-02-13 00:30:32
title: docker 多阶段构建
index_img: /gallery/2021-08-23-21-34-28.png
tags: 
   - Docker

categories:
   -  Docker

---

## 多阶段构建

- 简单介绍
  - 多阶段构建主要是为了方便 dockerfile 的阅读和管理
  - [参考文献](https://blog.alexellis.io/mutli-stage-docker-builds/)

## 在使用多阶段构建之前

- 注意事项
  - 构建镜像尽可能保证镜像尽可能小
  - 每一个在 dockerfile 中的指令就会添加一层镜像
  - 所以需要记得清理在下一层用不到的东西
  - 传统上，您需要使用 Shell 技巧和其他逻辑来使各层尽可能小
  - 并确保每一层都具有上一层所需的工件，而没有其他任何东西。
- 闲谈
  - 实际上，通常只有一个 Dockerfile 用于构建
  - 精简的 Dockerfile 用于构建时，它仅包含您的应用程序以及运行它所需的内容
  - 护两个 Dockerfile 是不理想的。

### 写为行例子

- 这里有一个欠佳的例子
<!-- ######## VIM Start########  -->

```vim
FROM golang:1.7.3
WORKDIR /go/src/github.com/alexellis/href-counter/
COPY app.go .
RUN go get -d -v golang.org/x/net/html \
  && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
```

<!-- ######## VIM End  ########  -->

- 介绍
  - 最后的 RUN 通过添加&&来把多条语句合为一行
  - 这样避免了多余层
  - 但是这样是容易出错而且难以维护
    - 很混乱
    - 不好看

### 关于为什么要用多阶段构建

<!-- ######## VIM Start########  -->

```vim
#!/bin/sh
echo Building alexellis2/href-counter:build

docker build --build-arg https_proxy=$https_proxy --build-arg http_proxy=$http_proxy \
    -t alexellis2/href-counter:build . -f Dockerfile.build

docker container create --name extract alexellis2/href-counter:build
docker container cp extract:/go/src/github.com/alexellis/href-counter/app ./app
docker container rm -f extract

echo Building alexellis2/href-counter:latest

docker build --no-cache -t alexellis2/href-counter:latest .
rm ./app
```

<!-- ######## VIM End  ########  -->

- 介绍
  - 上述是用一个 shell 来构建连个镜像
    - alexellis2/href-counter:build
    - alexellis2/href-counter:latest
  - 更加主要的问题
    - 第二阶段会用到第一阶段的生成文件
    - 第一阶段可能只是构建依赖环境
    - 其第一阶段可能不需要在最终结果出现
- 参考
  - [多阶段构建](https://docker_practice.gitee.io/zh-cn/image/multistage-builds/)
  - [Docker 多阶段构建最佳实践](http://www.dockone.io/article/8179)
  - [Docker 多阶段构建实战(multi-stage builds)](https://blog.csdn.net/boling_cavalry/article/details/90742657)

## 使用多阶段构建

- 介绍
  - 在同一个 dockerfile 中使用多个 FROM 语句
  - 每一个 FROM 语句使用不同的基镜像
  - 而且从其开始，是开始新的一个构建阶段
- 优点
  - 可以选择性的把某一阶段生成的文件传送到另一个构建阶段
  - 而在最后生成的镜像中不包含这些内容

### 例子

<!-- ######## VIM Start########  -->

```vim
FROM golang:1.7.3
WORKDIR /go/src/github.com/alexellis/href-counter/
RUN go get -d -v golang.org/x/net/html
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/alexellis/href-counter/app .
CMD ["./app"]
```

<!-- ######## VIM End  ########  -->

- 构建过程 - 你只需要运行这个命令
<!-- ######## BASH Start########  -->

```bash
docker build -t alexellis2/href-counter:latest .
```

<!-- ######## BASH End  ########  -->

- 结果
  - 降低了构建的复杂性
  - 构建了和之前一样的镜像
  - 无需再创建中间层镜像
  - 也不需要从本地去重新传送上下文
- 解释
  - 第二个构建阶段以 alpine:latest 为基镜像
    - 注意 alpine 是一个[操作系统](https://yeasy.gitbook.io/docker_practice/os/alpine)
    - 第二个阶段是构建运行镜像
    - 第一个阶段是纯 golang 来构建出 app 这个可执行文件
    - 然后复制到第二个真正执行的镜像
    - 用户使用到的 app 是在第二个阶段构建出的！
  - 最后多阶段构建也只是会输出一个最终镜像
    - 除了指定阶段停止

## 其他构建操作

### 为构建阶段命名

- 介绍
  - 默认情况下构建阶段是没有名字额的
  - 但你可以用 index 来引用指定的构建阶段 （0 为第一个）
- 语法

  - 可以在 FROM 语句后添加 AS \<name\> 语句

- 意义

#### 一个例子

<!-- ######## VIM Start########  -->

```vim
FROM golang:1.7.3 AS builder
WORKDIR /go/src/github.com/alexellis/href-counter/
RUN go get -d -v golang.org/x/net/html
COPY app.go    .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/alexellis/href-counter/app .
CMD ["./app"]
```

<!-- ######## VIM End  ########  -->

- 第一阶段
  - 第一句为第一个构建阶段
    - 其名字命名为 builder
  - 第二句定义镜像工作目录
  - 第三句在构建阶段运行语句
  - 第四句复制源代码
  - 第五句构建出的程序（app）文件到工作目录
  - 到第六行第一个构建阶段结束
- 第二阶段
  - 第二个阶段以 alpine:latest 为基镜像
  - 运行构建内语句
  - 定义工作目录
  - 从第一个构建阶段引用其构建出的程序复制到工作目录
  - 定义镜像启动时候运行的指令

### 在指定的构建阶段停止

- 介绍
  - 在多阶段构建的时候你不一定每次都要全部构建完
  - 所以可以在指定的构建阶段停止
- 例子 - 假设接下来的命令使用上个例子的 dockerfile - 那么下面的命令就会在第一个构建阶段结束的时候停止
<!-- ######## BASH Start########  -->

```bash
 docker build --target builder -t alexellis2/href-counter:latest .
```

<!-- ######## BASH End  ########  -->

- 好处
  - 对指定的构建阶段 debugg
  - Using a debug stage with all debugging symbols or tools enabled, and a lean production stage
  - 测试构建阶段的时候可以使用测试的数据
  - 但是可能在最终结果会使用不同构建阶段的真实数据
    - Using a testing stage in which your app gets populated with test data, but building for production using a different stage which uses real data

### 使用外部镜像作为一个阶段

- 解释
  - 在使用多阶段构建的时候你不仅仅可以使用在先前构建过的镜像
  - 可以使用 COPY --from 来从一个单独的镜像中复制文件
    - 使用本地映像名称
    - 本地或 Docker 注册仓库上可用的标签或标签 ID
- 例子
<!-- ######## VIM Start########  -->

```vim
COPY --from=nginx:latest /etc/nginx/nginx.conf /nginx.conf
```

<!-- ######## VIM End  ########  -->

- 理解
  - 这个从 nginx:latest 镜像中复制了/etc/nginx/nginx.conf
  - 文件到目前构建阶段的工作目录
  - 其 nginx:latest 可能不是在本 dockerfile 构建的

### 构建结果作为新构建阶段

- 例子
<!-- ######## VIM Start########  -->

```vim
FROM alpine:latest as builder
RUN apk --no-cache add build-base

FROM builder as build1
COPY source1.cpp source.cpp
RUN g++ -o /binary source.cpp

FROM builder as build2
COPY source2.cpp source.cpp
RUN g++ -o /binary source.cpp
```

<!-- ######## VIM End  ########  -->

- 解释
  - 很明显 hhhh

## 实际使用的体会

### COPY --from

- 理解
  - 这个是为了从某个构建阶段提取出指定的文件
  - 不是！！上下文！！！
  - 是存在与那个构建阶段或者镜像中的文件
    - 可以是镜像中的某些文件
    - 可以是构建阶段生成的程序
- 目的

  - 可能其最终环境只需要编译好的 app
  - 不需要构建 app 时候的所有依赖
  - 这样可以使构建的环境不存在于生产环境

- 注意
  - 上下文文件是互通的
    - 不需要--from 语句指定！
  - 如果是在镜像、构建阶段中的文件
    - 需要--from 语句指定！！
  - 最后只是生成了一个最后需要放在运行环境中的一个！！镜像

### AS \<name\>

- 理解
  - 这时为了指定其构建阶段
  - 方面在下面引用的时候快速找到对应的构建阶段

### 最后生成

- 最后也只是产生一个镜像
  - 一个 dockerfile 构建一个镜像
  - 目的还是为了减小体积
  - 方便文件在不同阶段传输
- 减小复杂度

### 构建例子

<!-- ######## VIM Start########  -->

```vim
FROM ubuntu AS ver1
WORKDIR /home/
COPY ver1/ver1  .
RUN cat ver1&&echo from ver1>ver1&&cat ver1

FROM ver1 AS ver2
WORKDIR /home/
COPY --from=ver1 /home/ver1 .
RUN cat ver1&&echo ver2 edit>/home/ver1

FROM busybox
WORKDIR /home/
COPY --from=ver2 /home/ver1 .
RUN cat ver1
```

<!-- ######## VIM End  ########  -->

<!-- ######## BASH Start########  -->

```bash
docker % docker build --no-cache --file=dockerfile/mul  content
Sending build context to Docker daemon  4.655kB
Step 1/12 : FROM ubuntu AS ver1
 ---> f643c72bc252
Step 2/12 : WORKDIR /home/
 ---> Running in 0f220bbe5e77
Removing intermediate container 0f220bbe5e77
 ---> 04808c2aac75
Step 3/12 : COPY ver1/ver1  .
 ---> fc258435889a
Step 4/12 : RUN cat ver1&&echo from ver1>ver1&&cat ver1
 ---> Running in caf20b702d90
ver1
from ver1
Removing intermediate container caf20b702d90
 ---> 6fb6cbf0de0e
Step 5/12 : FROM ver1 AS ver2
 ---> 6fb6cbf0de0e
Step 6/12 : WORKDIR /home/
 ---> Running in 3382b38a09fe
Removing intermediate container 3382b38a09fe
 ---> d7598ed4f8ed
Step 7/12 : COPY --from=ver1 /home/ver1 .
 ---> 0faa1f50e73b
Step 8/12 : RUN cat ver1&&echo ver2 edit>/home/ver1
 ---> Running in 0a6b241f05fa
from ver1
Removing intermediate container 0a6b241f05fa
 ---> 383ce05a43ca
Step 9/12 : FROM busybox
 ---> b97242f89c8a
Step 10/12 : WORKDIR /home/
 ---> Running in 84e6677b97b1
Removing intermediate container 84e6677b97b1
 ---> 122261d36eba
Step 11/12 : COPY --from=ver2 /home/ver1 .
 ---> aeb4ac1649e1
Step 12/12 : RUN cat ver1
 ---> Running in e339a6ed2dcf
ver2 edit
Removing intermediate container e339a6ed2dcf
 ---> 191c0961439d
Successfully built 191c0961439d

docker %  docker images
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
<none>       <none>    191c0961439d   40 minutes ago   1.23MB

```

<!-- ######## BASH End  ########  -->

- 解释
  - 第一个阶段从上下文文件复制了 ver1 文件作为第一阶段构建输入
    - 然后对 ver1 这个文件做出了修改
      - 可以理解为生成了个新文件
      - 或者是构建出了运行程序
  - 第二个阶段以第一个构建阶段为基镜像
    - 复制在第一个阶段构建出的文件 ver1
    - ver1 在第一个构建阶段修改过
    - 输出其阅读到的文件然后对其修改
  - 第三个阶段以 busybox 为基镜像
    - 其复制从第二个极端构建出来的文件
    - 显示其阅读到的内荣

```bash
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
<none>       <none>    191c0961439d   40 minutes ago   1.23MB
<none>       <none>    383ce05a43ca   40 minutes ago   72.9MB

```

- 理解
  - 最后看生成的镜像只有 busybox 镜像大小 1。2mb 左右
  - 但是其显示的文件确是由两个构建阶段构建过来的
