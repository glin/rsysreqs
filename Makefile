PACKAGES = rsysreqs/...

all: test build

build:
	@GOPATH=$(PWD) go install $(PACKAGES)

PORT ?= 3000
start:
	@PORT=$(PORT) GOPATH=$(PWD) ./bin/api

test:
	@GOPATH=$(PWD) go test $(PACKAGES)

clean:
	go clean
	rm -rf bin

vendor-ensure:
	@cd ./src/rsysreqs; GOPATH=$(PWD) dep ensure