PACKAGES = ./src/rsysreqs/...

all: test build

build:
	GOPATH=$(CURDIR) go install $(PACKAGES)

test:
	GOPATH=$(CURDIR) go test -v -cover $(PACKAGES)

clean:
	GOPATH=$(CURDIR) go clean -i $(PACKAGES)

vendor-ensure:
	cd src/rsysreqs; GOPATH=$(CURDIR) dep ensure
