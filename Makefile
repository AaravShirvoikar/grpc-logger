CLIENT_BINARY=clientApp
LOGGER_BINARY=loggerApp

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started"

up_build: build_client build_logger
	@echo "Stopping docker images..."
	docker-compose down
	@echo "Building and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done"

build_client:
	@echo "Building client binary..."
	cd ./logger-client && env GOOS=linux CGO_ENABLED=0 go build -o ${CLIENT_BINARY} ./cmd/api
	@echo "Done"

build_logger:
	@echo "Building logger binary..."
	cd ./logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done"