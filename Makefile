
init:
	@go mod tidy

server: init
	@go build -o bin/server cmd/main.go
	@JAEGER_SERVICE_NAME=memez ./bin/server

dev: server

test: 
	@mkdir -p out
	@go test -v --tags=integration,unit -cover -coverprofile=out/coverage.out ./... 

test-unit: 
	@go test -v --tags=unit -cover ./...

test-integration:
	@go test -v --tags=integration -cover ./...

coverage: test
	@go tool cover -html=out/coverage.out -o out/coverage.html

.PHONY: server init dev test coverage test-unit test-integration