# syntax=docker/dockerfile:experimental
FROM golang:1.16 AS builder

ENV GOPRIVATE github.com/MyOrg

COPY go.mod go.sum /go/app/
WORKDIR /go/app

# Install ssh client and git
RUN apt-get update && apt-get install -y openssh-client git

# Download public key for github.com
RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

# Download dependencies
RUN --mount=type=ssh go mod download -x

COPY . /go/app
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest as app

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 \
 && wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 \
 && chmod +x /bin/grpc_health_probe

COPY --from=builder /go/app/app /app/app

# database migrations
COPY --from=builder /go/app/db /app/db

RUN apk add --no-cache ca-certificates

WORKDIR /app
CMD ["./app"]