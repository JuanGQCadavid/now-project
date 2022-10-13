#!/bin/bash

export GOOS="linux"
export CGO_ENABLED="0"
export GOARCH="amd64"

echo "Golang version"
go version

echo "Golang envs"
go env

echo "Building service"
go build -o main ./cmd/http/main.go