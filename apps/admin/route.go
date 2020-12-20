package admin

/**
 * @Author: SimonWang00
 * @Description:
 * @File:  route.go
 * @Version: 1.0.0
 * @Date: 2020/12/20 9:53
 */


import (
	middleware "gin.blog/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

// 加载blog后台的路由
func LoadAdmin(g *gin.Engine)  {
	// 请求后台管理 限制每秒钟5000次请求
	g.GET("/admin/index.html",middleware.RateLimiter(1*time.Minute, 3000000), AdminHomeHandler)
	g.POST("/admin/index.html",middleware.RateLimiter(1*time.Minute, 3000000), AdminHomeHandler)
	// 登录博客后台管理页面
	g.GET("/admin/login",middleware.RateLimiter(1*time.Minute, 3000000), AdminLoginHandler)
	g.POST("/admin/login",middleware.RateLimiter(1*time.Minute, 300), AdminLoginHandler)
}