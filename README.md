# Welcome to stack-exchange-graphql-server

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg?cacheSeconds=2592000)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

> GraphQL implementation to expose stack overflow resources (Comments, Posts, Answers, Votes, etc)

### 🏠 [Try out the live API](https://stack-exchange-graphql-server.azurewebsites.net/)

[Stack Exchange](https://api.stackexchange.com/docs?tab=category#docs) doesn't have a GraphQL endpoint so I made one. For now, the live API above is in POC mode only as the schema is built. It is currently serving content from [https://anime.stackexchange.com/](https://anime.stackexchange.com/) and is updated weekly with more to come as the [pipeline & infra](https://github.com/snimmagadda1/stackexchange-dump-to-mysql) are beefed up & built out.

#### [🚀 CURRENT SCHEMA HERE 🚀](./graph/schema.graphqls)

## Compile

```sh
go build server.go
```

## Usage

The server is available as a docker container:

```
docker run -e SERVER='<dbhost>' -e UNAME='<dbUname>' -e PASS='<dbPass>' snimmagadda/stack-exchange-graphql-server:latest
```

or to run from source:

```sh
go run server.go
```

Some environment variables must be set to run locally

| Key      | Type   | Description          | Example        |
| -------- | ------ | -------------------- | -------------- |
| `SERVER` | String | DB host              | localhost      |
| `UNAME`  | String | app DB username cred | appuser        |
| `PASS`   | String | app DB password cred | supersecret123 |

`SERVER` should be the hostname of a MySQL DB. In order to serve content, the GraphQL server expects a populated `stacke` [schema](https://github.com/snimmagadda1/stack-exchange-dump-to-mysql/blob/master/src/main/resources/schema-base.sql).

TODO: configurable params...

## Run tests

```sh
TODO
```

## Built with

-   Go
-   [gqlgen](https://github.com/99designs/gqlgen)
-   [gorm](https://github.com/go-gorm/gorm)
-   [gqlparser](https://github.com/vektah/gqlparser)
-   [Azure](https://azure.microsoft.com/en-us/)
-   [Excalidraw](https://github.com/excalidraw/excalidraw)

### Current data pipeline

I'm probably going to do some refinement and swap to something like an Elastic backend but for now here's the early setup. XML dumps published by Stack Exchange are imported into a relational backend using a custom job written with [Spring Batch](https://github.com/spring-projects/spring-batch) on a schedule. The graphql server reads from this backend to expose Stack Exchange data. Expect a minimal amount of latency because this is currently hosted in an App Service, which will spin down during periods of low-usage. If this gets some traction, availability will be increased.

![Diagram of current processing pipeline](pipeline_current.png)

## Author

👤 **Sai Nimmagadda**

-   Website: s11a.com
-   Github: [@snimmagadda1](https://github.com/snimmagadda1)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](https://github.com/snimmagadda1/stack-exchange-graphql/issues).

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [Sai Nimmagadda](https://github.com/snimmagadda1).

This project is [MIT](LICENSE.md) licensed.

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
