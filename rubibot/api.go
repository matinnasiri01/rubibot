package rubibot

type API interface {
	Reply(to, mes, what string) error
	Send(to, what string) error
}
