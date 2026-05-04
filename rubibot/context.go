package rubibot

type HandlerFunc func(Context) error

type Context interface {
	Bot() API
	Update() Update
	Send(what string) error
	Reply(what string) error
}

type nativeContext struct {
	b     API
	u     Update
	store map[string]interface{}
}

func NewContext(b API, u Update) Context {
	return &nativeContext{
		b: b,
		u: u,
	}
}

func (c *nativeContext) Bot() API {
	return c.b
}

func (c *nativeContext) Update() Update {
	return c.u
}

func (c *nativeContext) Send(what string) error {
	err := c.b.Send(c.u.ChatID, what)
	return err
}

func (c *nativeContext) Reply(what string) error {
	err := c.b.Reply(c.u.ChatID, c.u.Message.MessageID, what)
	return err
}
