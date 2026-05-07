package rubibot

func (b *Bot) ProcessUpdate(u Update) {
	b.ProcessContext(b.NewContext(u))
}

func (b *Bot) ProcessContext(c Context) {
	u := c.Update()
	m := u.Message
	if m.Text != "" {

		match := cmdRx.FindAllStringSubmatch(m.Text, -1)
		if match != nil {
			command := match[0][1]
			if b.handle(command, c) {
				return
			}
		}

		b.handle(OnText, c)
		return
	}
}

func (b *Bot) handle(end string, c Context) bool {
	if handler, ok := b.handlers[end]; ok {
		b.runHandler(handler, c)
		return true
	}
	return false
}

func (b *Bot) runHandler(h HandlerFunc, c Context) {
	f := func() {
		if err := h(c); err != nil {
			panic(err)
		}
	}
	go f()
}
