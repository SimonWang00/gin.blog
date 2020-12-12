package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"gin.blog/apps/blog"
	"gin.blog/config"
)

// Register 注册路由和中间件
func SetRegisterRouters() *gin.Engine {
	// 加载不同app的路由, 加载bijiav1
	Include(blog.LoadBlog)
	// 加载路由启动
	g := Init()
	// ---------------------------------- 注册全局中间件 ----------------------------------
	// 使用 Recovery 中间件
	g.Use(gin.Recovery())
	// 使用 Logger 中间件,根据配置
	if config.AppConfig.RunMode != config.RunmodeRelease {
		g.Use(gin.Logger())
	}

	// +++++++++++++++++++ swagger +++++++++++++++++++
	// 需全局安装 go get -u github.com/swaggo/swag/cmd/swag 然后 swag init 生成文档
	if config.AppConfig.RunMode != config.RunmodeRelease {
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return g
}


