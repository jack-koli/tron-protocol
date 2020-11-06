#!/bin/bash
# gen pb.go
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. core/contract/*.proto
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. core/*.proto
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. api/*.proto