
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

test-load:
	@mkdir -p out
	@k6 run -e BASE_URL=http://localhost:1323 --out json=out/load.json ./tests/load/load-test.js

test-stress:
	@mkdir -p out
	@k6 run -e BASE_URL=http://localhost:1323 --out json=out/stress.json ./tests/load/stress-test.js

test-spike:
	@mkdir -p out
	@k6 run -e BASE_URL=http://localhost:1323 --out json=out/spike.json ./tests/load/spike-test.js

.PHONY: server init dev test coverage test-unit test-integration test-load test-stress test-spike