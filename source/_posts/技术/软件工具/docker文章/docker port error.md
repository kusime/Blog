---
title: Docker Port Error 
date: 2024-03-01 18:20:28
tags:
  - Docker
categories:
  - Docker
---

# Docker Port Error 

[DockerがErrorで起動しない(Ports are not available: listen ～An attempt was made to access a socket in a way forbidden by its access permissions.) - Qiita](https://qiita.com/Quantum/items/8891fa9c94d03b388555)


Administrator

```cmd
Restart-Service -Name winnat
```


```cmd
netstat -ano | Select-String ":port"

taskkill /F /PID pid
```