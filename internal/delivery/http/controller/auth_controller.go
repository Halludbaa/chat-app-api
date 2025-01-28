package controller

import (
	"chatross-api/internal/entity"
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/model"
	"chatross-api/internal/model/converter"
	"chatross-api/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthController struct {
    UseCase *usecase.UserUseCase
}

func NewAuthController(useCase *usecase.UserUseCase) *AuthController{
    return &AuthController{
        UseCase: useCase,
    }
}


func (c *AuthController) Verify(ctx *gin.Context) {
    response, _ := ctx.Get("auth")
    
    if response == nil {
        ctx.JSON(401, rerror.ErrUnauthorized)
        return
    }

    ctx.JSON(200, model.WebResponse[model.UserResponse]{
        Status: 200,
        Data: *converter.UserToResponse(response.(*entity.User)),
    })

}

func (c *AuthController) Register(ctx *gin.Context){
    request := new(model.RegisterUserRequest)

    err := ctx.ShouldBindJSON(request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, rerror.ErrInternalServer)
        return
    }

    response, err := c.UseCase.Create(ctx, request)
    if err != nil {
        ctx.JSON(err.(*rerror.ResponseError).GetCode(), err)
        return
    }

    ctx.JSON(200, model.WebResponse[*model.UserResponse]{
        Status: 200,
        Message: "success to register",
        Data: response,
    })
}

func (c *AuthController) Refresh(ctx *gin.Context) {
    request := new(model.TokenRequest)

    if err := ctx.ShouldBindJSON(request); err != nil {
        ctx.JSON(500, rerror.ErrInternalServer)
    }

    response, err := c.UseCase.Refresh(ctx, request)
    if err != nil {
        ctx.JSON(err.(*rerror.ResponseError).GetCode(), err)
        return
    }

    ctx.JSON(200, model.WebResponse[*model.TokenResponse]{
        Status: 200,
        Message: "success to register",
        Data: response,
    })
}


func (c *AuthController) Login(ctx *gin.Context) {
    request := new(model.LoginUserRequest)

    err := ctx.ShouldBindJSON(request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, rerror.ErrInternalServer)
        return
    }

    response, err := c.UseCase.Login(ctx, request)
    if err != nil {
        ctx.JSON(err.(*rerror.ResponseError).GetCode(), err)
        return
    }

    ctx.JSON(200, model.WebResponse[*model.TokenResponse]{
        Status: 200,
        Message: "success to login",
        Data: response,
    })
}