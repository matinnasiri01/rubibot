package common

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var NewMessages = make(chan UpdateEvent)

func (hc *HTCL) Update(duration time.Duration) {
	t := time.NewTicker(duration)
	off := SaveOff("")
	var apiResponse ApiResponse
	go func() {
		for range t.C {
			hc.POST(
				"getUpdates",
				GetUpdatesPayload{OffsetID: off},
				func(b []byte) {
					_ = json.Unmarshal(b, &apiResponse)
					off = SaveOff(apiResponse.Data.NextOffsetID)
					newMess(apiResponse.Data.Updates)
				},
			)
		}

	}()
}

func SaveOff(newOff string) string {
	file, _ := os.OpenFile("off.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if newOff != "" {
		fmt.Fprint(file, newOff)
	}
	s, _ := os.ReadFile(file.Name())
	return string(s)
}

func newMess(s []UpdateEvent) {
	for _, i := range s {
		if i.Type == "NewMessage" {
			NewMessages <- i
		}
	}
}

func (hc *HTCL) SendMessage(chat_id, text string, reply *string) {
	hc.POST(
		"sendMessage",
		SendMessagePayload{Text: text, ChatID: chat_id, ReplyTo: reply},
		func(b []byte) {
			println("Send!")
		},
	)
}
