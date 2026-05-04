package rubibot

type Poller interface {
	Poll(b *Bot, data chan Update, stop chan struct{})
}

type LongPoller struct {
	LastUpdateID string
}

func (p *LongPoller) Poll(b *Bot, dest chan Update, stop chan struct{}) {
	for {
		select {
		case <-stop:
			return
		default:
		}

		data, err := b.getUpdates(p.LastUpdateID)
		if err != nil {
			continue
		}

		for _, update := range data.Updates {
			dest <- update
		}
		p.LastUpdateID = data.NextOffsetID
	}
}
