package main

import (
	"log"
	"os"
	_ "chapter_2/matchers"
	"chapter_2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	//获取数据   --   数据源
	//执行搜索   goroutine
		//使用接口进行匹配
		//发送结果
		//报告任务完成
	//跟踪结果   goroutine

	//显示结果
	search.Run("")
}


