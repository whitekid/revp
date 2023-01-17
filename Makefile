TARGET=bin/revp bin/revps
SRC=$(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*_test.go")
PROTO_DEFS := $(shell find . -not -path "./vendor/*" -type f -name '*.proto' -print)
PROTO_GOS := pb/v1alpha1/revp_grpc.pb.go

GO?=go
BUILD_FLAGS?=-v

.PHONY: clean test dep tidy

all: build
build: $(TARGET)

$(TARGET): $(SRC) ${PROTO_GOS}
	@mkdir -p bin/
	${GO} build -o bin/ ${BUILD_FLAGS} ./cmd/...

clean:
	rm -rf bin/

test:
	${GO} test ./... --count=1

# update modules & tidy
dep:
	rm -f go.mod go.sum
	${GO} mod init revp

	@$(MAKE) tidy

tidy:
	${GO} mod tidy -v

pb/v1alpha1/revp_grpc.pb.go: pb/revp.proto
	protoc -I=./pb \
      --go_out=./pb/v1alpha1 \
      --go_opt=paths=source_relative \
      --go-grpc_out=./pb/v1alpha1 \
      --go-grpc_opt=paths=source_relative \
	  pb/revp.proto

	cd ./pb/v1alpha1 && mockery --name RevpClient
