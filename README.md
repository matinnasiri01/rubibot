# 🎧 Rubibot

A lightweight, fast, and type-safe Go client library for building bots on the **Rubika messaging platform** using the [official Bot API](https://rubika.ir/botapi).

---

## Features

- 🚀 Long Polling support (`getUpdates`)
- 🔄 Automatic offset management
- ⚡ Simple command handler system
- 🧠 Type-safe API responses
- 📦 Zero external dependencies
- 🛠 Easy to extend and customize

---

## 📦 Installation

```bash
go get github.com/matinnasiri01/rubibot
```

## ⚡ Quick Start
```bash
package main

import (
	"log"

	"github.com/matinnasiri01/rubibot"
)

func main() {

	bot, err := rubibot.NewBot(rubibot.Settings{
		Token: "YOUR_BOT_TOKEN_FROM_BOTFATHER",
	})

	if err != nil {
		log.Fatal(err)
	}

	// /start command
	bot.Handle("/start", func(c rubibot.Context) error {
		return c.Send("Welcome 👋")
	})

	// custom command
	bot.Handle("/setch", func(c rubibot.Context) error {
		log.Println("setch command triggered")
		return nil
	})

	// fallback handler (all text messages)
	bot.Handle(rubibot.OnText, func(c rubibot.Context) error {
		return c.Reply("Hello! 👋")
	})

	bot.Start()
}
```

## ⚙ Configuration
```bash
type Settings struct {
	URL     string
	Token   string
	Client  *http.Client
	Poller  Poller
	Updates int
}
```

## 🤖 Bot Structure
```bash
type Bot struct {
	URL      string
	Token    string
	Poller   Poller
	Updates  chan Update
	handlers map[string]HandlerFunc
	stop     chan struct{}
	client   *http.Client
}
```

## 🔌 API Layer
```bash
type API interface {
	Reply(to, mes, what string) error
	Send(to, what string) error

	// TODO: extend with inline keyboards, media support, etc.
}
```

## 🎧 Spotify
Vibe while building bots:
👉 [open.spotify](https://open.spotify.com/track/12ciTjJXFmroanclpa7syE?si=rer9XwD4TYy9MCS9XV5PYw)

<br/>

#### Made with 🤬 by Matin Nasiri