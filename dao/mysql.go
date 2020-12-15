package dao

//File  : mysql.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/9

import (
	"errors"
	"gin.blog/config"
	"gin.blog/models"
	"gin.blog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// DB gorm
var db *gorm.DB

// init 初始化数据库
func init() {
	var err error
	db, err = gorm.Open(config.DBConfig.Connection, config.DBConfig.URL)
	if err != nil {
		log.Fatal("Database connection failed. Database url: "+config.DBConfig.URL+" error: ", err)
	} else {
		log.Print("\n\n------------------------------------------ GORM OPEN SUCCESS! -----------------------------------------------\n\n")
	}
	db.LogMode(config.DBConfig.Debug)
	// 连接池最大连接数100
	db.DB().SetMaxOpenConns(100)
	// 最大空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 创建所有表
	createTable()
}

// 创建表
func createTable()  {
	if !db.HasTable(&models.Article{}){
		err := db.Set("gorm:table_options","ENGINE=InnoDB DEFAULT CHARSET=utf8mb4" ).
			CreateTable(&models.Article{}).Error
		if err != nil{
			log.Fatalf("create table AriticleTable error(%v)!", err.Error())
		}
		log.Println("AriticleTable create suecess!")
	}
	log.Println("AriticleTable has been created!")
	if !db.HasTable(&models.Message{}){
		err := db.Set("gorm:table_options","ENGINE=InnoDB DEFAULT CHARSET=utf8" ).
			CreateTable(&models.Message{}).Error
		if err != nil{
			log.Fatalf("create table MessageTable error(%v)!", err.Error())
		}
		log.Println("MessageTable create suecess!")
	}
	log.Println("MessageTable has been created!")
	if !db.HasTable(&models.Work{}){
		err := db.Set("gorm:table_options","ENGINE=InnoDB DEFAULT CHARSET=utf8" ).
			CreateTable(&models.Work{}).Error
		if err != nil{
			log.Fatalf("create table WorkTable error(%v)!", err.Error())
		}
		log.Println("WorkTable create suecess!")
	}
	log.Println("WorkTable has been created!")
	log.Println("all tables are ready!")
}

// QueryLatestAriticle 主页拉取最新的五篇文章
func QueryLatestAriticle()  *[]models.Article {
	var ariticles []models.Article
	// SELECT title, summary, content, classify, tag, create_time FROM `ariticles`   ORDER BY create_time desc
	db.Select([]string{"id","title","summary","content","classify","tag","create_time"}).Order("create_time desc").Limit(5).Find(&ariticles)
	return &ariticles
}

// QueryAllAriticle 查询所有文章
func QueryAllAriticle(page int, pagesize int) (*[]models.Article, int){
	var ariticles []models.Article
	var totalblogs int
	db.Model(models.Article{}).Select([]string{"id","title","summary","content","classify","tag","create_time"}).
		Order("create_time desc").Limit(pagesize).Count(&totalblogs).
		Offset((page-1)*pagesize).Find(&ariticles)
	return &ariticles, totalblogs
}

// QueryAriticleByclassify 通过分类查询所有文章
func QueryAriticleByclassify(classify string, page int, pagesize int) (*[]models.Article, int) {
	var ariticles []models.Article
	var totalblogs int
	db.Model(models.Article{}).Where("classify = ?", classify).Count(&totalblogs).
		Select([]string{"id","title","summary","content","classify","tag","create_time"}).
		Order("create_time desc").Limit(pagesize).Offset((page-1)*pagesize).Find(&ariticles)
	return &ariticles, totalblogs
}

// QueryAriticleById 通过id 查询文章
func QueryAriticleById(id int) *models.Article {
	var ariticle models.Article
	db.Where("id = ?", id).Select([]string{"id","title","summary","content","classify","tag","create_time"}).Find(&ariticle)
	return &ariticle
}

// QueryLatestMessage 主页拉取最新的评论
func QueryLatestMessage()  *[]models.Message {
	var message []models.Message
	// SELECT title, summary, content, classify, tag, create_time FROM `ariticles`   ORDER BY create_time desc
	db.Select([]string{"name","mail","content","create_time"}).Order("create_time desc").Limit(5).Find(&message)
	return &message
}

// QueryAllMessage 查询所有留言消息
func QueryAllMessage(page int, pagesize int) (*[]models.Message, int) {
	var messages [] models.Message
	var num int
	db.Table("messages").Count(&num).Order("create_time desc").
		Limit(pagesize).Offset((page - 1) * pagesize).Find(&messages)
	return &messages, num
}


// QueryLatestWork 主页拉取最新的3个作品
func QueryLatestWork()  *[]models.Work {
	var works []models.Work
	db.Select([]string{"title","about","star_num","fork_num","language","url"}).Limit(3).Find(&works)
	return &works
}

// 新增留言
func AddMesage(name, email,content string ) error {
	if name == "" || email == "" || content == ""{
		return errors.New("请求参数为空！")
	}
	mess := &models.Message{
		Name:name,
		Mail:email,
		Content:content,
		CreateTime:time.Now(),

	}
	if err := db.Create(mess).Error; err != nil{
		log.Printf("insert message error(%v)", err.Error())
		return err
	}
	return nil
}

// 新增作品
func AddWork(title, about, starNum, forkNum, language,url string ) error {
	if title == "" || about == "" || starNum == "" || forkNum == "" || language == "" {
		return errors.New("请求参数为空！")
	}
	work := &models.Work{
		Title:title,
		About:about,
		StarNum:starNum,
		ForkNum:forkNum,
		Language:language,
		Url:url,
	}
	if err := db.Create(work).Error; err != nil{
		log.Printf("insert message error(%v)", err.Error())
		return err
	}
	return nil
}

// QueryAllWork 查询所有作品
func QueryAllWork(page int, pagesize int) (*[]models.Work, int){
	var works []models.Work
	var totalblogs int
	db.Model(models.Work{}).Select([]string{"id","title","about","star_num","fork_num","language","url"}).
		Limit(pagesize).Count(&totalblogs).
		Offset((page-1)*pagesize).Find(&works)
	return &works, totalblogs
}

// AddBlog 新增博文
func AddBlog(title, content, classify string)  error{
	if title == "" ||  content == "" || classify == "" {
		return errors.New("请求参数为空！")
	}
	summary := utils.ContentSummary(content)
	tag := utils.ContentTag(content)
	blog := &models.Article{
		Title:      title,
		Summary:    summary,
		Content:    content,
		Classify:   classify,
		Tag:        tag,
		CreateTime: time.Now(),
	}
	if err := db.Create(blog).Error; err != nil{
		log.Printf("insert blog error(%v)", err.Error())
		return err
	}
	return nil
}