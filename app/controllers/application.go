package controllers

import (
	v1 "baozun.com/leak/app/controllers/v1"
	"baozun.com/leak/app/pkgs/response"
	e "errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// initialize routing information
func Init() *gin.Engine {
	engine := gin.New()
	// 设置路由中间件
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// 404
	engine.NoRoute(func(c *gin.Context) {
		g := response.Gin{Ctx: c}
		g.ApiFail(http.StatusNotFound, e.New("请求方法不存在"))
	})
	// ping
	engine.GET("/ping", func(c *gin.Context) {
		g := response.Gin{Ctx: c}
		g.ApiSuccess("pong")
	})
	// api v1
	v1.Resources(engine.Group("/api/v1"))

	return engine
}
