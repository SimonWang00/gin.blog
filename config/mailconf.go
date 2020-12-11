package config

import "github.com/spf13/viper"

// 邮件配置
type mailConfig struct {
	Host     string // 邮箱的服务器地址
	Port     int    // 邮箱的服务器端口
	From     string // 发送者邮箱
	Password string // 授权码或密码
	To string 		// 邮件接收者
	CC string 		// 抄送
}

func NewMailConfig() *mailConfig {
	// 默认配置
	viper.SetDefault("MAIL.MAIL_HOST", "smtp.qq.com")
	viper.SetDefault("MAIL.MAIL_PORT", "763646402@qq.com")
	viper.SetDefault("MAIL.MAIL_FROM", 465)
	viper.SetDefault("MAIL.MAIL_PASSWORD", "ojyeijohshodbbfh")
	viper.SetDefault("MAIL.MAIL_TO", "bw_wangxiaomeng@whty.com.cn,763646402@qq.com")
	viper.SetDefault("MAIL.MAIL_CC", "")

	return &mailConfig{
		Host:     viper.GetString("MAIL.MAIL_HOST"),
		Port:     viper.GetInt("MAIL.MAIL_PORT"),
		From:     viper.GetString("MAIL.MAIL_FROM"),
		Password: viper.GetString("MAIL.MAIL_PASSWORD"),
		To: viper.GetString("MAIL.MAIL_TO"),
		CC: viper.GetString("MAIL.MAIL_CC"),
	}
}