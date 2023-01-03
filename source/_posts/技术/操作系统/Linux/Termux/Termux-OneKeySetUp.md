---
date: 2022-04-12 16:20:06
title: Termx-OneSetUp
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Termux

categories:
  - Termux
---

current is brower sync .
termux supply hexoserver and sync signal emitter...
termux also can provide sql service ..lets checkout ..
# 直接配置代码 

```bash
#!/data/data/com.termux/files/usr/bin/bash
PREFIX=/data/data/com.termux/files/usr
ZHSRC_FILE=$PREFIX/etc/zshrc
### SYSTEM
echo SYSTEM
termux-change-repo # change mirror
pkg install curl -y
pkg install wget -y
pkg install openssh -y
pkg install openssl-tool -y
pkg install termux-auth -y
pkg install vim -y
cp -r $PREFIX/lib/openssl-1.1/* $PREFIX/lib/
passwd
echo SYSTEM
### SYSTEM

## PROGRAMING
pkg install python -y
pkg install golang -y
pkg install clang -y
## PROGRAMING


### zsh configrue
echo zsh configrue
pkg install unzip zsh -y
unzip -d ~/ ./zsh-easy-config.zip
cat ~/zshrc >> $ZHSRC_FILE
cd ~/.oh-my-zsh/custom/plugins/autojump/
python3 install.py
chsh -s $(command -v zsh)
echo 'alias gostorate=cd /storage/emulated/0/' >> $ZHSRC_FILE
echo "run (source ~/.zshrc) if nessary"
echo zsh configrue
### zsh configrue



### DATA BASE
echo start configure data base
pkg install postgresql -y
mkdir -p $PREFIX/var/lib/postgresql
initdb $PREFIX/var/lib/postgresql
echo 'alias startdb=pg_ctl -D $PREFIX/var/lib/postgresql start' >> $ZHSRC_FILE
echo 'alias stopdb=pg_ctl -D $PREFIX/var/lib/postgresql stop' >> $ZHSRC_FILE
echo 'alias opendb=psql mydb' >> $ZHSRC_FILE
createuser --superuser --pwprompt kusime
pg_ctl -D $PREFIX/var/lib/postgresql start #start
createdb mydb
echo data base configure done
### DATA BASE

### NODEJS
echo NODEJS
pkg install nodejs  -y
npm install hexo-cli -g
npm install hexo-server -g
echo NODEJS
### NODEJS

### BLOG CONFUGURE
echo  BLOG CONFUGURE
pkg install git -y
cd ~ # return home dir
git clone ssh://kusime@192.168.1.218/home/kusime/Desktop/Blog/My-Blog
git clone ssh://kusime@192.168.1.218/home/kusime/Desktop/Blog/My-Daily
echo BLOG CONFUGURE
### BLOG CONFUGURE



##NOTE
echo PLEASE Configure Battery in 

echo 'App & service > App launch > termux > Manage manually'

echo Congiguration done
##NOTE

```

上述配置需要 压缩包也在哈。目前架构使用的是 Aidlux 搭建code－server然后使用Termux来搭建HEXOSERVER的环境，因为Aid以及终端和本机的链接状态不是特别的好，本质上Aid一切都是一个web页面，然后他自定义以及终端，包，等等的支持没有termux好。所以使用的termux来进行数据库搭建什么的，甚至还可以跑quem等等所以直接一些灵活性要求比较高的话。还是正常使用termux就好了，然后脚本什么的和Linux的环境更加的像。

然后我使用转移的方式部署了相关的一些部署脚本全部放到了termux上面，然后比如这文章部署就直接到