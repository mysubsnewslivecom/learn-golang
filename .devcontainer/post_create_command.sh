#!/bin/bash

source $HOME/common_functions.sh

# log_info "Sleeping for 60s"

# spinner=(
# "-"
# "\\"
# "|"
# "/"
# )

# max=$((SECONDS + 10))

# while [[ ${SECONDS} -le ${max} ]]
# do
#     for item in ${spinner[*]}
#     do
#         echo -en "\r$item"
#         sleep 0.1
#         echo -en "\r              \r"
#     done
# done

# spin[0]="-"
# spin[1]="\\"
# spin[2]="|"
# spin[3]="/"

# ping -i 10 -c 2 -w 2 localhost 2> /dev/null &
# pid=$!
# trap "kill $pid 2> /dev/null" EXIT

# # While 'speedtest' is running:
# while kill -0 $pid 2> /dev/null; do
#     for i in "${spin[@]}"
#     do
#         echo -ne "\b$i"
#         sleep 0.1
#         echo -en "\r              \r"
#     done
# done
# log_error "Post create command"

sudo DEBIAN_FRONTEND=noninteractive apt update && sudo apt upgrade

sudo apt install -y netcat iputils-ping nginx certbot openssh-client \
    --install-recommends \
    && sudo apt autoremove -y \
    && sudo rm -rvf /var/lib/apt/lists/* \
    && sudo apt -y clean
