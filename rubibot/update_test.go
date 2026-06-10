package rubibot

import (
	"testing"
)

func newTestBot(t *testing.T) *Bot {
	t.Helper()
	bot, err := NewBot(Settings{Token: "test-token"})
	if err != nil {
		t.Fatalf("failed to create bot: %v", err)
	}
	return bot
}

func TestProcessUpdate_Command(t *testing.T) {
	bot := newTestBot(t)
	handled := false

	bot.Handle("/start", func(c Context) error {
		handled = true
		return nil
	})

	u := makeUpdate("chat1", "msg1", "/start")
	bot.ProcessUpdate(u)

	// handler runs in a goroutine; give it a moment
	for i := 0; i < 100 && !handled; i++ {
		// spin briefly
	}
}

func TestProcessUpdate_OnText(t *testing.T) {
	bot := newTestBot(t)
	called := false

	bot.Handle(OnText, func(c Context) error {
		called = true
		return nil
	})

	u := makeUpdate("chat1", "msg1", "just some text")
	bot.ProcessUpdate(u)

	for i := 0; i < 100 && !called; i++ {
		// spin briefly
	}
}

func TestProcessUpdate_NoHandlerForCommand(t *testing.T) {
	// Should not panic when no handler is registered for the command or OnText
	bot := newTestBot(t)
	u := makeUpdate("chat1", "msg1", "/unknown")
	bot.ProcessUpdate(u) // should not panic
}

func TestProcessUpdate_EmptyText(t *testing.T) {
	// Message with empty text — ProcessContext returns early; no panic expected
	bot := newTestBot(t)
	u := Update{ChatID: "chat1", Message: Message{MessageID: "m1", Text: ""}}
	bot.ProcessUpdate(u) // should not panic
}

func TestProcessContext_CommandWithArgument(t *testing.T) {
	bot := newTestBot(t)
	var receivedText string

	bot.Handle("/echo", func(c Context) error {
		receivedText = c.Update().Message.Text
		return nil
	})

	u := makeUpdate("chat1", "msg1", "/echo hello")
	bot.ProcessContext(bot.NewContext(u))

	for i := 0; i < 1000 && receivedText == ""; i++ {
	}

	if receivedText != "/echo hello" {
		t.Errorf("expected full message text, got %q", receivedText)
	}
}

func TestHandle_CommandFallsBackToOnText(t *testing.T) {
	bot := newTestBot(t)
	var onTextCalled bool

	// Register OnText but NOT the /unknown command
	bot.Handle(OnText, func(c Context) error {
		onTextCalled = true
		return nil
	})

	u := makeUpdate("c", "m", "/unknown")
	bot.ProcessContext(bot.NewContext(u))

	for i := 0; i < 1000 && !onTextCalled; i++ {
	}

	if !onTextCalled {
		t.Error("expected OnText handler to be called as fallback")
	}
}
