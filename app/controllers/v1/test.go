package v1

import (
	"baozun.com/leak/app/pkgs/response"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	g := response.Gin{Ctx: c}
	g.ApiSuccess(nil)
}
