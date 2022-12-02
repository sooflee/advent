.PHONY: test tools vet fmt clean lint install

test: 
	go test -v ./...

tools:
	go get -u github.com/stretchr/testify/assert

vet:
	go vet ./...

fmt:
	go fmt ./...

clean:
	go clean

lint:
	brew install golangci-lint
	brew upgrade golangci-lint
	golangci-lint run ./...

install:
	go install -v ./...
