package rubibot

import (
	"os"
)

type Poller interface {
	Poll(b *Bot, stop chan struct{})
}

type LongPoller struct{}

func (p *LongPoller) Poll(b *Bot, stop chan struct{}) {
	for {
		select {
		case <-stop:
			return
		default:
		}

		data, err := b.getUpdates(readOrCreateCard())
		if err != nil {
			continue
		}
		for _, update := range data.Updates {
			b.Updates <- update
		}
		if data.NextOffsetID != "" {
			saveCardNumber(data.NextOffsetID)

		}
	}
}

// HOT TEST!
const filename = "next.offset"

func readOrCreateCard() string {
	data, err := os.ReadFile(filename)
	if err == nil {
		return string(data)
	}
	return ""
}

func saveCardNumber(number string) {
	os.WriteFile(filename, []byte(number), 0644)
}
