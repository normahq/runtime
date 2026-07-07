# Runtime v2 Migration

`github.com/normahq/runtime/v2` is the standalone runtime module for Go ADK v2 based applications.

## Import Changes

- `github.com/normahq/norma/pkg/runtime/...` -> `github.com/normahq/runtime/v2/...`
- `github.com/normahq/runtime/...` -> `github.com/normahq/runtime/v2/...`
- `google.golang.org/adk/...` -> `google.golang.org/adk/v2/...`
- `github.com/normahq/runtime/acpagent` was removed; use `github.com/normahq/go-adk-acpagent/v2` directly.
- `providererror` was removed.

## Local Consumer Audit

These repositories were identified as consumers that need separate migration work:

- `balda`: imports `github.com/normahq/norma/pkg/runtime/...`
- `relay`: imports `github.com/normahq/norma/pkg/runtime/...`
- `diffpal`: imports `github.com/normahq/norma/pkg/runtime/...` and removed `providererror`
- `aida`: imports `github.com/normahq/norma/pkg/runtime/...`
- `alatooguide/assistant`: imports old `github.com/normahq/runtime/...`

This change does not update those consumers.
