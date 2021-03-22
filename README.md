![go-dgraph-starter](https://raw.githubusercontent.com/graph-gophers/graphql-go/master/docs/img/logo.png)

## Table of Contents
1. [Introduction](#introduction)
    1. [In Progress](#in-progress)
    1. [Tools](#tools)
        1. [Protocol Buffers](#protocol-buffers)
        1. [gRPC](#grpc)
        1. [GraphQL](#graphql)
        1. [Dgraph](#dgraph)
        1. [Sqlboiler](#sqlboiler)
        1. [Meilisearch](#meilisearch)
    1. [Patterns](#patterns)
        1. [Cursor pagination](#cursor-pagination)
        1. [Garden](#garden)
        1. [Change Data Capture](#change-data-capture)

TOC generated with
```
docker run -v $PWD:/app -w /app --rm -it pbzweihander/markdown-toc README.md --min-depth 1 
```

## Introduction

[![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://forthebadge.com)

This is a Todo List app built with [Dgraph](https://dgraph.io/)-backed [gRPC](https://grpc.io/) and [GraphQL](https://graphql.org/) APIs, combined with a [NextJS](https://nextjs.org/) + [Chakra-UI](https://chakra-ui.com/) powered front-end.

### In Progress

This project is still very much in progress.

https://github.com/kevinmichaelchen/go-dgraph-starter/projects/1

### Tools

#### Protocol Buffers

Protocol buffers are a great choice for a language-neutral representation
of your data models.

They are binary, lean, and fast to serialize when compared with JSON.

Their extensibility comes from the fact that you can
add or remove fields in a way while maintaining backwards compatibilty.

Code generation makes it easy to generate your models in any language.

#### gRPC

gRPC is a lean transport system (typically paired with HTTP/2) meant to
reduce latency and payload size.

HTTP/2 allows for long-lived connections (fewer handshakes).

Unlike REST, gRPC uses an HTTP POST method for all calls, and doesn't specify the resource in its URL, which means edge caching is impossible.

gRPC is also binary, so it's not as easy to debug than something human-readable, like GraphQL. But it is excellent for internal calls between microservices.

#### GraphQL

GraphQL is an ideal query language for APIs, especially when consumed by web clients. It's typically human-readable (JSON). Clients can specify exactly which fields they want. GraphiQL is an amazing tool for gaining insight into what an API offers.

It also can make you think about your data as a graph.

#### Dgraph

Dgraph is a popular open-source, fast, distributed graph database written in Golang.
You get high throughput and low latency for deep joins and complex traversals.
It offers a query language that is a superset of GraphQL.

Graph databases in general prioritize relationships between data points, rather than the data points themselves.

Performance: SQL performance suffers the more joins you ask it to do.

Flexibility / Complexity: Like NoSQL, the schema can be modified easily with time. Adding new relationships is as simple as adding a predicate. No need to create join tables. Ultimately, the graph model is more simpler / more intuitive.

#### Sqlboiler

It's not currently used in this project, but it's worth mentioning [Sqlboiler](https://github.com/volatiletech/sqlboiler)
since I believe it is by far the best SQL ORM for Golang due its "data-first"
approach: it auto-generates Go code based on your existing schema and as a
result you get extreme type safety.

For services that aren't expected to have a whole lot of data relationships, SQL is still an excellent choice, and sqlboiler is an ideal ORM for keeping your persistence layer code strongly typed.

#### Meilisearch

We use [MeiliSearch](https://www.meilisearch.com/) as our "open source, blazingly fast and hyper relevant search-engine."

Per their site, MeiliSearch is effective and accessible:

> Efficient search engines are often only accessible to companies with the financial means and resources necessary to develop a search solution adapted to their needs. The majority of other companies that do not have the means or do not realize that the lack of relevance of a search greatly impacts the pleasure of navigation on their application, end up with poor solutions that are more frustrating than effective, for both the developer and the user.

### Patterns

#### Cursor pagination

- https://uxdesign.cc/why-facebook-says-cursor-pagination-is-the-greatest-d6b98d86b6c0
- https://relay.dev/graphql/connections.htm

#### Garden

The problem: developers will not bother with running CI tests and integration tests locally; instead they'll push to their GitHub PR and let the CI system take care of it.

Not only are they deprived of a production-like system on their local machines, but the feedback loop is too slow: having CircleCI run all the end-to-end tests takes too long.

[Garden](https://garden.io/) fixes all of this and breaks down the barrier between dev, testing, and CI.

For shops with hundreds or thousands of microservices, I can see the argument that it's excessive (or even impossible) to run the whole system on your local machine. That said, maybe there's a term for a subset of services that are related, and maybe it's worth running those on Garden.

#### Change Data Capture

In a distributed system, when events occur in one service, they need to eventually be broadcast to other services. Too often, I've seen an event get emitted inside a database transaction, regardless of whether that transaction succeeds or fails. The [transactional outbox](https://microservices.io/patterns/data/transactional-outbox.html) pattern offers a way to keep data across services in sync.
