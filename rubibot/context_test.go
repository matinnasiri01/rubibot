package rubibot

import (
	"errors"
	"testing"
)

// mockAPI implements the API interface for testing.
type mockAPI struct {
	replyCalled bool
	sendCalled  bool
	replyTo     string
	replyMes    string
	replyWhat   string
	sendTo      string
	sendWhat    string
	returnErr   error
}

func (m *mockAPI) Reply(to, mes, what string) error {
	m.replyCalled = true
	m.replyTo = to
	m.replyMes = mes
	m.replyWhat = what
	return m.returnErr
}

func (m *mockAPI) Send(to, what string) error {
	m.sendCalled = true
	m.sendTo = to
	m.sendWhat = what
	return m.returnErr
}

func makeUpdate(chatID, messageID, text string) Update {
	return Update{
		ChatID: chatID,
		Message: Message{
			MessageID: messageID,
			Text:      text,
		},
	}
}

func TestNewContext(t *testing.T) {
	api := &mockAPI{}
	u := makeUpdate("chat1", "msg1", "hello")
	ctx := NewContext(api, u)

	if ctx.Bot() != api {
		t.Error("Bot() should return the provided API")
	}
	if ctx.Update().ChatID != "chat1" {
		t.Errorf("Update() ChatID mismatch: got %s", ctx.Update().ChatID)
	}
}

func TestContext_Send(t *testing.T) {
	api := &mockAPI{}
	u := makeUpdate("chat42", "msg1", "hi")
	ctx := NewContext(api, u)

	err := ctx.Send("world")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !api.sendCalled {
		t.Error("expected Send to be called on API")
	}
	if api.sendTo != "chat42" {
		t.Errorf("expected sendTo 'chat42', got %s", api.sendTo)
	}
	if api.sendWhat != "world" {
		t.Errorf("expected sendWhat 'world', got %s", api.sendWhat)
	}
}

func TestContext_Send_Error(t *testing.T) {
	api := &mockAPI{returnErr: errors.New("send failed")}
	ctx := NewContext(api, makeUpdate("c1", "m1", "t"))
	err := ctx.Send("msg")
	if err == nil {
		t.Error("expected error from Send")
	}
}

func TestContext_Reply(t *testing.T) {
	api := &mockAPI{}
	u := makeUpdate("chat99", "msg99", "hello")
	ctx := NewContext(api, u)

	err := ctx.Reply("a reply")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !api.replyCalled {
		t.Error("expected Reply to be called on API")
	}
	if api.replyTo != "chat99" {
		t.Errorf("expected replyTo 'chat99', got %s", api.replyTo)
	}
	if api.replyMes != "msg99" {
		t.Errorf("expected replyMes 'msg99', got %s", api.replyMes)
	}
	if api.replyWhat != "a reply" {
		t.Errorf("expected replyWhat 'a reply', got %s", api.replyWhat)
	}
}

func TestContext_Reply_Error(t *testing.T) {
	api := &mockAPI{returnErr: errors.New("reply failed")}
	ctx := NewContext(api, makeUpdate("c1", "m1", "t"))
	err := ctx.Reply("msg")
	if err == nil {
		t.Error("expected error from Reply")
	}
}
