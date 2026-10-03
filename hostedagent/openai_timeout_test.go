package hostedagent

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"google.golang.org/adk/v2/model"
	"google.golang.org/genai"
)

func TestOpenAIModelTimeout(t *testing.T) {
	defaultModel, err := NewOpenAIModel("key", "model")
	if err != nil {
		t.Fatal(err)
	}
	if defaultModel.client.Timeout != 30*time.Second {
		t.Fatalf("default timeout = %v", defaultModel.client.Timeout)
	}
	configuredModel, err := NewOpenAIModelWithTimeout("key", "model", 180*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if configuredModel.client.Timeout != 180*time.Second {
		t.Fatalf("configured timeout = %v", configuredModel.client.Timeout)
	}
	if _, err := NewOpenAIModelWithTimeout("key", "model", 0); err == nil {
		t.Fatal("zero timeout accepted")
	}
	if _, err := NewOpenAIModelWithOptions("key", "model", OpenAIModelOptions{ReasoningEffort: "high"}); err == nil {
		t.Fatal("high reasoning effort accepted without reasoning-content round trip")
	}
}

func TestOpenAIModelReasoningEffortWire(t *testing.T) {
	for _, tc := range []struct {
		reasoningEffort string
		want            bool
	}{
		{reasoningEffort: "", want: false},
		{reasoningEffort: "none", want: true},
	} {
		m, err := NewOpenAIModelWithOptions("key", "deepseek-v4-pro", OpenAIModelOptions{ReasoningEffort: tc.reasoningEffort})
		if err != nil {
			t.Fatal(err)
		}
		m.client.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
			var payload map[string]any
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				return nil, err
			}
			thinking, present := payload["thinking"].(map[string]any)
			if present != tc.want || (present && thinking["type"] != "disabled") {
				t.Errorf("thinking payload = %v; effort %q", payload["thinking"], tc.reasoningEffort)
			}
			return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"choices":[{"message":{"role":"assistant","content":"ok"}}]}`)), Header: make(http.Header), Request: r}, nil
		})
		_, err = m.generate(context.Background(), &model.LLMRequest{Contents: []*genai.Content{genai.NewContentFromText("hello", genai.RoleUser)}})
		if err != nil {
			t.Fatal(err)
		}
	}
}
