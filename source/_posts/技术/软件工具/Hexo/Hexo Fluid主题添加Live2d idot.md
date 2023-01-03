---
title: Hexo Fluid主题添加 Live2d 人偶
date: 2021-08-23 14:09:27
updated: 2021-08-23 14:09:27
index_img: /gallery/2021-08-23-17-55-48.png
tags:
  - HEXO
categories:
  - HEXO
---


# Hexo Fluid主题添加 Live2d-人偶

- 这里采用的是 [Live2D Widget](https://github.com/stevenjoezhang/live2d-widget)

- 这里需要注意的就是，默认加载成功的看板娘为

![不是自定义的](/gallery/2021-08-23-14-23-00.png)


## 基本加载

1. 下载对应的发行版本，并解压


![解压后的文件](/gallery/2021-08-23-14-25-06.png)

2. 然后修改 autoload.js

![autoload.js](/gallery/2021-08-23-14-25-36.png)

3. 在里面修改 live2d_path 这里决定的启动项目的autoload修改生效

![live2d_path](/gallery/2021-08-23-14-30-05.png)


4. 在里面修改 cdnPath ,这里决定着模型的加载

![cdnPath](/gallery/2021-08-23-14-29-37.png)

5. 修改head ，这里需要找到hexo主题的 head.ejs文件

![head.ejs 位置](/gallery/2021-08-23-14-32-55.png)

6. 合适的位置添加下面代码

```html
    <!--自定义看板娘-->
  <script src="/live2d-widget/autoload.js"></script>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/font-awesome/css/font-awesome.min.css">
```

![head.ejs 添加代码](/gallery/2021-08-23-14-35-23.png)

7. 把上述文件夹放到 主题文件夹的 source 中


## 关于 cdnPath 的修改

- 可以直接按照[这个仓库](https://github.com/fghrsh/live2d_api)来自己搭建一个api

- 或者可以和我一样操作（我这里可以添加更多的模型但是需要按照api的格式进行修改）

1. 在主题的 source 文件夹里面创建一个model-api文件夹，然后再到里面创建一个model文件夹
2. 然后在里面创建 model_list.json
3. 根据里面的子文件夹创建写入对应的json内容
4. 再创建一个undefined文件夹来保证能够正常的现实模型文件
5. 模型文件必须名字要改为index.json

![index.json](/gallery/2021-08-23-14-48-34.png)

7. 最后cdnPath就可以修改为 urltoSite/model-api