package controllers

import (
	v1 "baozun.com/leak/app/controllers/v1"
	"baozun.com/leak/app/pkgs/response"
	e "errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 设置路由中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) { // 404 route
		g := response.Gin{Ctx: c}
		g.ApiFail(http.StatusNotFound, e.New("找不到路由"))
	})
	router.NoMethod(func(c *gin.Context) { // 404 method
		g := response.Gin{Ctx: c}
		g.ApiFail(http.StatusNotFound, e.New("找不到方法"))
	})
	router.GET("/ping", func(c *gin.Context) { // ping
		g := response.Gin{Ctx: c}
		g.ApiSuccess("pong")
	})

	// api v1
	v1.RegisterApiRouter(router.Group("/api/v1"))

	return router
}
