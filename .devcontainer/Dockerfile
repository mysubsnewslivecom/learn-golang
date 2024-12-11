ARG BASE_IMAGE_NAME=docker.io/golang
ARG BASE_IMAGE_TAG=1.23.4-bookworm

FROM ${BASE_IMAGE_NAME}:${BASE_IMAGE_TAG}

ARG USERNAME=linux
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN apt-get update \
    && apt-get -y install --no-install-recommends \
    jq \
    sudo \
    build-essential \
    vim \
    apt-transport-https \
    apt-utils \
    ca-certificates \
    curl \
    locales \
    lsb-release \
    openssh-client \
    bash-completion \
    sqlite3 \
    fzf \
    unzip \
    fontconfig \
    && update-ca-certificates \
    && apt-get autoremove -yqq --purge \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -s /usr/bin/bash -m $USERNAME \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

RUN DEBIAN_FRONTEND=noninteractive \
    sh -c "$(curl -sS https://starship.rs/install.sh)" -- -y

USER $USERNAME

WORKDIR /apps

ENV \
    PATH="/home/$USERNAME/.local/bin:$PATH" \
    DEBIAN_FRONTEND=noninteractive \
    TZ=UTC \
    LANGUAGE=C.UTF-8 \
    LANG=C.UTF-8 \
    LC_ALL=C.UTF-8 \
    LC_CTYPE=C.UTF-8 \
    LC_MESSAGES=C.UTF-8 \
    LD_LIBRARY_PATH=/usr/local/lib \
    PIP_CACHE_DIR=/tmp/.cache/pip \
    HOME=/home/${USERNAME}

ADD https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh $HOME/git-prompt.sh

RUN echo '[[ -f ~/.local/bin/activate ]] && source ~/.local/bin/activate '| tee -a ~/.bashrc \
    && echo '[[ -f /usr/share/bash-completion/completions/git ]] && source /usr/share/bash-completion/completions/git' | tee -a ~/.bashrc \
    && echo '[[ -f \$HOME/git-prompt.sh ]] && source \$HOME/git-prompt.sh' | tee -a ~/.bashrc \
    && sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin

RUN echo 'eval "$(starship init bash)"'| tee -a ~/.bashrc \
    && echo 'eval "$(task --completion bash)"'| tee -a ~/.bashrc \
    && mkdir -p ~/.config ~/.local/share/fonts  \
    && wget https://github.com/ryanoasis/nerd-fonts/releases/download/v3.2.1/FiraCode.zip -P /tmp \
    && unzip /tmp/FiraCode.zip -d ~/.local/share/fonts \
    && fc-cache -fv

COPY --chown=$USERNAME:$USERNAME starship.toml /home/${USERNAME}/.config/starship.toml

ARG CI_COMMIT_SHA
ARG CI_COMMIT_REF_NAME
ARG CI_COMMIT_TAG
ARG CI_PROJECT_URL
ARG CI_PIPELINE_URL
ARG GITLAB_USER_EMAIL
ARG GITLAB_USER_LOGIN
ARG BASE_IMAGE_NAME
ARG BASE_IMAGE_TAG
ARG CI_PIPELINE_CREATED_AT
LABEL org.opencontainers.image.source=$CI_PROJECT_URL \
    org.opencontainers.image.created=$CI_PIPELINE_CREATED_AT \
    org.opencontainers.image.authors=$GITLAB_USER_EMAIL \
    org.opencontainers.image.url=$CI_PIPELINE_URL \
    org.opencontainers.image.documentation=$CI_PROJECT_URL/-/wikis/home \
    org.opencontainers.image.version=$CI_COMMIT_TAG \
    org.opencontainers.image.revision=$CI_COMMIT_SHA \
    org.opencontainers.image.vendor=$GITLAB_USER_LOGIN \
    org.opencontainers.image.licenses="MIT" \
    org.opencontainers.image.ref.name=$CI_COMMIT_REF_NAME \
    org.opencontainers.image.title="Golang" \
    org.opencontainers.image.description="Golang image for devcontainer based on ${BASE_IMAGE_NAME}:${BASE_IMAGE_TAG}"
