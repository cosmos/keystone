.phony: proto install clean build

INSTALL_DIR="$(shell go env GOPATH)/bin"
PROTO_SRC=proto
PROTO_DIR=keystone

# build is first target so that lone 'make' calls 'make build'
build: proto *.go
	go build -o keystoned .

install: ./keystoned
	cp keystoned $(INSTALL_DIR)

proto: $(PROTO_DIR)/keystone_base.pb.go $(PROTO_DIR)/keystone2.pb.go $(PROTO_DIR)/keystone2_grpc.pb.go $(PROTO_DIR)/keystone2_admin.pb.go $(PROTO_DIR)/keystone2_admin_grpc.pb.go

$(PROTO_DIR)/%.pb.go: proto/%.proto
	protoc --proto_path=$(PROTO_SRC) --go_out=$(PROTO_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_DIR) --go-grpc_opt=paths=source_relative $< --experimental_allow_proto3_optional

clean:
	rm -rf $(PROTO_DIR)/*.pb.go
	rm keystoned
