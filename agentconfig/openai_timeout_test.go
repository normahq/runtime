package agentconfig

import (
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

const testReasoningEffortNone = "none"

func TestOpenAITimeoutNormalization(t *testing.T) {
	for _, tc := range []struct {
		value string
		want  time.Duration
		bad   bool
	}{
		{value: "", want: 0},
		{value: "180s", want: 180 * time.Second},
		{value: "0s", bad: true},
		{value: "-1s", bad: true},
		{value: "abc", bad: true},
	} {
		cfg := Config{Type: AgentTypeOpenAI, OpenAI: &LocalAPIConfig{APIKey: "key", Model: "model", Timeout: tc.value}}
		got, err := NormalizeConfig(cfg, "")
		if tc.bad {
			if err == nil || cfg.Validate() == nil {
				t.Fatalf("timeout %q accepted", tc.value)
			}
			continue
		}
		if err != nil || got.Timeout != tc.want {
			t.Fatalf("timeout %q: got %v, err %v; want %v", tc.value, got.Timeout, err, tc.want)
		}
	}
}

func TestOpenAIReasoningEffortYAMLField(t *testing.T) {
	var cfg Config
	if err := yaml.Unmarshal([]byte("type: openai\nopenai:\n  api_key: key\n  model: model\n  reasoning_effort: none\n"), &cfg); err != nil {
		t.Fatal(err)
	}
	if cfg.OpenAI == nil || cfg.OpenAI.ReasoningEffort != testReasoningEffortNone {
		t.Fatalf("openai.reasoning_effort = %+v", cfg.OpenAI)
	}
}

func TestOpenAIReasoningEffortNormalization(t *testing.T) {
	for _, reasoningEffort := range []string{"", testReasoningEffortNone, "low", "enabled"} {
		cfg := Config{Type: AgentTypeOpenAI, OpenAI: &LocalAPIConfig{APIKey: "key", Model: "model", ReasoningEffort: reasoningEffort}}
		got, normalizeErr := NormalizeConfig(cfg, "")
		validateErr := cfg.Validate()
		if reasoningEffort != "" && reasoningEffort != testReasoningEffortNone {
			if normalizeErr == nil || validateErr == nil {
				t.Fatalf("invalid reasoning_effort %q accepted", reasoningEffort)
			}
			continue
		}
		if normalizeErr != nil || validateErr != nil || got.ReasoningEffort != reasoningEffort {
			t.Fatalf("reasoning_effort %q: got %q, normalize=%v, validate=%v", reasoningEffort, got.ReasoningEffort, normalizeErr, validateErr)
		}
	}
}
