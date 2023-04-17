#!/bin/bash

export GOOS="linux"
export CGO_ENABLED="0"
export GOARCH="amd64"
export GOPATH="$(pwd)"
export GOMOD="$(pwd)/go.mod"

echo "Golang version"
go version

echo "Golang envs"
go env

# echo "Installing dependencies"
# go get 

echo "Building service"
go build -o main ./cmd/lambda/main.go