FROM ubuntu:20.04
LABEL authors="kr-yeon"

# update apt-get
RUN apt-get update && apt-get install -y \
    wget \
    git \
    gcc-x86-64-linux-gnu \
    libc6-dev-amd64-cross \
    make \
    vim \
    curl

# install go
RUN wget https://golang.org/dl/go1.21.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
RUN rm go1.21.5.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/root/go"
ENV GOROOT="/usr/local/go"

# arm gcc
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc

# load project
RUN git clone https://github.com/kr-yeon/fiber-study
WORKDIR /fiber-study

# install package
COPY go.mod go.sum ./
RUN go mod download

# build
CMD ["swg", "init"]
CMD ["go", "run", "main.go"]

EXPOSE 3000
