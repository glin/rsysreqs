PACKAGE = rsysreqs
BINARY_NAME = rsysreqs

all: test build

build:
	@GOPATH=$(PWD) go build -o ./bin/$(BINARY_NAME) $(PACKAGE)

test:
	@GOPATH=$(PWD) go test $(PACKAGE)/...

clean:
	go clean
	rm -f bin/$(BINARY_NAME)

run:
	@./bin/$(BINARY_NAME)

vendor-ensure:
	@cd ./src/rsysreqs; GOPATH=$(PWD) dep ensure