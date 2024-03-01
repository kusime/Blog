---
title: Ubuntu Setup
date: 2024-02-29 20:16:23
tags:
  - Linux
categories:
  - Linux
---

# Ubuntu Setup

```bash
sudo apt update

# install usable packages
sudo apt install telnet ufw 
sudo systemctl enable ufw
sudo systemctl start ufw


# install docker
sudo apt install -y docker.io
sudo systemctl enable docker
sudo systemctl start docker

# japan sources
sudo sed -i 's/us\.archive\.ubuntu\.com/jp.archive.ubuntu.com/g' /etc/apt/sources.list
sudo apt update


```


# zsh Setup

```bash
# install zsh 
sudo apt install zsh

```

```bash
# install script
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"


```

## install plugins

`vim .zshrc`

```vim
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
plugins=(git docker python autojump  zsh-autosuggestions)
```

```bash
# install autojump

git clone https://github.com/wting/autojump.git
cd autojump
python install.py

```

`vim .zshrc`

```vim
[[ -s /root/.autojump/etc/profile.d/autojump.sh ]] && source /root/.autojump/etc/profile.d/autojump.sh
autoload -U compinit && compinit -u
```

## alias

```vim
alias zshconfig="vim ~/.zshrc"

alias apply='source ~/.zshrc' 
```


