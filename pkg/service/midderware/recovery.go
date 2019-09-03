package midderware

import (
	"github.com/gin-gonic/gin"
	logger "go.uber.org/zap"
	"net/http/httputil"
)

// Recovery 用于页面出现 panic 时候的恢复
func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			httpRequest, _ := httputil.DumpRequest(c.Request, false)
			logger.L().Error("recovery", logger.String("http request", string(httpRequest)), logger.Error(err.(error)))

			c.Redirect(307, "/errors")
			return
		}
	}()
	c.Next()
}

// NotFound 当页面 404 时的处理
func NotFound(c *gin.Context) {
	c.HTML(404, "errors/notfound.html", c.Keys)
}

// Errors 是错误的页面
func Errors(c *gin.Context) {
	c.HTML(500, "errors/errors.html", c.Keys)
	return
}
