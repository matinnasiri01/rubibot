package rubibot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

type Settings struct {
	URL     string
	Token   string
	Client  *http.Client
	Poller  Poller
	Updates int
}

type Bot struct {
	URL      string
	Token    string
	Poller   Poller
	Updates  chan Update
	handlers map[string]HandlerFunc
	stop     chan chan struct{}
	client   *http.Client
}

var (
	cmdRx = regexp.MustCompile(`^(/\w+)(?:\s+(\S+))?`)
)

const OnText = "On-Text"

func NewBot(pref Settings) (*Bot, error) {

	if pref.Updates == 0 {
		pref.Updates = 100
	}

	client := pref.Client
	if client == nil {
		client = &http.Client{Timeout: time.Minute}
	}

	if pref.URL == "" {
		pref.URL = "https://botapi.rubika.ir/v3/"
	}

	if pref.Token == "" {
		return nil, Error("Token!!!")
	}

	if pref.Poller == nil {
		pref.Poller = &LongPoller{}
	}

	return &Bot{

		URL:    pref.URL,
		Token:  pref.Token,
		Poller: pref.Poller,
		client: client,

		Updates:  make(chan Update, pref.Updates),
		handlers: make(map[string]HandlerFunc),
		stop:     make(chan chan struct{}),
	}, nil

}

func (b *Bot) Start() {
	if b.Poller == nil {
		panic("rubibot: can't start without a poller")
	}

	stop := make(chan struct{})
	stopConfirm := make(chan struct{})

	go func() {
		b.Poller.Poll(b, stop)
		close(stopConfirm)
	}()
	fmt.Println("Bot Successfully started!")
	for {
		select {
		case upd := <-b.Updates:
			b.ProcessUpdate(upd)
		case confirm := <-b.stop:
			close(stop)
			<-stopConfirm
			close(confirm)
			return
		}
	}
}

func (b *Bot) Handle(endpoint string, h HandlerFunc) {
	b.handlers[endpoint] = h
}

func (b *Bot) Raw(method string, payload interface{}) ([]byte, error) {
	url := b.URL + b.Token + "/" + method

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	resp.Close = true
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (b *Bot) getUpdates(offset string) (Data, error) {

	params := map[string]string{
		"offset_id": offset,
	}

	var emp Data
	data, err := b.Raw("getUpdates", params)
	if err != nil {
		return emp, err
	}

	var resp Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return emp, err
	}

	return resp.Data, nil
}

func (b *Bot) sendText(to, text, mesId string) bool {
	params := map[string]string{
		"chat_id": to,
		"text":    text,
	}
	if mesId != "" {
		params["reply_to_message_id"] = mesId
	}

	_, err := b.Raw("sendMessage", params)

	if err != nil {
		return false
	}

	return true
}

func (b *Bot) Reply(to, mes, what string) error {
	err := b.sendText(to, what, mes)
	if err == true {
		return Error("Bot/Reply")
	}
	return nil
}

func (b *Bot) Send(to, what string) error {
	err := b.sendText(to, what, "")
	if err == true {
		return Error("Bot/Send")
	}
	return nil
}

func (b *Bot) NewContext(u Update) Context {
	return NewContext(b, u)
}

func Error(text string) error {
	return errors.New(text)
}
