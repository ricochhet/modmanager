LDFLAGS=-X 'main.buildDate=$(shell date)' -X 'main.gitHash=$(shell git rev-parse HEAD)' -X 'main.buildOn=$(shell go version)' -w -s

GO_BUILD=go build -trimpath -ldflags "$(LDFLAGS)"

.PHONY: all fmt mod lint test deadcode syso modmanager-linux modmanager-linux-arm modmanager-darwin modmanager-darwin-arm modmanager-windows clean

all: modmanager-linux modmanager-linux-arm modmanager-darwin modmanager-darwin-arm modmanager-windows 

fmt:
	gofumpt -l -w .

mod:
	go get -u
	go mod tidy

lint:
	golangci-lint run

test:
	go test ./...

deadcode:
	deadcode ./...

syso:
	windres modmanager.rc -O coff -o modmanager.syso

modmanager-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o modmanager-linux

modmanager-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO_BUILD) -o modmanager-linux-arm

modmanager-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o modmanager-darwin

modmanager-darwin-arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GO_BUILD) -o modmanager-darwin-arm

modmanager-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO_BUILD) -o modmanager-windows.exe

clean:
	rm -f modmanager-linux modmanager-linux-arm modmanager-darwin modmanager-darwin-arm modmanager-windows.exe