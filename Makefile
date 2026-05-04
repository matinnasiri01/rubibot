# Build the application
all: build test

build:
	@echo "Building..."
	
	@go build -o main.exe main.go

# Run the application
run:
	@echo "Runing..."

	@go run main.go

# Format
fmt:
	@echo "Formating..."

	@go fmt main.go
	@go fmt rubibot/*.go


	@echo "Done!"

# Create DB container
docker-run:
	@docker compose up --build

# Shutdown DB container
docker-down:
	@docker compose down

clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command air -ErrorAction SilentlyContinue) { \
		clear; \
		air; \
		Write-Output 'Watching...'; \
	} else { \
		Write-Output 'Installing air...'; \
		go install github.com/air-verse/air@latest; \
		air; \
		Write-Output 'Watching...'; \
	}"

.PHONY: all build run fmt clean watch docker-run docker-down
