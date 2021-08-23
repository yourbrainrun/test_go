package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yourbrainrun/test_go/test_realip/realip"
)

func main() {
	router := gin.Default()
	router.GET("/test", test)
	_ = router.Run(":777")
}

func test(c *gin.Context) {
	ips := realip.RealIP(c.Request)
	fmt.Println(ips)
	fmt.Println(c.ClientIP())
}
