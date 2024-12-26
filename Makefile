ORDER_BINARY=orderApp
DELIVERY_BINARY=deliveryApp
NOTIFICATION_BINARY=notificationApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_order build_delivery build_notification
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the order binary as a linux executable
build_order:
	@echo "Building order-service binary..."
	cd ./order-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ORDER_BINARY} ./cmd
	@echo "Done!"

## build_broker: builds the delivery binary as a linux executable
build_delivery:
	@echo "Building delivery-service binary..."
	cd ./delivery-service && env GOOS=linux CGO_ENABLED=0 go build -o ${DELIVERY_BINARY} ./cmd
	@echo "Done!"

## build_broker: builds the notification binary as a linux executable
build_notification:
	@echo "Building notification-service binary..."
	cd ./notification-service && env GOOS=linux CGO_ENABLED=0 go build -o ${NOTIFICATION_BINARY} ./cmd
	@echo "Done!"
