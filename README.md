# gRPC Demo

A minimal Go application demonstrating RPC communication using gRPC.

## Concepts Explained

**RPC (Remote Procedure Call)**  
A protocol that lets a program execute a function/procedure on another server as if it were local.

**gRPC**  
Google's modern RPC framework using:
- HTTP/2 for transport
- Protocol Buffers (binary format) for data serialization
- Strong typing via `.proto` files

## Implementation

**Server** (`server/main.go`)
- Creates gRPC server
- Implements service methods
- Listens for client requests

**Client** (`client/main.go`)
- Establishes connection to server
- Calls remote methods via generated stubs

**Protocol Buffers** (`proto/greet.proto`)
- Define services and message formats
- Auto-generate Go code for serialization/deserialization

## Prerequisites
- Go 1.16+
- protoc compiler
- Go plugins:  
  `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`  
  `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

## Installation
1. Clone repository
2. Install dependencies:  
   `go mod download`

## Usage
1. Generate Protobuf code:  
   `protoc --go_out=. --go-grpc_out=. proto/greet.proto`

2. Start server:
   `cd server`
   `go run *.go`

3. Run client (in separate terminal):
   `cd client`
   `go run *.go`

## Services
- `SayHello`: Unary RPC (single request → single response)
- `SayHelloServerStreaming`: Server streaming RPC
- `SayHelloClientStreaming`: Client streaming RPC 
- `SayHelloBidirectionalStreaming`: Bidirectional streaming

## Development
- Update `.proto` file for service/message changes
- Regenerate code with `protoc` command above