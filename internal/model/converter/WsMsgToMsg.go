package converter

import (
	"chatross-api/internal/entity"
	"chatross-api/internal/model/wsmodel"
)

func WsMsgToMsg(msg *wsmodel.Message) *entity.Message {
	return &entity.Message{
		ChatID: msg.ChatID,
		From: msg.From,
		To: msg.To,
		Content: msg.Content,
	}
}

func MsgToWsMsg(msg *entity.Message) *wsmodel.Message {
	return &wsmodel.Message{
		ChatID: msg.ChatID,
		From: msg.From,
		To: msg.To,
		Content: msg.Content,
	}
}