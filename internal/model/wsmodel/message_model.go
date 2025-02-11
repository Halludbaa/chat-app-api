package wsmodel



type Message struct {
	ChatID string `json:"chat_id"`
	From string `json:"from"`
	To string	`json:"to"`
	Content string `json:"content"`
}