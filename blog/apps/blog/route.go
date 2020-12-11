package blog

//File  : route.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/8
import (
	"github.com/gin-gonic/gin"
	middleware "mygit.com/SimonWang00/blog/middleware"
	"net/http"
	"time"
)

// 加载blog的路由
func LoadBlog(g *gin.Engine)  {
	// 限制每秒钟5000次请求
	g.GET("/",middleware.RateLimiter(1*time.Minute, 3000000), HomeHandler)
	g.GET("/blog",middleware.RateLimiter(1*time.Minute, 3000000), BlogHandler)
	g.GET("/about", AboutHandler)
	g.GET("/works", WorkHandler)
	g.GET("/links", LinkHandler)
	g.GET("/contact", ContactHandler)
	g.POST("/contact", InsertMessage)
	g.GET("/test", func(c *gin.Context) {
		xv:=[][]int{}
		c.JSON(http.StatusOK,
			gin.H{
				"code":        http.StatusOK,
				"status":      "请输入有效的搜索词！",
				"total_count": 0,
				"data":        xv,
			})
	})
}