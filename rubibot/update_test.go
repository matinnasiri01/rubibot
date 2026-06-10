package rubibot

import (
	"testing"
	"time"
)

func newTestBot(t *testing.T) *Bot {
	t.Helper()
	bot, err := NewBot(Settings{Token: "test-token"})
	if err != nil {
		t.Fatalf("failed to create bot: %v", err)
	}
	return bot
}

func waitFor(t *testing.T, ch <-chan struct{}) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(2 * time.Second):
		t.Error("timed out waiting for handler to be called")
	}
}

func TestProcessUpdate_Command(t *testing.T) {
	bot := newTestBot(t)
	done := make(chan struct{})

	bot.Handle("/start", func(c Context) error {
		close(done)
		return nil
	})

	bot.ProcessUpdate(makeUpdate("chat1", "msg1", "/start"))
	waitFor(t, done)
}

func TestProcessUpdate_OnText(t *testing.T) {
	bot := newTestBot(t)
	done := make(chan struct{})

	bot.Handle(OnText, func(c Context) error {
		close(done)
		return nil
	})

	bot.ProcessUpdate(makeUpdate("chat1", "msg1", "just some text"))
	waitFor(t, done)
}

func TestProcessUpdate_NoHandlerForCommand(t *testing.T) {
	// Should not panic when no handler is registered
	bot := newTestBot(t)
	bot.ProcessUpdate(makeUpdate("chat1", "msg1", "/unknown"))
}

func TestProcessUpdate_EmptyText(t *testing.T) {
	// Empty text — ProcessContext returns early; no panic expected
	bot := newTestBot(t)
	u := Update{ChatID: "chat1", Message: Message{MessageID: "m1", Text: ""}}
	bot.ProcessUpdate(u)
}

func TestProcessContext_CommandWithArgument(t *testing.T) {
	bot := newTestBot(t)
	done := make(chan struct{})
	var receivedText string

	bot.Handle("/echo", func(c Context) error {
		receivedText = c.Update().Message.Text
		close(done)
		return nil
	})

	bot.ProcessContext(bot.NewContext(makeUpdate("chat1", "msg1", "/echo hello")))
	waitFor(t, done)

	if receivedText != "/echo hello" {
		t.Errorf("expected '/echo hello', got %q", receivedText)
	}
}

func TestHandle_CommandFallsBackToOnText(t *testing.T) {
	bot := newTestBot(t)
	done := make(chan struct{})

	bot.Handle(OnText, func(c Context) error {
		close(done)
		return nil
	})

	bot.ProcessContext(bot.NewContext(makeUpdate("c", "m", "/unknown")))
	waitFor(t, done)
}
