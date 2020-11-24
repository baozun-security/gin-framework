package v1

import "github.com/gin-gonic/gin"

// api v1 resources
func RegisterApiRouter(v1 *gin.RouterGroup) {
	v1.GET("", Test)
}
