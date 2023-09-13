.DEFAULT_GOAL := lint
APPNAME=$(shell basename "$(PWD)")
GOPKG=$(shell head -n 1 ./go.mod)
BIN_LINUX="${APPNAME}.linux"
BIN_DARWIN_ARM="./$(APPNAME).darwin-arm64"
BIN_DARWIN_X8664="./$(APPNAME).darwin-x86_64"
OS:=$(shell uname)

.PHONY: help

all: help

## install-tools: intall all golang tools
install-tools:
	go install github.com/mgechev/revive@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	go install golang.org/x/tools/gopls@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
ifeq ($(OS), Linux)
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
	pip install --user yamllint
endif

help: Makefile
	@echo
	@echo " Choose a command run in:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## lint: fast check all sources for errors
lint: 
	revive ./...
	golangci-lint  run ./...

## lintall: full check
lintall: 
	revive ./...
	golangci-lint  run ./...
	gosec -color=false -exclude=G104 ./...

## compile: build binaries for all supported os
compile: 
	~/go/bin/go build .

## build: cleanning and build new binaries
build: clean lint
	gosec -color=false -exclude=G104  ./...
	GOOS=linux  GOARCH=amd64 go build -o "${BIN_LINUX}" 
	GOOS=darwin GOARCH=arm64 go build -o ${BIN_DARWIN_ARM}
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_DARWIN_X8664}

## clean: delete compiled binaries
clean:
	rm -f ${APPNAME} ${BIN_DARWIN_ARM} ${BIN_DARWIN_X8664} ${BIN_LINUX}
