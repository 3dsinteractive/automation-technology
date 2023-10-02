FROM jenkins/jenkins:2.425-slim

USER root

RUN \
    apt-get update && \
    apt-get install -y build-essential && \
    apt-get install -y libltdl7 && \
    apt-get install -y vim && \
    apt-get install uuid-runtime

# Install nodejs 18
RUN \
    curl -fsSL https://deb.nodesource.com/setup_18.x | bash && \
    apt-get install -y nodejs

# Install yarn
RUN \
    apt-get remove cmdtest && \
    curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --dearmor | tee /usr/share/keyrings/yarnkey.gpg >/dev/null && \
    echo "deb [signed-by=/usr/share/keyrings/yarnkey.gpg] https://dl.yarnpkg.com/debian stable main" | tee /etc/apt/sources.list.d/yarn.list && \
    apt-get update && \
    apt-get install -y yarn

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN curl -O https://storage.googleapis.com/golang/go1.20.linux-amd64.tar.gz
RUN tar xvf go1.20.linux-amd64.tar.gz
RUN mv go /usr/local
ENV GOPATH /root/go
ENV PATH="/usr/local/go/bin:${PATH}"

# set permission use docker host agent
CMD DOCKER_GID=$(stat -c '%g' /var/run/docker.sock) && \
    groupadd -for -g ${DOCKER_GID} docker && \
    usermod -aG docker jenkins && \
    /usr/local/bin/jenkins.sh
