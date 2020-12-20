package models

/**
 * @Author: SimonWang00
 * @Description:
 * @File:  admin.go
 * @Version: 1.0.0
 * @Date: 2020/12/20 9:59
 */

// 管理员账号
type Admin struct {
	Id int				`gorm:"primary_key"`
	Username string		`gorm:"type:varchar(50);not null"`		//管理员名称
	Password string		`gorm:"type:char(32);not null"`			//管理员密码
	CreateTime string	`gorm:"type:datetime;not null"`			//管理员创建时间
}

// 登录日志
type Log struct {
	Id int				`gorm:"primary_key"`
	Username string		`gorm:"type:varchar(50);not null"`		//管理员名称
	Ipport string		`gorm:"type:varchar(20);not null"`		//管理员登录IP
	Status string		`gorm:"type:varchar(10);not null"`		//管理员登录状态
	LoginTime string	`gorm:"type:datetime;not null"`			//管理员登录时间
}