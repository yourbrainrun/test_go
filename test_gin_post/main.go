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
	router.Run(":999")
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

	go sonLine()
	fmt.Println("work over")
}

func sonLine() {
	fmt.Println("son line start run!")
	time.Sleep(time.Second * 5)
	fmt.Println("son line end run!")
}
