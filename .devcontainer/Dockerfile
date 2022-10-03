ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/devcontainers/go:1.19-bullseye

ENV PYTHONUNBUFFERED=1
ENV INSTALL_DIR /installs
ENV TFENV_VERSION 1.3.1

# nvm -----------------------------------------------------------------------------

ARG NODE_VERSION="lts/*"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c ". /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# base ----------------------------------------------------------------------------

RUN apt update \
    && apt -y install python3-dev python3-pip nodejs npm gcc musl-dev acl
RUN bash -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b $INSTALL_DIR

RUN ln -sf python3 /usr/bin/python
RUN ln -sf pip3 /usr/bin/pip
RUN ln -sf $INSTALL_DIR/task /usr/bin/task

# pip -----------------------------------------------------------------------------

COPY ./requirements.txt /usr/
RUN pip install --no-cache --upgrade pip setuptools
RUN pip install -r /usr/requirements.txt

# kubectl -------------------------------------------------------------------------

RUN curl -sSL -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
    && chmod +x /usr/local/bin/kubectl

# helm ----------------------------------------------------------------------------

RUN curl -s https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash -

# skaffold ------------------------------------------------------------------------

RUN curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
sudo install skaffold /usr/local/bin/

# ko ------------------------------------------------------------------------------

RUN go install github.com/google/ko@latest

# gcloud --------------------------------------------------------------------------

RUN sudo apt-get install apt-transport-https ca-certificates gnupg
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | tee /usr/share/keyrings/cloud.google.gpg && apt-get update -y && apt-get install google-cloud-sdk -y

# .minikube -----------------------------------------------------------------------

COPY kube-config.sh /usr/local/share/
RUN echo "source /usr/local/share/kube-config.sh" | tee -a /home/vscode/.bashrc >> /home/vscode/.zshrc

# task init -----------------------------------------------------------------------

RUN echo "task init" | tee -a /home/vscode/.bashrc >> /home/vscode/.zshrc
