---
date:  2021-03-14 10:10:13
updated:  2021-03-14 10:10:13
title:  vscode 基本配置
index_img: /gallery/2021-08-23-18-47-02.png
tags:
  - Vscode

categories:
  - Vscode
---


# linux 安装 vscode

```bash
sudo apt install firefox -y
firefox https://code.visualstudio.com/Download
#在弹出的页面选择linux版本，kali，ubuntu就选择deb格式就好
cd [你的下载路径]
sudo dpkg -i [你的安装包名字]
```

# 基本语言环境配置

## 安装中文

1. 打开 vscode
2. 在侧边栏找到 ![error_loading](/gallery/2021-03-14-10-17-28.png)
3. 然后在搜索框查找 Chinese
4. 找到排名第一的，直接点击安装就好
5. 需要点击一下重启 vscode

## vscode 配置 Python 环境

1. 打开 vscode
2. 在侧边栏找到 ![error_loading](/gallery/2021-03-14-10-17-28.png)
3. 然后在搜索框查找 python
4. 找到排名第一的，直接点击安装就好

## 安装一键运行插件

1. 打开 vscode
2. 在侧边栏找到 ![error_loading](/gallery/2021-03-14-10-17-28.png)
3. 然后在搜索框查找 code runner
4. 找到排名第一的，直接点击安装就好
5. 然后你会看见一个小箭头 ![error_loading](/gallery/2021-03-14-10-22-28.png)
6. 把程序写好后点击这个小箭头就可以直接运行（Python 可以）

## vscode 配置 C/C++环境

1. 打开 vscode
2. 在侧边栏找到 ![error_loading](/gallery/2021-03-14-10-17-28.png)
3. 然后在搜索框查找 C/C++
4. 找到排名第一的，直接点击安装就好
5. 在 linux 自带了 gcc 所以不需要配置环境变量
6. 结合上面的一键运行插件就可以直接运行 C 语言的代码了

### 美化插件推荐

#### Better Comment

- 美好注释

#### Material Icon Theme

- 图标主题

#### Winter is Coming Theme

- 蓝色高冷主题

#### Bracket Pair Colorizer 2

- 彩虹括号

### 实用插件推荐

#### shellman

- （主要用于编写 shell 脚本的）

#### Docker

- (结合 docker 使用的)

# [settings.json 备份](/posts/code/json/settings.json)
