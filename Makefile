TARGET=bin/revp bin/revps
GO_PKG=github.com/whitekid/revp
SRC=$(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*_test.go")
PROTO_DEFS := $(shell find . -not -path "./vendor/*" -type f -name '*.proto' -print)
PROTO_GOS := $(patsubst %.proto,%.pb.go,$(PROTO_DEFS))

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
	${GO} mod init ${GO_PKG}

	@$(MAKE) tidy

tidy:
	${GO} mod tidy

%.pb.go: $(patsubst %.pb.go,%.proto,$@)
	protoc -I=./$(@D) --go_out=./$(@D) --go_opt=paths=source_relative \
		--go-grpc_out=./$(@D) --go-grpc_opt=paths=source_relative ./$(patsubst %.pb.go,%.proto,$@)
