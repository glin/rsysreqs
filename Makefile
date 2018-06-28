PACKAGE = rsysreqs

all: test build

build:
	@GOPATH=$(PWD) go install $(PACKAGE)/...

test:
	@GOPATH=$(PWD) go test $(PACKAGE)/...

clean:
	go clean
	rm -rf bin

vendor-ensure:
	@cd ./src/rsysreqs; GOPATH=$(PWD) dep ensure