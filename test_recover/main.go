package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(Handler)
	router.GET("/test", test)

	router.Run(":999")
}

func test(c *gin.Context) {
	dealData(c)
	c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
		"ok": 123,
	})
}

type Api struct {
	Code    int
	Message string
	Retry   int
}

func Handler(c *gin.Context) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)

			switch t := p.(type) {
			case *Api:
				c.JSON(http.StatusInternalServerError, t)
			default:
				c.JSON(http.StatusInternalServerError, t)
			}

			c.Abort()
		}
	}()

	c.Next()
}

func dealData(c *gin.Context) {
	defer func() {
		if p := recover(); p != nil {
			panic(&Api{Code: 111, Message: "error", Retry: 1})
		}
	}()
	c.MustGet("none")
}
