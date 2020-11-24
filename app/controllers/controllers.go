package controllers

import (
	v1 "baozun.com/leak/app/controllers/v1"
	"baozun.com/leak/app/middlewares"
	"baozun.com/leak/app/pkgs/response"
	e "errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// apply panic recovery for server fallback
	router.Use(middlewares.Recovery()) // 异常处理
	router.Use(middlewares.Logger())   // 请求审计

	router.NoRoute(func(c *gin.Context) { // 404 route
		g := response.Gin{Ctx: c}
		g.ApiFail(http.StatusNotFound, e.New("not found route"))
	})
	router.NoMethod(func(c *gin.Context) { // 404 method
		g := response.Gin{Ctx: c}
		g.ApiFail(http.StatusNotFound, e.New("not found method"))
	})
	router.GET("/ping", func(c *gin.Context) { // ping
		g := response.Gin{Ctx: c}
		g.ApiSuccess("pong")
	})

	// api v1
	v1.RegisterApiRouter(router.Group("/api/v1"))

	return router
}
