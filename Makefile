
server:
	@echo "Running Web Server..."
	@echo -e "\n"
	@go run cmd/server/main.go

format:
	@echo "Formatting..."
	@go fmt ./...

venv:
	@echo "Creating virtual environment..."
	@python3 -m venv venv


.PHONY: server format