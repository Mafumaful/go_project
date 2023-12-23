FROM golang:1.16 as base

FROM base as dev

# >>>>>>>>>>>>>>>>>>>>>>>>> user settitngs >>>>>>>>>>>>>>>>>>>>>>>>>
# Update and install sudo (if not already installed in your base image)
ENV USER=docker
RUN apt-get update && apt-get install -y sudo zsh wget git

# Create a new user 'myuser' and add it to the 'sudo' group
RUN useradd -m ${USER} && echo "${USER}:${USER}" | chpasswd && adduser ${USER} sudo

# Give '${USER}' passwordless sudo privileges
RUN echo "${USER} ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

# Set zsh as the default shell for the user
RUN chsh -s /bin/zsh ${USER}

# Switch to the new user
USER ${USER}:${USER}
# <<<<<<<<<<<<<<<<<<<<<<<<< user settitngs <<<<<<<<<<<<<<<<<<<<<<<<<


# >>>>>>>>>>>>>>>>>>>>>>>>> golangci-lint >>>>>>>>>>>>>>>>>>>>>>>>>
RUN mkdir /home/${USER}/code
WORKDIR /home/${USER}/code

# golangci-lint
RUN go env -w GO111MODULE=auto\
    && go mod init v2\
    && go mod tidy
# <<<<<<<<<<<<<<<<<<<<<<<<< golangci-lint <<<<<<<<<<<<<<<<<<<<<<<<<

# >>>>>>>>>>>>>>>>>>>>>>>>> oh-my-zsh >>>>>>>>>>>>>>>>>>>>>>>>>
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

WORKDIR /home/${USER}/code

# set the default shell to zsh
CMD [ "zsh" ]
# <<<<<<<<<<<<<<<<<<<<<<<<< oh-my-zsh <<<<<<<<<<<<<<<<<<<<<<<<<