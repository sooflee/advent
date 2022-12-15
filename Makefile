.PHONY: test tools vet fmt clean lint install

test: 
	go test ./... -count=1 -v

tools:
	go get -u github.com/stretchr/testify
	go get -u github.com/davecgh/go-spew/spew
	go get -u github.com/golang-collections/collections

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
