---

date: 2021-02-09 00:30:32
updated: 2021-02-09 00:30:32
title: dockerfile 学习
index_img: /gallery/2021-08-23-21-34-28.png
tags: 
   - Docker

categories:
   -  Docker

---

- [Dockerfile](#dockerfile)
  - [命令查看](#命令查看)
    - [一些小提示](#一些小提示)
    - [Docker file 是什么](#docker-file-是什么)
    - [注释](#注释)
      - [Parser directives](#parser-directives)
        - [syntax](#syntax)
        - [escape](#escape)
    - [.dockerignore 文件](#dockerignore-文件)
    - [FROM](#from)
      - [FROM 与 ARG 作用域问题](#from-与-arg-作用域问题)
    - [RUN](#run)
      - [RUN 缓存机制的小提示](#run-缓存机制的小提示)
      - [ISSUES](#issues)
    - [CMD](#cmd)
    - [LABEL](#label)
    - [使用实例](#使用实例)
      - [MAINTAINER](#maintainer)
      - [EXPOSE](#expose)
    - [ENV](#env)
      - [另类声明方式](#另类声明方式)
      - [环境变量的生命周期](#环境变量的生命周期)
      - [一些建议](#一些建议)
    - [ARG](#arg)
      - [一些使用事项](#一些使用事项)
      - [默认值](#默认值)
      - [一个阶段构建阶段的作用域](#一个阶段构建阶段的作用域)
        - [一些解释](#一些解释)
      - [多个构建阶段的作用域](#多个构建阶段的作用域)
      - [对于以上的一些解释](#对于以上的一些解释)
      - [使用 ARG 参数变量](#使用-arg-参数变量)
        - [一个有传入的重写例子](#一个有传入的重写例子)
        - [一个无传入的重写例子](#一个无传入的重写例子)
      - [默认值拓展语法](#默认值拓展语法)
        - [一些实例](#一些实例)
        - [调用变量方式补充](#调用变量方式补充)
        - [对于支持替换的指令](#对于支持替换的指令)
    - [ADD](#add)
      - [对于权限处理](#对于权限处理)
      - [关于 dockerfile 的来源问题](#关于-dockerfile-的来源问题)
      - [ADD 遵从的规则](#add-遵从的规则)
    - [COPY](#copy)
    - [COPY 遵循的规则](#copy-遵循的规则)
    - [ENTRYPOINT](#entrypoint)
    - [VOLUME](#volume)
      - [VOLUME 的小提示](#volume-的小提示)
    - [USER](#user)
      - [一些注意事项](#一些注意事项)
      - [使用例子](#使用例子)
    - [WORKDIR](#workdir)
      - [WORKDIR 使用实例](#workdir使用实例)
    - [ONBUILD](#onbuild)
      - [实例](#实例)
    - [STOPSIGNAL](#stopsignal)
    - [HEALTHCHECK](#healthcheck)
    - [SHELL](#shell)
    - [文件匹配规则](#文件匹配规则)
      - [ADD 使用实例](#add使用实例)
      - [权限处理规则](#权限处理规则)
      - [.dockerignore 匹配规则](#dockerignore-匹配规则)
    - [RUN 和 CMD 的小提示](#run-和-cmd-的小提示)
      - [环境变量处理](#环境变量处理)
      - [Json 字串](#json-字串)

# Dockerfile

## 命令查看

<!-- #### VIM Start####  -->

```vim
FROM  [镜像名称]   定制的镜像都是基于 FROM 的镜像

RUN  <cmd-line>/["exe" "arg1" "arg2" "arg3" ....] 后者等价于 exe arg1 arg2 ...

COPY [--chown-<user>:<group>] <源路径1>...  <目标路径>

ADD [--chown-<user>:<group>] <源路径1>...  <目标路径>

CMD   <cmd-line>/["exe" "arg1" "arg2" "arg3" ....] 后者等价于 exe arg1 arg2 ...

ENTRYPOINT ["<executeable>","<param1>","<param2>",...]  ["<param1>","<param2>",...]

ENV <key1>-<value1> <key2>-<value2>... 在后续的指令中可以引用

ARG <参数名>[-<默认值>] ARG 设置的环境变量仅对 Dockerfile 内有效

VOLUME ["<路径1>", "<路径2>"...] 在启动容器 docker run 的时候,我们可以通过 -v 参数修改挂载点.

EXPOSE 帮助镜像使用者理解这个镜像服务的守护端口

WORKDIR <工作目录路径>  指定的工作目录

USER <用户名>[:<用户组>]  用于指定执行后续命令的用户和用户组

HEALTHCHECK   用于指定某个程序或者指令来监控 docker 容器服务的运行状态.

ONBUILD <其它指令>
```

<!-- #### VIM End  ####  -->

### 一些小提示

- Docker file 中的每一条命令就会在 docker 镜像新建一层，所以尽量要把多条命令合并为一条命令 (&&) 命令
- 上下文路径：是指 docker 在构建镜像，有时候想要使用到本机的文件（比如复制）,docker build 命令得知这个路径后，会将路径下的所有内容打包。
- 上下文路径下不要放无用的文件，因为会一起打包发送给 docker 引擎，如果文件过多会造成过程缓慢。
- ADD 的优点：在执行 《源文件》 为 tar 压缩文件的话，压缩格式为 gzip, bzip2 以及 xz 的情况下，会自动复制并解压到 《目标路径》.
- ADD 的缺点：在不解压的前提下，无法复制 tar 压缩文件。会令镜像构建缓存失效，从而可能会令镜像构建变得比较缓慢。具体是否使用，可以根据是否需要自动解压来决定。
- CMD 与 RUN 的区别

  - CMD 是在构建完成后，启动之前
  - RUN 是在 docker build.（可以理解为是这个镜像的加工过程，是构建中）

  - 如果 Dockerfile 中如果存在多个 CMD 指令，仅最后一个生效

- ENTRYPOINT
  - CMD 会被当做 ENTRYPOINT 的参数传入
  - 如果 Dockerfile 中如果存在多个 ENTRYPOINT 指令，仅最后一个生效。
  - 在执行 docker run 的时候可以指定 ENTRYPOINT 运行所需的参数。

### Docker file 是什么

- dockerfile 用于自动按照指令逐层构建镜像，可用的 file 源种类有 2 种
  - 本机文件
  - URL

---

- docker cli 会把构建目录下的所有文件以及子目录传递给 dockerdeamon, 如果是 git url 那么子模组也会被传入

---

- Docker 构建镜像的过程是在 docker deamon 中进行的，构建镜像的第一件事情就是把构建所需要的文件递归的传送给 docker deamon, 所以最好的情况就是把 dockerfile 放在一个空的文件夹，然后放入构建所必须的文件

---

- _如果构建目录在根目录。..docker cli 会把整个系统盘全部给搬过去：(_

---

- 为了提高构建的性能,可以手动添加[.dockerignore](#dockerignore-%E6%96%87%E4%BB%B6)文件,来选择性的忽略构建目录下的指定文件

---

### 注释

<!-- #### VIM Start####  -->

```vim

只要是以 # 开头的行,就会被当做是注释行


RUN echo hello \
# comment
world

--------  上下等价

RUN echo hello \
world

---------------------------


        # this is a comment-line
    RUN echo hello
RUN echo world

--------   上下等价


# this is a comment-line
RUN echo hello
RUN echo world


```

<!-- #### VIM End  ####  -->

注意，如果注释行是以持续符结尾那个符号是不会生效的，下一行该是啥就是啥

#### Parser directives

Parser directives 是一种类注释的特殊语法，现在支持的有两种 - syntax - escape

---

约定的规则

- 不能重复
- 大小写不敏感，建议小写
- 行内空格忽略，但不能跨行
- 与后续的指令之间插入空行

##### syntax

用法就是 ：注意这个特性只有在 BuildKit 启用后有用
\\# syntax-name:tag
The syntax directive defines the location of the Dockerfile builder that is used for building the current Dockerfile. The BuildKit backend allows to seamlessly use external implementations of builders that are distributed as Docker images and execute inside a container sandbox environment.

Custom Dockerfile implementation allows you to:

- Automatically get bugfixes without updating the daemon
- Make sure all users are using the same implementation to build your Dockerfile
- Use the latest features without updating the daemon
  Try out new experimental or third-party features

##### escape

用法：

<!-- ######## VIM Start########  -->

```vim
escape-` (backtick)
escape-\ (backslash)
```

<!-- ######## VIM End  ########  -->

---

这个的作用主要是定义在 docker file 中的转移符，这个主要就是处理 Windows 和 Linux 的转移符的一个兼容性

---

比如下面的例子

<!-- #### VIM Start####  -->

```vim

FROM microsoft/nanoserver
COPY testfile.txt c:\\
RUN dir c:\

------

PS C:\John> docker build -t cmd .
Sending build context to Docker daemon 3.072 kB
Step 1/2 : FROM microsoft/nanoserver
 ---> 22738ff49c6d
Step 2/2 : COPY testfile.txt [[[ c:\RUN dir c: ]]]]]]
GetFileAttributesEx c:RUN: The system cannot find the file specified.
PS C:\John>

*可以看见上面的变成了
c:\RUN dir c:
RUN本来应该是在dockerfile中运行的
因为COPY的最后的 \ 被理解为行中继符了,
导致后面的RUN被传递到COPY一行

----

```

<!-- #### VIM End  ####  -->

---

正确的实践

<!-- #### VIM Start####  -->

```vim
# escape-`

FROM microsoft/nanoserver
COPY testfile.txt c:\
RUN dir c:\

------------

PS C:\John> docker build -t succeeds --no-cache-true .
Sending build context to Docker daemon 3.072 kB
Step 1/3 : FROM microsoft/nanoserver
 ---> 22738ff49c6d
Step 2/3 : COPY testfile.txt c:\
 ---> 96655de338de
Removing intermediate container 4db9acbb1682
Step 3/3 : RUN dir c:\
 ---> Running in a2c157f842f5
 Volume in drive C has no label.
 Volume Serial Number is 7E6D-E0F7

 Directory of c:\

10/05/2016  05:04 PM             1,894 License.txt
10/05/2016  02:22 PM    <DIR>          Program Files
```

<!-- #### VIM End  ####  -->

可以看见指令正确的运行了。因为 第一行定义了中继符为 `
所以遇到、之后是不会被理解为一行的，那么直接传入的就是 c:\

- 这个主要就是为了解决 LINUX 和 WINDOWS 的持续符的差别
  - Linux 是 \
  - WINDOWS 是 _`_
- 主要原因还是因为 window 把 \ 这个符号当做是盘符目录

### .dockerignore 文件

- 时间：在 docker client 传递文件给 docker deamon 之前
- 行为：docker 会在上下文（构建镜像所需文件目录）, 去寻找名字叫做。dockerignore 的文件
- 动作：根据这个文件内容去[匹配](#dockerignore匹配规则)上下文的所有文件,然后剔除之
- 目的：避免把不必要或者敏感文件或目录传送给 docker deamon 让其有潜在的可能性让 ADD / COPY 把那些文件添加在镜像之中

### FROM

三种调用格式

<!-- ######## VIM Start########  -->

```vim
- FROM [--platform-<platform>] <image> [AS <name>]
- FROM [--platform-<platform>] <image>[:<tag>] [AS <name>]
- FROM [--platform-<platform>] <image>[@<digest>] [AS <name>]
```

<!-- ######## VIM End  ########  -->

- 为之后的指令指定基镜像
- FROM 可以在同一个 dockerfile 出现多次
- 每个 FROM 指令清除由先前指令创建的任何状态
- 在执行下一个 FORM 指令之前输出现在这个构建中的 ID
- tag / digest 这两个参数是可选的，如果不填入的话就默认使用最新的来构建
- 如果 FROM 引用了多平台镜像，则可选的 --platform 标志可用于指定镜像的平台

#### FROM 与 ARG 作用域问题

[多构建阶段的作用域](#多个构建阶段的作用域)
[多阶段构建介绍](http://www.ming-cloud.com/zh/posts/docker-5/)

### RUN

形式

<!-- #### VIM Start####  -->

```vim
RUN <command> [shell形式]
RUN ["executable", "param1", "param2"]
```

<!-- #### VIM End  ####  -->

- RUN 指令是在 **构建过程中** 运行，然后提交运行结果

- Docker 产生镜像的过程是分层的，提交只添加增量

- 第二种形式可以避免破坏命令字符串

- 可以使用[SHELL](#shell)命令更改默认 shell.

- 第一种模式可以用持续符 ( **\\** ) 来把多行变一行

<!-- #### VIM Start####  -->

```vim
RUN /bin/bash -c 'source $HOME/.bashrc; \
echo $HOME'

---- 它们在一起等效于以下这一行

RUN /bin/bash -c 'source $HOME/.bashrc; echo $HOME'

---  要使用'/ bin / sh'以外的其他shell,请使用exec形式传入所需的shell.例如：

RUN ["/bin/bash", "-c", "echo hello"]

----

exec表单被解析为JSON数组,这意味着您必须在单词而非单引号(')周围使用双引号(“)

-----
```

<!-- #### VIM End  ####  -->

#### RUN 缓存机制的小提示

- 运行指令的缓存在下次生成期间不会自动失效，就比如说
- **RUN apt-get dist-upgrade -y** 会被下次构建时候重用，
- 如果构建的时候使用了 --no-cache 那么这个缓存就会失效

#### ISSUES

Known issues (RUN)
Issue [783](https://github.com/docker/docker/issues/783) is about file permissions problems that can occur when using the AUFS file system. You might notice it during an attempt to rm a file, for example.

For systems that have recent aufs version (i.e., dirperm1 mount option can be set), docker will attempt to fix the issue automatically by mounting the layers with dirperm1 option. More details on dirperm1 option can be found at aufs man page

If your system doesn’t have support for dirperm1, the issue describes a workaround.

### CMD

三种形式

<!-- ######## VIM Start########  -->

```vim
- CMD ["executable","param1","param2"]
- CMD ["param1","param2"] (as default parameters to ENTRYPOINT)
- CMD command param1 param2 (shell form)
```

<!-- ######## VIM End  ########  -->

- Dockerfile 中只能有一个 CMD 指令。如果列出多个 CMD, 则只有最后一个 CMD 生效。
- CMD 的主要目的是为执行容器提供默认值。这些默认值可以包括可执行文件，也可以省略可执行文件，在这种情况下，还必须指定 ENTRYPOINT 指令。
- 如果 CMD 用于为 ENTRYPOINT 指令提供默认参数,则 CMD 和 ENTRYPOINT 指令都应使用 [JSON 数组](#json字串)格式指定.
- 这个 CMD 指令会出现和之前的 RUN 指令一样的处理环境变量的 [规则](#环境变量处理)

### LABEL

形式

<!-- #### VIM Start####  -->

```vim
    LABEL <key>=<value> <key>=<value> <key>=<value> ...
```

<!-- #### VIM End  ####  -->

作用和介绍

- LABEL 指令的作用主要就是给镜像添加元数据，
- 是为了给镜像的使用者提供一些信息，比如说镜像的版本，种类，作用，注意事项。..

### 使用实例

<!-- ######## VIM Start########  -->

```vim
LABEL multi.label1="value1" multi.label2="value2" other="value3"


LABEL multi.label1="value1" \
      multi.label2="value2" \
      other="value3"

查看信息 docker image inspect --format='' myimage
{
  "com.example.vendor": "ACME Incorporated",
  "com.example.label-with-value": "foo",
  "version": "1.0",
  "description": "This text illustrates that label-values can span multiple lines.",
  "multi.label1": "value1",
  "multi.label2": "value2",
  "other": "value3"
}
```

<!-- ######## VIM End  ########  -->

#### MAINTAINER

<!-- #### VIM Start####  -->

```vim
MAINTAINER <name>

----推荐用法
LABEL maintainer=<name>
```

<!-- #### VIM End  ####  -->

- 这个本质上就是声明一种特殊的 LABEL, 即指明这个镜像的维护者是谁
- 这个语法现在已经过时了，推荐使用 LABEL 来定义 maintainer

#### EXPOSE

使用格式

<!-- #### VIM Start####  -->

```vim
EXPOSE <port> [<port>/<protocol>...]

--- 默认情况下暴露的是TCP端口可以指定暴露的协议
EXPOSE 80/udp

--- 如果需要两个协议暴露在同一端口
EXPOSE 80/tcp
EXPOSE 80/udp


- 在这种情况下,如果将-P与docker run配合使用,则该端口仅对TCP公开一次,对于UDP公开一次.
- 请记住,-P在主机上使用临时的高阶主机端口,因此该端口对于TCP和UDP将是不同的.
- 上面的估计是和docker run 命令相关的.
```

<!-- #### VIM End  ####  -->

- 注意，这里指定的不是真的暴露这个端口，本质上这个是类似于 LABEL 的一种文档了类型，是用来和使用镜像的人来做一个小提示的。正式使端口暴露的时候是在 docekr run 的时候 -p 指定的一个端口。
- 因为是文档性质类型的，所以及时被使用者重写也是没有问题的
- 第一点就是写在 Dockerfile 中进行声明，能让运维人员或者后来者知道我们开启了容器的哪些端口
- EXPOSE 指令通知 Docker 该容器在运行时监听指定的网络端口。您可以指定端口是侦听 TCP 还是 UDP, 如果未指定协议，则默认值为 TCP.
- 当我们声明了 EXPOSE 端口之后，我们使用 -P 命令进行随机映射的时候，是会对这个端口进行映射的
- 下面是涉及到的其他知识
  - [EXPOSE in doceker run](https://docs.docker.com/engine/reference/run/#expose-incoming-ports)
  - [docker network](https://docs.docker.com/network/)

### ENV

语法

<!-- #### VIM Start####  -->

```vim
ENV <key>=<value> ...
ENV MY_VAR my-value
ENV ONE TWO= THREE=world #不推荐
```

作用

---

- ENV 指令设置了环境变量，这些变量和会在之后的构建阶段作为一种代替量
- （比如说 shell 里面用户可以自定义变量，然后用 **$** 应用之前设置的量）

---

#### 另类声明方式

<!-- #### VIM Start####  -->

```vim
与命令行解析一样,引号和反斜杠可用于在值中包含空格.
ENV MY_NAME="John Doe"
ENV MY_DOG=Rex\ The\ Dog
ENV MY_CAT=fluffy

---对于调用方式的补充
ENV指令允许一次设置多个key: value ...变量,下面的示例将在最终镜像中产生相同的最终结果：
ENV MY_NAME="John Doe" MY_DOG=Rex\ The\ Dog \   <----这里是持续符起了作用
    MY_CAT=fluffy


```

<!-- #### VIM End  ####  -->

#### 环境变量的生命周期

- 使用 ENV 指令声明的环境变量会一直存在，就算是输出了最终的镜像，运行在容器中
- 可以通过 docker inspect 来查看这些环境变量
- 并通过 docker run --env \<key\>=\<value\> 来改变对应的值

#### 一些建议

- 有些环境变量可能会改变某些命令的运行结果

- 对于只要在构建过程中使用的环境变量，而且不需要在最终镜像存在的情况下
  - 推荐直接把这个变量设置为 SHELL 中的变量
  - RUN DEBIAN_FRONTEND=noninteractive apt-get update
  - 对于上面的来说其实就是把那个变量设置到了 sh 的环境变量中
  - 而且 RUN 的话只存在与构建过程中，不会存在于最后产生的镜像中|区别 CMD
  - 或者采用 ARG 指令来声明环境变量

### ARG

语法

<!-- ######## VIM Start########  -->

```vim
ARG <name>[=<default value>]

- 如果只有名字的话,那么这个变量就是由命令行 来定义的
- 或者是延续这个变量的生命周期

```

<!-- ######## VIM End  ########  -->

- 注意：不推荐在构建期间使用 ARG 传递密码、个人认证之类的信息，因为任何人都可以通过 docker history 查看到构建期间传递的变量值
- 如果 ARG 指令有默认值并且在构建期间没有接收到参数、则使用默认值

#### 一些使用事项

- 一个 dockerfile 中可以含有多个 ARG 声明的变量
- ARG 主要是定义了在构建过程中使用到的环境变量，通过 docker build --build-arg 来传入
- 如果在用命令行传入，但是 dockerfile 却没有声明的时候，docker 会报错
  - [Warning] One or more build-args [foo] were not consumed.
  - 来说明传入的这个变量没有没使用
- 不推荐在构建的时候来传入一些私密信息，因为使用 [docker history] 命令可以查看构建过程中使用的环境变量
- 可以参考 [“build images with BuildKit”](https://docs.docker.com/develop/develop-images/build_enhancements/#new-docker-build-secret-information)文献来查询如何安全的构建一个镜像

#### 默认值

<!-- #### VIM Start####  -->

```vim
ARG 指令可以选择包含默认值：
FROM busybox
ARG user1=someuser
ARG buildno=1

如果 ARG 指令具有默认值,而且构建的时候没有通过
 docker build --build-arg来传入值,那么对应的
 值会使用默认值

```

<!-- #### VIM End  ####  -->

#### 一个阶段构建阶段的作用域

<!-- #### VIM Start####  -->

```vim
FROM busybox
USER ${user:-some_user}
ARG user
USER $user
```

<!-- #### VIM End  ####  -->

##### 一些解释

<!-- ######## BASH Start########  -->

```bash
docker build --build-arg user=what_user .
```

<!-- ######## BASH End  ########  -->

- 其处理流程
  - 第二句 user 因为没有被声明，所以传入的值没有生效
  - 这个时候第二句解释为 USER some_user
  - 第三句 ARG user , 声明了 user 变量，
  - 命令行的值被传入，user 这个时候其值为 what_user
  - 第四句 USER $user 使用第三句话的值
  - 这个时候被解释为 USER what_user

#### 多个构建阶段的作用域

<!-- #### VIM Start####  -->

```vim
FROM busybox
ARG SETTINGS
RUN ./run/setup $SETTINGS
FROM busybox
ARG SETTINGS
RUN ./run/other $SETTINGS
```

<!-- #### VIM End  ####  -->

#### 对于以上的一些解释

- 第二句没有默认值，那么就应该是从命令行得到该值
- 第三句使用其第二句得到的值
- 第四句 FROM busybox 是第二次出现的 FROM 语句
  - 这个时候构建进入了第二个阶段
  - 在这个时候第一个 ARG 声明的值已经失效了
- 第五句重新声明 SETTINGS 语句
- 第六句使用在第二句从命令行的到值

#### 使用 ARG 参数变量

##### 一个有传入的重写例子

<!-- ######## VIM Start########  -->

```vim
FROM ubuntu
ARG CONT_IMG_VER
ENV CONT_IMG_VER=v1.0.0
RUN echo $CONT_IMG_VER
```

<!-- ######## VIM End  ########  -->

如果这个时候用户运行

<!-- ######## BASH Start########  -->

```bash
docker build --build-arg CONT_IMG_VER=v2.0.1 .
```

<!-- ######## BASH End  ########  -->

- 第二句 ARG 得到从命令行得到的 v2.0.1 值并赋予 CONT_IMG_VER
- 第三句 ENV 声明一个和 ARG 声明的同名环境变量 CONT_IMG_VER - 默认给其值 v1.0.0 - 这个时候 ENV 重写 ARG 的 v2.0.1 - 这个时候 CONT_IMG_VER 被重写为 v1.0.0
- 第四句输出 CONT_IMG_VER - 最后值被解析为 v1.0.0

##### 一个无传入的重写例子

<!-- ######## VIM Start########  -->

```vim
FROM ubuntu
ARG CONT_IMG_VER
ENV CONT_IMG_VER=${CONT_IMG_VER:-v1.0.0}
RUN echo $CONT_IMG_VER

```

<!-- ######## VIM End  ########  -->

如果这个时候用户运行

<!-- ######## BASH Start########  -->

```bash
docker build .
```

<!-- ######## BASH End  ########  -->

- 第二句 ARG 声明一个环境变量 CONT_IMG_VER
- 第三句 ENV 声明一个 CONT_IMG_VER 环境变量，并且如果其 CONT_IMG_VER 为空时，默认值为 v1.0.0
  - 关于这个${CONT_IMG_VER:-v1.0.0}语法的补充看[这里](#默认值拓展语法)
  - 这个时候 ENV 声明的 v1.0.0 成为 CONT_IMG_VER 的值
- 第四句 输出 v1.0.0

#### 默认值拓展语法

两种语法格式

<!-- ######## VIM Start########  -->

```vim
${variable:-word}
${variable:+word}
```

<!-- ######## VIM End  ########  -->

- 解释
  - ${variable:-word}表示如果设置了变量，那么结果将是该值。如果未设置变量，则结果为 word.
  - ${variable:+word}表示如果设置了变量，则 word 将为结果，否则结果为空字符串。

##### 一些实例

<!-- ######## VIM Start########  -->

```vim
${variable:-word}示例
[root@localhost ~]# echo ${NAME:-tom}
tom
[root@localhost ~]# NAME-jerry
[root@localhost ~]# echo ${NAME:-tom}
jerry

-----  ${variable:-word} vs ${variable:+word}

${variable:+word示例
[root@localhost ~]# echo $NAME
jerry
[root@localhost ~]# echo ${NAME:+tom}
tom
[root@localhost ~]# unset NAME
[root@localhost ~]# echo ${NAME:+tom}
[root@localhost ~]# echo $NAME
```

<!-- ######## VIM End  ########  -->

##### 调用变量方式补充

<!-- #### VIM Start####  -->

```vim
FROM busybox
ENV FOO-/bar
WORKDIR ${FOO}   # WORKDIR /bar
ADD . $FOO       # ADD . /bar
COPY \$FOO /quux # COPY $FOO /quux

${NAME} <--> $NAME  <-->  \$NAME 这三个调用方式等价
===========

---- 环境变量支持嵌套使用

ENV abc-hello
ENV abc-bye def-$abc
ENV ghi-$abc

```

<!-- #### VIM End  ####  -->

##### 对于支持替换的指令

ADD
COPY
ENV
EXPOSE
FROM
LABEL
STOPSIGNAL
USER
VOLUME
WORKDIR
ONBUILD (when combined with one of the supported instructions above)

### ADD

形式

<!-- #### VIM Start####  -->

```vim
ADD [--chown=<user>:<group>] <src>... <dest>
ADD [--chown=<user>:<group>] ["<src>",... "<dest>"]
包含空格的路径需要后一种形式.
```

<!-- #### VIM End  ####  -->

作用

- ADD 指令从 src 复制新文件、目录或远程文件 URL, 并将它们添加到路径 dest 处的镜像中（虚拟系统）
- src 可以是多个文件，但是其不许要存在于 上下文路径中
- 每个 src 可能包含通配符,其遵循[匹配规则](#add-1)

#### 对于权限处理

- --chown 功能仅在用于生成 Linux 容器的 Dockerfile 上受支持，并且在 Windows 容器上不起作用。
- 由于用户和组所有权概念不在 Linux 和 Windows 之间转换
  - 因此使用 /etc/passwd 和 /etc/Group 将用户和组名称转换为 ID 会限制此功能仅对基于 Linux OS 的容器可行。

#### 关于 dockerfile 的来源问题

- 如果 dockerfile 是通过 STDIN 传入的 (docker build - < somefile)
  - 那么这个时候是没有指定上下文目录的
  - 所以在文件中的 ADD 指令就只能有 URL
- 当然也可以直接传入一个压缩文档到 STDIN 中 (docker build - < archive.tar.gz)
  - 那么这个时候会自动解压
  - 解压后的文件就作为 dockerfile 构建镜像的上下文目录
  - 那么 ADD 就可以包含在压缩文档中的文件
- 如果 URL 文件有相关的验证手续

  - 这个时候就需要使用 RUN 命令来调用 wget 或者 curl
  - 因为 ADD 指令是不支持验证的

- 如果 ADD 引用的 URL 的源失效了，那么在 ADD 之后的指令都会失效，
  - 这包括由 RUN 指令产生的镜像缓存，也会失效

#### ADD 遵从的规则

- src 路径必须位于生成上下文内
- src 是 URL, 并且 dest 是类文件 (/etc/dest), 那么会被下载并复制到 dest 处
- src 是 URL, 但是 dest 是类文件夹 (/etc/dest/), 那么会被下载到、<dest\>/\<filename\>
- src 是目录，则复制目录的全部内容
- src 是可识别压缩文件，则将其解压缩为目录
  - 来自远程 URL 的资源不会解压缩。
  - 复制或解压缩目录时，它的行为与 tar -x 相同
- 如果 src 是任何其他种类的文件，则将其及其元数据单独复制，如果目标是目录它将被视为一个目录，src 的内容将写在 dest/base(src).
- 如果直接指定了多个 src 资源，或者由于使用通配符，则 dest 必须是目录，并且它必须以斜杠 / 结束。
- 如果 dest 不以尾随斜杠结束，则它将被视为常规文件，src 的内容将写入 dest .
- 如果 dest 不存在，则与路径中所有缺少的目录一起创建。

### COPY

形式

<!-- #### VIM Start####  -->

```vim
COPY [--chown=<user>:<group>] <src>... <dest>
COPY [--chown=<user>:<group>] ["<src>",... "<dest>"]
包含空格的路径需要后一种形式,
```

<!-- #### VIM End  ####  -->

- COPY 的[权限处理](#权限处理规则)和上面的 ADD 权限处理相一致
- COPY 文件的[匹配规则](#文件匹配规则)和上面的 ADD 的匹配规则相一致

### COPY 遵循的规则

- src 路径必须位于生成上下文内
- 如果 src 是目录，则复制目录的全部内容，包括文件系统元数据
- 如果 src 是任何其他种类的文件，则将其及其元数据单独复制。在这种情况下，如果 dest 以尾随斜杠 / 结尾，则它将被视为目录，src 的内容将写入 dest /base ( src )
- 如果直接指定了多个 src 资源，或者由于使用通配符，则 dest 必须是目录，并且它必须以斜杠 / 结束。
- 如果 dest 不以尾随斜杠结束，则它将被视为常规文件，src 的内容将写入 dest .
- 如果 dest 不存在，则与路径中所有缺少的目录一起创建。

### ENTRYPOINT

形式

<!-- #### VIM Start####  -->

```vim
ENTRYPOINT ["executable", "param1", "param2"]  ----The exec form
ENTRYPOINT command param1 param2       ---- The shell form:
```

<!-- #### VIM End  ####  -->

作用

所以通过上面的实例可以理解

- 如果在 docker file 中指定了 ENTRYPOINT, 那么入口点所指定的程序或者脚本才是真正运行的实体，其他 CMD 输入，或者文件中指定的 CMD 会作为参数传入
- 可以通过 --entrypoint 来重写入口点
- 执行窗体解析为 [JSON 数组](#json字串),这意味着您必须围绕非单引号单词 (')使用双引号 (')
- 对于 ENTRTPOINT 的环境变量处理和 RUN 的[处理方式](#环境变量处理)一样
- (docker run \<image-id\> \<exec\>)
  - 对于运行的命令会作为参数传入到 ENTRYPOINT 中

---

The shell form prevents any CMD or run command line arguments from being used, but has the disadvantage that your ENTRYPOINT will be started as a subcommand of /bin/sh -c, which does not pass signals. This means that the executable will not be the container’s PID 1 - and will not receive Unix signals - so your executable will not receive a SIGTERM from docker stop \<container\>

---

### VOLUME

格式

<!-- ######## VIM Start########  -->

```vim
VOLUME ["/data"]
```

<!-- ######## VIM End  ########  -->

作用

- VOLUME 的主要作用还是为了能在一种具有 “暂时性”的容器中能保存一种悠久的信息，这个信息的流向是 容器 -> 主机
- 对于运行容器来说，其 VOLUME 被理解为一种卷，但是相对于主机来说就是为了一种 文件

#### VOLUME 的小提示

- 基于 Windows 的容器上的卷：使用基于 Windows 的容器时，容器中的卷的目标必须是：
  - 不存在或空目录
  - C 以外的驱动器：
- 所有想要永久保存的操作（操作产生数据）, 通过卷的方式保存到主机，那么就需要所有的操作都在卷的声明之前完成

<!-- #### VIM Start####  -->

```vim
FROM ubuntu
RUN mkdir /myvol
RUN echo "hello world" > /myvol/greeting
VOLUME /myvol
RUN echo "hello world" > /myvol/discard

===========
对于上面的例子来说,只会保存 VOLUME /myvol前 写入myvol文件夹的内容
```

<!-- #### VIM End  ####  -->

- JSON 格式：列表解析为 JSON 数组。您必须用双引号 (') 而不是单引号 (') 括好单词。

- VOLUME 不是真的挂在一个主机的真实卷到镜像中
  - 为了保证可移植性
  - 而是在镜像里面虚拟了一个盘
  - 这个盘的文件会保存在运行这个镜像的主机里

### USER

格式

<!-- #### VIM Start####  -->

```vim
USER <user>[:<group>]
USER <UID>[:<GID>]
```

<!-- #### VIM End  ####  -->

作用

为后续的指令使用指定的用户、用户组，如果用户不存在于机器里面是需要手动创建的，而且默认 docker 会默认拿 root 用户来运行指令

#### 一些注意事项

- 用户指令设置用户名（或 UID) 和用户组（或 GID) 以在运行映像时使用，以及用于 Dockerfile 中遵循的任何 RUN、CMD 和 ENTRYPOINT 指令。
- 请注意，为用户指定组时，用户将只有指定的组成员身份。将忽略任何其他配置的组成员身份。
- 当用户没有主组时，图像（或下一个指令）将与 root 组一起运行。
- 在 Windows 上，如果用户不是内置帐户，则必须先创建该用户。这可以通过作为 Dockerfile 的一部分调用的 net 用户命令完成。

#### 使用例子

使用例子

<!-- #### VIM Start####  -->

```vim
FROM ubuntu
RUN whoami

USER ming
RUN whoami
```

<!-- #### BASH Start####  -->

```bash
➜    docker build .
Sending build context to Docker daemon  4.096kB
Step 1/5 : FROM ubuntu
 ---> f643c72bc252
Step 2/5 : RUN whoami
 ---> Using cache
 ---> eb93247d28ca
Step 3/5 : USER ming
 ---> Running in 9d02b483e8ca
Removing intermediate container 9d02b483e8ca
 ---> e558d5d7752b
Step 4/5 : RUN whoami
 ---> Running in f163296f2b23
unable to find user ming: no matching entries in passwd file



0--------------0
可以看见上面在USER ming之后,之后的RUN运行的命令就尝试以 ming的身份来运行了,因为没有ming这个用户那么构建也失败了
```

<!-- #### BASH End  ####  -->

<!-- #### VIM End  ####  -->

### WORKDIR

- 作用
  - 为之后的命令指定工作目录

#### WORKDIR 使用实例

<!-- ######## VIM Start########  -->

```vim
WORKDIR /a
WORKDIR b
WORKDIR c
RUN pwd
-------------
The output of the final pwd command in this Dockerfile would be /a/b/c.
-------------
ENV DIRPATH=/path
WORKDIR $DIRPATH/$DIRNAME
RUN pwd
--------------
The output of the final pwd command in this Dockerfile would be /path/$DIRNAME
--------------

```

<!-- ######## VIM End  ########  -->

### ONBUILD

- 作用
  - 这个指令定义的指令不会在构建的过程中直接运行
  - 在构建完成后，生成的镜像作为另一个镜像的基镜像时候
    - 会运行 ONBUILD 的指令

#### 实例

<!-- #### VIM Start####  -->

```vim
我们来看一个简单例子.

1、先编写一个Dockerfile文件,内容如下：

#test
FROM ubuntu
MAINTAINER hello
ONBUILD RUN mkdir mydir
利用上面的dockerfile文件构建镜像： docker build -t imagea .
利用imagea镜像创建容器： docker run --name test1 -it imagea /bin/bash

我们发现test1容器的根目录下并没有mydir目录.说明ONBUILD指令指定的指令并不会在自己的构建中执行.

2、再编写一个新的Dockerfile文件,内容 如下

#test
FROM imagea
MAINTAINER hello1
注意,该构建准备使用的基础镜像是上面构造出的镜像imagea
利用上面的dockerfile文件构建镜像： docker build -t imageb .
利用imagea镜像创建容器： docker run --name test2 -it imageb /bin/bash

我们发现test2容器的根目录下有mydir目录,说明触发器执行了. 这个其实从构建imageb的输出日志就可看出.日志如下：

复制代码
xxx@ubuntu:~/myimage$ docker build -t imageb .
Sending build context to Docker daemon 15.87 kB
Step 1 : FROM imagea
# Executing 1 build trigger...
Step 1 : RUN mkdir mydir
 ---> Running in e16c35c94b03
 ---> 4b393d1610a6
Removing intermediate container e16c35c94b03
Step 2 : MAINTAINER hello1
 ---> Running in c7b0312516ea
 ---> 0f63b8e04d82
Removing intermediate container c7b0312516ea
Successfully built 0f63b8e04d82
复制代码
我们可以看出,FROM指令执行之后,就立即执行的是触发器(ONBUILD指令指定的指令)

copy from  https://www.cnblogs.com/51kata/p/5265107.html
```

<!-- #### VIM End  ####  -->

### STOPSIGNAL

<!-- #### VIM Start####  -->

```vim
STOPSIGNAL 指令设置将被发送到容器退出的系统调用信号.该信号可以是与内核 syscall 表中的位置匹配的有效无符号数字(例如9),也可以是 SIGNAME 格式的信号名称(例如 SIGKILL).

使用这个指令允许用户自定义应用在收到 docker stop 所发送的信号,是通过重写 signal 库内的 stopsignal 来支持自定义信号的传递,在上层调用时则将用户自定义的信号传入底层函数.

https://stackoverflow.com/questions/34260858/how-to-use-stopsignal-instruction-within-docker
```

<!-- #### VIM End  ####  -->

### HEALTHCHECK

<!-- #### VIM Start####  -->

```vim
https://blog.csdn.net/tech_salon/article/details/77255915
https://www.bookstack.cn/read/docker-practice/image-dockerfile-healthcheck.md
```

<!-- #### VIM End  ####  -->

### SHELL

<!-- #### VIM Start####  -->

```vim
SHELL ["executable", "parameters"]

Dockerfile文件SHELL指令可以覆盖命令的shell模式所使用的默认shell.Linux的默认shell是[“/bin/sh”, “-c”],Windows的是[“cmd”, “/S”, “/C”].SHELL指令必须以JSON格式编写.

用于改变默认的SHELL 作用RUN等

```

<!-- #### VIM End  ####  -->

### 文件匹配规则

#### ADD 使用实例

<!-- #### VIM Start####  -->

```vim

ADD hom* /mydir/     ----添加所有以hom开头的文件

----这里的正则表达和上面的dockerignore文件的匹配规则是差不多的

ADD hom?.txt /mydir/ ---？号指代任意单个字母


---  dest 是一个绝对路径,或相对于 WORKDIR 的路径,源将复制到目标容器中. 下面的示例使用相对路径

ADD test.txt relativeDir/

--- 此示例使用绝对路径

ADD test.txt /absoluteDir/

--- 添加包含特殊字符(如 + 和 +)的文件或目录时,需要按照 Golang 规则转义这些路径,以防止它们被视为匹配的模式.例如,若要添加名为 arr{0}.txt的文件,请使用以下内容;

ADD arr[[]0].txt /mydir/



============================================================

```

<!-- #### VIM End  ####  -->

#### 权限处理规则

<!-- ######## VIM Start########  -->

```vim
--- 所有新文件和目录都使用 UID 和 GID 0 创建,除非可选的 --chown 标志指定给定用户名、组名或 UID/GID 组合,以请求对添加的内容的特定所有权

--- --chown 标志的格式允许用户名和组名字符串或任意组合中的直接整数 UID 和 GID



--- 提供不带组名的用户名或没有 GID 的 UID 将使用与 GID 相同的数字 UID


----  如果提供了用户名或组名,则容器的根文件系统 /etc/passwd 和 /etc/组文件将分别用于执行从名称到整数 UID 或 GID 的转换

---->以下示例显示 --chown 标志的有效定义：

ADD --chown=55:mygroup files* /somedir/
ADD --chown=bin files* /somedir/
ADD --chown=1 files* /somedir/
ADD --chown=10:11 files* /somedir/



----- 如果容器根文件系统不包含 /etc/passwd 或 /etc/组文件,并且用户或组名称在 --chown 标志中使用,则 BUILD 将在 ADD 操作中失败

---- 使用数字 ID 不需要查找,也不需要依赖于容器根文件系统内容.


--- 如果 src 是远程文件 URL,则目标的权限为 600.如果正在检索的远程文件具有 HTTP 上次修改标头,则来自该标头的时间戳将用于设置目标文件上的 mtime

---但是,与在 ADD 期间处理的任何其他文件一样,mtime 将不包括在确定文件是否已更改以及缓存是否应更新时.



```

<!-- ######## VIM End  ########  -->

#### .dockerignore 匹配规则

<!-- ######## VIM Start########  -->

```vim
- # comment   这个会被当做注释处理,忽略

---

- */temp*     会忽略匹配文件名的文件以及 这个文件的直接子目录
-  /somedir/temporary.txt is excluded


---

- */*/temp*   会忽略匹配文件名的文件以及 这个文件所在文件夹的上两级
- /somedir/subdir/temporary.txt is excluded.

---

- temp? 一个文件就是包含其任意的一个字母的文件
- 比如说 /tempa   /tempb

---

- **/*.go
- 排除所有以 .go 结尾的文件
- 包括根目录下的

---

!README*.md
README-secret.md

- 第一句匹配到了README-secret.md,所以应该是不被排除(前面有！)
- 但是第二句指明了README-secret.md,那么规则被改写,最后还是被排除了


---

- dockerfile不会被ADD/COPY 传入镜像中

---
```

<!-- ######## VIM End  ########  -->

### RUN 和 CMD 的小提示

#### 环境变量处理

- 当使用 exec 窗体并直接执行 shell 时（如 shell 窗体中的情况）时，执行环境变量扩展的外壳不是 docker.

- RUN [ "echo", "$HOME" ] 执行的时候 HOME 不会被替换

- RUN [ "sh", "-c", "echo $HOME" ] 这样会被 sh 里面的 HOME 环境变量替换

- 变量替换存在于 shell 中

---

<!-- #### VIM Start####  -->

```vim
如果使用 CMD的shell形式,默认执行的终端为 /bin/sh -c:

FROM ubuntu
CMD echo "This is a test." | wc -

如果你想要无shell运行程序,那么就必须使用json数组的形式(第二种形式),并给定运行程序的确切目录
这种数组的形式一般推荐在CMD里面使用. 所有参数都要单独的传入到数组里面

FROM ubuntu
CMD ["/usr/bin/wc","--help"]
```

<!-- #### VIM End  ####  -->

#### Json 字串

- 在 JSON 窗体中，必须转义反斜杠。这在反斜杠是路径分隔符的 Windows 上尤其相关。否则，由于 JSON 无效，以下行将被视为 shell 窗体，并且以意外的方式失败：

- 执行窗体解析为 JSON 数组，这意味着您必须围绕非单引号单词 (') 使用双引号 (').

<!-- #### VIM Start####  -->

```vim
RUN ["c:\windows\system32\tasklist.exe"]


=====  The correct syntax for this example is:


RUN ["c:\\windows\\system32\\tasklist.exe"]
```

<!-- #### VIM End  ####  -->
