package route

import (
	"chatross-api/internal/delivery/http/controller"
	"chatross-api/internal/delivery/http/middleware"
	"chatross-api/internal/delivery/websockets"
	"chatross-api/internal/model"

	"github.com/gin-gonic/gin"
)

type  RouteConfig struct {
	App 				*gin.Engine
	AuthController 		*controller.AuthController
	AuthMiddleware		gin.HandlerFunc
	Hub 				*websockets.Hub
	WebsocketHandler 	*controller.WsController
}

func (c *RouteConfig) Setup(){
	c.SetupMiddleware()
	c.SetupGuestRoute()
	c.SetupForTest()
	c.SetupWebsocket()
}

func (c *RouteConfig) SetupMiddleware() {
	c.App.Use(gin.Recovery(), gin.Logger(), middleware.CORSMiddleware())
}

func (c *RouteConfig) SetupWebsocket() {
	websocket := c.App.Group("")
	websocket.Use(c.AuthMiddleware)
	{
		websocket.GET("/ws", c.WebsocketHandler.Connect)
		websocket.GET("/ws/_client", c.WebsocketHandler.GetClient)
	}
}

func (c *RouteConfig) SetupAuthRoute(parent *gin.RouterGroup) {
	authorized := parent.Group("")
	authorized.Use(c.AuthMiddleware)
	{
		authorized.GET("/_verify", c.AuthController.Verify)
	}
}

func (c *RouteConfig) SetupForTest() {
	c.App.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, model.WebResponse[string]{
			Status: 200,
			Data: "pong",
		})
	})
}

func (c *RouteConfig) SetupGuestRoute() {
	v1 := c.App.Group("/api")
	{
		v1.POST("/user", c.AuthController.Register)
		v1.POST("/_login", c.AuthController.Login)
		v1.POST("/_refresh", c.AuthController.Refresh)
		c.SetupAuthRoute(v1)
	}

	
}