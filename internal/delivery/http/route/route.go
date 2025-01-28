package route

import (
	"chatross-api/internal/delivery/http/controller"
	"chatross-api/internal/delivery/http/middleware"
	"chatross-api/internal/model"

	"github.com/gin-gonic/gin"
)

type  RouteConfig struct {
	App 	*gin.Engine
	AuthController *controller.AuthController
	AuthMiddleware gin.HandlerFunc
}

func (c *RouteConfig) Setup(){
	c.SetupMiddleware()
	c.SetupGuestRoute()
	c.SetupForTest()
}

func (c *RouteConfig) SetupMiddleware() {
	c.App.Use(middleware.CORSMiddleware())
	c.App.Use(gin.Recovery())
	c.App.Use(gin.Logger())
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
		v1.POST("/user/_login", c.AuthController.Login)
		v1.POST("/_refresh", c.AuthController.Refresh)
		c.SetupAuthRoute(v1)
	}

	
}