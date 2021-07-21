package main

import (
	"other/modeltools/dbtools"
	"other/modeltools/generate"
)

func main() {
	//初始化数据库
	dbtools.Init()
	generate.Genertate() //生成所有表信息
	//generate.Genertate("lian_v1_article") //生成指定表信息，可变参数可传入多个表名
}
