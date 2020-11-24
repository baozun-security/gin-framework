package middlewares

import (
	"baozun.com/leak/app/pkgs/logger"
	"baozun.com/leak/app/pkgs/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				// where does panic occur? try max 20 depths
				pcs := make([]uintptr, 20)
				max := runtime.Callers(2, pcs)

				if max == 0 {
					logger.Logger.Warn("No pcs available")
				} else {
					frames := runtime.CallersFrames(pcs[:max])
					for {
						frame, more := frames.Next()

						// To keep this example's output stable
						// even if there are changes in the testing package,
						// stop unwinding when we leave package runtime.
						if strings.Contains(frame.Function, "runtime.") {
							if more {
								continue
							} else {
								break
							}
						}

						tmp := strings.SplitN(frame.File, "/src/", 2)
						if len(tmp) == 2 {
							logger.Logger.Errorf("(src/%s:%d: %v)", tmp[1], frame.Line, panicErr)
						} else {
							logger.Logger.Errorf("(%s:%d: %v)", frame.File, frame.Line, panicErr)
						}

						break
					}
				}
				debug.PrintStack() // 打印错误堆栈信息

				g := response.Gin{Ctx: c}
				g.ApiFail(http.StatusInternalServerError, fmt.Errorf("system error, %s", errorToString(panicErr)))
				// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		c.Next()
	}
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
