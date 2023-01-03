---
date:  2021-02-28  19:10:25
updated:  2021-02-28  19:10:25
title:  搭建 FTP 服务器
index_img: /gallery/2021-08-23-19-18-20.png
tags:
  - Linux

categories:
  - Linux
---
## FTP ON LINUX

### 安装 vsftpd

```bash
sudo apt install vsftpd
```

### 配置 vsftpd

```bash
sudo rm /etc/vsftpd.conf
sudo vim /etc/vsftpd.conf
```

---

```vim
listen=YES
anonymous_enable=NO
local_enable=YES
write_enable=YES
dirmessage_enable=YES
use_localtime=YES
xferlog_enable=YES
connect_from_port_20=YES
chown_uploads=YES
chown_username=yangyiming
chroot_local_user=YES
secure_chroot_dir=/var/run/vsftpd/empty
pam_service_name=vsftpd
rsa_cert_file=/etc/ssl/certs/ssl-cert-snakeoil.pem
rsa_private_key_file=/etc/ssl/private/ssl-cert-snakeoil.key
ssl_enable=NO
local_root=/
```

### vsftpd 控制

```bash
sudo service vsftpd status
sudo service vsftpd start
sudo service vsftpd restart
sudo service vsftpd stop
```
