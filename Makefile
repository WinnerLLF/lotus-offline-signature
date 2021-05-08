
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	go build -o mwallet-signature -ldflags "-X main.VERSION=$version -X 'main.BUILD_TIME=`date`' " main.go

.PHONY: test
test:
	go build -o mwallet-signature-test -ldflags "-X main.VERSION=$version -X 'main.BUILD_TIME=`date`' " main.go


