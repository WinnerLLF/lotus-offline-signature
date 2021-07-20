
GOPATH:=$(shell go env GOPATH)

.PHONY: build
.PHONY: test

build:
	go build -o wallet-signature -ldflags "-X main.VERSION=$version -X 'main.BUILD_TIME=`date`' " ./cmd/server/main.go

test:
	go build -o wallet-signature-test -ldflags "-X main.VERSION=$version -X 'main.BUILD_TIME=`date`' " ./cmd/server/main.go