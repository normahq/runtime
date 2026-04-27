# AGENTS Guidelines

## Scope
- This repository is a standalone Go library module for reusable runtime packages extracted from `norma`.
- Preserve public package APIs unless a change is explicitly required and documented in `README.md`.

## Workflow
- Use the Go version declared in `go.mod`.
- Prefer project-local tools via `go tool`.
- Run `go test ./...`, `go test -race ./...`, and `go tool golangci-lint run ./...` before finalizing changes.
- Keep ACP integration tests opt-in; they require external binaries and auth.

## Code Style
- Follow the [Google Go Style Guide](https://google.github.io/styleguide/go/).
- Keep exported behavior stable and document any import-path or config-shape changes.
- Use descriptive typed context keys in tests and avoid stringly-typed context values.
- Prefer small, focused examples in docs that compile against the current module path.
