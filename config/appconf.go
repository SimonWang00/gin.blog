package config

import (
	 easyconfig "github.com/spf13/viper"
)

// 应用程序配置
type appConfig struct {
	// 应用名称
	Name string
	// 运行模式: debug, release, gotest
	RunMode string
	// 运行 addr
	Addr string
	// 完整 url
	URL string
	// secret key
	Key string

	// 静态资源存放路径
	PublicPath string
	// 模板等前端源码文件存放路径
	ResourcesPath string
	// 模板文件存放的路径
	ViewsPath string

	// 是否开启 csrf
	EnableCsrf bool
	// csrf param name
	CsrfParamName string
	// csrf header
	CsrfHeaderName string

	// auth session key
	AuthSessionKey string
	// Context 中当前用户数据的 key
	ContextCurrentUserDataKey string
}

func newAppConfig() *appConfig {
	// 默认配置
	easyconfig.SetDefault("APP.NAME", "blog")
	easyconfig.SetDefault("APP.RUNMODE", "release")
	easyconfig.SetDefault("APP.ADDR", ":8080")
	easyconfig.SetDefault("APP.KEY", "weqtrwyrererererererererereresdfbsg")
	easyconfig.SetDefault("APP.ENABLE_CSRF", true)

	return &appConfig{
		Name:    easyconfig.GetString("APP.NAME"),
		RunMode: easyconfig.GetString("APP.RUNMODE"),
		Addr:    easyconfig.GetString("APP.ADDR"),
		URL:     easyconfig.GetString("APP.URL"),
		Key:     easyconfig.GetString("APP.KEY"),

		PublicPath:    "public",
		ResourcesPath: "resources",
		ViewsPath:     "resources/views",

		EnableCsrf:     easyconfig.GetBool("APP.ENABLE_CSRF"),
		CsrfParamName:  "_csrf",
		CsrfHeaderName: "X-CsrfToken",

		AuthSessionKey:            "blog_session",
		ContextCurrentUserDataKey: "currentUserData",
	}
}
