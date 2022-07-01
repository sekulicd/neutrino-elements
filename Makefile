## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## fmt: Go Format
fmt:
	@echo "Gofmt..."
	@gofmt -w -l .

test:
	@echo "Testing..."
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -rf ./bin

build-n:
	@echo "Building neutrino cli..."
	@go build -v -o bin/neutrino ./cmd/neutrino

## build neutrino daemon
build-nd:
	@echo "Building neutrino daemon..."
	@export GO111MODULE=on; \
	env go build -tags netgo -ldflags="-s -w" -o bin/neutrinod ./cmd/neutrinod/main.go
