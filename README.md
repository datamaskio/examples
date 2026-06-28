# DataMask Examples

Runnable examples demonstrating DataMask integration patterns.

[![Go Report Card](https://goreportcard.com/badge/github.com/datamaskio/examples)](https://goreportcard.com/report/github.com/datamaskio/examples)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Examples

| Example | Demonstrates |
|---|---|
| [`basic/`](basic/) | Struct tags + `datamask.Redact` |
| [`slog/`](slog/) | `maskslog.NewHandler` — secret detection in logs |
| [`otel/`](otel/) | `maskotel.NewProcessor` — secret warnings in spans |

## Run

```bash
cd basic    && go run main.go
cd slog     && go run main.go
cd otel     && go run main.go
```

Or from anywhere:

```bash
go run github.com/datamaskio/examples/basic@v0.1.0
```

## License

MIT
