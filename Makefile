BINARY_NAME ?= uconverter

build:
	go build -o $(BINARY_NAME)

release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "-s -w" -o $(BINARY_NAME)

test:
	go test ./...
