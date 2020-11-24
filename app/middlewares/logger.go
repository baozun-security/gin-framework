package middlewares

import (
	"baozun.com/framework/app/pkgs/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Logger.WithFields(logrus.Fields{
			"client_ip": c.ClientIP(),
			"params":    c.Request.URL.Query(),
		}).Infof("%s %s", c.Request.Method, c.Request.URL.Path)

		c.Next()
	}
}
