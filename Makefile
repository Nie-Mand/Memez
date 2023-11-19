
init:
	@go mod tidy

server: init
	@go build -o bin/server cmd/main.go
	@JAEGER_SERVICE_NAME=memez ./bin/server

dev: server

.PHONY: server init dev