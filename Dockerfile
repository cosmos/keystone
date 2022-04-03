FROM debian:latest
MAINTAINER jk@stabledomain.net

RUN apt -y update && apt -y install build-essential wget git
RUN wget https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz
RUN tar -xvf go1.17.7.linux-amd64.tar.gz
RUN mv go /usr/local
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH="${GOPATH}/bin:$GOROOT/bin:${PATH}"
RUN git clone https://github.com/frumioj/cosmos-sdk.git
RUN git clone https://github.com/frumioj/keystone.git
WORKDIR keystone
RUN go mod tidy
RUN make build
WORKDIR /keystone/plugin/file
RUN go build -buildmode=plugin -o file_keys.so file.go
WORKDIR /keystone/plugin/pkcs11
RUN go build -buildmode=plugin -o pkcs11_keys.so pkcs11.go
