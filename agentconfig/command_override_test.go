package agentconfig

import (
	"reflect"
	"testing"
)

func TestNormalizeConfig_ACPAliasCommandOverride(t *testing.T) {
	cases := []struct {
		name string
		cfg  Config
	}{
		{"codex", Config{Type: AgentTypeCodexACP, CodexACP: &ACPConfig{Cmd: []string{"custom-codex"}, ExtraArgs: []string{"--stdio"}, Model: "model", ReasoningEffort: "high", Mode: "review"}}},
		{"claude", Config{Type: AgentTypeClaudeCodeACP, ClaudeCodeACP: &ACPConfig{Cmd: []string{"custom-claude"}, ExtraArgs: []string{"--stdio"}, Model: "model", ReasoningEffort: "high", Mode: "review"}}},
		{"opencode", Config{Type: AgentTypeOpenCodeACP, OpenCodeACP: &ACPConfig{Cmd: []string{"custom-opencode"}, ExtraArgs: []string{"--stdio"}, Model: "model", ReasoningEffort: "high", Mode: "review"}}},
		{"copilot", Config{Type: AgentTypeCopilotACP, CopilotACP: &ACPConfig{Cmd: []string{"custom-copilot"}, ExtraArgs: []string{"--stdio"}, Model: "model", ReasoningEffort: "high", Mode: "review"}}},
		{"grok", Config{Type: AgentTypeGrokACP, GrokACP: &ACPConfig{Cmd: []string{"custom-grok"}, ExtraArgs: []string{"--stdio"}, Model: "model", ReasoningEffort: "high", Mode: "review"}}},
		{"agy", Config{Type: AgentTypeAgyACP, AgyACP: &ACPConfig{Cmd: []string{"custom-agy"}, ExtraArgs: []string{"--stdio"}, Model: "model", ReasoningEffort: "high", Mode: "review"}}},
		{"antigravity", Config{Type: AgentTypeAntigravityACP, AntigravityACP: &ACPConfig{Cmd: []string{"custom-antigravity"}, ExtraArgs: []string{"--stdio"}, Model: "model", ReasoningEffort: "high", Mode: "review"}}},
		{"registry", Config{Type: AgentTypeRegistryACP, RegistryACP: &ACPConfig{Cmd: []string{"custom-registry"}, ExtraArgs: []string{"--stdio"}, Model: "model", ReasoningEffort: "high", Mode: "review"}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NormalizeConfig(tc.cfg, "")
			if err != nil {
				t.Fatal(err)
			}
			if got.Type != AgentTypeGenericACP {
				t.Fatalf("type = %q", got.Type)
			}
			if !reflect.DeepEqual(got.Command, []string{"custom-" + tc.name, "--stdio"}) {
				t.Fatalf("command = %#v", got.Command)
			}
			if got.Model != "model" || got.ReasoningEffort != "high" || got.Mode != "review" {
				t.Fatalf("metadata = %#v", got)
			}
		})
	}
}
