---
title: plt 绘画出现问题的解决
date: 2023-02-20 18:41:17
tags:
  - AI
categories:
  - AI
---

#  plt 绘画出现问题的解决


```vim
import os
os.environ["KMP_DUPLICATE_LIB_OK"]="TRUE"
```

# 或者

把这个k-v值加到系统环境变量里面