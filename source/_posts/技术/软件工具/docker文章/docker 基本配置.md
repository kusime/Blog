---

date: 2021-02-17 00:30:32
updated: 2021-02-17 00:30:32
title: docker 安装和配置
index_img: /gallery/2021-08-23-21-34-28.png
tags: 
   - Docker

categories:
   -  Docker

---


# Docker 的安装的配置

### 安装

```bash
sudo apt install git curl
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
```

### 切换阿里镜像源

```bash
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://ebrsf6fn.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
```

### 添加用户到 dokcer 组

```bash
 sudo usermod -aG docker kusime
 reboot
```

### 一些操作

1. 登录阿里云 Docker Registry
   $ sudo docker login --username=yangyiminghaking registry.cn-hangzhou.aliyuncs.com
   用于登录的用户名为阿里云账号全名，密码为开通服务时设置的密码。

2. 从 Registry 中拉取镜像
   $ sudo docker pull registry.cn-hangzhou.aliyuncs.com/kusime/kusime:[镜像版本号]
3. 将镜像推送到 Registry
   $ sudo docker login --username=yangyiminghaking registry.cn-hangzhou.aliyuncs.com
   $ sudo docker tag [ImageId] registry.cn-hangzhou.aliyuncs.com/kusime/kusime:[镜像版本号]
   $ sudo docker push registry.cn-hangzhou.aliyuncs.com/kusime/kusime:[镜像版本号]
   请根据实际镜像信息替换示例中的[ImageId]和[镜像版本号]参数。
