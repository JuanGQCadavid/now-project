FROM golang:latest

ENV GOOS="linux"
ENV CGO_ENABLED="0"
ENV GOARCH="amd64"

ARG CMD="http"

WORKDIR /app

COPY tmp/pkgs /pkgs
COPY cmd/${CMD} cmd/${CMD}
COPY internal internal

COPY go.mod go.mod
COPY go.sum go.sum

RUN go build -o main ./cmd/${CMD}/main.go

EXPOSE 8000

CMD  ["./main"]