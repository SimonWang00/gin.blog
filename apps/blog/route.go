package blog

//File  : route.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/8
import (
	middleware "gin.blog/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

// 加载blog的路由
func LoadBlog(g *gin.Engine)  {
	// 请求主页 限制每秒钟5000次请求
	g.GET("/",middleware.RateLimiter(1*time.Minute, 3000000), HomeHandler)
	// 请求博客页面
	g.GET("/blog",middleware.RateLimiter(1*time.Minute, 3000000), BlogHandler)
	// 请求关于我
	g.GET("/about", AboutHandler)
	// 外链
	g.GET("/links", LinkHandler)
	// 留言
	g.GET("/contact", ContactHandler)
	// 新增留言
	g.POST("/contact", InsertMessage)
	// 我的作品
	g.GET("/works", WorkHandler)
	// 新增作品
	g.POST("/works", InsertWork)
	// 发布作品
	g.GET("/public/work", PublicWork)
}