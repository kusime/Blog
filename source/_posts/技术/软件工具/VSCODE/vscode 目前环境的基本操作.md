---
date:  2021-03-17 21:45:25
updated:  2021-03-17 21:45:25
title:  vscode 目前的快捷键的操作
index_img: /gallery/2021-08-23-18-47-02.png
tags:
  - Vscode

categories:
  - Vscode
---


# 基本操作

- ctrl + h 查找替换
- ctrl + j 查找
- ctrl + shift + v 黏贴图片
- ctrl + F1 创建目录
- ctrl + F2 更新目录
- ctrl + shift + i 格式化代码
- ctrl + shift + c 创建任务列表
- ctrl + shift + enter 完成任务列表

# 自定义快捷补全

- 设置->用户代码片段
- 然后可以就在这里搜索对应的语言
- 然后修改对应的东西
  - 在里面要注意 json 文件解析语法
  - 有时候要使用\来取消转义

![error_loading](/gallery/2021-03-17-21-52-41.png)

## 默认模板

```json
//  这里记得格式化一下这个代码
 "这个补全的名字":{
 "prefix": "触发补全的名字",
 "body": [
 "补全的内容",
 "如果是多行注意用数组的形式补全",
 "$0 是最后光标停下来的位置 $1 是第一次tab停留的位置 $2 是第二次tab停留的位置",
 ], // << 注意这里有个小逗号
 "description": "这个代码片段的描述"
 }
```

## 手动开启 markdown 的用户补全

- 设置 -> settings.json
- 然后添加下面的代码
- 然后我自己的测试，其他语言就不需要这样单独的测试

```json
    "[markdown]": {
        "editor.wordWrap": "on",
        "editor.quickSuggestions": true,
        "editor.defaultFormatter": "esbenp.prettier-vscode"
    },
```
