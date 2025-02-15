package controller

import (
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	useCase		*usecase.ChatUsecase
}

func NewChatController(useCase *usecase.ChatUsecase) *ChatController {
	return &ChatController{
		useCase,
	}
}

func (cc *ChatController) GetChatMessage(ctx *gin.Context) {
	chatID, _ := strconv.Atoi(ctx.Param("chat_id")) 

	response, err := cc.useCase.GetChatwithMessage(ctx, int64(chatID))
	if err != nil {
		ctx.JSON(err.(*rerror.ResponseError).Code, err)
	}

	ctx.JSON(http.StatusOK, response)
}