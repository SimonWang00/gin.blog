package blog

import "time"

//File  : model.go
//Author: Simon
//Describe: define mysql table
//Date  : 2020/12/9

// 文章
type Ariticle struct {
	Id int				`gorm:"primary_key"`
	Title string		`gorm:"type:varchar(50);not null"`	//文章标题
	Summary string		`gorm:"type:varchar(255);not null"`	//文章概要
	Content string		`gorm:"type:longtext;not null"`		//文章内容
	Classify string		`gorm:"type:varchar(10);not null"`	//文章分类
	Tag string			`gorm:"type:varchar(20);not null"`	//文章标签
	CreateTime time.Time									//发表时间
}


// 留言
type Message struct {
	Id int					`gorm:"primary_key"`
	Name string				`gorm:"type:varchar(50);not null"`		//留言者昵称
	Mail string				`gorm:"type:varchar(50);not null"`		//留言者邮箱
	Content string			`gorm:"type:varchar(50);not null"`		//留言内容
	CreateTime time.Time											//留言时间
}


// 作品
type Work struct {
	Id int				`gorm:"primary_key"`
	Title string		`gorm:"type:varchar(50);not null"`	//作品标题
	About string		`gorm:"type:varchar(255);not null"`	//作品简介
	StarNum string		`gorm:"type:int;default 0"`			//star数量
	ForkNum string		`gorm:"type:int;default 0"`			//fork数量
	Language string		`gorm:"type:varchar(10);not null"`	//编程语言
	Url string			`gorm:"type:varchar(200);not null"`	//作品链接
}