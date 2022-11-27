# Build
FROM golang:1.19-alpine3.16 AS builder

ARG BASE_PATH=/go/src/now-project/services
ARG APP_HOME=app
ARG GOOS="linux"
ARG CGO_ENABLED="0"
ARG GOARCH="amd64"

# Moving pkgs
WORKDIR ${BASE_PATH}
COPY .tempWorkDir ./

# Moving app
WORKDIR ${BASE_PATH}/${APP_HOME}
COPY ./cmd ./cmd
COPY ./internal ./internal 
COPY ./pkg ./pkg
COPY go.mod go.mod
COPY go.sum go.sum

RUN echo "Building service" \
    && go build -o main ./cmd/httpserver/main.go

# RUN adduser -S -D -h /usr/app/src appuser
# RUN chown -R appuser /usr/app
# USER appuser


ENTRYPOINT ["./main"]


# # Deploy
# FROM amazon/aws-lambda-go
# COPY --from=builder /go/src/now-project/services/app/main ${LAMBDA_TASK_ROOT}
# CMD [ "main" ]
