proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        pkg/**/proto/*.proto

server:
	go run cmd/main.go

build:
	go build -o bin/api-gateway cmd/main.go

run:
	bin/api-gateway

gateway:
	bash ./scripts/boot