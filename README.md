# runtime

`runtime` is a standalone Go module containing reusable agent runtime packages extracted from `norma`.

The repo stays as one Go module. The root import path `github.com/normahq/runtime/v2` is documentation-only; functional APIs live in the subpackages.

## Start Here

- Use `agentfactory` when you want to build runtime agents from validated provider config.
- Use `appconfig` when you need config-file loading, profile overlays, and runtime validation.
- Use `agentconfig` when you already have decoded config and need schema validation or normalization.
- Use `agentfactory` with ACP provider config when you need ACP subprocess-backed agents.
- Use `hostedagent` when you want to wrap a local API-backed model as an ADK agent.
- Use `poolagent` to provide ordered failover across providers.
- Use `structuredagent` to enforce JSON-schema-constrained I/O.
- Use `mcpregistry` and `sessionstate` as low-level support packages.

## Installation

```bash
go get github.com/normahq/runtime/v2
```

## Packages

- `github.com/normahq/runtime/v2/agentconfig`
- `github.com/normahq/runtime/v2/agentfactory`
- `github.com/normahq/runtime/v2/appconfig`
- `github.com/normahq/runtime/v2/hostedagent`
- `github.com/normahq/runtime/v2/mcpregistry`
- `github.com/normahq/runtime/v2/poolagent`
- `github.com/normahq/runtime/v2`
- `github.com/normahq/runtime/v2/sessionstate`
- `github.com/normahq/runtime/v2/structuredagent`
- `github.com/normahq/runtime/v2/taskmaster`
- `github.com/normahq/runtime/v2/taskmaster/adk`

## Usage

### Validate runtime settings

```go
package main

import (
	"fmt"

	"github.com/normahq/runtime/v2/appconfig"
)

func main() {
	settings := map[string]any{
		"providers": map[string]any{
			"codex": map[string]any{
				"type": "codex_acp",
				"codex_acp": map[string]any{
					"bridge_version": "latest",
				},
			},
		},
	}

	if err := appconfig.ValidateSettings(settings); err != nil {
		panic(err)
	}

	fmt.Println("runtime config is valid")
}
```

### Build an agent from provider config

```go
package main

import (
	"context"

	"github.com/normahq/runtime/v2/agentconfig"
	"github.com/normahq/runtime/v2/agentfactory"
	"github.com/normahq/runtime/v2/mcpregistry"
)

func main() {
	providers := map[string]agentconfig.Config{
		"codex": {
			Type: agentconfig.AgentTypeCodexACP,
			CodexACP: &agentconfig.ACPConfig{
				BridgeVersion: "latest",
			},
		},
	}

	factory := agentfactory.New(providers, mcpregistry.New(nil))
	_, err := factory.Build(context.Background(), agentfactory.BuildRequest{
		AgentID:          "codex",
		WorkingDirectory: "/workspace",
	})
	if err != nil {
		panic(err)
	}
}
```

For `codex_acp`, `bridge_version` accepts an npm version or dist-tag for
`@normahq/codex-acp-bridge`; empty defaults to `latest`.

`reasoning_effort` is a generic ACP config field and may be set on any ACP-backed
provider config (`generic_acp`, `codex_acp`, `opencode_acp`, `copilot_acp`,
`claude_code_acp`). Runtime forwards it to the ACP agent config and leaves
backend-specific handling to the concrete ACP implementation.

`gemini_acp` is deprecated and no longer supported. Use another supported ACP
provider type, or configure Gemini explicitly through `generic_acp` if you need
a custom ACP command path.

Direct ACP agent construction lives in `github.com/normahq/go-adk-acpagent/v2`.
Runtime v2 uses that package internally through `agentfactory`.

## Migration Notes

- `github.com/normahq/runtime/...` imports move to `github.com/normahq/runtime/v2/...`.
- `github.com/normahq/norma/pkg/runtime/...` imports move to `github.com/normahq/runtime/v2/...`.
- ADK imports must use `google.golang.org/adk/v2/...`.
- The old `github.com/normahq/runtime/acpagent` package is removed; use `github.com/normahq/go-adk-acpagent/v2`.
- The old `providererror` package is removed.

See [MIGRATION.md](MIGRATION.md) for the local consumer audit.

## Development

```bash
go mod tidy
go test ./...
go test -race ./...
go tool golangci-lint run ./...
task docs:check
task docs:generate
```

## Go Doc

The source comments and example tests are the source of truth for package documentation.

Generate local docs from source:

```bash
task docs:generate
```

Validate that every public package has package docs and renders with `go doc`:

```bash
task docs:check
```

Generated output is written to `.cache/godoc/` and is not committed.

See `AGENTS.md` for repository-specific guidance.
