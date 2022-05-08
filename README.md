# Introduction

Keystone is a key management system. It provides a key management server, offering a gRpc interface, where new keys can be added to a `keyring`, and then used for things like cryptographic signatures. Keystone supports keys stored in several different ways -- currently: traditional filesystem-based keys and PKCS11 (HSMs both hardware and cloud-based). It offers a plugin API for implementing support for other key storage types.

# Prerequisites

1. Install the basics for building software with go-lang (minimum version as of today is 1.17)

`apt -y update && apt -y install build-essential wget git`
`wget https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz`
`tar -xvf go1.17.7.linux-amd64.tar.gz`
`mv go /usr/local`

2. Setup your go-lang environment reasonably

```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH="${GOPATH}/bin:$GOROOT/bin:${PATH}"
```

3. Clone the repo (which also means for now first cloning the Cosmos SDK)

```
git clone https://github.com/frumioj/cosmos-sdk.git
git clone https://github.com/frumioj/keystone.git
```

# Build Keystone (I will make this easier one day - promise!)

`cd keystone`, and then
`go mod tidy` to ensure dependencies are in order

Keystone requires building at least one key-providing plugin. There (as of today) two plugins:

```
cd plugin/file
go build -buildmode=plugin -o file_keys.so file.go
cd ../pkcs11
go build -buildmode=plugin -o pkcs11_keys.so pkcs11.go
```

Back up to the top level:

```
cd ../..
make build
```

# Running the keystone server

This line will run Keystone with support for filesystem-based keys where the keys are stored in the given directory:

`./keystoned -chain-id foo -chain-rpc none -key-addr none -keyring-dir none -keyring-type none -key-plugin ./plugin/file/file_keys.so -file-cfg /home/johnk/src/keystoned2/plugin/file/keys/`

# Dockerfile

There is a Dockerfile in this directory, which implements most of the instructions in here, and can be used to build a Keystone server with:

`docker build .`

# Using keystone

In the `test` directory is a test program that can be used once you have started the keystone server (see above).

```
cd test
go build keytester.go
./keytester --help
```

You can create a new key (make sure you copy the label you is displayed on the screen if you wish to use the key!)

```
[johnk@fedora test]$ ./keytester -create
New key: eb1966c335055ded66febcfb277f221
```
