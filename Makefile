.phony: proto install clean build

INSTALL_DIR="$(shell go env GOPATH)/bin"

# build is first target so that lone 'make' calls 'make build'
build: proto *.go
	go build -o keystoned .

install: ./keystoned
	cp keystoned $(INSTALL_DIR)

proto: keystone/keystone_base.pb.go keystone/keystone2.pb.go keystone/keystone2_grpc.pb.go

keystone/%.pb.go: proto/%.proto
	protoc --proto_path=./proto --go_out=. --go-grpc_out=. ./proto/*.proto

clean:
	rm -rf keystone/*.pb.go
	rm keystoned
