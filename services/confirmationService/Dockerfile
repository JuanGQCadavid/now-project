FROM golang:alpine3.20 as rest-builder

ENV GOOS="linux"
ENV CGO_ENABLED="0"
ARG GOARCH="arm64"

# ARG CMD="http"

WORKDIR /app

COPY cmd/${CMD} cmd/${CMD}
COPY internal internal

COPY go.mod go.mod
COPY go.sum go.sum

# RUN go build -o main ./cmd/${CMD}/main.go
RUN go build -o main ./cmd/lambda/main.go

FROM golang:alpine3.20 as sqs-builder

ENV GOOS="linux"
ENV CGO_ENABLED="0"
ARG GOARCH="arm64"

WORKDIR /app

COPY cmd/${CMD} cmd/${CMD}
COPY internal internal

COPY go.mod go.mod
COPY go.sum go.sum

RUN go build -o main ./cmd/sqs/main.go

FROM alpine:3.19

EXPOSE 8000

COPY --from=rest-builder /app/main /bin/rest/
COPY --from=sqs-builder /app/main /bin/sqs/

ENTRYPOINT  ["/bin/rest/main"]
