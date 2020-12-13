package blog

//File  : handles.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/8

import (
	"fmt"
	"gin.blog/dao"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

// 主页
func HomeHandler(c *gin.Context)  {
	articles := dao.QueryLatestAriticle()
	messages := dao.QueryLatestMessage()
	works := dao.QueryLatestWork()
	c.HTML(http.StatusOK, "index.html",gin.H{
		"articles":articles,
		"messages":messages,
		"works":works,
	})
}

// gin.blog
func BlogHandler(c *gin.Context)  {
	var page, pagesize int
	var path  = "/blog"
	pagestr := c.Query("page")
	page, err := strconv.Atoi(pagestr)
	if err != nil{
		page = 1
	}
	pagesizestr := c.Query("pagesize")
	pagesize, err = strconv.Atoi(pagesizestr)
	if err != nil{
		pagesize = 10
	}
	classify := c.Query("classify")
	// 根据分类查询所有文章
	if classify != ""{
		articles, totalblogs := dao.QueryAriticleByclassify(classify, page, pagesize)
		page_map := generatePageLink(path, classify, totalblogs, page, pagesize)
		//pages := int(math.Ceil(float64(totalblogs)/10))
		c.HTML(http.StatusOK, "blog.html",gin.H{
			"classify":classify,
			"articles":articles,
			"pages":page_map,
			"currentpage":page,
		})
		return
	}
	id := c.Query("id")
	// 查询文章详情
	if id != ""{
		id, _ := strconv.Atoi(id)
		article := dao.QueryAriticleById(id)
		c.HTML(http.StatusOK, "content.html",gin.H{
			"article":article,
		})
		return
	}
	articles, totalblogs := dao.QueryAllAriticle(page, pagesize)
	page_map := generatePageLink(path, classify, totalblogs, page, pagesize)
	//pages := int(math.Ceil(float64(totalblogs)/10))
	c.HTML(http.StatusOK, "blog.html",gin.H{
		"classify":"",
		"articles":articles,
		"pages":page_map,
		"currentpage":page,
	})
}

// 只显示五页
func limitPage(path string, classify string, totalblogs int, currentpage int, pagesize int) *map[int]string {
	page_map := make(map[int]string, 6)
	pages := int(math.Ceil(float64(totalblogs)/float64(pagesize)))
	for i:=1; i < pages+1; i++ {
		if i >= currentpage -1{
			if i == currentpage{
				page_map[i] = "#"
			} else {
				pagestr := fmt.Sprintf("%v?classify=%v&page=%v&pagesize=%v", path ,classify, i,pagesize)
				page_map[i] = pagestr
			}
		}
		if len(page_map) >= 5{
			break
		}
	}
	return &page_map
}

// 生成翻页链接
// /blog?classify=classify&page=1&pagesize=10
// /blog?page=1&pagesize=10
func generatePageLink(path string, classify string,totalblogs int, currentpage int , pagesize int) *map[int]string {
	page_map := limitPage(path, classify, totalblogs, currentpage, pagesize)
	return page_map
}

// 关于页面
func AboutHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "about.html","")
}

// 作品页面
func WorkHandler(c *gin.Context)  {
	var page, pagesize int
	var path = "/works"
	pagestr := c.Query("page")
	page, err := strconv.Atoi(pagestr)
	if err != nil{
		page = 1
	}
	pagesizestr := c.Query("pagesize")
	pagesize, err = strconv.Atoi(pagesizestr)
	if err != nil{
		pagesize = 9
	}
	works, pages := dao.QueryAllWork(page, pagesize)
	page_map := generatePageLink(path, "", pages, page, pagesize)
	c.HTML(http.StatusOK, "works.html",gin.H{
		"works":works,
		"pages":page_map,
	})
}

// 链接页面
func LinkHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "links.html","")
}

// 联系页面
func ContactHandler(c *gin.Context)  {
	pagestr := c.Query("page")
	pagesizestr := c.Query("pagesize")
	page, err := strconv.Atoi(pagestr)
	if err != nil{
		page = 1
	}
	pagesize, err := strconv.Atoi(pagesizestr)
	if err != nil{
		pagesize = 10
	}
	messages ,num := dao.QueryAllMessage(page, pagesize)
	c.HTML(http.StatusOK, "contact.html",gin.H{
		"messages":messages,
		"nums":num,
	})
}

// 新增留言
func InsertMessage(c *gin.Context)  {
	contact_name, _ := c.GetPostForm("contact_name")
	contact_email,_ := c.GetPostForm("contact_email")
	contact_con,_ := c.GetPostForm("contact_con")
	if contact_name == "" || contact_email == "" || contact_con == "" {
		c.HTML(http.StatusBadRequest, "reload.html", gin.H{
			"title":"留言添加失败",
			"url":"/contact",
			"message":"留言请求参数不能为空， 3秒钟后跳转到留言页面",
		})
		return
	}
	err := dao.AddMesage(contact_name, contact_email, contact_con)
	if err != nil{
		c.HTML(http.StatusInternalServerError, "reload.html", gin.H{
			"title":"留言添加失败",
			"url":"/contact",
			"message":"留言新增失败，请重试！ 3秒钟后跳转到留言页面",
		})
		return
	}
	c.HTML(http.StatusOK, "reload.html", gin.H{
		"title":"留言添加成功！",
		"url":"/contact",
		"message":"留言成功！博主看到会第一时间回复你的哦~~， 3秒钟后跳转到留言页面",
	})
}


// 新增作品
func InsertWork(c *gin.Context)  {
	title, _ := c.GetPostForm("title")
	about, _ := c.GetPostForm("about")
	starNum, _ := c.GetPostForm("starNum")
	forkNum, _ := c.GetPostForm("forkNum")
	language, _ := c.GetPostForm("language")
	url, _ := c.GetPostForm("url")
	if title == "" || about == "" || starNum == "" || forkNum == "" || language == "" {
		c.HTML(http.StatusBadRequest, "reload.html", gin.H{
			"title":"作品添加失败",
			"url":"/works",
			"message":"作品信息录入不完整， 3秒钟后跳转到作品页面",
		})
		return
	}
	err := dao.AddWork(title, about, starNum, forkNum, language, url)
	if err != nil{
		c.HTML(http.StatusInternalServerError, "reload.html", gin.H{
			"title":"作品添加失败",
			"url":"/works",
			"message":"作品新增失败，请重试！ 3秒钟后跳转到作品页面",
		})
		return
	}
	c.HTML(http.StatusOK, "reload.html", gin.H{
		"title":"作品添加成功！",
		"url":"/works",
		"message":"作品添加成功哦~~， 3秒钟后跳转到作品页面",
	})
}

// 发布作品页面
func PublicWork(c *gin.Context)  {
	c.HTML(http.StatusOK, "publicwork.html", "")
}