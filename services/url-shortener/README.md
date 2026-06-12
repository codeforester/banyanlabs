# URL Shortener Service

This is the first Banyan Labs service. It starts as a local-first Go service
with SQLite storage, startup migrations, structured JSON logging, and a health
endpoint.

## Run Locally

From the repo root after `basectl setup banyanlabs`:

```bash
basectl run banyanlabs dev
```

That starts the service in the background. Check or stop it with:

```bash
basectl run banyanlabs status
basectl run banyanlabs stop
```

To run the service in the foreground:

```bash
basectl run banyanlabs dev -- --foreground
```

Or directly from this service directory:

```bash
go run ./cmd/url-shortener
```

Defaults:

- address: `127.0.0.1:8080`
- database: repo-local `var/url-shortener/url-shortener.sqlite3` when started
  through `basectl run banyanlabs dev`
- log level: `info`

Configuration environment variables:

```text
BANYAN_URL_SHORTENER_ADDR
BANYAN_URL_SHORTENER_DATABASE
BANYAN_LOG_LEVEL
```

## Validate

From the repo root:

```bash
basectl test banyanlabs
basectl run banyanlabs api-test
```

Or directly from this service directory:

```bash
CGO_ENABLED=0 go test ./...
CGO_ENABLED=0 go vet ./...
CGO_ENABLED=0 go build ./...
```

## Health Check

```bash
curl http://127.0.0.1:8080/healthz
```

## API Contract

The OpenAPI contract for this service is tracked at
[`openapi.yaml`](openapi.yaml). Black-box API smoke tests are tracked under
[`../../tests/api/url-shortener`](../../tests/api/url-shortener).
