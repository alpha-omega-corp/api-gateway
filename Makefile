server:
	go run cmd/main.go

build:
	go build -o bin/api-gateway cmd/main.go

run:
	bin/api-gateway

gateway:
	bash ./scripts/boot