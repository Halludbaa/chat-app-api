package middleware

import (
	"chatross-api/internal/helper"
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewAuth(userUseCase *usecase.UserUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(401, rerror.ErrUnauthorized)
			return
		}
		
		userID, err := helper.ValidateAccessToken(token)
		if err != nil {
			ctx.JSON(err.(*rerror.ResponseError).GetCode(), 
							err)
			ctx.Abort()
			return
		}
		
		auth, err := userUseCase.Verify(ctx, userID);
		if err != nil {
			ctx.JSON(err.(*rerror.ResponseError).GetCode(), err)
			ctx.Abort()
			return
		}
		fmt.Println(auth)

		ctx.Set("auth", auth)
		ctx.Next()
	}
} 