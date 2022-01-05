package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	sign "github.com/yourbrainrun/test_go/test_gin_post/sign"
	"time"
)

func main() {
	router := gin.New()
	router.POST("/test", post)
	router.GET("/work", work)
	router.Run(":9900")
}

func post(c *gin.Context) {
	// 127.0.0.1:999/test

	c.Request.ParseMultipartForm(12)
	data := c.Request.Form
	paramsSlice := sign.GetSign(data, "secret")
	fmt.Println(paramsSlice)

}

func work(c *gin.Context) {
	// http://127.0.0.1:999/work

	fmt.Println("start work... ")
	str := ""
	for index, value := range c.Request.Header {
		fmt.Println(index, "=>", value)
		str += fmt.Sprintf("%s => %s", index, value)
	}

	if c.Request.Header.Get("proxy") == "true" {
		fmt.Println("proxy", "true")
		str += "proxy:true"
	} else {
		fmt.Println("proxy", "false")
		str += "proxy:false"
	}

	c.JSON(200, str)
	//go sonLine()
	fmt.Println("work over")
}

func sonLine() {
	fmt.Println("son line start run!")
	time.Sleep(time.Second * 5)
	fmt.Println("son line end run!")
}
