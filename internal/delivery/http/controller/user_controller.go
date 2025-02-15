package controller

import (
	"chatross-api/internal/entity"
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase		*usecase.UserUsecase
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase,
	}
}

func (uc *UserController) GetUserChat(ctx *gin.Context) {
	user, _ := ctx.Get("auth")

	response, err := uc.userUsecase.GetChat(ctx, user.(*entity.User))
	if err != nil {
		ctx.JSON(err.(*rerror.ResponseError).Code, err)
	}

	ctx.JSON(http.StatusOK, response)
}