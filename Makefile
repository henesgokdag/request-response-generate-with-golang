PROTO_GO_DIR := ./go.proto
PROTO_GOGO_DIR := gogo_proto
PROTO_GO_FILES := $(shell find $(PROTO_GO_DIR) -name '*.proto')
PROTO_GOGO_FILES := $(shell find $(PROTO_GOGO_DIR) -name '*.proto')
GENERATED_GO_DIR := ./generated_go
GENERATED_GOGO_REQUEST_DIR := ./generated_gogo/request
GENERATED_GOGO_RESPONSE_DIR := ./generated_gogo/response
GENERATED_GO_FILES := $(patsubst $(PROTO_GO_DIR)/%.proto,$(GENERATED_GO_DIR)/%.pb.go,$(PROTO_GO_FILES))
PROTOC := $(shell command -v protoc 2> /dev/null)
PROTOC_GEN_GO := $(shell command -v protoc-gen-go 2> /dev/null)

.PHONY: all clean install-protoc

all: install-protoc $(GENERATED_GO_FILES)
$(GENERATED_GO_DIR)/%.pb.go: $(PROTO_GO_DIR)/%.proto
	@mkdir -p $(dir $@)
	@tmp_file=$$(echo "$(dir $<)gen_$(notdir $<)"); \
	cp $< $$tmp_file; \
	echo 'option go_package = "$(patsubst %/,%,$(patsubst $(PROTO_GO_DIR)/%,%,$(dir $<)));";' >> $$tmp_file; \
	$(PROTOC) --go_out=$(GENERATED_GO_DIR) --go_opt=paths=source_relative -I$(PROTO_GO_DIR) $$tmp_file; \
	rm $$tmp_file

clean:
	rm -rf $(GENERATED_GO_DIR)

install-protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31

.PHONY: install
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	go get github.com/gogo/protobuf/gogoproto
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/jsonpb
	go get github.com/gogo/protobuf/protoc-gen-gogo
	go get -u github.com/gogo/protobuf/protoc-gen-gogo/descriptor

regenerate: install
	(protoc -I=. -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf -I=$(GOPATH)/src --gogofaster_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:./generated_gogo/response gogo_proto/response.proto)