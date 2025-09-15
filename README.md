# Go URL Shortener

A minimal URL shortener written in Go with a Redis-backed store and a tiny Gin HTTP API.

## Overview
- Shortens long URLs deterministically per user using SHA-256 + Base58, truncated to 8 chars.
- Persists mappings in Redis with a 6-hour TTL.
- Exposes two endpoints: create a short URL and redirect via the short code.

## Project Structure
- `main.go` — Gin HTTP server entrypoint (listens on `:9808`).
- `handler/` — HTTP handlers.
- `shortener/` — Short link generator.
- `store/` — Redis-backed storage service and tests.
- `go.mod`, `go.sum` — module and dependencies.
- `Makefile` — common dev tasks.

## Requirements
- Go installed
- Redis 7+ (local Docker is fine)
- Optional: Docker to run Redis locally

## Docker
- Build and run app + Redis with Docker Compose:

```bash
make compose-up
# or
docker compose up --build -d
```

- Follow logs:

```bash
make compose-logs
# or
docker compose logs -f app
```

- Stop and remove containers:

```bash
make compose-down
# or
docker compose down
```

Notes:
- The app listens on `localhost:9808` and Redis on `localhost:6379` via published ports.
- The application reads `REDIS_ADDR` from the environment; in Compose it is set to `redis:6379`.

## Quickstart
1) Start Redis locally

```bash
make redis
# or
docker run --rm -p 6379:6379 --name urlshort-redis redis:7
```

2) Run the server

```bash
make run
# or
go run main.go
```

3) Create a short URL

```bash
curl -sS -X POST http://localhost:9808/short-urls \
  -H 'Content-Type: application/json' \
  -d '{"long_url":"https://example.com","user_id":"user123"}'
# {"message":"short url created successfully","short_url":"http://localhost:9808/XXXXXXXX"}
```

4) Follow the redirect

```bash
curl -i http://localhost:9808/XXXXXXXX
# HTTP/1.1 302 Found
# Location: https://example.com
```

## API
- `POST /short-urls`
  - Body: `{ "long_url": string, "user_id": string }` (both required)
  - Response: `{ "message": string, "short_url": string }`
- `GET /:shortURL`
  - Redirects with HTTP 302 to the original URL

## Configuration
The app reads environment variables from the process (no dotenv loader):
- `REDIS_ADDR` — default `localhost:6379`
- `REDIS_PASSWORD` — default empty
- `REDIS_DB` — default `0`

Example (bash):

```bash
export REDIS_ADDR="localhost:6379"
export REDIS_PASSWORD=""
export REDIS_DB="0"
make run
```

## Development & Testing
Common tasks via `make`:

- `make run` — run the server
- `make build` — build binary to `./server`
- `make test` — run tests
- `make cover` — run tests with coverage
- `make fmt` — format code
- `make vet` — static analysis
- `make tidy` — tidy modules (`go mod tidy`)
- `make redis` — run Redis locally via Docker
- `make compose-up` — start app and Redis via Docker Compose
- `make compose-logs` — tail app logs
- `make compose-down` — stop Compose stack
- `make clean` — remove built artifacts

Notes:
- Store tests in `store/` require Redis available at `localhost:6379`.
- Run `go test -v ./store` to focus Redis-dependent tests.

## Implementation Notes
- Shortening algorithm: `sha256(long_url + user_id)` → Base58 (Bitcoin alphabet) → first 8 chars.
- Storage: Redis `SET shortCode originalURL` with a 6-hour expiration.
- Error handling: current code uses `panic` on Redis/setup failures; in production, prefer returning errors and robust logging.

## Security
- Default Redis has no auth and is intended only for local dev. For deployments, set `REDIS_ADDR`, `REDIS_PASSWORD`, and `REDIS_DB` appropriately and secure the instance.
