package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	visit_chan chan Visit = make(chan Visit, 1) // use no bufferred channel
)

func main() {
	config := GetConfig() // parse config.yaml
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "用法错误：./visit_analytics config.yaml")
		os.Exit(1)
	}
	config.parse(os.Args[1])

	InitDB() // init database

	// generateRandomRecords() // for generate fake daily_record

	router := gin.Default()

	// for recording visit
	router.GET("/analytics.js", func(c *gin.Context) {
		go analyse(c)
		c.Header("Content-Type", "application/javascript")
		c.String(http.StatusOK, "console.log(\"网站统计系统\")")
	})

	// for testing
	//router.StaticFile("/test", "./www/test.html")
	//router.StaticFile("/test2", "./www/test.html")

	// for web displaying
	MakeRoutes(router)

	//handle database manipulation of visit record
	go func() {
		for {
			visit := <-visit_chan
			handleVisit(visit)
		}
	}()

	router.Run(config.RunAt)
}

// record according to Request.Referer()
func analyse(c *gin.Context) {
	referer := TrimUrl(c.Request.Referer()) //过滤/和/?符号
	host_name := GetHostName(referer)       //获取主机名
	if len(referer) == 0 || host_name == "" {
		return
	}
	title := GetTitle(referer) //获取网站标题

	visit_chan <- Visit{
		ClientIp:  c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Referer:   referer,
		Title:     title,
		Host:      host_name,
	}
}

// database manipulation of visit record
func handleVisit(visit Visit) {
	recordHost(visit)
	recordPage(visit)
	recordDailyRecord(visit)
	recordMonthlyRecord(visit)
}
