package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	sign "github.com/yourbrainrun/test_go/test_gin_post/sign"
)

func main() {
	router := gin.New()
	router.POST("/test", post)
	router.Run(":999")
}

func post(c *gin.Context) {
	c.Request.ParseMultipartForm(12)
	data := c.Request.Form
	paramsSlice := sign.GetSign(data, "secret")
	fmt.Println(paramsSlice)
}
