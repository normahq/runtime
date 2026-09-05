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
		{"codex", Config{Type: AgentTypeCodexACP, CodexACP: testACPConfig("custom-codex")}},
		{"claude", Config{Type: AgentTypeClaudeCodeACP, ClaudeCodeACP: testACPConfig("custom-claude")}},
		{"opencode", Config{Type: AgentTypeOpenCodeACP, OpenCodeACP: testACPConfig("custom-opencode")}},
		{"copilot", Config{Type: AgentTypeCopilotACP, CopilotACP: testACPConfig("custom-copilot")}},
		{"grok", Config{Type: AgentTypeGrokACP, GrokACP: testACPConfig("custom-grok")}},
		{"agy", Config{Type: AgentTypeAgyACP, AgyACP: testACPConfig("custom-agy")}},
		{"antigravity", Config{Type: AgentTypeAntigravityACP, AntigravityACP: testACPConfig("custom-antigravity")}},
		{"registry", Config{Type: AgentTypeRegistryACP, RegistryACP: testACPConfig("custom-registry")}},
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
			if got.ModelConfigID != "provider_model" || got.ReasoningEffortConfigID != "thought_level" {
				t.Fatalf("config option IDs = %#v", got)
			}
		})
	}
}

func testACPConfig(command string) *ACPConfig {
	return &ACPConfig{
		Cmd:                     []string{command},
		ExtraArgs:               []string{"--stdio"},
		Model:                   "model",
		ModelConfigID:           "provider_model",
		ReasoningEffort:         "high",
		ReasoningEffortConfigID: "thought_level",
		Mode:                    "review",
	}
}
