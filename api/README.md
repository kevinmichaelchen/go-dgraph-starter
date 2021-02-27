# api

## Getting started
Running
```
make
```
should spin up everything.

## Dependencies
This is a Go-based back-end using:
* [gRPC](https://grpc.io/) for high-performance data transport
* [Protocol Buffers](https://developers.google.com/protocol-buffers), a language-neutral binary serialization tool for the domain's data structures
* [github.com/go-redis/redis](https://github.com/go-redis/redis), a Redis SDK
* [rs/xid](https://github.com/rs/xid) for efficient, globally unique, k-ordered ID generation
* [rs/zerolog](https://github.com/rs/zerolog) for performant structured logging
* [spf13/pflag](https://github.com/spf13/pflag) and [spf13/viper](https://github.com/spf13/viper) for config binding
* [otel](https://go.opentelemetry.io/otel) for [OpenTelemetry](https://opentelemetry.io/)
* [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation) for struct validation