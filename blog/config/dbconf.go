package config

import (
	"fmt"
	easyconfig "github.com/spf13/viper"
)

// mysql数据库配置
type dbConfig struct {
	Connection string
	Host       string
	Port       int
	Database   string
	Username   string
	Password   string

	URL string
	Debug bool
}

// newDBConfig 是数据库默认配置
func newDBConfig() *dbConfig {
	// 默认配置
	easyconfig.SetDefault("DB.CONNECTION", "mysql")
	easyconfig.SetDefault("DB.HOST", "127.0.0.1")
	easyconfig.SetDefault("DB.PORT", 3306)
	easyconfig.SetDefault("DB.DATABASE", easyconfig.GetString("APP.NAME"))
	easyconfig.SetDefault("DB.USERNAME", "root")
	easyconfig.SetDefault("DB.PASSWORD", "000000")

	username := easyconfig.GetString("DB.USERNAME")
	//password := easyconfig.GetString("DB.PASSWORD")
	password := "000000"
	host := easyconfig.GetString("DB.HOST")
	port := easyconfig.GetInt("DB.PORT")

	database := easyconfig.GetString("DB.DATABASE")
	database = database + "_" + AppConfig.RunMode
	url := createDBURL(username, password, host, port, database)

	return &dbConfig{
		Connection: easyconfig.GetString("DB.CONNECTION"),
		Host:       host,
		Port:       port,
		Database:   database,
		Username:   username,
		Password:   password,
		URL:        url,
		Debug:      AppConfig.RunMode == RunmodeDebug,
	}
}

// createDBURL 创建mysql连接符
func createDBURL(uname string, pwd string, host string, port int, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		uname, pwd,
		host, port,
		name, true, "Local")
}