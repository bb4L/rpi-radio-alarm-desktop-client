# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.158.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.15, 1.14
ARG VARIANT="1.15"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi


RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN apt-get update
RUN apt-get install -y build-essential curl file git libgtk-3-dev libwebkit2gtk-4.0-dev build-essential binutils-mingw-w64

RUN go get -u github.com/guark/guark/cmd/guark