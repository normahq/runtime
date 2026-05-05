# runtime

`runtime` is a standalone Go module containing reusable agent runtime packages extracted from `norma`.

The repo stays as one Go module. The root import path `github.com/normahq/runtime` is documentation-only; functional APIs live in the subpackages.

## Start Here

- Use `agentfactory` when you want to build runtime agents from validated provider config.
- Use `appconfig` when you need config-file loading, profile overlays, and runtime validation.
- Use `agentconfig` when you already have decoded config and need schema validation or normalization.
- Use `acpagent` when you need direct ACP subprocess control.
- Use `hostedagent` when you want to wrap a local API-backed model as an ADK agent.
- Use `poolagent` to provide ordered failover across providers.
- Use `structuredagent` to enforce JSON-schema-constrained I/O.
- Use `mcpregistry` and `sessionstate` as low-level support packages.

## Installation

```bash
go get github.com/normahq/runtime
```

## Packages

- `github.com/normahq/runtime/acpagent`
- `github.com/normahq/runtime/agentconfig`
- `github.com/normahq/runtime/agentfactory`
- `github.com/normahq/runtime/appconfig`
- `github.com/normahq/runtime/hostedagent`
- `github.com/normahq/runtime/mcpregistry`
- `github.com/normahq/runtime/poolagent`
- `github.com/normahq/runtime`
- `github.com/normahq/runtime/sessionstate`
- `github.com/normahq/runtime/structuredagent`

## Usage

### Validate runtime settings

```go
package main

import (
	"fmt"

	"github.com/normahq/runtime/appconfig"
)

func main() {
	settings := map[string]any{
		"providers": map[string]any{
			"codex": map[string]any{
				"type": "codex_acp",
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

	"github.com/normahq/runtime/agentconfig"
	"github.com/normahq/runtime/agentfactory"
	"github.com/normahq/runtime/mcpregistry"
)

func main() {
	providers := map[string]agentconfig.Config{
		"codex": {
			Type: agentconfig.AgentTypeCodexACP,
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

### Start an ACP-backed agent directly

```go
package main

import (
	"context"
	"os"

	"github.com/normahq/runtime/acpagent"
)

func main() {
	agent, err := acpagent.New(acpagent.Config{
		Context:    context.Background(),
		Command:    []string{"opencode", "acp"},
		WorkingDir: ".",
		Stderr:     os.Stderr,
	})
	if err != nil {
		panic(err)
	}
	defer func() { _ = agent.Close() }()
}
```

ACP `session/update.plan` notifications are exposed through ADK event state at
`event.Actions.StateDelta[acpagent.PlanStateKey]`. Each update replaces the
full plan snapshot and is not emitted as a content part.

## Integration Tests

Optional ACP integration tests are available for real runtimes:

```bash
go test -tags='integration,opencode' -count=1 ./acpagent
go test -tags='integration,codex' -count=1 ./acpagent
```

Requirements:
- `opencode` must be available on `PATH` for the OpenCode integration suite.
- `npx` must be available on `PATH` for the Codex ACP bridge suite.
- External runtime auth and local environment setup must already be configured.

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
