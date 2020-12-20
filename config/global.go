package config

import (
	"log"

	easyconfig "github.com/spf13/viper"
)

const (
	// RunmodeDebug -
	RunmodeDebug = "debug"
	// RunmodeRelease -
	RunmodeRelease = "release"
	// RunmodeTest -
	RunmodeTest = "gotest"

	// 配置文件路径
	configFilePath = "./config.yaml"
	// 日志文件路径
	//logFilePath = "./logs/api.log"
	LogFilePath = "./logs/api.log"
	// 配置文件格式
	configFileType = "yaml"
)

var (
	// AppConfig 应用配置
	AppConfig *appConfig
	// DBConfig 数据库配置
	DBConfig *dbConfig
	// MailConfig 邮件配置
	MailConfig *mailConfig
)

func init()  {
	InitConfig("", true)
}

// InitConfig 初始化配置
func InitConfig(c string, hasLog bool) {
	if c == "" {
		c = configFilePath
	}
	// 初始化 viper 配置
	easyconfig.SetConfigFile(c)
	easyconfig.SetConfigType(configFileType)

	if err := easyconfig.ReadInConfig(); err != nil {
		log.Fatal("读取配置文件失败，请检查 config.yaml 配置文件是否存在: %v", err)
	}

	// 初始化 apps 配置
	AppConfig = newAppConfig()
	// 初始化数据库配置,通过GORM初始化参数为全局变量
	DBConfig = newDBConfig()
	// 初始化邮件配置
	MailConfig = NewMailConfig()
}


