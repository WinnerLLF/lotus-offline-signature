
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	go build -o wallet-signature -ldflags "-X main.VERSION=$version -X 'main.BUILD_TIME=`date`' " main.go

.PHONY: test
test:
	go build -o wallet-signature-test -ldflags "-X main.VERSION=$version -X 'main.BUILD_TIME=`date`' " main.go


