.PHONY: build test lint vet clean

build:
	go build -o gorganizer .

test:
	go test -race -count=1 ./...

lint:
	golangci-lint run ./...

vet:
	go vet ./...

clean:
	rm -f gorganizer
