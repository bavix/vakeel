# vakeel

vakeel is a service that allows you to send updates from multiple clients to a single server.

## Description of the service

The vakeel service is built using [gRPC](https://grpc.io/) and [zerolog](https://github.com/rs/zerolog). 

### Agent
The agent is a client that sends updates to the server.

### Server
The server is a gRPC server that receives updates from multiple clients and stores them in memory.

### Protocol
The protocol used by the service is defined in the [protobuf definition](https://github.com/bavix/vakeel-way/blob/master/api/vakeel_way/state.proto).

## Run the service
```
LOG_LEVEL=info go run main.go agent --id=224f8a59-6705-4f3e-b7de-177757932aad
```
