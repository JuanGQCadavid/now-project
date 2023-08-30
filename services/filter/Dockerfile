FROM golang:latest

ENV GOOS="linux"
ENV CGO_ENABLED="0"
ENV GOARCH="amd64"

ARG CMD = "http"

WORKDIR /app
COPY cmd cmd
COPY internal/${CMD} internal/${CMD}
COPY go.mod go.mod
COPY go.sum go.sum

RUN go get

EXPOSE 8000

CMD  ["go", "run", "./cmd/${CMD}/main.go"]