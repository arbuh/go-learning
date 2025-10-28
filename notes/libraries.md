# Libraries and packages

A list of useful libraries and frameworks

## Utility

* `fmt` - strings formatting and printing
* `encoding/json` - JSON serialization/deserialization

## REST API

* `net/http` - the standard Go library that provides HTTP client and server. It is usually enough for building a simple API.
* `gorilla/mux` - a library that provides more advanced routing options comparing to `het/http`
* `swaggo/swag` - a Swagger generator that heavily relies on annotation

## Local apps

* `github.com/spf13/cobra` - a framework for building CLI applications

## Logging and monitoring

* `log` - an old simple Logging
* `log/slog` - a modern package that supports structured logging
* `context` - defines the Context type can be used to carry the context values or managing the requests

Links:
* [Logging in Go with Slog: The Ultimate Guide ](https://betterstack.com/community/guides/logging/logging-in-go/)
