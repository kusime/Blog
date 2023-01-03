---
title: Nginx 压缩图片
date: 2021-08-26 18:55:57
updated: 2021-08-26 18:55:57
index_img: /gallery/2021-08-23-18-03-58.png
tags:
  - Nginx
categories:
  - Nginx

---


{% note success %}

如果已经有模块了直接配置就好了

{% endnote %}

# 安装必要模块

```conf
./configure arguments: --prefix=/usr/local/nginx --with-http_image_filter_module # 我用apt安装的ngixn好像自带emm
```







# 配置 location 块

```conf
	location /gallery {
		
		image_filter resize 550 550 ;
           		 image_filter_jpeg_quality 75;
           		 image_filter_buffer 100M;
	}
```
