package rubibot

type Data struct {
	Updates      []Update `json:"updates"`
	NextOffsetID string   `json:"next_offset_id"`
}

type Update struct {
	Type       string  `json:"type"`
	ChatID     string  `json:"chat_id"`
	Message    Message `json:"new_message"`
	UpdateTime int64   `json:"update_time"`
}

type Message struct {
	MessageID        string `json:"message_id"`
	Text             string `json:"text"`
	Time             string `json:"time"`
	IsEdited         bool   `json:"is_edited"`
	SenderType       string `json:"sender_type"`
	SenderID         string `json:"sender_id"`
	ReplyToMessageID string `json:"reply_to_message_id,omitempty"`
}

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
