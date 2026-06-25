module github.com/datamaskio/examples

go 1.26.4

require (
	github.com/datamaskio/datamask v0.1.0
	github.com/datamaskio/maskotel v0.1.0
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/sdk v1.28.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace (
	github.com/datamaskio/datamask => ../datamask
	github.com/datamaskio/maskotel => ../maskotel
)
