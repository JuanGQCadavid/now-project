FROM golang:alpine3.19 as builder

ENV GOOS="linux"
ENV CGO_ENABLED="0"
ENV GOARCH="arm64"

ARG CMD="http"

WORKDIR /app

COPY cmd/${CMD} cmd/${CMD}
COPY internal internal

COPY go.mod go.mod
COPY go.sum go.sum

RUN go build -o main ./cmd/${CMD}/main.go

EXPOSE 8000

FROM alpine:3.19 as prod
COPY --from=builder /app/main /bin/
ENTRYPOINT  ["/bin/main"]

# FROM public.ecr.aws/lambda/provided:al2023 as lambda
# COPY --from=build /helloworld/main ./main
# ENTRYPOINT [ "./main" ]