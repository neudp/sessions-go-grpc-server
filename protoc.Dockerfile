FROM golang:1.25-alpine AS builder

ARG VERSION=32.0

RUN apk add --no-cache unzip curl

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${VERSION}/protoc-${VERSION}-linux-x86_64.zip && \
    unzip protoc-${VERSION}-linux-x86_64.zip -d /usr/local

# Install protoc-gen-go and protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

FROM alpine:latest

# Copy the Go plugins from builder
COPY --from=builder /usr/local/bin/protoc /usr/local/bin/
COPY --from=builder /go/bin/protoc-gen-go /usr/local/bin/
COPY --from=builder /go/bin/protoc-gen-go-grpc /usr/local/bin/
