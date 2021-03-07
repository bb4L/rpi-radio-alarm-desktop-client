# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.158.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.15, 1.14
ARG VARIANT="1.15"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# RUN npm install
# RUN yarn global add @vue/cli

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    # && apt-get -y install --no-install-recommends libgtk-3-dev libwebkit2gtk-4.0-dev build-essential
# RUN apt-get update && apt-get -y install libgtk-3-dev libwebkit2gtk-4.0-dev build-essential binutils-mingw-w64
# RUN su root -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN apt-get update
RUN apt-get install -y build-essential curl file git libgtk-3-dev libwebkit2gtk-4.0-dev build-essential binutils-mingw-w64
# RUN git clone https://github.com/Homebrew/brew ~/.linuxbrew/Homebrew && mkdir ~/.linuxbrew/bin && ln -s ~/.linuxbrew/Homebrew/bin/brew ~/.linuxbrew/bin && eval $(~/.linuxbrew/bin/brew shellenv)

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>
# WORKDIR /workspaces/rpi-radio-alarm-desktop-client
# WORKDIR /workspaces/rpi-radio-alarm-desktop-client/
RUN go get -u github.com/guark/guark/cmd/guark
# RUN go install

# WORKDIR /workspaces/rpi-radio-alarm-desktop-client

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
# RUN su root -c "source /usr/local/share/nvm/nvm.sh && npm install" 2>&1
# RUN npm install
# RUN npm install -g @vue/cli @vue/cli-service 
# RUN npm install -g @vue/cli-service


# RUN /usr/local/share/nvm/nvm.sh && npm install 2>&1