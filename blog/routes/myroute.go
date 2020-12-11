package routes

import (
	"github.com/gin-gonic/gin"
)

// 定义同参数的匿名函数
type option func(c *gin.Engine)

// MyRoute -
type MyRoute struct {
	Router gin.IRouter
}

var (
	// RouterMap : 存放路由 name 和它的 path map
	RouterMap = make(map[string]string, 0)
	methodMap = make(map[string]string, 0)
	// 匿名函数切片
	options [] option
)

// router初始化
func Init() *gin.Engine {
	r := gin.New()
	// oper 是匿名函数，参数是engine
	for _, oper := range options{
		oper(r)
	}
	return r
}

// 将不同app下的路由加载的匿名函数放入分片
func Include(ops ...option)  {
	options = append(options, ops...)
}

// Middleware 注册中间件
func (r *MyRoute) Middleware(middlewares ...gin.HandlerFunc) *MyRoute {
	return &MyRoute{
		Router: r.Router.Group("/", middlewares...),
	}
}

// Group 注册路由组
func (r *MyRoute) Group(relativePath string, handlers ...gin.HandlerFunc) *MyRoute {
	return &MyRoute{
		Router: r.Router.Group(relativePath, handlers...),
	}
}

// Register 注册路由
func (r *MyRoute) Register(method string, name string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	Name(r.Router, name, method, relativePath)
	return r.Router.Handle(method, relativePath, handlers...)
}

// Name : 注册路由
// 动态参数目前只支持 :xx 形式
func Name(g gin.IRoutes, name string, method string, path string) {
	s := path
	group, ok := g.(*gin.RouterGroup)
	if ok {
		s = group.BasePath() + path
	}

	if RouterMap[name] != "" {
		panic("该路由已经命名过了: [" + name + "] " + s)
	}
	RouterMap[name] = s
	methodMap[name] = method
}
