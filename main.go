package main

import (
	"gr/common"
	"gr/rubika"
	"time"
)

func main() {
	new := rubika.RubikaHC()
	println("~>Connect!")
	new.Update(5 * time.Second)
	for num := range common.NewMessages {
		println("~>", num.NewMessage.Text)
		new.SendMessage(num.ChatID, num.NewMessage.Text, &num.NewMessage.MessageID)
	}

	select {}
}
