---
title: Docker Email Server Setup
date: 2024-02-29 22:31:33
tags:
  - Docker
categories:
  - Docker
---

# Docker Email Server Setup


## MTA-STS


mta-sts.FQDN will neet to server an txt file

```vim
_mta-sts.domain.com IN TXT "v=STSv1; id=20240229"
```

```vim
version: STSv1
mode: enforce
mx: mail.domain.com
max_age: 86400
```

## mailserver.env

[mailserver.env](/code/env/mailserver.env)


## DNS Setup

[DKIM, DMARC & SPF - Docker Mailserver](https://docker-mailserver.github.io/docker-mailserver/latest/config/best-practices/dkim_dmarc_spf/)

### DMARC
```vim
v=DMARC1; p=quarantine; rua=mailto:username@domain.com; ruf=mailto:username@domain.com; sp=quarantine; ri=86400
```

### SPF

```vim
v=spf1 a mx ip4:mailserver_address ~all
```

### DKIM

[Network Tools: DNS,IP,Email](https://mxtoolbox.com/SuperTool.aspx?action=dkim%3adomain.com%3amail&run=toolpage)

```bash
docker exec -it mailserver setup config dkim

docker exec -it mailserver setup config dkim help
```

```vim
v=DKIM1; h=sha256; k=rsa; p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAstXd80VIfwtLFSeAY5M4jmMMRj+75U1ExLar+3b6KSeTZB9lQ1QNvlycMhJfQqgdxTgIHFiVLZmf9iddpcwZYX/Wr7LhBh4sXEBaknGsupMD6wJZBshKbZaoy0ovBi4Yf0YZFHPjPRsV9UFt8AKVuGKmckuY97hQfPD8rpB4lVZw4Ffb58zsOQRHs7n2ZvroLwhfgoX+OSLcGKr1cHeJLwJf4OLmTGKjfzXptUI0Vdyd7HzFOpQ1fKhvsz0492XJxpZaUs5GsrukxUwE3n7P5Lxnu9zjCu1hPbJEfmzzuM4ZvMC1b6WyX3zyj4Yjzvrpq7YKS7roSHPtgdgbx7K/WQIDAQAB
```


## Email Relay

if your vps support the 25 port outgoing , that fine 

[Mail Forwarding | Relay Hosts - Docker Mailserver](https://docker-mailserver.github.io/docker-mailserver/latest/config/advanced/mail-forwarding/relay-hosts/)


# TLS Setup

`compose.yml`

```vim
services:
  mailserver:
    image: ghcr.io/docker-mailserver/docker-mailserver:latest
    container_name: mailserver
    # Provide the FQDN of your mail server here (Your DNS MX record should point to this value)
    hostname: FQDN
    env_file: mailserver.env
    # More information about the mail-server ports:
    # https://docker-mailserver.github.io/docker-mailserver/latest/config/security/understanding-the-ports/
    # To avoid conflicts with yaml base-60 float, DO NOT remove the quotation marks.
    ports:
      - "25:25"    # SMTP  (explicit TLS => STARTTLS, Authentication is DISABLED => use port 465/587 instead)
      - "143:143"  # IMAP4 (explicit TLS => STARTTLS)
      - "465:465"  # ESMTP (implicit TLS)
      - "587:587"  # ESMTP (explicit TLS => STARTTLS)
      - "993:993"  # IMAP4 (implicit TLS)
    volumes:
      - ./docker-data/dms/mail-data/:/var/mail/
      - ./docker-data/dms/mail-state/:/var/mail-state/
      - ./docker-data/dms/mail-logs/:/var/log/mail/
      - ./docker-data/dms/config/:/tmp/docker-mailserver/
      - /etc/localtime:/etc/localtime:ro
      # TLS configuration
      - ./docker-data/certbot/certs/:/etc/letsencrypt
    restart: always
    stop_grace_period: 1m
    # Uncomment if using `ENABLE_FAIL2BAN=1`:
    # cap_add:
    #   - NET_ADMIN
    healthcheck:
      test: "ss --listening --tcp | grep -P 'LISTEN.+:smtp' || exit 1"
      timeout: 3s
      retries: 0
```

## ACME get certs


`get certs`
```bash
docker run --rm -it \
  -v "${PWD}/docker-data/certbot/certs/:/etc/letsencrypt/" \
  -v "${PWD}/docker-data/certbot/logs/:/var/log/letsencrypt/" \
  -p 80:80 \
  certbot/certbot certonly --standalone -d domail.name

```

`renew certs`

```bash
docker run --rm -it \
  -v "${PWD}/docker-data/certbot/certs/:/etc/letsencrypt/" \
  -v "${PWD}/docker-data/certbot/logs/:/var/log/letsencrypt/" \
  -p 80:80 \
  -p 443:443 \
  certbot/certbot renew
```



# FAIL2BAN

```bash
docker exec mailserver setup fail2ban status 

docker exec mailserver setup fail2ban [<ban|unban> <IP>]
```


# Aliases

for example , if sending error , the google mail will info this postmaster@domain.com for some reporting or the detail

```bash
docker exec -ti mailserver setup alias add postmaster@domain.com username@domain.com
```


# UserManagement

```bash
docker exec -ti mailserver setup email add username@domain.com
```