// Example: maskslog detects secrets (API keys, tokens, JWTs) in log
// attributes and replaces them with [TYPENAME] placeholders.
//
// For PII redaction (emails, phones, SSNs), use struct tags + datamask.Redact
// before logging. See the basic example.
//
//	go run github.com/datamaskio/examples/slog
package main

import (
	"log/slog"
	"os"

	"github.com/datamaskio/datamask/pkg/maskslog"
)

func main() {
	// One-line adoption: wrap any slog.Handler.
	logger := slog.New(maskslog.NewHandler(
		slog.NewJSONHandler(os.Stdout, nil),
	))

	// secrets are detected and replaced, PII passes through.
	logger.Info("api request",
		"user", "alice",
		"token", "ghp_1234567890abcdefghijklmnopqrstuvwxyz", // ← redacted
		"endpoint", "/users",
	)
}
