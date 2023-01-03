---

date: 2021-02-24 00:30:32
updated: 2021-02-24 00:30:32
title: docker build 学习
index_img: /gallery/2021-08-23-21-34-28.png
tags: 
   - Docker

categories:
   -  Docker

---

# Docker build

### 命令介绍

用法

```vim
docker build [OPTIONS] PATH | URL | -
```

扩展描述 - dockerbuild 命令使用 dockerfile 和上下文来构建镜像 - 上下文由本机文件或者 URL 来指定 - 在构建过程中可以调用任何上下文文件 - 比如可用 COPY 指令来复制上下文文件到镜像中

- URL 目前支持的种类
  - GIT 仓库
  - 压缩包
  - 纯文本

### 目前 URL 支持的种类

#### Git repositories

- 这时仓库本身作为构建镜像的上下文
- 系统会采用其子模组以及仓库本身
  - 提交历史不会被采用
- 处理流程
  - 首先仓库文件会被拉取到本地临时目录
  - 然后这个目录被传送给 docker deamon
  - 本地副本使您能够使用本地用户凭据，VPN 等访问私有存储库
- docker 拉取的过程好比使用 git clone --recursive

##### GIT URL 片段拓展

- 介绍

  - 这个部分主要是和 giturl 的参数设置有关的
  - 同一个仓库，后面添加不同的参数会对构建内容有影响
    - 比如使用哪个分支上的提交
    - 采用哪个分支提交上的哪个文件夹
  - 注意这些参数不会影响 git 显示页面
  - 这些参数的处理是交给 docker deamon 爬虫的：》

- 如果直接使用源仓库（myrepo.git）
  - 使用 refs/heads/master 提交
  - 上下文包含整个提交内容
- 如果使用锚点指定 tag/branch

  - 提交使用对于的标签/分支
  - 上下文包含整个提交内容

- myrepo.git#:myfolder
  - refs/heads/master
  - /myfolder
- myrepo.git#master:myfolder

  - refs/heads/master
  - /myfolder

- myrepo.git#mytag:myfolder
  - refs/tags/mytag
  - /myfolder
- myrepo.git#mybranch:myfolder

  - refs/heads/mybranch
  - /myfolder

- tips
  - 如果使用构建工具，不能指定上下文路径

#### 压缩文档

- 例子 "docker build context.tar.gz"
- 处理流程
  - 下载流程会在运行 docker deamon 的主机上进行
  - 运行 docker deamon 的主机和 docker cli 可以不是同一台
  - 压缩文档中的内容会作为构建镜像上下文
- 注意事项
  - 其压缩文档必须是可是别的
    - xz
    - bzip2
    - gzip

#### 纯文本

- 介绍
  - 除了传递整个上下文文件
  - 还可以通过 URL 或这从标准输入中传入 df 的内容
- 例子
  - 构建管道 （$ docker build - < Dockerfile）
- 注意事项
  - 不管是使用 URL 传递的文本还是直接从管道输入的
  - 都会被保存到名叫 docker file 的文件中
  - 任何 -f （指定本机 dockerfile）的选项都会失效
  - 这个情况下，dockerdeamon 不会收到上下文目录
    - （因为根本没有传入到 dockerdeamon 中）

### 对于构建流程的介绍

- 默认情况下 docker build 会在上下文目录查找 dockerfile
  - 也可以使用 [--file 选项](#docker-build--f)来指定其位置
- 因为 docker build 构建镜像的时候首先会传送文件到 dockerdeamon
  - 所以可以遵循[这个建议](#docker-file-是什么)
  - 或者使用[.dockerignore 文件](http://www.ming-cloud.com/zh/posts/docker-2-/#dockerignore-%E6%96%87%E4%BB%B6)
- 错误处理
  - 如果 dockerclient 和 docker 守护进程失去联系，那么构建会被取消
  - docker 如果那个时候正在拉取镜像，那么拉取过程也会被取消

### Docker build 选项速查

```vim
--add-host               -- 添加一个自定义的主机-IP的映射 (host:ip)
--build-arg              -- 添加构建过程的变量
--cache-from             -- 把镜像作为缓存源
--cgroup-parent          -- 指定容器的父级控制组
--compress               -- 使用gzip压缩构建环境
--cpu-period             -- 限制CPU CFS优先级
--cpu-quota              -- 限制CPU CFS配额 必须不小于1ms，即 >= 1000
--cpu-rt-period          -- 限制CPU运行时钟期间优先级 范围从 100ms~1s，即[1000, 1000000]
--cpu-rt-runtime         -- 限制CPU运行时钟期间runtime 时长
--cpuset-cpus            -- 设置容器使用固定的 CPU
--cpuset-mems            -- 允许在上执行的内存节点（MEMs），只对 NUMA 系统有效
--cpu-shares             -- 设置使用CPU权重
--disable-content-trust  -- 取消镜像的内容校验，默认为真
--file                   -- 指定Dockerfile文件名称
--force-rm               -- 移除中间容器（构建不成功也移除中间容器）
--help                   -- Print usage
--isolation              -- 使用容器隔离技术；
--label                  -- 设置镜像使用的元数据；
--memory                 -- 设置内存最大值；
--memory-swap            -- 设置Swap的最大值为内存+swap，"-1"表示不限swap；
--network                -- 默认 default。在构建期间设置RUN指令的网络模式
--no-cache               -- 创建镜像的过程不使用缓存；
--pull                   -- 尝试去更新镜像的新版本；
--quiet                  -- 安静模式，成功后只输出镜像 ID；
--rm                     -- 成功构建镜像后，移除中间容器（构建不成功就不移除中间容器）
--shm-size               -- 设置容器 /dev/shm 分区的大小  (format is '<number><unit>')
--squash                 -- 将 Dockerfile 中所有的操作压缩为一层。
--tag                    -- 镜像的名字及标签，通常 name:tag 或者 name 格式；可以在一次构建中为一个镜像设置多个标签
--target                 -- 指定构建的目标，达成构建目标之后就不会继续构建了
--ulimit                 -- Ulimit配置。
--userns                 -- Container user namespace

```

## 常用的

### 指定 dockerfile 位置

- 例子

```bash
docker build -f  [path/to/dockerfile]  [other option] [contex path]
```

- 作用介绍
  - -f [filepath] 指定 dockerfile 位置
  - 区别上下文和 dockerfile

### 给镜像打标签

- 例子

```bash
docker build -t vieux/apache:2.0 .
https://docs.docker.com/engine/reference/commandline/tag/
```

- 说明
  - 作用就是给构建后的镜像打上标签
  - 冒号前面就是名字（不用纠结**_/_**）
  - 一个镜像文件可以有多个标签
- 额外涉及到的
  - [docker tag](https://docs.docker.com/engine/reference/commandline/tag/)

### 定义构建时参数

- 例子

```bash
docker build --build-arg VAR1=1234 --build-arg VA2=567 .
```

- 这个部分和 ARG 指令有关

## 不常用的参数

### --cgroup-parent

- 作用
  - 这个涉及[LINUX 的控制组](https://xinqiu.gitbooks.io/linux-insides-cn/content/Cgroups/linux-cgroups-1.html)概念
  - 这个只是是涉及到 Linux 内核的，之后会填坑

### --ulimit

- 作用
  - 在容器中设置 ulimit
  - 涉及到[docker run](https://docs.docker.com/engine/reference/commandline/run/#set-ulimits-in-container---ulimit)

### --security-opt

- 作用
  - This flag is only supported on a daemon running on Windows, and only supports the credentialspec option.
  - The credentialspec must be in the format file://spec.txt or registry://keyname.

### --isolation

- 介绍
  - 在 Windows 上运行 Docker 容器的情况下，此选项很有用。
  - 指定容器虚拟化的技术
  - 在 Linux 上，唯一受支持的是使用 Linux 名称空间的默认选项。
- 可选值
  - default
  - Use the value specified by the Docker daemon’s --exec-opt . If the daemon does not specify an isolation technology,
  - Microsoft Windows uses process as its default value.
  - process
  - hyperv

## 一些奇葩的构建方式

- [Pipe Dockerfile through stdin](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#pipe-dockerfile-through-stdin)
