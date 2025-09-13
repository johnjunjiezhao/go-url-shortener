# Repository Guidelines

## Project Structure & Module Organization
- `main.go`: Gin HTTP server entrypoint (listens on `:9808`).
- `store/`: Redis-backed storage package.
  - `store_service.go`: client init and URL mapping helpers.
  - `store_service_test.go`: tests for storage behavior.
- `go.mod`, `go.sum`: module `github.com/johnjunjiezhao/go-url-shortener` and deps.

## Build, Test, and Development Commands
- Run server: `go run main.go` or `make run` (requires Redis).
- Build binary: `go build -o server .` or `make build` (outputs `./server`).
- Run tests: `go test ./...` or `make test`; verbose: `go test -v ./store`.
- Coverage: `go test -cover ./...` or `make cover`.
- Format & vet: `go fmt ./...` / `make fmt` and `go vet ./...` / `make vet` (run before pushing).
- Start Redis locally: `docker run --rm -p 6379:6379 redis:7` or `make redis`.

## Configuration
- Env vars: `REDIS_ADDR` (default `localhost:6379`), `REDIS_PASSWORD` (default empty), `REDIS_DB` (default `0`).
- Example file: see `.env.example` (note: app reads env via process, not a dotenv loader).

## Coding Style & Naming Conventions
- Go style: use `gofmt` (tabs, standard imports). Prefer `goimports` if available.
- Packages: short, all lowercase (`store`). Files: lowercase with `_` if needed; tests end with `_test.go`.
- Identifiers: exported `CamelCase`, unexported `camelCase`. Keep functions small and purposeful.
- Errors: return and handle errors; avoid `panic` in request paths. Log with context and propagate upstream.

## Testing Guidelines
- Frameworks: standard `testing` with `github.com/stretchr/testify/assert`.
- Names: `TestXxx(t *testing.T)` per package; table tests encouraged for variants.
- Integration: tests in `store/` require Redis on `localhost:6379`.
- Run all: `go test ./...`; with coverage: `go test -cover ./...`.

## Commit & Pull Request Guidelines
- Commits: concise, imperative present tense (e.g., "add store service"). Group related changes; keep diffs focused.
- Optional: Conventional Commits (`feat:`, `fix:`, `chore:`) welcomed for clarity.
- PRs: include a clear description, rationale, testing steps/commands, and any screenshots of responses (e.g., curl output). Link related issues. Ensure `go fmt`, `go vet`, and tests pass.

## Security & Configuration Tips
- Redis defaults to no auth here; use only in local dev. Do not commit secrets or `.env` files.
- For deployments, parameterize Redis via env vars (e.g., `REDIS_ADDR`, `REDIS_PASSWORD`) and add proper error handling and timeouts.
