package admin

import (
	"gin.blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

/**
 * @Author: SimonWang00
 * @Description: 博客后台管理
 * @File:  handles.go
 * @Version: 1.0.0
 * @Date: 2020/12/20 9:53
 */

var wg sync.WaitGroup

// 后台管理页面
func AdminHomeHandler(c *gin.Context)  {
	if c.Request.Method == "POST"{
		username,_ := c.GetPostForm("username")
		useragent := c.Request.UserAgent()
		if username != "" {
			blog, totalBlogNums := dao.GetLatestBlog()
			message, messageNums := dao.GetLatestMessage()
			latestlogin, logincount , usernum:= dao.GetLatestLogin(username)
			c.HTML(http.StatusOK, "admin/index.html",gin.H{
				"username":username,
				"logincount":logincount,
				"usernum":usernum,
				"useragent":useragent,
				"latestLoginIpport":latestlogin.Ipport,
				"latestLoginLoginTime":latestlogin.LoginTime,
				"latestBlog":blog.Title,
				"latestCreateTime":blog.CreateTime,
				"totalBlogNums":totalBlogNums,
				"messageNums":messageNums,
				"messageContent":message.Content,
				"messageUser":message.Name,
				"messageCreateTime":message.CreateTime,
			})
			return
		}
	}
	c.HTML(http.StatusOK, "admin/index.html","")
}

// 后台登录页面
func AdminLoginHandler(c *gin.Context)  {
	if c.Request.Method == "GET"{
		c.HTML(http.StatusOK, "login.html", "")
		return
	} else if c.Request.Method == "POST"{
		username,_ := c.GetPostForm("username")
		userpwd,_ := c.GetPostForm("userpwd")
		admin := dao.SelectAdmin(username)
		ipport := c.ClientIP()
		if admin.Password == userpwd{
			status := "success"
			dao.AddLog(username, ipport, status)
			//c.HTML(http.StatusOK, "admin/index.html", "")
			c.Redirect(http.StatusPermanentRedirect,"/admin/index.html")
			return
		}
		c.HTML(http.StatusOK, "admin/reload.html", gin.H{
			"url":"/admin/login",
			"message":"账号或者密码错误！3秒后跳转至登录页面~~",
		})

	}
}

