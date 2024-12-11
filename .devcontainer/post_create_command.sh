#!/usr/bin/env bash

# sudo DEBIAN_FRONTEND=noninteractive apt update && sudo apt upgrade

# sudo apt install -y netcat iputils-ping nginx certbot openssh-client \
#     --install-recommends \
#     && sudo apt autoremove -y \
#     && sudo rm -rvf /var/lib/apt/lists/* \
#     && sudo apt -y clean

go install -v golang.org/x/tools/gopls@latest \
  && go mod download \
  && go mod tidy \
  && go mod vendor
