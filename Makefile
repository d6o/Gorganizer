.PHONY: build test lint vet clean

VERSION ?= dev

build:
	go build -ldflags "-X main.version=$(VERSION)" -o gorganizer .

test:
	go test -race -count=1 ./...

lint:
	golangci-lint run ./...

vet:
	go vet ./...

clean:
	rm -f gorganizer
