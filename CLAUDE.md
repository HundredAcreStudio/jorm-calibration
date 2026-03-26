# CLAUDE.md

## Build & Development

```bash
go build ./...        # Build all packages
go test ./...         # Run all tests
go vet ./...          # Static analysis
```

## Architecture

Simple HTTP API service for user management. Used as a calibration test suite for jorm pipeline evaluation.

- **`cmd/server/main.go`** — HTTP server entry point with mux routing
- **`internal/model/user.go`** — User struct
- **`internal/store/store.go`** — `Store` interface + `MemoryStore` implementation (thread-safe via RWMutex)
- **`internal/handler/user.go`** — HTTP handlers: ListUsers, GetUser, CreateUser, DeleteUser
- **`internal/cache/cache.go`** — Simple in-memory cache (NOT thread-safe — intentional bug)
- **`internal/middleware/`** — Empty, for future middleware

## Code Style

- Standard Go conventions (gofmt, go vet)
- Error handling: return errors, wrap with `fmt.Errorf("context: %w", err)`
- Table-driven tests preferred
- Handler pattern: `func HandlerName(deps) http.HandlerFunc`
