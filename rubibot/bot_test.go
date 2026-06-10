package rubibot

import (
	"net/http"
	"testing"
)

func TestNewBot_DefaultSettings(t *testing.T) {
	bot, err := NewBot(Settings{Token: "test-token"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if bot.Token != "test-token" {
		t.Errorf("expected token 'test-token', got %q", bot.Token)
	}
	if bot.URL != "https://botapi.rubika.ir/v3/" {
		t.Errorf("unexpected URL: %s", bot.URL)
	}
	if bot.client == nil {
		t.Error("expected default http client, got nil")
	}
	if bot.Poller == nil {
		t.Error("expected default poller, got nil")
	}
	if cap(bot.Updates) != 100 {
		t.Errorf("expected updates channel cap 100, got %d", cap(bot.Updates))
	}
}

func TestNewBot_MissingToken(t *testing.T) {
	_, err := NewBot(Settings{})
	if err == nil {
		t.Fatal("expected error for missing token, got nil")
	}
}

func TestNewBot_CustomSettings(t *testing.T) {
	customClient := &http.Client{}
	customPoller := &LongPoller{}

	bot, err := NewBot(Settings{
		Token:   "my-token",
		URL:     "https://custom.api/",
		Client:  customClient,
		Poller:  customPoller,
		Updates: 50,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bot.URL != "https://custom.api/" {
		t.Errorf("expected custom URL, got %s", bot.URL)
	}
	if cap(bot.Updates) != 50 {
		t.Errorf("expected updates cap 50, got %d", cap(bot.Updates))
	}
	if bot.client != customClient {
		t.Error("expected custom http client")
	}
}

func TestBot_Handle(t *testing.T) {
	bot, _ := NewBot(Settings{Token: "tok"})
	called := false
	bot.Handle("/start", func(c Context) error {
		called = true
		return nil
	})
	if _, ok := bot.handlers["/start"]; !ok {
		t.Error("handler for /start not registered")
	}
	_ = called
}

func TestBot_Reply_And_Send(t *testing.T) {
	// sendText returns false on network error — Reply/Send should return nil (no error) when sendText returns false
	// and return an error only when sendText returns true (success path is inverted — see bot.go)
	bot, _ := NewBot(Settings{Token: "tok", URL: "http://127.0.0.1:0/"})

	// With an unreachable URL, sendText returns false → Reply returns an error
	err := bot.Reply("chat1", "msg1", "hello")
	if err == nil {
		t.Error("expected error from Reply with unreachable URL")
	}

	err = bot.Send("chat1", "hello")
	if err == nil {
		t.Error("expected error from Send with unreachable URL")
	}
}

func TestError(t *testing.T) {
	err := Error("something went wrong")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if err.Error() != "something went wrong" {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}
