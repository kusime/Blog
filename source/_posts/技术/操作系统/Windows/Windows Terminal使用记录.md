---
date: 2021-06-04 22:10:26
updated: 2021-06-04 22:10:26
index_img: /gallery/2021-08-23-18-42-52.png
title: Windows Terminal 的使用
tags:
  - Windows
categories:
  - Windows
---

# 基本配置备份

- 因为最近更新了可以直接在图形界面进行设置
- 基本只要知道基本应用设置好 PATH 和命令运行的参数
- 就这样就可以直接或者间接的配置
- 但是美化的话,为了简单还是直接使用现有的配置文件就好了
- [基本配置](/posts/code/json/settingsOfWINODWSTERMINAL.json)

# 基本试图操作

- 对于已经打开了的选项卡可以直接(_冲突了,不推荐使用_)

![error_loading](/gallery/2021-06-04-22-47-35.png)

- 对于在全局切换

  - 使用 `Ctrl+Shift+Tab` 的时候按住然后使用上下来选择也是可以的
  - 使用 `Ctrl+Shift+W` 这个就是关闭切割出来的小窗口
  - 使用 `Alt+F4` 是直接关闭所有的会话
  - 使用 `Ctrl+Alt+W` 是关闭一个小标签(前提就是标签页要大于 1)
    ![error_loading](/gallery/2021-06-04-22-44-35.png)

- 拆分窗口

![图片描述](/gallery/2021-06-04-23-18-52.png)



![图片描述](/gallery/2021-06-04-23-25-43.png)


![error_loading](/gallery/2021-06-04-22-54-43.png)

- 关于焦点切换
  - 这里只有对于已经拆分了小窗口的会话才有意义

![error_loading](/gallery/2021-06-04-22-48-50.png)

![error_loading](/gallery/2021-06-04-22-56-16.png)

- 代替滚轮翻页(远离鼠标好耶!)
  - 一页就是方向键上面六个键最靠边上的那两个,使用的时候要注意按住

![error_loading](/gallery/2021-06-04-22-49-29.png)

- 直接打开指定的标签页
  - index 为 0 的不推荐切换
  - 展开小箭头可以看见对应标签的打开快捷方式

![error_loading](/gallery/2021-06-04-23-48-19.png)

![error_loading](/gallery/2021-06-04-22-50-02.png)

- 调整窗口大小 - 这个等同于直接 按住 Ctrl+滚轮

![error_loading](/gallery/2021-06-04-22-50-25.png)

# 对于功能性的操作

- _查找_ `Ctrl+Shift+F`

  - 这个查找操作是直接在终端进行操作
  - 这里和 VsCode 的通用

- _打开设置_ `Ctrl+,`

  - 这里和 VsCode 的通用

- 重复打开窗口(当前焦点) `Ctrl+Shift+D`
