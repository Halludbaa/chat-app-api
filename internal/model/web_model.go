package model

type WebResponse[T any] struct {
	Status 	int 		`json:"status"`
	Message string		`json:"message,omitempty"`
	Data 	T			`json:"data"`
}