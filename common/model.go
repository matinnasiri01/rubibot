package common

type Message struct {
	MessageID  string `json:"message_id"`
	Text       string `json:"text"`
	Time       string `json:"time"`
	IsEdited   bool   `json:"is_edited"`
	SenderType string `json:"sender_type"`
	SenderID   string `json:"sender_id"`
}

type UpdateEvent struct {
	Type       string   `json:"type"`
	ChatID     string   `json:"chat_id"`
	NewMessage *Message `json:"new_message,omitempty"`
	UpdateTime int64    `json:"update_time"`
}

type Data struct {
	Updates      []UpdateEvent `json:"updates"`
	NextOffsetID string        `json:"next_offset_id"`
}

type ApiResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type GetUpdatesPayload struct {
	OffsetID string `json:"offset_id"`
}

type SendMessagePayload struct {
	Text    string  `json:"text"`
	ChatID  string  `json:"chat_id"`
	ReplyTo *string `json:"reply_to_message_id"`
}

type Empty struct{}
