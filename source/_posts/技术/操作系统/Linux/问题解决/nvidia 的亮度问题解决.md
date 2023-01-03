---
date: 2021-03-18 16:16:07
updated: 2021-03-18 16:16:07
title: linux NVIDIA卡的亮度控制
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---

# 解决方法

```bash
sudo vim /usr/share/X11/xorg.conf.d/10-nvidia.conf

---复制下面的代码到上面的文件中--
Section "OutputClass"
    Identifier "nvidia"
    MatchDriver "nvidia-drm"
    Driver "nvidia"
    Option "AllowEmptyInitialConfiguration"
    ModulePath "/usr/lib/x86_64-linux-gnu/nvidia/xorg"
EndSection

Section "Device"
        Identifier "Device0"
        Driver "nvidia"
        VendorName "NVIDIA Corporation"
        Option "RegistryDwords" "EnableBrightnessControl=1"
        Option "NoLogo" "True"
EndSection
--------
```