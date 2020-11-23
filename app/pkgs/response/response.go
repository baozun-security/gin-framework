package response

import (
	"baozun.com/leak/app/pkgs/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) ApiSuccess(data interface{}) {
	g.Ctx.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
	return
}

func (g *Gin) ApiFail(httpCode int, err error) {
	businessError, ok := err.(*errors.BusinessError)
	if ok {
		g.Ctx.JSON(httpCode, Response{
			Code:    businessError.Code,
			Message: businessError.Error(),
		})
	} else {
		g.Ctx.JSON(httpCode, Response{
			Code:    httpCode,
			Message: err.Error(),
		})
	}
	return
}
