
init:
	@echo "Initializing..."
	@go mod tidy

server: init
	@echo "Running Web Server..."
	@echo -e "\n"
	@go run cmd/server/main.go

dev: server

# venv:
# 	@echo "Creating virtual environment..."
# 	@python3 -m venv venv

.PHONY: server format init dev