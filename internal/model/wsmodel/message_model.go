package wsmodel



type Message struct {
	Type 	string `json:"type,omitempty"`
	ChatID int64 `json:"chat_id,omitempty"`
	From string `json:"from"`
	To string	`json:"to"`
	Content string `json:"content"`
}