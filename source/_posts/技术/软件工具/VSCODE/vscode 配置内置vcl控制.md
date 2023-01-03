---
date:  2021-03-17 22:20:39
updated:  2021-03-17 22:20:39
title:  vscode 中控制 vlc 播放器
index_img: /gallery/2021-08-23-18-47-02.png
tags:
  - Vscode

categories:
  - Vscode
---

# 在 vscode 中的操作

- 下载 Media Player 插件 ![error_loading](/gallery/2021-03-17-22-23-30.png)
- 然后在 setting.json 文件插入下面的代码

```json
    "media.player": {
        "players": [
            {
                "name": "vlc player",
                "type": "vlc",
                "password": "admin",
                "port": 8080
            }
        ]
    }
```

# 在 vlc 的操作

- 打开 web telnet 这些界面

![error_loading](/gallery/2021-03-17-22-26-42.png)

- 然后进入 界面 ->> Lua

![error_loading](/gallery/2021-03-17-22-27-53.png)

- 然后在 passwd 里面写入 admin 就好了

![error_loading](/gallery/2021-03-17-22-28-42.png)

- 然后保持 vlc 的打开，然后可以事前播放一个文件夹音乐

# 最后实现的效果图

![error_loading](/gallery/2021-03-17-22-30-14.png)

- 然后这个插件比![error_loading](/gallery/2021-03-17-22-31-36.png)这个点
- 上面这个插件也可以实现控制，也提供基本的操作
- 但是上面的不可以正确的现实插件，而且也不能控制，查看播放列表
- 但是这个插件有可视化的配置，不用直接编辑 setting.json 文件
