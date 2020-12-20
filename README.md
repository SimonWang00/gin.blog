



# blog

我的个人博客：gin + gorm(反射、线程池)+ bootstrap + jquery全栈构建的个人博客



# 功能特性

- [x] 博客主页
- [x] 发布博客（markdown）
- [x] 发布作品（GitHub）
- [x] 发布留言
- [x] 后台管理系统
- [x] 管理员登录
- [x] 博客最新情况报告
- [ ] 编辑/删除文章
- [ ] 编辑/删除作品



# 环境依赖

- golang
- mysql



# 快速部署

1.访问到项目目录，编译镜像

```shell
docker build . -t gin.blog:v1.0.0
```

2.启动项目

```shell
docker run -p 8081:8081 -d "镜像"
```



# 在线访问

访问地址：http://106.52.120.86:8081/

后台账号密码：admin/admin