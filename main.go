package main

import (
	"gr/rubibot"
	"log"
)

func main() {

	b, e := rubibot.NewBot(rubibot.Settings{
		Token: "BAJEIC0KRWVJLABHOQGKDIRANRGQFTAELTGQDFYHVIGUAYADUNXUTAEPBGTJHVJZ",
	})

	if e != nil {
		panic("Error")
	}

	b.Handle("/start", func(c rubibot.Context) error {
		c.Send("Welcome!")
		return nil
	})

	b.Handle("/setch", func(c rubibot.Context) error {
		log.Println("/setch ~>")
		return nil
	})

	b.Handle(rubibot.OnText, func(c rubibot.Context) error {
		c.Reply("Hello!")
		return nil
	})

	b.Start()
}
