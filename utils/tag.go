package utils

import "strings"

//File  : tag.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/15

var (
	TAG_ARR = [30]string {"Python","python", "Go", "go", "Golang", "golang","人工智能","AI","Java","java","C++","大数据",
		"SQL","区块链","hadoop","zookeeper","elasticsearch","gin","微服务","hbase","tensorflow","数据挖掘","NLP","CV"}
)

// ContentTag 查询文章标签
func ContentTag(content string)  string{
	var tags string = ""
	for _, tag := range TAG_ARR {
		if strings.Contains(content, tag) && len(tag) <= 50{
			tags = tags + tag
		}
	}
	if tags == ""{
		tags += "随笔"
	}
	return tags
}

// ContentSummary 截取文章摘要
func ContentSummary(content string) string {
	var summary string
	runes := []rune(content)
	if len(runes) > 200{
		summary = string(runes[:200]) + "..."
	} else if len(runes) <=200{
		summary = content + "..."
	} else {
		summary = ""
	}
	return summary
}