package acpagent

// MCPServerType represents the transport type for an MCP server.
type MCPServerType string

const (
	// MCPServerTypeStdio is the stdio transport type.
	MCPServerTypeStdio MCPServerType = "stdio"
	// MCPServerTypeHTTP is the HTTP transport type.
	MCPServerTypeHTTP MCPServerType = "http"
	// MCPServerTypeSSE is the SSE (Server-Sent Events) transport type.
	MCPServerTypeSSE MCPServerType = "sse"
)

// MCPServerConfig describes how to connect to an MCP server.
type MCPServerConfig struct {
	Type       MCPServerType
	Cmd        []string
	Args       []string
	Env        map[string]string
	WorkingDir string
	URL        string
	Headers    map[string]string
}
