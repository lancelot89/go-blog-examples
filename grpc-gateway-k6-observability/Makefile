.PHONY: proto build up down clone-googleapis-initial update-googleapis

GOOGLEAPIS_DIR := vendor/googleapis

proto:
	protoc -I proto -I $(GOOGLEAPIS_DIR) --go_out=proto/echo --go_opt=paths=source_relative \
	--go-grpc_out=proto/echo --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=proto/echo --grpc-gateway_opt=paths=source_relative \
	proto/echo.proto

clone-googleapis-initial:
	@if [ ! -d "$(GOOGLEAPIS_DIR)" ]; then \
		echo "Cloning googleapis repository..."; \
		mkdir -p "$(GOOGLEAPIS_DIR)"; \
		git clone https://github.com/googleapis/googleapis "$(GOOGLEAPIS_DIR)"; \
	else \
		echo "googleapis repository already exists. Skipping initial clone."; \
	fi

update-googleapis:
	@if [ -d "$(GOOGLEAPIS_DIR)" ]; then \
		echo "Updating googleapis repository..."; \
		(cd "$(GOOGLEAPIS_DIR)" && git pull); \
	else \
		echo "googleapis repository not found. Please run 'make clone-googleapis-initial' first."; \
	fi

build:
	go build -o app ./cmd/main.go

up:
	docker-compose up

down:
	docker-compose down