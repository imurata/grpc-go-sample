all: build

CLIENT_ENTRY := client/main.go
SERVER_ENTRY := server/main.go

CLIENT_BINARY := ./client/client
SERVER_BINARY := ./server/server

build:
	go mod init github.com/example
	cd proto/ && protoc  --go_out=. --go-grpc_out=require_unimplemented_servers=false:. sample.proto
	go mod tidy
	@echo "Building gRPC client..."
	go build -o $(CLIENT_BINARY) $(CLIENT_ENTRY)
	@echo "Building gRPC server..."
	go build -o $(SERVER_BINARY) $(SERVER_ENTRY)

clean:

	rm $(CLIENT_BINARY) || true
	rm $(SERVER_BINARY) || true
	rm go.mod go.sum || true
	rm -r ./grpc_sample/* || true
	
