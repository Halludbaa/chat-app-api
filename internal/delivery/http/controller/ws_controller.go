package controller

import (
	"chatross-api/internal/delivery/websockets"
	"chatross-api/internal/entity"
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/model"

	"github.com/gin-gonic/gin"
)

type WsController struct {
	Hub *websockets.Hub
}

func NewWsController(hub *websockets.Hub) *WsController {
	return &WsController{
		Hub: hub,
	}
}

func(c *WsController) Connect(ctx *gin.Context) {
	user, exist := ctx.Get("auth")
	if !exist {
		ctx.JSON(401, rerror.ErrUnauthorized)
	}
	userIDstr := user.(*entity.User).ID

	// userIDstr := strconv.Itoa((userID))

	websockets.ServeWS(c.Hub, ctx, &userIDstr)
}

func (c *WsController) GetClient(ctx *gin.Context) {
	data := make([]string, 0, len(c.Hub.Clients))
	for key := range c.Hub.Clients{
		data = append(data, key)
	}

	ctx.JSON(200, model.WebResponse[any]{
		Status: 200,
		Data: data,
	})
}