#选择体积依赖较小的版本
FROM golang:1.15.0-alpine3.12

LABEL maintainer=simon_wang00@163.com
COPY . /$GOPATH/src/gin.blog
WORKDIR /$GOPATH/src/gin.blog

#设置环境变量，开启go module和设置下载代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

#会在当前目录生成一个go.mod文件用于包管理
RUN go mod init
#增加缺失的包，移除没用的包
RUN go mod vendor

RUN go build blogserver.go
EXPOSE 8080:8080
CMD ["./blogserver"]