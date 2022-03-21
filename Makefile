TARGET=bin/revp bin/revps
GO_PKG=github.com/whitekid/revp
SRC=$(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*_test.go")

GO?=go
BUILD_FLAGS?=-v

.PHONY: clean test dep tidy

all: build
build: $(TARGET)

$(TARGET): $(SRC)
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
