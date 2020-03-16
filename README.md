# Simple gRPC client and server example
Basic runnable example of gRPC implementation in Go

## How to:
1. Create a proto file like `my_service/my_service.proto`
2. Execute `make proto` to generate `my_service.pb.go` file based on proto file definition
3. Create server and implement interfaces
4. Create client and use the interfaces for communication

## Start
1. Start client and server
```bash
# From project directory
go run server/main.go
go run client/main.go
```
2. Make a curl request to test
```bash
curl localhost:8080/add/2/6
# and
curl localhost:8080/multiply/2/6
```