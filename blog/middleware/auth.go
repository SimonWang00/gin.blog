package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleWare 进行客户端cookie校验
// 校验参数示例userid
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("userid"); err == nil {
			if cookie == "123" {
				// 执行函数
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
