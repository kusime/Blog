---
title: OVER40779122 VR Oculus 报错的解决策略
index_img: /gallery/2021-11-23-14-48-09.png
date: 2023-01-10 19:38:42
tags:
  - Saying
categories:
  - Saying
---

# OVER40779122 VR Oculus 报错的解决策略

{% note success %}
[OVER40779122 VR Oculus 报错的解决策略](https://blog.csdn.net/qq_23369807/article/details/106224314)
{% endnote %}

改host,这个是中国的DNS污染

`C:\Windows\System32\drivers\etc`


```vim
157.240.3.49 http://graph.oculus.com
157.240.3.49 http://www2.oculus.com
157.240.8.49 http://scontent.oculuscdn.com
157.240.8.49 http://securecdn.oculus.com
157.240.3.49 http://www.oculus.com
157.240.7.49 http://id.oculus.com
157.240.7.49 http://secure.oculus.com
157.240.3.29 http://static.xx.fbcdn.net
```


