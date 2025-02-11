package wsmodel



type Message struct {
	ChatID int64 `json:"chat_id,omitempty"`
	From string `json:"from"`
	To string	`json:"to"`
	Content string `json:"content"`
}