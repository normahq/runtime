# runtime

`runtime` is a standalone Go module containing reusable agent runtime packages extracted from `norma`.

It provides:
- ACP-backed agents via `acpagent`
- runtime config validation via `agentconfig` and `appconfig`
- ADK-compatible agent construction via `agentfactory`
- MCP server resolution via `mcpregistry`
- hosted and pooled agent helpers

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
```

See `AGENTS.md` for repository-specific guidance.
