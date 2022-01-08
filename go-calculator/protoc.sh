#!/bin/zsh
protoc -I api --go_out=. api/calculator.proto
protoc -I api --go-grpc_out=. api/calculator.proto
