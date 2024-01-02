FROM ubuntu:20.04
LABEL authors="kr-yeon"

# golang install
RUN apt-get update && apt-get install -y \
    wget \
    git \
    gcc \
    make \
    vim \
    curl
RUN wget https://golang.org/dl/go1.21.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
RUN rm go1.21.5.linux-amd64.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
RUN echo "export GOPATH=$HOME/go" >> ~/.bashrc
RUN echo "export GOROOT=/usr/local/go" >> ~/.bashrc
SHELL ["/bin/bash", "-c"]
RUN git clone https://github.com/kr-yeon/fiber-study
WORKDIR /fiber-study
FROM docker:dind
RUN apk add --no-cache go
COPY go.mod go.sum ./
RUN go mod download
FROM docker:dind
WORKDIR /fiber-study
RUN apk add --no-cache swag
RUN swag init
RUN go build -o main .

ENTRYPOINT ["top", "-b"]