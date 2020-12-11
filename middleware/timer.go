package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// MyTimerMiddleware中间件，计时器
// 获取计时方法c.Get("timer")即可读取
func MyTimerMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		t0 := time.Now()
		c.Next()
		t1 := time.Now()
		delta := t1.Sub(t0)
		c.Set("timer", delta)
		log.Printf("hand fun cost time %v s", delta)
	}
}
