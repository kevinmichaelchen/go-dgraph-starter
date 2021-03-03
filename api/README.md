# api

## Getting started
Running
```
make
```
should spin up everything.

## Dgraph UI
You can access the Dgraph UI at [localhost:8000](http://localhost:8000/).

Query examples:
```
{
  todos(func: eq(dgraph.type, "Todo")) {
    title
    created_at
    is_done
    creator {
      name
      created_at
    }
  }
}
```

## Dependencies
This is a Go-based back-end using:
* [gRPC](https://grpc.io/) for high-performance data transport
* [Protocol Buffers](https://developers.google.com/protocol-buffers), a language-neutral binary serialization tool for the domain's data structures
* [go-redis/redis](https://github.com/go-redis/redis), a Redis SDK
* [rs/xid](https://github.com/rs/xid) for efficient, globally unique, k-ordered ID generation
* [rs/zerolog](https://github.com/rs/zerolog) for performant structured logging
* [spf13/pflag](https://github.com/spf13/pflag) and [spf13/viper](https://github.com/spf13/viper) for config binding
* [otel](https://go.opentelemetry.io/otel) for [OpenTelemetry](https://opentelemetry.io/)
* [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation) for struct validation