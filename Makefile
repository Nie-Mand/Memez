
init:
	@go mod tidy

server: init
	@go build -o bin/server cmd/main.go
	@JAEGER_SERVICE_NAME=memez ./bin/server

dev: server

test: 
	@mkdir -p out
	@go test -v -cover -coverprofile=out/coverage.out ./... 

coverage:
	@mkdir -p out
	@go tool cover -html=coverage.out -o out/coverage.html

.PHONY: server init dev test coverage