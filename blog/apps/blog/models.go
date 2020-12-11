package blog

import "time"

//File  : models.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/8


// 文章
type Ariticle struct {
	Title string			//文章标题
	CreateTime time.Time		//发表时间
	Summary string			//文章概要
	Classify string			//文章分类
	Tag string				//文章标签
}

// 留言
type Message struct {
	Name string				//留言者昵称
	Mail string				//留言者邮箱
	Content string			//留言内容
	CreateTime time.Time		//留言时间
}

// 作品
type Work struct {
	Title string			//作品标题
	About string			//作品简介
	StarNum string			//star数量
	ForkNum string			//fork数量
	Language string			//编程语言
}