PACKAGES = ./...

all: test build

build:
	@go install $(PACKAGES)

test:
	@go test -v -cover $(PACKAGES)

clean:
	@go clean -i $(PACKAGES)

vendor-ensure:
	dep ensure