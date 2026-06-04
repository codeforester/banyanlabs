# URL Shortener Service

This is the first Banyan Labs service. It starts as a local-first Go service
with SQLite storage, startup migrations, structured JSON logging, and a health
endpoint.

## Run Locally

```bash
go run ./cmd/url-shortener
```

Defaults:

- address: `127.0.0.1:8080`
- database: `var/url-shortener/url-shortener.sqlite3`
- log level: `info`

Configuration environment variables:

```text
BANYAN_URL_SHORTENER_ADDR
BANYAN_URL_SHORTENER_DATABASE
BANYAN_LOG_LEVEL
```

## Validate

```bash
CGO_ENABLED=0 go test ./...
CGO_ENABLED=0 go vet ./...
CGO_ENABLED=0 go build ./...
```

## Health Check

```bash
curl http://127.0.0.1:8080/healthz
```
