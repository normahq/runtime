package taskmaster

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"
)

// ErrUnsupported reports that a target cannot handle the requested locator operation.
var ErrUnsupported = errors.New("unsupported locator operation")

// Target dispatches messages to supported locators.
type Target interface {
	Supports(locator Locator) bool
	DispatchMessage(ctx context.Context, msg Message) error
}

type targetRegistry struct {
	targets []Target
}

func newTargetRegistry(targets []Target) targetRegistry {
	cloned := make([]Target, 0, len(targets))
	for _, target := range targets {
		if target != nil {
			cloned = append(cloned, target)
		}
	}
	return targetRegistry{targets: cloned}
}

func (r targetRegistry) supports(locator Locator) bool {
	for _, target := range r.targets {
		if target.Supports(locator) {
			return true
		}
	}
	return false
}

func (r targetRegistry) dispatchMessage(ctx context.Context, msg Message) error {
	supported := false
	for _, target := range r.targets {
		if !target.Supports(msg.To) {
			continue
		}
		supported = true
		err := target.DispatchMessage(ctx, msg)
		if err == nil {
			return nil
		}
		if errors.Is(err, ErrUnsupported) {
			continue
		}
		return err
	}
	if supported {
		return fmt.Errorf("unsupported dispatch for locator %s", msg.To)
	}
	return fmt.Errorf("no target for locator %s", msg.To)
}

// CLILogTarget writes CLI log messages through zerolog.
type CLILogTarget struct {
	logger zerolog.Logger
}

// NewCLILogTarget constructs a target for the built-in CLI log locator.
func NewCLILogTarget(logger zerolog.Logger) Target {
	return CLILogTarget{logger: logger}
}

// Supports reports whether the target handles the locator.
func (t CLILogTarget) Supports(locator Locator) bool {
	return locator.Class == LocatorClassIntegration &&
		locator.Transport == LocatorTransportCLI &&
		locator.Key == CLILogKey
}

// DispatchMessage writes a message to the configured logger.
func (t CLILogTarget) DispatchMessage(_ context.Context, msg Message) error {
	message := msg.Content
	if message == "" {
		message = "(empty message content)"
	}
	t.logger.Info().
		Str("message_id", msg.ID).
		Str("session_id", msg.SessionID).
		Str("kind", msg.Kind).
		Str("locator", msg.To.String()).
		Str("message_text", message).
		Msg("cli log delivered")
	return nil
}
